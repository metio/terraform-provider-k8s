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

type AppsKubeedgeIoEdgeApplicationV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*AppsKubeedgeIoEdgeApplicationV1Alpha1Resource)(nil)
)

type AppsKubeedgeIoEdgeApplicationV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppsKubeedgeIoEdgeApplicationV1Alpha1GoModel struct {
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
		WorkloadScope *struct {
			TargetNodeGroups *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Overriders *struct {
					ImageOverriders *[]struct {
						Component *string `tfsdk:"component" yaml:"component,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						Predicate *struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"predicate" yaml:"predicate,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"image_overriders" yaml:"imageOverriders,omitempty"`

					Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`
				} `tfsdk:"overriders" yaml:"overriders,omitempty"`
			} `tfsdk:"target_node_groups" yaml:"targetNodeGroups,omitempty"`
		} `tfsdk:"workload_scope" yaml:"workloadScope,omitempty"`

		WorkloadTemplate *struct {
			Manifests *[]map[string]string `tfsdk:"manifests" yaml:"manifests,omitempty"`
		} `tfsdk:"workload_template" yaml:"workloadTemplate,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppsKubeedgeIoEdgeApplicationV1Alpha1Resource() resource.Resource {
	return &AppsKubeedgeIoEdgeApplicationV1Alpha1Resource{}
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apps_kubeedge_io_edge_application_v1alpha1"
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "EdgeApplication is the Schema for the edgeapplications API",
		MarkdownDescription: "EdgeApplication is the Schema for the edgeapplications API",
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
				Description:         "Spec represents the desired behavior of EdgeApplication.",
				MarkdownDescription: "Spec represents the desired behavior of EdgeApplication.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"workload_scope": {
						Description:         "WorkloadScope represents which node groups the workload will be deployed in.",
						MarkdownDescription: "WorkloadScope represents which node groups the workload will be deployed in.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"target_node_groups": {
								Description:         "TargetNodeGroups represents the target node groups of workload to be deployed.",
								MarkdownDescription: "TargetNodeGroups represents the target node groups of workload to be deployed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name represents the name of target node group",
										MarkdownDescription: "Name represents the name of target node group",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"overriders": {
										Description:         "Overriders represents the override rules that would apply on workload.",
										MarkdownDescription: "Overriders represents the override rules that would apply on workload.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"image_overriders": {
												Description:         "ImageOverriders represents the rules dedicated to handling image overrides.",
												MarkdownDescription: "ImageOverriders represents the rules dedicated to handling image overrides.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"component": {
														Description:         "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - k8s.gcr.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",
														MarkdownDescription: "Component is part of image name. Basically we presume an image can be made of '[registry/]repository[:tag]'. The registry could be: - k8s.gcr.io - fictional.registry.example:10443 The repository could be: - kube-apiserver - fictional/nginx The tag cloud be: - latest - v1.19.1 - @sha256:dbcc1c35ac38df41fd2f5e4130b32ffdb93ebae8b3dbe638c23575912276fc9c",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Registry", "Repository", "Tag"),
														},
													},

													"operator": {
														Description:         "Operator represents the operator which will apply on the image.",
														MarkdownDescription: "Operator represents the operator which will apply on the image.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("add", "remove", "replace"),
														},
													},

													"predicate": {
														Description:         "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: spec/containers/<N>/image   - ReplicaSet: spec/template/spec/containers/<N>/image   - Deployment: spec/template/spec/containers/<N>/image   - StatefulSet: spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",
														MarkdownDescription: "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: spec/containers/<N>/image   - ReplicaSet: spec/template/spec/containers/<N>/image   - Deployment: spec/template/spec/containers/<N>/image   - StatefulSet: spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"path": {
																Description:         "Path indicates the path of target field",
																MarkdownDescription: "Path indicates the path of target field",

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

													"value": {
														Description:         "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
														MarkdownDescription: "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",

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

											"replicas": {
												Description:         "Replicas will override the replicas field of deployment",
												MarkdownDescription: "Replicas will override the replicas field of deployment",

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

					"workload_template": {
						Description:         "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",
						MarkdownDescription: "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"manifests": {
								Description:         "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",
								MarkdownDescription: "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",

								Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

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

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apps_kubeedge_io_edge_application_v1alpha1")

	var state AppsKubeedgeIoEdgeApplicationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsKubeedgeIoEdgeApplicationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.kubeedge.io/v1alpha1")
	goModel.Kind = utilities.Ptr("EdgeApplication")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeedge_io_edge_application_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apps_kubeedge_io_edge_application_v1alpha1")

	var state AppsKubeedgeIoEdgeApplicationV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsKubeedgeIoEdgeApplicationV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.kubeedge.io/v1alpha1")
	goModel.Kind = utilities.Ptr("EdgeApplication")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apps_kubeedge_io_edge_application_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
