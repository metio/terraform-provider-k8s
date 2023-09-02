/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeedge_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource{}
)

func NewAppsKubeedgeIoEdgeApplicationV1Alpha1DataSource() datasource.DataSource {
	return &AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource{}
}

type AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AppsKubeedgeIoEdgeApplicationV1Alpha1DataSourceData struct {
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

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeedge_io_edge_application_v1alpha1"
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
											Required:            false,
											Optional:            false,
											Computed:            true,
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
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents the operator which will apply on the image.",
																MarkdownDescription: "Operator represents the operator which will apply on the image.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"predicate": schema.SingleNestedAttribute{
																Description:         "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: /spec/containers/<N>/image   - ReplicaSet: /spec/template/spec/containers/<N>/image   - Deployment: /spec/template/spec/containers/<N>/image   - StatefulSet: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",
																MarkdownDescription: "Predicate filters images before applying the rule.  Defaults to nil, in that case, the system will automatically detect image fields if the resource type is Pod, ReplicaSet, Deployment or StatefulSet by following rule:   - Pod: /spec/containers/<N>/image   - ReplicaSet: /spec/template/spec/containers/<N>/image   - Deployment: /spec/template/spec/containers/<N>/image   - StatefulSet: /spec/template/spec/containers/<N>/image In addition, all images will be processed if the resource object has more than one containers.  If not nil, only images matches the filters will be processed.",
																Attributes: map[string]schema.Attribute{
																	"path": schema.StringAttribute{
																		Description:         "Path indicates the path of target field",
																		MarkdownDescription: "Path indicates the path of target field",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"value": schema.StringAttribute{
																Description:         "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
																MarkdownDescription: "Value to be applied to image. Must not be empty when operator is 'add' or 'replace'. Defaults to empty and ignored when operator is 'remove'.",
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

												"replicas": schema.Int64Attribute{
													Description:         "Replicas will override the replicas field of deployment",
													MarkdownDescription: "Replicas will override the replicas field of deployment",
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

					"workload_template": schema.SingleNestedAttribute{
						Description:         "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",
						MarkdownDescription: "WorkloadTemplate contains original templates of resources to be deployed as an EdgeApplication.",
						Attributes: map[string]schema.Attribute{
							"manifests": schema.ListAttribute{
								Description:         "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",
								MarkdownDescription: "Manifests represent a list of Kubernetes resources to be deployed on the managed node groups.",
								ElementType:         types.MapType{ElemType: types.StringType},
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
	}
}

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *AppsKubeedgeIoEdgeApplicationV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_apps_kubeedge_io_edge_application_v1alpha1")

	var data AppsKubeedgeIoEdgeApplicationV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apps.kubeedge.io", Version: "v1alpha1", Resource: "EdgeApplication"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse AppsKubeedgeIoEdgeApplicationV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("apps.kubeedge.io/v1alpha1")
	data.Kind = pointer.String("EdgeApplication")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
