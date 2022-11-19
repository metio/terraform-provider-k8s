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

type ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource struct{}

var (
	_ resource.Resource = (*ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource)(nil)
)

type ServicebindingIoClusterWorkloadResourceMappingV1Alpha3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ServicebindingIoClusterWorkloadResourceMappingV1Alpha3GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Versions *[]struct {
			Annotations *string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Containers *[]struct {
				Env *string `tfsdk:"env" yaml:"env,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				VolumeMounts *string `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
			} `tfsdk:"containers" yaml:"containers,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`

			Volumes *string `tfsdk:"volumes" yaml:"volumes,omitempty"`
		} `tfsdk:"versions" yaml:"versions,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource() resource.Resource {
	return &ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource{}
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicebinding_io_cluster_workload_resource_mapping_v1alpha3"
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterWorkloadResourceMapping is the Schema for the clusterworkloadresourcemappings API",
		MarkdownDescription: "ClusterWorkloadResourceMapping is the Schema for the clusterworkloadresourcemappings API",
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
				Description:         "ClusterWorkloadResourceMappingSpec defines the desired state of ClusterWorkloadResourceMapping",
				MarkdownDescription: "ClusterWorkloadResourceMappingSpec defines the desired state of ClusterWorkloadResourceMapping",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"versions": {
						Description:         "Versions is the collection of versions for a given resource, with mappings.",
						MarkdownDescription: "Versions is the collection of versions for a given resource, with mappings.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"annotations": {
								Description:         "Annotations is a Restricted JSONPath that references the annotations map within the workload resource. These annotations must end up in the resulting Pod, and are generally not the workload resource's annotations. Defaults to '.spec.template.metadata.annotations'.",
								MarkdownDescription: "Annotations is a Restricted JSONPath that references the annotations map within the workload resource. These annotations must end up in the resulting Pod, and are generally not the workload resource's annotations. Defaults to '.spec.template.metadata.annotations'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"containers": {
								Description:         "Containers is the collection of mappings to container-like fragments of the workload resource. Defaults to mappings appropriate for a PodSpecable resource.",
								MarkdownDescription: "Containers is the collection of mappings to container-like fragments of the workload resource. Defaults to mappings appropriate for a PodSpecable resource.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"env": {
										Description:         "Env is a Restricted JSONPath that references the slice of environment variables for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.envs'.",
										MarkdownDescription: "Env is a Restricted JSONPath that references the slice of environment variables for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.envs'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name is a Restricted JSONPath that references the name of the container with the container-like workload resource fragment. If not defined, container name filtering is ignored.",
										MarkdownDescription: "Name is a Restricted JSONPath that references the name of the container with the container-like workload resource fragment. If not defined, container name filtering is ignored.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path is the JSONPath within the workload resource that matches an existing fragment that is container-like.",
										MarkdownDescription: "Path is the JSONPath within the workload resource that matches an existing fragment that is container-like.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"volume_mounts": {
										Description:         "VolumeMounts is a Restricted JSONPath that references the slice of volume mounts for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.volumeMounts'.",
										MarkdownDescription: "VolumeMounts is a Restricted JSONPath that references the slice of volume mounts for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.volumeMounts'.",

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

							"version": {
								Description:         "Version is the version of the workload resource that this mapping is for.",
								MarkdownDescription: "Version is the version of the workload resource that this mapping is for.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"volumes": {
								Description:         "Volumes is a Restricted JSONPath that references the slice of volumes within the workload resource. Defaults to '.spec.template.spec.volumes'.",
								MarkdownDescription: "Volumes is a Restricted JSONPath that references the slice of volumes within the workload resource. Defaults to '.spec.template.spec.volumes'.",

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

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3")

	var state ServicebindingIoClusterWorkloadResourceMappingV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ServicebindingIoClusterWorkloadResourceMappingV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("servicebinding.io/v1alpha3")
	goModel.Kind = utilities.Ptr("ClusterWorkloadResourceMapping")

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

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3")
	// NO-OP: All data is already in Terraform state
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3")

	var state ServicebindingIoClusterWorkloadResourceMappingV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ServicebindingIoClusterWorkloadResourceMappingV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("servicebinding.io/v1alpha3")
	goModel.Kind = utilities.Ptr("ClusterWorkloadResourceMapping")

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

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3")
	// NO-OP: Terraform removes the state automatically for us
}
