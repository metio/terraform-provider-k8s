/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package servicebinding_io_v1alpha3

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
	_ datasource.DataSource = &ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Manifest{}
)

func NewServicebindingIoClusterWorkloadResourceMappingV1Alpha3Manifest() datasource.DataSource {
	return &ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Manifest{}
}

type ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Manifest struct{}

type ServicebindingIoClusterWorkloadResourceMappingV1Alpha3ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Versions *[]struct {
			Annotations *string `tfsdk:"annotations" json:"annotations,omitempty"`
			Containers  *[]struct {
				Env          *string `tfsdk:"env" json:"env,omitempty"`
				Name         *string `tfsdk:"name" json:"name,omitempty"`
				Path         *string `tfsdk:"path" json:"path,omitempty"`
				VolumeMounts *string `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"containers" json:"containers,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
			Volumes *string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"versions" json:"versions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest"
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterWorkloadResourceMapping is the Schema for the clusterworkloadresourcemappings API",
		MarkdownDescription: "ClusterWorkloadResourceMapping is the Schema for the clusterworkloadresourcemappings API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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

			"spec": schema.SingleNestedAttribute{
				Description:         "ClusterWorkloadResourceMappingSpec defines the desired state of ClusterWorkloadResourceMapping",
				MarkdownDescription: "ClusterWorkloadResourceMappingSpec defines the desired state of ClusterWorkloadResourceMapping",
				Attributes: map[string]schema.Attribute{
					"versions": schema.ListNestedAttribute{
						Description:         "Versions is the collection of versions for a given resource, with mappings.",
						MarkdownDescription: "Versions is the collection of versions for a given resource, with mappings.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"annotations": schema.StringAttribute{
									Description:         "Annotations is a Restricted JSONPath that references the annotations map within the workload resource. These annotations must end up in the resulting Pod, and are generally not the workload resource's annotations. Defaults to '.spec.template.metadata.annotations'.",
									MarkdownDescription: "Annotations is a Restricted JSONPath that references the annotations map within the workload resource. These annotations must end up in the resulting Pod, and are generally not the workload resource's annotations. Defaults to '.spec.template.metadata.annotations'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"containers": schema.ListNestedAttribute{
									Description:         "Containers is the collection of mappings to container-like fragments of the workload resource. Defaults to mappings appropriate for a PodSpecable resource.",
									MarkdownDescription: "Containers is the collection of mappings to container-like fragments of the workload resource. Defaults to mappings appropriate for a PodSpecable resource.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"env": schema.StringAttribute{
												Description:         "Env is a Restricted JSONPath that references the slice of environment variables for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.envs'.",
												MarkdownDescription: "Env is a Restricted JSONPath that references the slice of environment variables for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.envs'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is a Restricted JSONPath that references the name of the container with the container-like workload resource fragment. If not defined, container name filtering is ignored.",
												MarkdownDescription: "Name is a Restricted JSONPath that references the name of the container with the container-like workload resource fragment. If not defined, container name filtering is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "Path is the JSONPath within the workload resource that matches an existing fragment that is container-like.",
												MarkdownDescription: "Path is the JSONPath within the workload resource that matches an existing fragment that is container-like.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"volume_mounts": schema.StringAttribute{
												Description:         "VolumeMounts is a Restricted JSONPath that references the slice of volume mounts for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.volumeMounts'.",
												MarkdownDescription: "VolumeMounts is a Restricted JSONPath that references the slice of volume mounts for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.volumeMounts'.",
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

								"version": schema.StringAttribute{
									Description:         "Version is the version of the workload resource that this mapping is for.",
									MarkdownDescription: "Version is the version of the workload resource that this mapping is for.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"volumes": schema.StringAttribute{
									Description:         "Volumes is a Restricted JSONPath that references the slice of volumes within the workload resource. Defaults to '.spec.template.spec.volumes'.",
									MarkdownDescription: "Volumes is a Restricted JSONPath that references the slice of volumes within the workload resource. Defaults to '.spec.template.spec.volumes'.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3_manifest")

	var model ServicebindingIoClusterWorkloadResourceMappingV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("servicebinding.io/v1alpha3")
	model.Kind = pointer.String("ClusterWorkloadResourceMapping")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
