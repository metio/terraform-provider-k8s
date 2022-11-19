/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type ChaosMeshOrgIOChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgIOChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgIOChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgIOChaosV1Alpha1GoModel struct {
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
		Action *string `tfsdk:"action" yaml:"action,omitempty"`

		Attr *struct {
			Atime *struct {
				Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

				Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
			} `tfsdk:"atime" yaml:"atime,omitempty"`

			Blocks *int64 `tfsdk:"blocks" yaml:"blocks,omitempty"`

			Ctime *struct {
				Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

				Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
			} `tfsdk:"ctime" yaml:"ctime,omitempty"`

			Gid *int64 `tfsdk:"gid" yaml:"gid,omitempty"`

			Ino *int64 `tfsdk:"ino" yaml:"ino,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Mtime *struct {
				Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

				Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
			} `tfsdk:"mtime" yaml:"mtime,omitempty"`

			Nlink *int64 `tfsdk:"nlink" yaml:"nlink,omitempty"`

			Perm *int64 `tfsdk:"perm" yaml:"perm,omitempty"`

			Rdev *int64 `tfsdk:"rdev" yaml:"rdev,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

			Uid *int64 `tfsdk:"uid" yaml:"uid,omitempty"`
		} `tfsdk:"attr" yaml:"attr,omitempty"`

		ContainerNames *[]string `tfsdk:"container_names" yaml:"containerNames,omitempty"`

		Delay *string `tfsdk:"delay" yaml:"delay,omitempty"`

		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		Errno *int64 `tfsdk:"errno" yaml:"errno,omitempty"`

		Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

		Mistake *struct {
			Filling *string `tfsdk:"filling" yaml:"filling,omitempty"`

			MaxLength *int64 `tfsdk:"max_length" yaml:"maxLength,omitempty"`

			MaxOccurrences *int64 `tfsdk:"max_occurrences" yaml:"maxOccurrences,omitempty"`
		} `tfsdk:"mistake" yaml:"mistake,omitempty"`

		Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

		Path *string `tfsdk:"path" yaml:"path,omitempty"`

		Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`

		Selector *struct {
			AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

			ExpressionSelectors *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

			FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

			LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

			Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

			NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

			Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

			PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

			Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`

		Value *string `tfsdk:"value" yaml:"value,omitempty"`

		VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgIOChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgIOChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgIOChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_io_chaos_v1alpha1"
}

func (r *ChaosMeshOrgIOChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "IOChaos is the Schema for the iochaos API",
		MarkdownDescription: "IOChaos is the Schema for the iochaos API",
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
				Description:         "IOChaosSpec defines the desired state of IOChaos",
				MarkdownDescription: "IOChaosSpec defines the desired state of IOChaos",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"action": {
						Description:         "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",
						MarkdownDescription: "Action defines the specific pod chaos action. Supported action: latency / fault / attrOverride / mistake",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("latency", "fault", "attrOverride", "mistake"),
						},
					},

					"attr": {
						Description:         "Attr defines the overrided attribution",
						MarkdownDescription: "Attr defines the overrided attribution",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"atime": {
								Description:         "Timespec represents a time",
								MarkdownDescription: "Timespec represents a time",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"nsec": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"sec": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"blocks": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ctime": {
								Description:         "Timespec represents a time",
								MarkdownDescription: "Timespec represents a time",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"nsec": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"sec": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gid": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ino": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "FileType represents type of file",
								MarkdownDescription: "FileType represents type of file",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mtime": {
								Description:         "Timespec represents a time",
								MarkdownDescription: "Timespec represents a time",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"nsec": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"sec": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"nlink": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"perm": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rdev": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"uid": {
								Description:         "",
								MarkdownDescription: "",

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

					"container_names": {
						Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
						MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"delay": {
						Description:         "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Delay defines the value of I/O chaos action delay. A delay string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"duration": {
						Description:         "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Duration represents the duration of the chaos action. It is required when the action is 'PodFailureAction'. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"errno": {
						Description:         "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",
						MarkdownDescription: "Errno defines the error code that returned by I/O action. refer to: https://www-numi.fnal.gov/offline_software/srt_public_context/WebDocs/Errors/unix_system_errors.html",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"methods": {
						Description:         "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",
						MarkdownDescription: "Methods defines the I/O methods for injecting I/O chaos action. default: all I/O methods.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mistake": {
						Description:         "Mistake defines what types of incorrectness are injected to IO operations",
						MarkdownDescription: "Mistake defines what types of incorrectness are injected to IO operations",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"filling": {
								Description:         "Filling determines what is filled in the mistake data.",
								MarkdownDescription: "Filling determines what is filled in the mistake data.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("zero", "random"),
								},
							},

							"max_length": {
								Description:         "Max length of each wrong data segment in bytes",
								MarkdownDescription: "Max length of each wrong data segment in bytes",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"max_occurrences": {
								Description:         "There will be [1, MaxOccurrences] segments of wrong data.",
								MarkdownDescription: "There will be [1, MaxOccurrences] segments of wrong data.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mode": {
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
						},
					},

					"path": {
						Description:         "Path defines the path of files for injecting I/O chaos action.",
						MarkdownDescription: "Path defines the path of files for injecting I/O chaos action.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"percent": {
						Description:         "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",
						MarkdownDescription: "Percent defines the percentage of injection errors and provides a number from 0-100. default: 100.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "Selector is used to select pods that are used to inject chaos action.",
						MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotation_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"expression_selectors": {
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

							"field_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"label_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespaces": {
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_selectors": {
								Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"nodes": {
								Description:         "Nodes is a set of node name and objects must belong to these nodes.",
								MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_phase_selectors": {
								Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pods": {
								Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

								Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"value": {
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_path": {
						Description:         "VolumePath represents the mount path of injected volume",
						MarkdownDescription: "VolumePath represents the mount path of injected volume",

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
		},
	}, nil
}

func (r *ChaosMeshOrgIOChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_io_chaos_v1alpha1")

	var state ChaosMeshOrgIOChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgIOChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("IOChaos")

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

func (r *ChaosMeshOrgIOChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_io_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgIOChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_io_chaos_v1alpha1")

	var state ChaosMeshOrgIOChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgIOChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("IOChaos")

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

func (r *ChaosMeshOrgIOChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_io_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
