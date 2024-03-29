/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &AppsKubeblocksIoComponentResourceConstraintV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoComponentResourceConstraintV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoComponentResourceConstraintV1Alpha1Manifest{}
}

type AppsKubeblocksIoComponentResourceConstraintV1Alpha1Manifest struct{}

type AppsKubeblocksIoComponentResourceConstraintV1Alpha1ManifestData struct {
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
		ComponentSelector *[]struct {
			ComponentDefRef *string   `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
			Rules           *[]string `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"component_selector" json:"componentSelector,omitempty"`
		Rules *[]struct {
			Cpu *struct {
				Max   *string   `tfsdk:"max" json:"max,omitempty"`
				Min   *string   `tfsdk:"min" json:"min,omitempty"`
				Slots *[]string `tfsdk:"slots" json:"slots,omitempty"`
				Step  *string   `tfsdk:"step" json:"step,omitempty"`
			} `tfsdk:"cpu" json:"cpu,omitempty"`
			Memory *struct {
				MaxPerCPU  *string `tfsdk:"max_per_cpu" json:"maxPerCPU,omitempty"`
				MinPerCPU  *string `tfsdk:"min_per_cpu" json:"minPerCPU,omitempty"`
				SizePerCPU *string `tfsdk:"size_per_cpu" json:"sizePerCPU,omitempty"`
			} `tfsdk:"memory" json:"memory,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Storage *struct {
				Max *string `tfsdk:"max" json:"max,omitempty"`
				Min *string `tfsdk:"min" json:"min,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		Selector *[]struct {
			ClusterDefRef *string `tfsdk:"cluster_def_ref" json:"clusterDefRef,omitempty"`
			Components    *[]struct {
				ComponentDefRef *string   `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
				Rules           *[]string `tfsdk:"rules" json:"rules,omitempty"`
			} `tfsdk:"components" json:"components,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoComponentResourceConstraintV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_component_resource_constraint_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoComponentResourceConstraintV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ComponentResourceConstraint is the Schema for the componentresourceconstraints API",
		MarkdownDescription: "ComponentResourceConstraint is the Schema for the componentresourceconstraints API",
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
				Description:         "ComponentResourceConstraintSpec defines the desired state of ComponentResourceConstraint",
				MarkdownDescription: "ComponentResourceConstraintSpec defines the desired state of ComponentResourceConstraint",
				Attributes: map[string]schema.Attribute{
					"component_selector": schema.ListNestedAttribute{
						Description:         "componentSelector is used to bind the resource constraint to components based on ComponentDefinition API.",
						MarkdownDescription: "componentSelector is used to bind the resource constraint to components based on ComponentDefinition API.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_def_ref": schema.StringAttribute{
									Description:         "In versions prior to KB 0.8.0, ComponentDefRef is the name of the component definition in the ClusterDefinition. In KB 0.8.0 and later versions, ComponentDefRef is the name of ComponentDefinition.",
									MarkdownDescription: "In versions prior to KB 0.8.0, ComponentDefRef is the name of the component definition in the ClusterDefinition. In KB 0.8.0 and later versions, ComponentDefRef is the name of ComponentDefinition.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"rules": schema.ListAttribute{
									Description:         "rules are the constraint rules that will be applied to the component.",
									MarkdownDescription: "rules are the constraint rules that will be applied to the component.",
									ElementType:         types.StringType,
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

					"rules": schema.ListNestedAttribute{
						Description:         "Component resource constraint rules.",
						MarkdownDescription: "Component resource constraint rules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cpu": schema.SingleNestedAttribute{
									Description:         "The constraint for vcpu cores.",
									MarkdownDescription: "The constraint for vcpu cores.",
									Attributes: map[string]schema.Attribute{
										"max": schema.StringAttribute{
											Description:         "The maximum count of vcpu cores, [Min, Max] defines a range for valid vcpu cores, and the value in this range must be multiple times of Step. It's useful to define a large number of valid values without defining them one by one. Please see the documentation for Step for some examples. If Slots is specified, Max, Min, and Step are ignored",
											MarkdownDescription: "The maximum count of vcpu cores, [Min, Max] defines a range for valid vcpu cores, and the value in this range must be multiple times of Step. It's useful to define a large number of valid values without defining them one by one. Please see the documentation for Step for some examples. If Slots is specified, Max, Min, and Step are ignored",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min": schema.StringAttribute{
											Description:         "The minimum count of vcpu cores, [Min, Max] defines a range for valid vcpu cores, and the value in this range must be multiple times of Step. It's useful to define a large number of valid values without defining them one by one. Please see the documentation for Step for some examples. If Slots is specified, Max, Min, and Step are ignored",
											MarkdownDescription: "The minimum count of vcpu cores, [Min, Max] defines a range for valid vcpu cores, and the value in this range must be multiple times of Step. It's useful to define a large number of valid values without defining them one by one. Please see the documentation for Step for some examples. If Slots is specified, Max, Min, and Step are ignored",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"slots": schema.ListAttribute{
											Description:         "The valid vcpu cores, it's useful if you want to define valid vcpu cores explicitly. If Slots is specified, Max, Min, and Step are ignored",
											MarkdownDescription: "The valid vcpu cores, it's useful if you want to define valid vcpu cores explicitly. If Slots is specified, Max, Min, and Step are ignored",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"step": schema.StringAttribute{
											Description:         "The minimum granularity of vcpu cores, [Min, Max] defines a range for valid vcpu cores and the value in this range must be multiple times of Step. For example: 1. Min is 2, Max is 8, Step is 2, and the valid vcpu core is {2, 4, 6, 8}. 2. Min is 0.5, Max is 2, Step is 0.5, and the valid vcpu core is {0.5, 1, 1.5, 2}.",
											MarkdownDescription: "The minimum granularity of vcpu cores, [Min, Max] defines a range for valid vcpu cores and the value in this range must be multiple times of Step. For example: 1. Min is 2, Max is 8, Step is 2, and the valid vcpu core is {2, 4, 6, 8}. 2. Min is 0.5, Max is 2, Step is 0.5, and the valid vcpu core is {0.5, 1, 1.5, 2}.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"memory": schema.SingleNestedAttribute{
									Description:         "The constraint for memory size.",
									MarkdownDescription: "The constraint for memory size.",
									Attributes: map[string]schema.Attribute{
										"max_per_cpu": schema.StringAttribute{
											Description:         "The maximum size of memory per vcpu core, [MinPerCPU, MaxPerCPU] defines a range for valid memory size per vcpu core. It is useful on GCP as the ratio between the CPU and memory may be a range. If SizePerCPU is specified, MinPerCPU and MaxPerCPU are ignored. Reference: https://cloud.google.com/compute/docs/general-purpose-machines#custom_machine_types",
											MarkdownDescription: "The maximum size of memory per vcpu core, [MinPerCPU, MaxPerCPU] defines a range for valid memory size per vcpu core. It is useful on GCP as the ratio between the CPU and memory may be a range. If SizePerCPU is specified, MinPerCPU and MaxPerCPU are ignored. Reference: https://cloud.google.com/compute/docs/general-purpose-machines#custom_machine_types",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_per_cpu": schema.StringAttribute{
											Description:         "The minimum size of memory per vcpu core, [MinPerCPU, MaxPerCPU] defines a range for valid memory size per vcpu core. It is useful on GCP as the ratio between the CPU and memory may be a range. If SizePerCPU is specified, MinPerCPU and MaxPerCPU are ignored. Reference: https://cloud.google.com/compute/docs/general-purpose-machines#custom_machine_types",
											MarkdownDescription: "The minimum size of memory per vcpu core, [MinPerCPU, MaxPerCPU] defines a range for valid memory size per vcpu core. It is useful on GCP as the ratio between the CPU and memory may be a range. If SizePerCPU is specified, MinPerCPU and MaxPerCPU are ignored. Reference: https://cloud.google.com/compute/docs/general-purpose-machines#custom_machine_types",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size_per_cpu": schema.StringAttribute{
											Description:         "The size of memory per vcpu core. For example: 1Gi, 200Mi. If SizePerCPU is specified, MinPerCPU and MaxPerCPU are ignore.",
											MarkdownDescription: "The size of memory per vcpu core. For example: 1Gi, 200Mi. If SizePerCPU is specified, MinPerCPU and MaxPerCPU are ignore.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "The name of the constraint.",
									MarkdownDescription: "The name of the constraint.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"storage": schema.SingleNestedAttribute{
									Description:         "The constraint for storage size.",
									MarkdownDescription: "The constraint for storage size.",
									Attributes: map[string]schema.Attribute{
										"max": schema.StringAttribute{
											Description:         "The maximum size of storage.",
											MarkdownDescription: "The maximum size of storage.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min": schema.StringAttribute{
											Description:         "The minimum size of storage.",
											MarkdownDescription: "The minimum size of storage.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"selector": schema.ListNestedAttribute{
						Description:         "selector is used to bind the resource constraint to cluster definitions based on ClusterDefinition API.",
						MarkdownDescription: "selector is used to bind the resource constraint to cluster definitions based on ClusterDefinition API.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cluster_def_ref": schema.StringAttribute{
									Description:         "clusterDefRef is the name of the cluster definition.",
									MarkdownDescription: "clusterDefRef is the name of the cluster definition.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"components": schema.ListNestedAttribute{
									Description:         "selector is used to bind the resource constraint to components.",
									MarkdownDescription: "selector is used to bind the resource constraint to components.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"component_def_ref": schema.StringAttribute{
												Description:         "In versions prior to KB 0.8.0, ComponentDefRef is the name of the component definition in the ClusterDefinition. In KB 0.8.0 and later versions, ComponentDefRef is the name of ComponentDefinition.",
												MarkdownDescription: "In versions prior to KB 0.8.0, ComponentDefRef is the name of the component definition in the ClusterDefinition. In KB 0.8.0 and later versions, ComponentDefRef is the name of ComponentDefinition.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"rules": schema.ListAttribute{
												Description:         "rules are the constraint rules that will be applied to the component.",
												MarkdownDescription: "rules are the constraint rules that will be applied to the component.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
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

func (r *AppsKubeblocksIoComponentResourceConstraintV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_component_resource_constraint_v1alpha1_manifest")

	var model AppsKubeblocksIoComponentResourceConstraintV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("ComponentResourceConstraint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
