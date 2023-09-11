/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package storage_k8s_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
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
	"time"
)

var (
	_ resource.Resource                = &StorageK8SIoStorageClassV1Resource{}
	_ resource.ResourceWithConfigure   = &StorageK8SIoStorageClassV1Resource{}
	_ resource.ResourceWithImportState = &StorageK8SIoStorageClassV1Resource{}
)

func NewStorageK8SIoStorageClassV1Resource() resource.Resource {
	return &StorageK8SIoStorageClassV1Resource{}
}

type StorageK8SIoStorageClassV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type StorageK8SIoStorageClassV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	AllowVolumeExpansion *bool `tfsdk:"allow_volume_expansion" json:"allowVolumeExpansion,omitempty"`
	AllowedTopologies    *[]struct {
		MatchLabelExpressions *[]struct {
			Key    *string   `tfsdk:"key" json:"key,omitempty"`
			Values *[]string `tfsdk:"values" json:"values,omitempty"`
		} `tfsdk:"match_label_expressions" json:"matchLabelExpressions,omitempty"`
	} `tfsdk:"allowed_topologies" json:"allowedTopologies,omitempty"`
	MountOptions      *[]string          `tfsdk:"mount_options" json:"mountOptions,omitempty"`
	Parameters        *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
	Provisioner       *string            `tfsdk:"k8s_provisioner" json:"provisioner,omitempty"`
	ReclaimPolicy     *string            `tfsdk:"reclaim_policy" json:"reclaimPolicy,omitempty"`
	VolumeBindingMode *string            `tfsdk:"volume_binding_mode" json:"volumeBindingMode,omitempty"`
}

func (r *StorageK8SIoStorageClassV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_storage_k8s_io_storage_class_v1"
}

func (r *StorageK8SIoStorageClassV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
		MarkdownDescription: "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"wait_for_upsert": schema.ListNestedAttribute{
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
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
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

			"allow_volume_expansion": schema.BoolAttribute{
				Description:         "allowVolumeExpansion shows whether the storage class allow volume expand.",
				MarkdownDescription: "allowVolumeExpansion shows whether the storage class allow volume expand.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"allowed_topologies": schema.ListNestedAttribute{
				Description:         "allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.",
				MarkdownDescription: "allowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"match_label_expressions": schema.ListNestedAttribute{
							Description:         "A list of topology selector requirements by labels.",
							MarkdownDescription: "A list of topology selector requirements by labels.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The label key that the selector applies to.",
										MarkdownDescription: "The label key that the selector applies to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"values": schema.ListAttribute{
										Description:         "An array of string values. One value must match the label to be selected. Each entry in Values is ORed.",
										MarkdownDescription: "An array of string values. One value must match the label to be selected. Each entry in Values is ORed.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"mount_options": schema.ListAttribute{
				Description:         "mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ['ro', 'soft']. Not validated - mount of the PVs will simply fail if one is invalid.",
				MarkdownDescription: "mountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ['ro', 'soft']. Not validated - mount of the PVs will simply fail if one is invalid.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"parameters": schema.MapAttribute{
				Description:         "parameters holds the parameters for the provisioner that should create volumes of this storage class.",
				MarkdownDescription: "parameters holds the parameters for the provisioner that should create volumes of this storage class.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"k8s_provisioner": schema.StringAttribute{
				Description:         "provisioner indicates the type of the provisioner.",
				MarkdownDescription: "provisioner indicates the type of the provisioner.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"reclaim_policy": schema.StringAttribute{
				Description:         "reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.",
				MarkdownDescription: "reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"volume_binding_mode": schema.StringAttribute{
				Description:         "volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
				MarkdownDescription: "volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},
		},
	}
}

func (r *StorageK8SIoStorageClassV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *StorageK8SIoStorageClassV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_storage_k8s_io_storage_class_v1")

	var model StorageK8SIoStorageClassV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("storage.k8s.io/v1")
	model.Kind = pointer.String("StorageClass")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse StorageK8SIoStorageClassV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.AllowVolumeExpansion = readResponse.AllowVolumeExpansion
	model.AllowedTopologies = readResponse.AllowedTopologies
	model.MountOptions = readResponse.MountOptions
	model.Parameters = readResponse.Parameters
	model.Provisioner = readResponse.Provisioner
	model.ReclaimPolicy = readResponse.ReclaimPolicy
	model.VolumeBindingMode = readResponse.VolumeBindingMode

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *StorageK8SIoStorageClassV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_storage_k8s_io_storage_class_v1")

	var data StorageK8SIoStorageClassV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse StorageK8SIoStorageClassV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.AllowVolumeExpansion = readResponse.AllowVolumeExpansion
	data.AllowedTopologies = readResponse.AllowedTopologies
	data.MountOptions = readResponse.MountOptions
	data.Parameters = readResponse.Parameters
	data.Provisioner = readResponse.Provisioner
	data.ReclaimPolicy = readResponse.ReclaimPolicy
	data.VolumeBindingMode = readResponse.VolumeBindingMode

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *StorageK8SIoStorageClassV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_storage_k8s_io_storage_class_v1")

	var model StorageK8SIoStorageClassV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("storage.k8s.io/v1")
	model.Kind = pointer.String("StorageClass")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse StorageK8SIoStorageClassV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.AllowVolumeExpansion = readResponse.AllowVolumeExpansion
	model.AllowedTopologies = readResponse.AllowedTopologies
	model.MountOptions = readResponse.MountOptions
	model.Parameters = readResponse.Parameters
	model.Provisioner = readResponse.Provisioner
	model.ReclaimPolicy = readResponse.ReclaimPolicy
	model.VolumeBindingMode = readResponse.VolumeBindingMode

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *StorageK8SIoStorageClassV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_storage_k8s_io_storage_class_v1")

	var data StorageK8SIoStorageClassV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *StorageK8SIoStorageClassV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
