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

type StorageK8SIoCSIDriverV1Resource struct{}

var (
	_ resource.Resource = (*StorageK8SIoCSIDriverV1Resource)(nil)
)

type StorageK8SIoCSIDriverV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type StorageK8SIoCSIDriverV1GoModel struct {
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
		AttachRequired *bool `tfsdk:"attach_required" yaml:"attachRequired,omitempty"`

		FsGroupPolicy *string `tfsdk:"fs_group_policy" yaml:"fsGroupPolicy,omitempty"`

		PodInfoOnMount *bool `tfsdk:"pod_info_on_mount" yaml:"podInfoOnMount,omitempty"`

		RequiresRepublish *bool `tfsdk:"requires_republish" yaml:"requiresRepublish,omitempty"`

		SeLinuxMount *bool `tfsdk:"se_linux_mount" yaml:"seLinuxMount,omitempty"`

		StorageCapacity *bool `tfsdk:"storage_capacity" yaml:"storageCapacity,omitempty"`

		TokenRequests *[]struct {
			Audience *string `tfsdk:"audience" yaml:"audience,omitempty"`

			ExpirationSeconds *int64 `tfsdk:"expiration_seconds" yaml:"expirationSeconds,omitempty"`
		} `tfsdk:"token_requests" yaml:"tokenRequests,omitempty"`

		VolumeLifecycleModes *[]string `tfsdk:"volume_lifecycle_modes" yaml:"volumeLifecycleModes,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewStorageK8SIoCSIDriverV1Resource() resource.Resource {
	return &StorageK8SIoCSIDriverV1Resource{}
}

func (r *StorageK8SIoCSIDriverV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_storage_k8s_io_csi_driver_v1"
}

func (r *StorageK8SIoCSIDriverV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.",
		MarkdownDescription: "CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.",
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
				Description:         "CSIDriverSpec is the specification of a CSIDriver.",
				MarkdownDescription: "CSIDriverSpec is the specification of a CSIDriver.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"attach_required": {
						Description:         "attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called.This field is immutable.",
						MarkdownDescription: "attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called.This field is immutable.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"fs_group_policy": {
						Description:         "Defines if the underlying volume supports changing ownership and permission of the volume before being mounted. Refer to the specific FSGroupPolicy values for additional details.This field is immutable.Defaults to ReadWriteOnceWithFSType, which will examine each volume to determine if Kubernetes should modify ownership and permissions of the volume. With the default policy the defined fsGroup will only be applied if a fstype is defined and the volume's access mode contains ReadWriteOnce.",
						MarkdownDescription: "Defines if the underlying volume supports changing ownership and permission of the volume before being mounted. Refer to the specific FSGroupPolicy values for additional details.This field is immutable.Defaults to ReadWriteOnceWithFSType, which will examine each volume to determine if Kubernetes should modify ownership and permissions of the volume. With the default policy the defined fsGroup will only be applied if a fstype is defined and the volume's access mode contains ReadWriteOnce.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_info_on_mount": {
						Description:         "If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeConext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. 'csi.storage.k8s.io/pod.name': pod.Name 'csi.storage.k8s.io/pod.namespace': pod.Namespace 'csi.storage.k8s.io/pod.uid': string(pod.UID) 'csi.storage.k8s.io/ephemeral': 'true' if the volume is an ephemeral inline volume                                defined by a CSIVolumeSource, otherwise 'false''csi.storage.k8s.io/ephemeral' is a new feature in Kubernetes 1.16. It is only required for drivers which support both the 'Persistent' and 'Ephemeral' VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver.This field is immutable.",
						MarkdownDescription: "If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeConext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. 'csi.storage.k8s.io/pod.name': pod.Name 'csi.storage.k8s.io/pod.namespace': pod.Namespace 'csi.storage.k8s.io/pod.uid': string(pod.UID) 'csi.storage.k8s.io/ephemeral': 'true' if the volume is an ephemeral inline volume                                defined by a CSIVolumeSource, otherwise 'false''csi.storage.k8s.io/ephemeral' is a new feature in Kubernetes 1.16. It is only required for drivers which support both the 'Persistent' and 'Ephemeral' VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver.This field is immutable.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"requires_republish": {
						Description:         "RequiresRepublish indicates the CSI driver wants 'NodePublishVolume' being periodically called to reflect any possible change in the mounted volume. This field defaults to false.Note: After a successful initial NodePublishVolume call, subsequent calls to NodePublishVolume should only update the contents of the volume. New mount points will not be seen by a running container.",
						MarkdownDescription: "RequiresRepublish indicates the CSI driver wants 'NodePublishVolume' being periodically called to reflect any possible change in the mounted volume. This field defaults to false.Note: After a successful initial NodePublishVolume call, subsequent calls to NodePublishVolume should only update the contents of the volume. New mount points will not be seen by a running container.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"se_linux_mount": {
						Description:         "SELinuxMount specifies if the CSI driver supports '-o context' mount option.When 'true', the CSI driver must ensure that all volumes provided by this CSI driver can be mounted separately with different '-o context' options. This is typical for storage backends that provide volumes as filesystems on block devices or as independent shared volumes. Kubernetes will call NodeStage / NodePublish with '-o context=xyz' mount option when mounting a ReadWriteOncePod volume used in Pod that has explicitly set SELinux context. In the future, it may be expanded to other volume AccessModes. In any case, Kubernetes will ensure that the volume is mounted only with a single SELinux context.When 'false', Kubernetes won't pass any special SELinux mount options to the driver. This is typical for volumes that represent subdirectories of a bigger shared filesystem.Default is 'false'.",
						MarkdownDescription: "SELinuxMount specifies if the CSI driver supports '-o context' mount option.When 'true', the CSI driver must ensure that all volumes provided by this CSI driver can be mounted separately with different '-o context' options. This is typical for storage backends that provide volumes as filesystems on block devices or as independent shared volumes. Kubernetes will call NodeStage / NodePublish with '-o context=xyz' mount option when mounting a ReadWriteOncePod volume used in Pod that has explicitly set SELinux context. In the future, it may be expanded to other volume AccessModes. In any case, Kubernetes will ensure that the volume is mounted only with a single SELinux context.When 'false', Kubernetes won't pass any special SELinux mount options to the driver. This is typical for volumes that represent subdirectories of a bigger shared filesystem.Default is 'false'.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_capacity": {
						Description:         "If set to true, storageCapacity indicates that the CSI volume driver wants pod scheduling to consider the storage capacity that the driver deployment will report by creating CSIStorageCapacity objects with capacity information.The check can be enabled immediately when deploying a driver. In that case, provisioning new volumes with late binding will pause until the driver deployment has published some suitable CSIStorageCapacity object.Alternatively, the driver can be deployed with the field unset or false and it can be flipped later when storage capacity information has been published.This field was immutable in Kubernetes <= 1.22 and now is mutable.",
						MarkdownDescription: "If set to true, storageCapacity indicates that the CSI volume driver wants pod scheduling to consider the storage capacity that the driver deployment will report by creating CSIStorageCapacity objects with capacity information.The check can be enabled immediately when deploying a driver. In that case, provisioning new volumes with late binding will pause until the driver deployment has published some suitable CSIStorageCapacity object.Alternatively, the driver can be deployed with the field unset or false and it can be flipped later when storage capacity information has been published.This field was immutable in Kubernetes <= 1.22 and now is mutable.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"token_requests": {
						Description:         "TokenRequests indicates the CSI driver needs pods' service account tokens it is mounting volume for to do necessary authentication. Kubelet will pass the tokens in VolumeContext in the CSI NodePublishVolume calls. The CSI driver should parse and validate the following VolumeContext: 'csi.storage.k8s.io/serviceAccount.tokens': {  '<audience>': {    'token': <token>,    'expirationTimestamp': <expiration timestamp in RFC3339>,  },  ...}Note: Audience in each TokenRequest should be different and at most one token is empty string. To receive a new token after expiry, RequiresRepublish can be used to trigger NodePublishVolume periodically.",
						MarkdownDescription: "TokenRequests indicates the CSI driver needs pods' service account tokens it is mounting volume for to do necessary authentication. Kubelet will pass the tokens in VolumeContext in the CSI NodePublishVolume calls. The CSI driver should parse and validate the following VolumeContext: 'csi.storage.k8s.io/serviceAccount.tokens': {  '<audience>': {    'token': <token>,    'expirationTimestamp': <expiration timestamp in RFC3339>,  },  ...}Note: Audience in each TokenRequest should be different and at most one token is empty string. To receive a new token after expiry, RequiresRepublish can be used to trigger NodePublishVolume periodically.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"audience": {
								Description:         "Audience is the intended audience of the token in 'TokenRequestSpec'. It will default to the audiences of kube apiserver.",
								MarkdownDescription: "Audience is the intended audience of the token in 'TokenRequestSpec'. It will default to the audiences of kube apiserver.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"expiration_seconds": {
								Description:         "ExpirationSeconds is the duration of validity of the token in 'TokenRequestSpec'. It has the same default value of 'ExpirationSeconds' in 'TokenRequestSpec'.",
								MarkdownDescription: "ExpirationSeconds is the duration of validity of the token in 'TokenRequestSpec'. It has the same default value of 'ExpirationSeconds' in 'TokenRequestSpec'.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_lifecycle_modes": {
						Description:         "volumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is 'Persistent', which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is 'Ephemeral'. In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future. This field is beta.This field is immutable.",
						MarkdownDescription: "volumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is 'Persistent', which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is 'Ephemeral'. In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future. This field is beta.This field is immutable.",

						Type: types.ListType{ElemType: types.StringType},

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

func (r *StorageK8SIoCSIDriverV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_storage_k8s_io_csi_driver_v1")

	var state StorageK8SIoCSIDriverV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel StorageK8SIoCSIDriverV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("storage.k8s.io/v1")
	goModel.Kind = utilities.Ptr("CSIDriver")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *StorageK8SIoCSIDriverV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_storage_k8s_io_csi_driver_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *StorageK8SIoCSIDriverV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_storage_k8s_io_csi_driver_v1")

	var state StorageK8SIoCSIDriverV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel StorageK8SIoCSIDriverV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("storage.k8s.io/v1")
	goModel.Kind = utilities.Ptr("CSIDriver")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *StorageK8SIoCSIDriverV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_storage_k8s_io_csi_driver_v1")
	// NO-OP: Terraform removes the state automatically for us
}
