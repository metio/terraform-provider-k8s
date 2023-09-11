/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package data_fluid_io_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &DataFluidIoDatasetV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &DataFluidIoDatasetV1Alpha1DataSource{}
)

func NewDataFluidIoDatasetV1Alpha1DataSource() datasource.DataSource {
	return &DataFluidIoDatasetV1Alpha1DataSource{}
}

type DataFluidIoDatasetV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type DataFluidIoDatasetV1Alpha1DataSourceData struct {
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
		AccessModes         *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
		DataRestoreLocation *struct {
			NodeName *string `tfsdk:"node_name" json:"nodeName,omitempty"`
			Path     *string `tfsdk:"path" json:"path,omitempty"`
		} `tfsdk:"data_restore_location" json:"dataRestoreLocation,omitempty"`
		Mounts *[]struct {
			EncryptOptions *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				ValueFrom *struct {
					SecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"encrypt_options" json:"encryptOptions,omitempty"`
			MountPoint *string            `tfsdk:"mount_point" json:"mountPoint,omitempty"`
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			Options    *map[string]string `tfsdk:"options" json:"options,omitempty"`
			Path       *string            `tfsdk:"path" json:"path,omitempty"`
			ReadOnly   *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
			Shared     *bool              `tfsdk:"shared" json:"shared,omitempty"`
		} `tfsdk:"mounts" json:"mounts,omitempty"`
		NodeAffinity *struct {
			Required *struct {
				NodeSelectorTerms *[]struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchFields *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_fields" json:"matchFields,omitempty"`
				} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
			} `tfsdk:"required" json:"required,omitempty"`
		} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
		Owner *struct {
			Gid   *int64  `tfsdk:"gid" json:"gid,omitempty"`
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Uid   *int64  `tfsdk:"uid" json:"uid,omitempty"`
			User  *string `tfsdk:"user" json:"user,omitempty"`
		} `tfsdk:"owner" json:"owner,omitempty"`
		Placement *string `tfsdk:"placement" json:"placement,omitempty"`
		Runtimes  *[]struct {
			Category       *string `tfsdk:"category" json:"category,omitempty"`
			MasterReplicas *int64  `tfsdk:"master_replicas" json:"masterReplicas,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Type           *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"runtimes" json:"runtimes,omitempty"`
		SharedEncryptOptions *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			ValueFrom *struct {
				SecretKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"shared_encrypt_options" json:"sharedEncryptOptions,omitempty"`
		SharedOptions *map[string]string `tfsdk:"shared_options" json:"sharedOptions,omitempty"`
		Tolerations   *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataFluidIoDatasetV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_data_fluid_io_dataset_v1alpha1"
}

func (r *DataFluidIoDatasetV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Dataset is the Schema for the datasets API",
		MarkdownDescription: "Dataset is the Schema for the datasets API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "DatasetSpec defines the desired state of Dataset",
				MarkdownDescription: "DatasetSpec defines the desired state of Dataset",
				Attributes: map[string]schema.Attribute{
					"access_modes": schema.ListAttribute{
						Description:         "AccessModes contains all ways the volume backing the PVC can be mounted",
						MarkdownDescription: "AccessModes contains all ways the volume backing the PVC can be mounted",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"data_restore_location": schema.SingleNestedAttribute{
						Description:         "DataRestoreLocation is the location to load data of dataset  been backuped",
						MarkdownDescription: "DataRestoreLocation is the location to load data of dataset  been backuped",
						Attributes: map[string]schema.Attribute{
							"node_name": schema.StringAttribute{
								Description:         "NodeName describes the nodeName of restore if Path is  in the form of local://subpath",
								MarkdownDescription: "NodeName describes the nodeName of restore if Path is  in the form of local://subpath",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"path": schema.StringAttribute{
								Description:         "Path describes the path of restore, in the form of  local://subpath or pvc://<pvcName>/subpath",
								MarkdownDescription: "Path describes the path of restore, in the form of  local://subpath or pvc://<pvcName>/subpath",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"mounts": schema.ListNestedAttribute{
						Description:         "Mount Points to be mounted on Alluxio.",
						MarkdownDescription: "Mount Points to be mounted on Alluxio.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"encrypt_options": schema.ListNestedAttribute{
									Description:         "The secret information",
									MarkdownDescription: "The secret information",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "The name of encryptOption",
												MarkdownDescription: "The name of encryptOption",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "The valueFrom of encryptOption",
												MarkdownDescription: "The valueFrom of encryptOption",
												Attributes: map[string]schema.Attribute{
													"secret_key_ref": schema.SingleNestedAttribute{
														Description:         "The encryptInfo obtained from secret",
														MarkdownDescription: "The encryptInfo obtained from secret",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The required key in the secret",
																MarkdownDescription: "The required key in the secret",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "The name of required secret",
																MarkdownDescription: "The name of required secret",
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
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"mount_point": schema.StringAttribute{
									Description:         "MountPoint is the mount point of source.",
									MarkdownDescription: "MountPoint is the mount point of source.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "The name of mount",
									MarkdownDescription: "The name of mount",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"options": schema.MapAttribute{
									Description:         "The Mount Options. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Mount Options</a>.  <br> The option has Prefix 'fs.' And you can Learn more from <a href='https://docs.alluxio.io/os/user/stable/en/ufs/S3.html'>The Storage Integrations</a>",
									MarkdownDescription: "The Mount Options. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Mount Options</a>.  <br> The option has Prefix 'fs.' And you can Learn more from <a href='https://docs.alluxio.io/os/user/stable/en/ufs/S3.html'>The Storage Integrations</a>",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"path": schema.StringAttribute{
									Description:         "The path of mount, if not set will be /{Name}",
									MarkdownDescription: "The path of mount, if not set will be /{Name}",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"read_only": schema.BoolAttribute{
									Description:         "Optional: Defaults to false (read-write).",
									MarkdownDescription: "Optional: Defaults to false (read-write).",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"shared": schema.BoolAttribute{
									Description:         "Optional: Defaults to false (shared).",
									MarkdownDescription: "Optional: Defaults to false (shared).",
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

					"node_affinity": schema.SingleNestedAttribute{
						Description:         "NodeAffinity defines constraints that limit what nodes this dataset can be cached to. This field influences the scheduling of pods that use the cached dataset.",
						MarkdownDescription: "NodeAffinity defines constraints that limit what nodes this dataset can be cached to. This field influences the scheduling of pods that use the cached dataset.",
						Attributes: map[string]schema.Attribute{
							"required": schema.SingleNestedAttribute{
								Description:         "Required specifies hard node constraints that must be met.",
								MarkdownDescription: "Required specifies hard node constraints that must be met.",
								Attributes: map[string]schema.Attribute{
									"node_selector_terms": schema.ListNestedAttribute{
										Description:         "Required. A list of node selector terms. The terms are ORed.",
										MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's labels.",
													MarkdownDescription: "A list of node selector requirements by node's labels.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

												"match_fields": schema.ListNestedAttribute{
													Description:         "A list of node selector requirements by node's fields.",
													MarkdownDescription: "A list of node selector requirements by node's fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The label key that the selector applies to.",
																MarkdownDescription: "The label key that the selector applies to.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"operator": schema.StringAttribute{
																Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"values": schema.ListAttribute{
																Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"owner": schema.SingleNestedAttribute{
						Description:         "The owner of the dataset",
						MarkdownDescription: "The owner of the dataset",
						Attributes: map[string]schema.Attribute{
							"gid": schema.Int64Attribute{
								Description:         "The gid to run the alluxio runtime",
								MarkdownDescription: "The gid to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"group": schema.StringAttribute{
								Description:         "The group name to run the alluxio runtime",
								MarkdownDescription: "The group name to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uid": schema.Int64Attribute{
								Description:         "The uid to run the alluxio runtime",
								MarkdownDescription: "The uid to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"user": schema.StringAttribute{
								Description:         "The user name to run the alluxio runtime",
								MarkdownDescription: "The user name to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"placement": schema.StringAttribute{
						Description:         "Manage switch for opening Multiple datasets single node deployment or not TODO(xieydd) In future, evaluate node resources and runtime resources to decide whether to turn them on",
						MarkdownDescription: "Manage switch for opening Multiple datasets single node deployment or not TODO(xieydd) In future, evaluate node resources and runtime resources to decide whether to turn them on",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"runtimes": schema.ListNestedAttribute{
						Description:         "Runtimes for supporting dataset (e.g. AlluxioRuntime)",
						MarkdownDescription: "Runtimes for supporting dataset (e.g. AlluxioRuntime)",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"category": schema.StringAttribute{
									Description:         "Category the runtime object belongs to (e.g. Accelerate)",
									MarkdownDescription: "Category the runtime object belongs to (e.g. Accelerate)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"master_replicas": schema.Int64Attribute{
									Description:         "Runtime master replicas",
									MarkdownDescription: "Runtime master replicas",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the runtime object",
									MarkdownDescription: "Name of the runtime object",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the runtime object",
									MarkdownDescription: "Namespace of the runtime object",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "Runtime object's type (e.g. Alluxio)",
									MarkdownDescription: "Runtime object's type (e.g. Alluxio)",
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

					"shared_encrypt_options": schema.ListNestedAttribute{
						Description:         "SharedEncryptOptions is the encryptOption to all mount",
						MarkdownDescription: "SharedEncryptOptions is the encryptOption to all mount",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "The name of encryptOption",
									MarkdownDescription: "The name of encryptOption",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "The valueFrom of encryptOption",
									MarkdownDescription: "The valueFrom of encryptOption",
									Attributes: map[string]schema.Attribute{
										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "The encryptInfo obtained from secret",
											MarkdownDescription: "The encryptInfo obtained from secret",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The required key in the secret",
													MarkdownDescription: "The required key in the secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "The name of required secret",
													MarkdownDescription: "The name of required secret",
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
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"shared_options": schema.MapAttribute{
						Description:         "SharedOptions is the options to all mount",
						MarkdownDescription: "SharedOptions is the options to all mount",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "If specified, the pod's tolerations.",
						MarkdownDescription: "If specified, the pod's tolerations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *DataFluidIoDatasetV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *DataFluidIoDatasetV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_data_fluid_io_dataset_v1alpha1")

	var data DataFluidIoDatasetV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "data.fluid.io", Version: "v1alpha1", Resource: "datasets"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse DataFluidIoDatasetV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("data.fluid.io/v1alpha1")
	data.Kind = pointer.String("Dataset")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
