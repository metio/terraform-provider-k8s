/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package argoproj_io_v1alpha1

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
	_ datasource.DataSource              = &ArgoprojIoArgoCdexportV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ArgoprojIoArgoCdexportV1Alpha1DataSource{}
)

func NewArgoprojIoArgoCdexportV1Alpha1DataSource() datasource.DataSource {
	return &ArgoprojIoArgoCdexportV1Alpha1DataSource{}
}

type ArgoprojIoArgoCdexportV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ArgoprojIoArgoCdexportV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Argocd   *string `tfsdk:"argocd" json:"argocd,omitempty"`
		Image    *string `tfsdk:"image" json:"image,omitempty"`
		Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
		Storage  *struct {
			Backend *string `tfsdk:"backend" json:"backend,omitempty"`
			Pvc     *struct {
				AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				DataSource  *struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"data_source" json:"dataSource,omitempty"`
				DataSourceRef *struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"pvc" json:"pvc,omitempty"`
			SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ArgoprojIoArgoCdexportV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_argoproj_io_argo_cd_export_v1alpha1"
}

func (r *ArgoprojIoArgoCdexportV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ArgoCDExport is the Schema for the argocdexports API",
		MarkdownDescription: "ArgoCDExport is the Schema for the argocdexports API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ArgoCDExportSpec defines the desired state of ArgoCDExport",
				MarkdownDescription: "ArgoCDExportSpec defines the desired state of ArgoCDExport",
				Attributes: map[string]schema.Attribute{
					"argocd": schema.StringAttribute{
						Description:         "Argocd is the name of the ArgoCD instance to export.",
						MarkdownDescription: "Argocd is the name of the ArgoCD instance to export.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image": schema.StringAttribute{
						Description:         "Image is the container image to use for the export Job.",
						MarkdownDescription: "Image is the container image to use for the export Job.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"schedule": schema.StringAttribute{
						Description:         "Schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.",
						MarkdownDescription: "Schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage defines the storage configuration options.",
						MarkdownDescription: "Storage defines the storage configuration options.",
						Attributes: map[string]schema.Attribute{
							"backend": schema.StringAttribute{
								Description:         "Backend defines the storage backend to use, must be 'local' (the default), 'aws', 'azure' or 'gcp'.",
								MarkdownDescription: "Backend defines the storage backend to use, must be 'local' (the default), 'aws', 'azure' or 'gcp'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pvc": schema.SingleNestedAttribute{
								Description:         "PVC is the desired characteristics for a PersistentVolumeClaim.",
								MarkdownDescription: "PVC is the desired characteristics for a PersistentVolumeClaim.",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"data_source": schema.SingleNestedAttribute{
										Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
										MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"data_source_ref": schema.SingleNestedAttribute{
										Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
										MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"selector": schema.SingleNestedAttribute{
										Description:         "A label query over volumes to consider for binding.",
										MarkdownDescription: "A label query over volumes to consider for binding.",
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
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

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

									"storage_class_name": schema.StringAttribute{
										Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volume_mode": schema.StringAttribute{
										Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
										MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volume_name": schema.StringAttribute{
										Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
										MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName is the name of a Secret with encryption key, credentials, etc.",
								MarkdownDescription: "SecretName is the name of a Secret with encryption key, credentials, etc.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"version": schema.StringAttribute{
						Description:         "Version is the tag/digest to use for the export Job container image.",
						MarkdownDescription: "Version is the tag/digest to use for the export Job container image.",
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

func (r *ArgoprojIoArgoCdexportV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *ArgoprojIoArgoCdexportV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_argoproj_io_argo_cd_export_v1alpha1")

	var data ArgoprojIoArgoCdexportV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "argocdexports"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ArgoprojIoArgoCdexportV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("argoproj.io/v1alpha1")
	data.Kind = pointer.String("ArgoCDExport")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
