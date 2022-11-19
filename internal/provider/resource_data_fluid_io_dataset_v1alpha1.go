/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type DataFluidIoDatasetV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*DataFluidIoDatasetV1Alpha1Resource)(nil)
)

type DataFluidIoDatasetV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type DataFluidIoDatasetV1Alpha1GoModel struct {
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
		AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

		DataRestoreLocation *struct {
			NodeName *string `tfsdk:"node_name" yaml:"nodeName,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`
		} `tfsdk:"data_restore_location" yaml:"dataRestoreLocation,omitempty"`

		Mounts *[]struct {
			EncryptOptions *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ValueFrom *struct {
					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"encrypt_options" yaml:"encryptOptions,omitempty"`

			MountPoint *string `tfsdk:"mount_point" yaml:"mountPoint,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Options *map[string]string `tfsdk:"options" yaml:"options,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			Shared *bool `tfsdk:"shared" yaml:"shared,omitempty"`
		} `tfsdk:"mounts" yaml:"mounts,omitempty"`

		NodeAffinity *struct {
			Required *struct {
				NodeSelectorTerms *[]struct {
					MatchExpressions *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

					MatchFields *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
					} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
				} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
			} `tfsdk:"required" yaml:"required,omitempty"`
		} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

		Owner *struct {
			Gid *int64 `tfsdk:"gid" yaml:"gid,omitempty"`

			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Uid *int64 `tfsdk:"uid" yaml:"uid,omitempty"`

			User *string `tfsdk:"user" yaml:"user,omitempty"`
		} `tfsdk:"owner" yaml:"owner,omitempty"`

		Placement *string `tfsdk:"placement" yaml:"placement,omitempty"`

		Runtimes *[]struct {
			Category *string `tfsdk:"category" yaml:"category,omitempty"`

			MasterReplicas *int64 `tfsdk:"master_replicas" yaml:"masterReplicas,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"runtimes" yaml:"runtimes,omitempty"`

		Tolerations *[]struct {
			Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

			TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewDataFluidIoDatasetV1Alpha1Resource() resource.Resource {
	return &DataFluidIoDatasetV1Alpha1Resource{}
}

func (r *DataFluidIoDatasetV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_data_fluid_io_dataset_v1alpha1"
}

func (r *DataFluidIoDatasetV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Dataset is the Schema for the datasets API",
		MarkdownDescription: "Dataset is the Schema for the datasets API",
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
				Description:         "DatasetSpec defines the desired state of Dataset",
				MarkdownDescription: "DatasetSpec defines the desired state of Dataset",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"access_modes": {
						Description:         "AccessModes contains all ways the volume backing the PVC can be mounted",
						MarkdownDescription: "AccessModes contains all ways the volume backing the PVC can be mounted",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"data_restore_location": {
						Description:         "DataRestoreLocation is the location to load data of dataset  been backuped",
						MarkdownDescription: "DataRestoreLocation is the location to load data of dataset  been backuped",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_name": {
								Description:         "NodeName describes the nodeName of restore if Path is  in the form of local://subpath",
								MarkdownDescription: "NodeName describes the nodeName of restore if Path is  in the form of local://subpath",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "Path describes the path of restore, in the form of  local://subpath or pvc://<pvcName>/subpath",
								MarkdownDescription: "Path describes the path of restore, in the form of  local://subpath or pvc://<pvcName>/subpath",

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

					"mounts": {
						Description:         "Mount Points to be mounted on Alluxio.",
						MarkdownDescription: "Mount Points to be mounted on Alluxio.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"encrypt_options": {
								Description:         "The secret information",
								MarkdownDescription: "The secret information",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "The name of encryptOption",
										MarkdownDescription: "The name of encryptOption",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_from": {
										Description:         "The valueFrom of encryptOption",
										MarkdownDescription: "The valueFrom of encryptOption",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_key_ref": {
												Description:         "The encryptInfo obtained from secret",
												MarkdownDescription: "The encryptInfo obtained from secret",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The required key in the secret",
														MarkdownDescription: "The required key in the secret",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "The name of required secret",
														MarkdownDescription: "The name of required secret",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount_point": {
								Description:         "MountPoint is the mount point of source.",
								MarkdownDescription: "MountPoint is the mount point of source.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(5),
								},
							},

							"name": {
								Description:         "The name of mount",
								MarkdownDescription: "The name of mount",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(0),
								},
							},

							"options": {
								Description:         "The Mount Options. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Mount Options</a>.  <br> The option has Prefix 'fs.' And you can Learn more from <a href='https://docs.alluxio.io/os/user/stable/en/ufs/S3.html'>The Storage Integrations</a>",
								MarkdownDescription: "The Mount Options. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Mount Options</a>.  <br> The option has Prefix 'fs.' And you can Learn more from <a href='https://docs.alluxio.io/os/user/stable/en/ufs/S3.html'>The Storage Integrations</a>",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "The path of mount, if not set will be /{Name}",
								MarkdownDescription: "The path of mount, if not set will be /{Name}",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_only": {
								Description:         "Optional: Defaults to false (read-write).",
								MarkdownDescription: "Optional: Defaults to false (read-write).",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"shared": {
								Description:         "Optional: Defaults to false (shared).",
								MarkdownDescription: "Optional: Defaults to false (shared).",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_affinity": {
						Description:         "NodeAffinity defines constraints that limit what nodes this dataset can be cached to. This field influences the scheduling of pods that use the cached dataset.",
						MarkdownDescription: "NodeAffinity defines constraints that limit what nodes this dataset can be cached to. This field influences the scheduling of pods that use the cached dataset.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"required": {
								Description:         "Required specifies hard node constraints that must be met.",
								MarkdownDescription: "Required specifies hard node constraints that must be met.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_selector_terms": {
										Description:         "Required. A list of node selector terms. The terms are ORed.",
										MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"match_expressions": {
												Description:         "A list of node selector requirements by node's labels.",
												MarkdownDescription: "A list of node selector requirements by node's labels.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The label key that the selector applies to.",
														MarkdownDescription: "The label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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

											"match_fields": {
												Description:         "A list of node selector requirements by node's fields.",
												MarkdownDescription: "A list of node selector requirements by node's fields.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The label key that the selector applies to.",
														MarkdownDescription: "The label key that the selector applies to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"operator": {
														Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
														MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"values": {
														Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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
										}),

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

					"owner": {
						Description:         "The owner of the dataset",
						MarkdownDescription: "The owner of the dataset",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"gid": {
								Description:         "The gid to run the alluxio runtime",
								MarkdownDescription: "The gid to run the alluxio runtime",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"group": {
								Description:         "The group name to run the alluxio runtime",
								MarkdownDescription: "The group name to run the alluxio runtime",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"uid": {
								Description:         "The uid to run the alluxio runtime",
								MarkdownDescription: "The uid to run the alluxio runtime",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"user": {
								Description:         "The user name to run the alluxio runtime",
								MarkdownDescription: "The user name to run the alluxio runtime",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"placement": {
						Description:         "Manage switch for opening Multiple datasets single node deployment or not TODO(xieydd) In future, evaluate node resources and runtime resources to decide whether to turn them on",
						MarkdownDescription: "Manage switch for opening Multiple datasets single node deployment or not TODO(xieydd) In future, evaluate node resources and runtime resources to decide whether to turn them on",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("Exclusive", "", "Shared"),
						},
					},

					"runtimes": {
						Description:         "Runtimes for supporting dataset (e.g. AlluxioRuntime)",
						MarkdownDescription: "Runtimes for supporting dataset (e.g. AlluxioRuntime)",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"category": {
								Description:         "Category the runtime object belongs to (e.g. Accelerate)",
								MarkdownDescription: "Category the runtime object belongs to (e.g. Accelerate)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"master_replicas": {
								Description:         "Runtime master replicas",
								MarkdownDescription: "Runtime master replicas",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the runtime object",
								MarkdownDescription: "Name of the runtime object",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the runtime object",
								MarkdownDescription: "Namespace of the runtime object",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Runtime object's type (e.g. Alluxio)",
								MarkdownDescription: "Runtime object's type (e.g. Alluxio)",

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

					"tolerations": {
						Description:         "If specified, the pod's tolerations.",
						MarkdownDescription: "If specified, the pod's tolerations.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"effect": {
								Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
								MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key": {
								Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
								MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"operator": {
								Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
								MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"toleration_seconds": {
								Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
								MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
								MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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
		},
	}, nil
}

func (r *DataFluidIoDatasetV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_data_fluid_io_dataset_v1alpha1")

	var state DataFluidIoDatasetV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DataFluidIoDatasetV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("data.fluid.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Dataset")

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

func (r *DataFluidIoDatasetV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_data_fluid_io_dataset_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *DataFluidIoDatasetV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_data_fluid_io_dataset_v1alpha1")

	var state DataFluidIoDatasetV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DataFluidIoDatasetV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("data.fluid.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Dataset")

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

func (r *DataFluidIoDatasetV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_data_fluid_io_dataset_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
