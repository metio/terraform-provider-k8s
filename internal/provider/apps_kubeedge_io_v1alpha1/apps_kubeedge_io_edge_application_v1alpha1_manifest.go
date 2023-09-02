/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeedge_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest{}
)

func NewAppsKubeedgeIoEdgeApplicationV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest{}
}

type AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest struct{}

type AppsKubeedgeIoEdgeApplicationV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		WorkloadScope *struct {
			TargetNodeGroups *[]struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Overriders *struct {
					ImageOverriders *[]struct {
						Component *string `tfsdk:"component" json:"component,omitempty"`
						Operator  *string `tfsdk:"operator" json:"operator,omitempty"`
						Predicate *struct {
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"predicate" json:"predicate,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"image_overriders" json:"imageOverriders,omitempty"`
					Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				} `tfsdk:"overriders" json:"overriders,omitempty"`
			} `tfsdk:"target_node_groups" json:"targetNodeGroups,omitempty"`
		} `tfsdk:"workload_scope" json:"workloadScope,omitempty"`
		WorkloadTemplate *struct {
			Manifests *[]map[string]string `tfsdk:"manifests" json:"manifests,omitempty"`
		} `tfsdk:"workload_template" json:"workloadTemplate,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeedge_io_edge_application_v1alpha1_manifest"
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EdgeApplication is the Schema for the edgeapplications API",
		MarkdownDescription: "EdgeApplication is the Schema for the edgeapplications API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "Spec represents the desired behavior of EdgeApplication.",
				MarkdownDescription: "Spec represents the desired behavior of EdgeApplication.",
				Attributes: map[string]schema.Attribute{
					"workload_scope": schema.SingleNestedAttribute{
						Description:         "WorkloadScope represents which node groups the workload will be deployed in.",
						MarkdownDescription: "WorkloadScope represents which node groups the workload will be deployed in.",
						Attributes: map[string]schema.Attribute{
							"target_node_groups": schema.ListNestedAttribute{
								Description:         "TargetNodeGroups represents the target node groups of workload to be deployed.",
								MarkdownDescription: "TargetNodeGroups represents the target node groups of workload to be deployed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name represents the name of target node group",
											MarkdownDescription: "Name represents the name of target node group",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"overriders": schema.SingleNestedAttribute{
											Description:         "Overriders represents the override rules that would apply on workload.",
											MarkdownDescription: "Overriders represents the override rules that would apply on workload.",
											Attributes: map[string]schema.Attribute{
												"image_overriders": schema.ListNestedAttribute{
													Description:         "ImageOverriders represents the rules dedicated to handling image overrides.",
													MarkdownDescription: "ImageOverriders represents the rules dedicated to handling image overrides.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"component": schema.StringAttribute{
																Description:         "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - k8s.gcr.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
																MarkdownDescription: "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - k8s.gcr.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Registry", "Repository", "Tag"),
																},
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents the operator which will apply on the image.",
																MarkdownDescription: "Operator represents the operator which will apply on the image.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("add", "remove", "replace"),
																},
															},

															"predicate": schema.SingleNestedAttribute{
																Description:         "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: /spec/containers/<N>/image   - ReplicaSet: /spec/template/spec/containers/<N>/image   - Deployment: /spec/template/spec/containers/<N>/image   - StatefulSet: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",
																MarkdownDescription: "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: /spec/containers/<N>/image   - ReplicaSet: /spec/template/spec/containers/<N>/image   - Deployment: /spec/template/spec/containers/<N>/image   - StatefulSet: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "Path indicates the path of target field",
																		MarkdownDescription: "Path indicates the path of target field",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": schema.StringAttribute{
																Description:         "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
																MarkdownDescription: "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
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

												"replicas": schema.Int64Attribute{
													Description:         "Replicas will override the replicas field of deployment",
													MarkdownDescription: "Replicas will override the replicas field of deployment",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"workload_template": schema.SingleNestedAttribute{
						Description:         "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",
						MarkdownDescription: "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",
						Attributes: map[string]schema.Attribute{
							"manifests": schema.ListAttribute{
								Description:         "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",
								MarkdownDescription: "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",
								ElementType:         types.MapType{ElemType: types.StringType},
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeedge_io_edge_application_v1alpha1_manifest")

	var model AppsKubeedgeIoEdgeApplicationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("apps.kubeedge.io/v1alpha1")
	model.Kind = pointer.String("EdgeApplication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
