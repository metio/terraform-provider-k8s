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

type StorageK8SIoStorageClassV1Resource struct{}

var (
	_ resource.Resource = (*StorageK8SIoStorageClassV1Resource)(nil)
)

type StorageK8SIoStorageClassV1TerraformModel struct {
	Id                   types.Int64  `tfsdk:"id"`
	YAML                 types.String `tfsdk:"yaml"`
	ApiVersion           types.String `tfsdk:"api_version"`
	Kind                 types.String `tfsdk:"kind"`
	Metadata             types.Object `tfsdk:"metadata"`
	AllowVolumeExpansion types.Bool   `tfsdk:"allow_volume_expansion"`
	AllowedTopologies    types.List   `tfsdk:"allowed_topologies"`
	MountOptions         types.List   `tfsdk:"mount_options"`
	Parameters           types.Map    `tfsdk:"parameters"`
	Provisioner          types.String `tfsdk:"provisioner"`
	ReclaimPolicy        types.String `tfsdk:"reclaim_policy"`
	VolumeBindingMode    types.String `tfsdk:"volume_binding_mode"`
}

type StorageK8SIoStorageClassV1GoModel struct {
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

	AllowVolumeExpansion *bool `tfsdk:"allow_volume_expansion" yaml:"allowVolumeExpansion,omitempty"`

	AllowedTopologies *[]struct {
		MatchLabelExpressions *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
		} `tfsdk:"match_label_expressions" yaml:"matchLabelExpressions,omitempty"`
	} `tfsdk:"allowed_topologies" yaml:"allowedTopologies,omitempty"`

	MountOptions *[]string `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

	Parameters *map[string]string `tfsdk:"parameters" yaml:"parameters,omitempty"`

	Provisioner *string `tfsdk:"provisioner" yaml:"provisioner,omitempty"`

	ReclaimPolicy *string `tfsdk:"reclaim_policy" yaml:"reclaimPolicy,omitempty"`

	VolumeBindingMode *string `tfsdk:"volume_binding_mode" yaml:"volumeBindingMode,omitempty"`
}

func NewStorageK8SIoStorageClassV1Resource() resource.Resource {
	return &StorageK8SIoStorageClassV1Resource{}
}

func (r *StorageK8SIoStorageClassV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage_k8s_io_storage_class_v1"
}

func (r *StorageK8SIoStorageClassV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
		MarkdownDescription: "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
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

			"allow_volume_expansion": {
				Description:         "AllowVolumeExpansion shows whether the storage class allow volume expand",
				MarkdownDescription: "AllowVolumeExpansion shows whether the storage class allow volume expand",

				Type: types.BoolType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"allowed_topologies": {
				Description:         "Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.",
				MarkdownDescription: "Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.",

				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

					"match_label_expressions": {
						Description:         "A list of topology selector requirements by labels.",
						MarkdownDescription: "A list of topology selector requirements by labels.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "The label key that the selector applies to.",
								MarkdownDescription: "The label key that the selector applies to.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"values": {
								Description:         "An array of string values. One value must match the label to be selected. Each entry in Values is ORed.",
								MarkdownDescription: "An array of string values. One value must match the label to be selected. Each entry in Values is ORed.",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
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

			"mount_options": {
				Description:         "Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ['ro', 'soft']. Not validated - mount of the PVs will simply fail if one is invalid.",
				MarkdownDescription: "Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ['ro', 'soft']. Not validated - mount of the PVs will simply fail if one is invalid.",

				Type: types.ListType{ElemType: types.StringType},

				Required: false,
				Optional: true,
				Computed: false,
			},

			"parameters": {
				Description:         "Parameters holds the parameters for the provisioner that should create volumes of this storage class.",
				MarkdownDescription: "Parameters holds the parameters for the provisioner that should create volumes of this storage class.",

				Type: types.MapType{ElemType: types.StringType},

				Required: false,
				Optional: true,
				Computed: false,
			},

			"provisioner": {
				Description:         "Provisioner indicates the type of the provisioner.",
				MarkdownDescription: "Provisioner indicates the type of the provisioner.",

				Type: types.StringType,

				Required: true,
				Optional: false,
				Computed: false,
			},

			"reclaim_policy": {
				Description:         "Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.",
				MarkdownDescription: "Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},

			"volume_binding_mode": {
				Description:         "VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
				MarkdownDescription: "VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",

				Type: types.StringType,

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *StorageK8SIoStorageClassV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_storage_k8s_io_storage_class_v1")

	var state StorageK8SIoStorageClassV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel StorageK8SIoStorageClassV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("storage.k8s.io/v1")
	goModel.Kind = utilities.Ptr("StorageClass")

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

func (r *StorageK8SIoStorageClassV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_storage_k8s_io_storage_class_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *StorageK8SIoStorageClassV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_storage_k8s_io_storage_class_v1")

	var state StorageK8SIoStorageClassV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel StorageK8SIoStorageClassV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("storage.k8s.io/v1")
	goModel.Kind = utilities.Ptr("StorageClass")

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

func (r *StorageK8SIoStorageClassV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_storage_k8s_io_storage_class_v1")
	// NO-OP: Terraform removes the state automatically for us
}
