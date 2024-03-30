/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &ChaosMeshOrgPodIochaosV1Alpha1Manifest{}
)

func NewChaosMeshOrgPodIochaosV1Alpha1Manifest() datasource.DataSource {
	return &ChaosMeshOrgPodIochaosV1Alpha1Manifest{}
}

type ChaosMeshOrgPodIochaosV1Alpha1Manifest struct{}

type ChaosMeshOrgPodIochaosV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Actions *[]struct {
			Atime *struct {
				Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
				Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
			} `tfsdk:"atime" json:"atime,omitempty"`
			Blocks *int64 `tfsdk:"blocks" json:"blocks,omitempty"`
			Ctime  *struct {
				Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
				Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
			} `tfsdk:"ctime" json:"ctime,omitempty"`
			Faults *[]struct {
				Errno  *int64 `tfsdk:"errno" json:"errno,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"faults" json:"faults,omitempty"`
			Gid     *int64    `tfsdk:"gid" json:"gid,omitempty"`
			Ino     *int64    `tfsdk:"ino" json:"ino,omitempty"`
			Kind    *string   `tfsdk:"kind" json:"kind,omitempty"`
			Latency *string   `tfsdk:"latency" json:"latency,omitempty"`
			Methods *[]string `tfsdk:"methods" json:"methods,omitempty"`
			Mistake *struct {
				Filling        *string `tfsdk:"filling" json:"filling,omitempty"`
				MaxLength      *int64  `tfsdk:"max_length" json:"maxLength,omitempty"`
				MaxOccurrences *int64  `tfsdk:"max_occurrences" json:"maxOccurrences,omitempty"`
			} `tfsdk:"mistake" json:"mistake,omitempty"`
			Mtime *struct {
				Nsec *int64 `tfsdk:"nsec" json:"nsec,omitempty"`
				Sec  *int64 `tfsdk:"sec" json:"sec,omitempty"`
			} `tfsdk:"mtime" json:"mtime,omitempty"`
			Nlink   *int64  `tfsdk:"nlink" json:"nlink,omitempty"`
			Path    *string `tfsdk:"path" json:"path,omitempty"`
			Percent *int64  `tfsdk:"percent" json:"percent,omitempty"`
			Perm    *int64  `tfsdk:"perm" json:"perm,omitempty"`
			Rdev    *int64  `tfsdk:"rdev" json:"rdev,omitempty"`
			Size    *int64  `tfsdk:"size" json:"size,omitempty"`
			Source  *string `tfsdk:"source" json:"source,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
			Uid     *int64  `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"actions" json:"actions,omitempty"`
		Container       *string `tfsdk:"container" json:"container,omitempty"`
		VolumeMountPath *string `tfsdk:"volume_mount_path" json:"volumeMountPath,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgPodIochaosV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_pod_io_chaos_v1alpha1_manifest"
}

func (r *ChaosMeshOrgPodIochaosV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PodIOChaos is the Schema for the podiochaos API",
		MarkdownDescription: "PodIOChaos is the Schema for the podiochaos API",
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

			"spec": schema.SingleNestedAttribute{
				Description:         "PodIOChaosSpec defines the desired state of IOChaos",
				MarkdownDescription: "PodIOChaosSpec defines the desired state of IOChaos",
				Attributes: map[string]schema.Attribute{
					"actions": schema.ListNestedAttribute{
						Description:         "Actions are a list of IOChaos actions",
						MarkdownDescription: "Actions are a list of IOChaos actions",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"atime": schema.SingleNestedAttribute{
									Description:         "Timespec represents a time",
									MarkdownDescription: "Timespec represents a time",
									Attributes: map[string]schema.Attribute{
										"nsec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"sec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"blocks": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ctime": schema.SingleNestedAttribute{
									Description:         "Timespec represents a time",
									MarkdownDescription: "Timespec represents a time",
									Attributes: map[string]schema.Attribute{
										"nsec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"sec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"faults": schema.ListNestedAttribute{
									Description:         "Faults represents the fault to inject",
									MarkdownDescription: "Faults represents the fault to inject",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"errno": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"weight": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
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

								"gid": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ino": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "FileType represents type of file",
									MarkdownDescription: "FileType represents type of file",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"latency": schema.StringAttribute{
									Description:         "Latency represents the latency to inject",
									MarkdownDescription: "Latency represents the latency to inject",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"methods": schema.ListAttribute{
									Description:         "Methods represents the method that the action will inject in",
									MarkdownDescription: "Methods represents the method that the action will inject in",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"mistake": schema.SingleNestedAttribute{
									Description:         "MistakeSpec represents the mistake to inject",
									MarkdownDescription: "MistakeSpec represents the mistake to inject",
									Attributes: map[string]schema.Attribute{
										"filling": schema.StringAttribute{
											Description:         "Filling determines what is filled in the mistake data.",
											MarkdownDescription: "Filling determines what is filled in the mistake data.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("zero", "random"),
											},
										},

										"max_length": schema.Int64Attribute{
											Description:         "Max length of each wrong data segment in bytes",
											MarkdownDescription: "Max length of each wrong data segment in bytes",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"max_occurrences": schema.Int64Attribute{
											Description:         "There will be [1, MaxOccurrences] segments of wrong data.",
											MarkdownDescription: "There will be [1, MaxOccurrences] segments of wrong data.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"mtime": schema.SingleNestedAttribute{
									Description:         "Timespec represents a time",
									MarkdownDescription: "Timespec represents a time",
									Attributes: map[string]schema.Attribute{
										"nsec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"sec": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"nlink": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "Path represents a glob of injecting path",
									MarkdownDescription: "Path represents a glob of injecting path",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"percent": schema.Int64Attribute{
									Description:         "Percent represents the percent probability of injecting this action",
									MarkdownDescription: "Percent represents the percent probability of injecting this action",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"perm": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rdev": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"size": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source": schema.StringAttribute{
									Description:         "Source represents the source of current rules",
									MarkdownDescription: "Source represents the source of current rules",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "IOChaosType represents the type of IOChaos Action",
									MarkdownDescription: "IOChaosType represents the type of IOChaos Action",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"uid": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"container": schema.StringAttribute{
						Description:         "TODO: support multiple different container to inject in one pod",
						MarkdownDescription: "TODO: support multiple different container to inject in one pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_mount_path": schema.StringAttribute{
						Description:         "VolumeMountPath represents the target mount path It must be a root of mount path now. TODO: search the mount parent of any path automatically. TODO: support multiple different volume mount path in one pod",
						MarkdownDescription: "VolumeMountPath represents the target mount path It must be a root of mount path now. TODO: search the mount parent of any path automatically. TODO: support multiple different volume mount path in one pod",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ChaosMeshOrgPodIochaosV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_pod_io_chaos_v1alpha1_manifest")

	var model ChaosMeshOrgPodIochaosV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("PodIOChaos")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
