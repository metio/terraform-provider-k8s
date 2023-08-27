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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &StorageK8SIoCSIDriverV1DataSource{}
	_ datasource.DataSourceWithConfigure = &StorageK8SIoCSIDriverV1DataSource{}
)

func NewStorageK8SIoCSIDriverV1DataSource() datasource.DataSource {
	return &StorageK8SIoCSIDriverV1DataSource{}
}

type StorageK8SIoCSIDriverV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type StorageK8SIoCSIDriverV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AttachRequired    *bool   `tfsdk:"attach_required" json:"attachRequired,omitempty"`
		FsGroupPolicy     *string `tfsdk:"fs_group_policy" json:"fsGroupPolicy,omitempty"`
		PodInfoOnMount    *bool   `tfsdk:"pod_info_on_mount" json:"podInfoOnMount,omitempty"`
		RequiresRepublish *bool   `tfsdk:"requires_republish" json:"requiresRepublish,omitempty"`
		SeLinuxMount      *bool   `tfsdk:"se_linux_mount" json:"seLinuxMount,omitempty"`
		StorageCapacity   *bool   `tfsdk:"storage_capacity" json:"storageCapacity,omitempty"`
		TokenRequests     *[]struct {
			Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
			ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
		} `tfsdk:"token_requests" json:"tokenRequests,omitempty"`
		VolumeLifecycleModes *[]string `tfsdk:"volume_lifecycle_modes" json:"volumeLifecycleModes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StorageK8SIoCSIDriverV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_storage_k8s_io_csi_driver_v1"
}

func (r *StorageK8SIoCSIDriverV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.",
		MarkdownDescription: "CSIDriver captures information about a Container Storage Interface (CSI) volume driver deployed on the cluster. Kubernetes attach detach controller uses this object to determine whether attach is required. Kubelet uses this object to determine whether pod information needs to be passed on mount. CSIDriver objects are non-namespaced.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "CSIDriverSpec is the specification of a CSIDriver.",
				MarkdownDescription: "CSIDriverSpec is the specification of a CSIDriver.",
				Attributes: map[string]schema.Attribute{
					"attach_required": schema.BoolAttribute{
						Description:         "attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called.This field is immutable.",
						MarkdownDescription: "attachRequired indicates this CSI volume driver requires an attach operation (because it implements the CSI ControllerPublishVolume() method), and that the Kubernetes attach detach controller should call the attach volume interface which checks the volumeattachment status and waits until the volume is attached before proceeding to mounting. The CSI external-attacher coordinates with CSI volume driver and updates the volumeattachment status when the attach operation is complete. If the CSIDriverRegistry feature gate is enabled and the value is specified to false, the attach operation will be skipped. Otherwise the attach operation will be called.This field is immutable.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"fs_group_policy": schema.StringAttribute{
						Description:         "Defines if the underlying volume supports changing ownership and permission of the volume before being mounted. Refer to the specific FSGroupPolicy values for additional details.This field is immutable.Defaults to ReadWriteOnceWithFSType, which will examine each volume to determine if Kubernetes should modify ownership and permissions of the volume. With the default policy the defined fsGroup will only be applied if a fstype is defined and the volume's access mode contains ReadWriteOnce.",
						MarkdownDescription: "Defines if the underlying volume supports changing ownership and permission of the volume before being mounted. Refer to the specific FSGroupPolicy values for additional details.This field is immutable.Defaults to ReadWriteOnceWithFSType, which will examine each volume to determine if Kubernetes should modify ownership and permissions of the volume. With the default policy the defined fsGroup will only be applied if a fstype is defined and the volume's access mode contains ReadWriteOnce.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pod_info_on_mount": schema.BoolAttribute{
						Description:         "If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeConext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. 'csi.storage.k8s.io/pod.name': pod.Name 'csi.storage.k8s.io/pod.namespace': pod.Namespace 'csi.storage.k8s.io/pod.uid': string(pod.UID) 'csi.storage.k8s.io/ephemeral': 'true' if the volume is an ephemeral inline volume                                defined by a CSIVolumeSource, otherwise 'false''csi.storage.k8s.io/ephemeral' is a new feature in Kubernetes 1.16. It is only required for drivers which support both the 'Persistent' and 'Ephemeral' VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver.This field is immutable.",
						MarkdownDescription: "If set to true, podInfoOnMount indicates this CSI volume driver requires additional pod information (like podName, podUID, etc.) during mount operations. If set to false, pod information will not be passed on mount. Default is false. The CSI driver specifies podInfoOnMount as part of driver deployment. If true, Kubelet will pass pod information as VolumeContext in the CSI NodePublishVolume() calls. The CSI driver is responsible for parsing and validating the information passed in as VolumeContext. The following VolumeConext will be passed if podInfoOnMount is set to true. This list might grow, but the prefix will be used. 'csi.storage.k8s.io/pod.name': pod.Name 'csi.storage.k8s.io/pod.namespace': pod.Namespace 'csi.storage.k8s.io/pod.uid': string(pod.UID) 'csi.storage.k8s.io/ephemeral': 'true' if the volume is an ephemeral inline volume                                defined by a CSIVolumeSource, otherwise 'false''csi.storage.k8s.io/ephemeral' is a new feature in Kubernetes 1.16. It is only required for drivers which support both the 'Persistent' and 'Ephemeral' VolumeLifecycleMode. Other drivers can leave pod info disabled and/or ignore this field. As Kubernetes 1.15 doesn't support this field, drivers can only support one mode when deployed on such a cluster and the deployment determines which mode that is, for example via a command line parameter of the driver.This field is immutable.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"requires_republish": schema.BoolAttribute{
						Description:         "RequiresRepublish indicates the CSI driver wants 'NodePublishVolume' being periodically called to reflect any possible change in the mounted volume. This field defaults to false.Note: After a successful initial NodePublishVolume call, subsequent calls to NodePublishVolume should only update the contents of the volume. New mount points will not be seen by a running container.",
						MarkdownDescription: "RequiresRepublish indicates the CSI driver wants 'NodePublishVolume' being periodically called to reflect any possible change in the mounted volume. This field defaults to false.Note: After a successful initial NodePublishVolume call, subsequent calls to NodePublishVolume should only update the contents of the volume. New mount points will not be seen by a running container.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"se_linux_mount": schema.BoolAttribute{
						Description:         "SELinuxMount specifies if the CSI driver supports '-o context' mount option.When 'true', the CSI driver must ensure that all volumes provided by this CSI driver can be mounted separately with different '-o context' options. This is typical for storage backends that provide volumes as filesystems on block devices or as independent shared volumes. Kubernetes will call NodeStage / NodePublish with '-o context=xyz' mount option when mounting a ReadWriteOncePod volume used in Pod that has explicitly set SELinux context. In the future, it may be expanded to other volume AccessModes. In any case, Kubernetes will ensure that the volume is mounted only with a single SELinux context.When 'false', Kubernetes won't pass any special SELinux mount options to the driver. This is typical for volumes that represent subdirectories of a bigger shared filesystem.Default is 'false'.",
						MarkdownDescription: "SELinuxMount specifies if the CSI driver supports '-o context' mount option.When 'true', the CSI driver must ensure that all volumes provided by this CSI driver can be mounted separately with different '-o context' options. This is typical for storage backends that provide volumes as filesystems on block devices or as independent shared volumes. Kubernetes will call NodeStage / NodePublish with '-o context=xyz' mount option when mounting a ReadWriteOncePod volume used in Pod that has explicitly set SELinux context. In the future, it may be expanded to other volume AccessModes. In any case, Kubernetes will ensure that the volume is mounted only with a single SELinux context.When 'false', Kubernetes won't pass any special SELinux mount options to the driver. This is typical for volumes that represent subdirectories of a bigger shared filesystem.Default is 'false'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"storage_capacity": schema.BoolAttribute{
						Description:         "If set to true, storageCapacity indicates that the CSI volume driver wants pod scheduling to consider the storage capacity that the driver deployment will report by creating CSIStorageCapacity objects with capacity information.The check can be enabled immediately when deploying a driver. In that case, provisioning new volumes with late binding will pause until the driver deployment has published some suitable CSIStorageCapacity object.Alternatively, the driver can be deployed with the field unset or false and it can be flipped later when storage capacity information has been published.This field was immutable in Kubernetes <= 1.22 and now is mutable.",
						MarkdownDescription: "If set to true, storageCapacity indicates that the CSI volume driver wants pod scheduling to consider the storage capacity that the driver deployment will report by creating CSIStorageCapacity objects with capacity information.The check can be enabled immediately when deploying a driver. In that case, provisioning new volumes with late binding will pause until the driver deployment has published some suitable CSIStorageCapacity object.Alternatively, the driver can be deployed with the field unset or false and it can be flipped later when storage capacity information has been published.This field was immutable in Kubernetes <= 1.22 and now is mutable.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"token_requests": schema.ListNestedAttribute{
						Description:         "TokenRequests indicates the CSI driver needs pods' service account tokens it is mounting volume for to do necessary authentication. Kubelet will pass the tokens in VolumeContext in the CSI NodePublishVolume calls. The CSI driver should parse and validate the following VolumeContext: 'csi.storage.k8s.io/serviceAccount.tokens': {  '<audience>': {    'token': <token>,    'expirationTimestamp': <expiration timestamp in RFC3339>,  },  ...}Note: Audience in each TokenRequest should be different and at most one token is empty string. To receive a new token after expiry, RequiresRepublish can be used to trigger NodePublishVolume periodically.",
						MarkdownDescription: "TokenRequests indicates the CSI driver needs pods' service account tokens it is mounting volume for to do necessary authentication. Kubelet will pass the tokens in VolumeContext in the CSI NodePublishVolume calls. The CSI driver should parse and validate the following VolumeContext: 'csi.storage.k8s.io/serviceAccount.tokens': {  '<audience>': {    'token': <token>,    'expirationTimestamp': <expiration timestamp in RFC3339>,  },  ...}Note: Audience in each TokenRequest should be different and at most one token is empty string. To receive a new token after expiry, RequiresRepublish can be used to trigger NodePublishVolume periodically.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"audience": schema.StringAttribute{
									Description:         "Audience is the intended audience of the token in 'TokenRequestSpec'. It will default to the audiences of kube apiserver.",
									MarkdownDescription: "Audience is the intended audience of the token in 'TokenRequestSpec'. It will default to the audiences of kube apiserver.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"expiration_seconds": schema.Int64Attribute{
									Description:         "ExpirationSeconds is the duration of validity of the token in 'TokenRequestSpec'. It has the same default value of 'ExpirationSeconds' in 'TokenRequestSpec'.",
									MarkdownDescription: "ExpirationSeconds is the duration of validity of the token in 'TokenRequestSpec'. It has the same default value of 'ExpirationSeconds' in 'TokenRequestSpec'.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"volume_lifecycle_modes": schema.ListAttribute{
						Description:         "volumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is 'Persistent', which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is 'Ephemeral'. In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future. This field is beta.This field is immutable.",
						MarkdownDescription: "volumeLifecycleModes defines what kind of volumes this CSI volume driver supports. The default if the list is empty is 'Persistent', which is the usage defined by the CSI specification and implemented in Kubernetes via the usual PV/PVC mechanism. The other mode is 'Ephemeral'. In this mode, volumes are defined inline inside the pod spec with CSIVolumeSource and their lifecycle is tied to the lifecycle of that pod. A driver has to be aware of this because it is only going to get a NodePublishVolume call for such a volume. For more information about implementing this mode, see https://kubernetes-csi.github.io/docs/ephemeral-local-volumes.html A driver can support one or more of these modes and more modes may be added in the future. This field is beta.This field is immutable.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *StorageK8SIoCSIDriverV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *StorageK8SIoCSIDriverV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_storage_k8s_io_csi_driver_v1")

	var data StorageK8SIoCSIDriverV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "CSIDriver"}).
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

	var readResponse StorageK8SIoCSIDriverV1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("storage.k8s.io/v1")
	data.Kind = pointer.String("CSIDriver")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
