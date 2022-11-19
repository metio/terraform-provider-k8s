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

type ChaosMeshOrgPodIOChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgPodIOChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgPodIOChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgPodIOChaosV1Alpha1GoModel struct {
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
		Actions *[]struct {
			Atime *struct {
				Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

				Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
			} `tfsdk:"atime" yaml:"atime,omitempty"`

			Blocks *int64 `tfsdk:"blocks" yaml:"blocks,omitempty"`

			Ctime *struct {
				Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

				Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
			} `tfsdk:"ctime" yaml:"ctime,omitempty"`

			Faults *[]struct {
				Errno *int64 `tfsdk:"errno" yaml:"errno,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
			} `tfsdk:"faults" yaml:"faults,omitempty"`

			Gid *int64 `tfsdk:"gid" yaml:"gid,omitempty"`

			Ino *int64 `tfsdk:"ino" yaml:"ino,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

			Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

			Mistake *struct {
				Filling *string `tfsdk:"filling" yaml:"filling,omitempty"`

				MaxLength *int64 `tfsdk:"max_length" yaml:"maxLength,omitempty"`

				MaxOccurrences *int64 `tfsdk:"max_occurrences" yaml:"maxOccurrences,omitempty"`
			} `tfsdk:"mistake" yaml:"mistake,omitempty"`

			Mtime *struct {
				Nsec *int64 `tfsdk:"nsec" yaml:"nsec,omitempty"`

				Sec *int64 `tfsdk:"sec" yaml:"sec,omitempty"`
			} `tfsdk:"mtime" yaml:"mtime,omitempty"`

			Nlink *int64 `tfsdk:"nlink" yaml:"nlink,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Percent *int64 `tfsdk:"percent" yaml:"percent,omitempty"`

			Perm *int64 `tfsdk:"perm" yaml:"perm,omitempty"`

			Rdev *int64 `tfsdk:"rdev" yaml:"rdev,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			Uid *int64 `tfsdk:"uid" yaml:"uid,omitempty"`
		} `tfsdk:"actions" yaml:"actions,omitempty"`

		Container *string `tfsdk:"container" yaml:"container,omitempty"`

		VolumeMountPath *string `tfsdk:"volume_mount_path" yaml:"volumeMountPath,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgPodIOChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgPodIOChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgPodIOChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_pod_io_chaos_v1alpha1"
}

func (r *ChaosMeshOrgPodIOChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "PodIOChaos is the Schema for the podiochaos API",
		MarkdownDescription: "PodIOChaos is the Schema for the podiochaos API",
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
				Description:         "PodIOChaosSpec defines the desired state of IOChaos",
				MarkdownDescription: "PodIOChaosSpec defines the desired state of IOChaos",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"actions": {
						Description:         "Actions are a list of IOChaos actions",
						MarkdownDescription: "Actions are a list of IOChaos actions",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

							"faults": {
								Description:         "Faults represents the fault to inject",
								MarkdownDescription: "Faults represents the fault to inject",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"errno": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"weight": {
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

							"latency": {
								Description:         "Latency represents the latency to inject",
								MarkdownDescription: "Latency represents the latency to inject",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"methods": {
								Description:         "Methods represents the method that the action will inject in",
								MarkdownDescription: "Methods represents the method that the action will inject in",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mistake": {
								Description:         "MistakeSpec represents the mistake to inject",
								MarkdownDescription: "MistakeSpec represents the mistake to inject",

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

							"path": {
								Description:         "Path represents a glob of injecting path",
								MarkdownDescription: "Path represents a glob of injecting path",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"percent": {
								Description:         "Percent represents the percent probability of injecting this action",
								MarkdownDescription: "Percent represents the percent probability of injecting this action",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
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

							"source": {
								Description:         "Source represents the source of current rules",
								MarkdownDescription: "Source represents the source of current rules",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "IOChaosType represents the type of IOChaos Action",
								MarkdownDescription: "IOChaosType represents the type of IOChaos Action",

								Type: types.StringType,

								Required: true,
								Optional: false,
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

					"container": {
						Description:         "TODO: support multiple different container to inject in one pod",
						MarkdownDescription: "TODO: support multiple different container to inject in one pod",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_mount_path": {
						Description:         "VolumeMountPath represents the target mount path It must be a root of mount path now. TODO: search the mount parent of any path automatically. TODO: support multiple different volume mount path in one pod",
						MarkdownDescription: "VolumeMountPath represents the target mount path It must be a root of mount path now. TODO: search the mount parent of any path automatically. TODO: support multiple different volume mount path in one pod",

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

func (r *ChaosMeshOrgPodIOChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_pod_io_chaos_v1alpha1")

	var state ChaosMeshOrgPodIOChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPodIOChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PodIOChaos")

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

func (r *ChaosMeshOrgPodIOChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_pod_io_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgPodIOChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_pod_io_chaos_v1alpha1")

	var state ChaosMeshOrgPodIOChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgPodIOChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("PodIOChaos")

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

func (r *ChaosMeshOrgPodIOChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_pod_io_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
