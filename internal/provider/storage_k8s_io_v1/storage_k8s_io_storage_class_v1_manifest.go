/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package storage_k8s_io_v1

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
	_ datasource.DataSource = &StorageK8SIoStorageClassV1Manifest{}
)

func NewStorageK8SIoStorageClassV1Manifest() datasource.DataSource {
	return &StorageK8SIoStorageClassV1Manifest{}
}

type StorageK8SIoStorageClassV1Manifest struct{}

type StorageK8SIoStorageClassV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

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

func (r *StorageK8SIoStorageClassV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_storage_k8s_io_storage_class_v1_manifest"
}

func (r *StorageK8SIoStorageClassV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned. StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
		MarkdownDescription: "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned. StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
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
				Description:         "volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
				MarkdownDescription: "volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},
		},
	}
}

func (r *StorageK8SIoStorageClassV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_storage_k8s_io_storage_class_v1_manifest")

	var model StorageK8SIoStorageClassV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("storage.k8s.io/v1")
	model.Kind = pointer.String("StorageClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
