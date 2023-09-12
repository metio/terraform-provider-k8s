/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package addons_cluster_x_k8s_io_v1alpha4

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
	_ datasource.DataSource              = &AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource{}
	_ datasource.DataSourceWithConfigure = &AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource{}
)

func NewAddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource() datasource.DataSource {
	return &AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource{}
}

type AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource struct {
	kubernetesClient dynamic.Interface
}

type AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSourceData struct {
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
		Bindings *[]struct {
			ClusterResourceSetName *string `tfsdk:"cluster_resource_set_name" json:"clusterResourceSetName,omitempty"`
			Resources              *[]struct {
				Applied         *bool   `tfsdk:"applied" json:"applied,omitempty"`
				Hash            *string `tfsdk:"hash" json:"hash,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				LastAppliedTime *string `tfsdk:"last_applied_time" json:"lastAppliedTime,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"bindings" json:"bindings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4"
}

func (r *AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterResourceSetBinding lists all matching ClusterResourceSets with the cluster it belongs to.  Deprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "ClusterResourceSetBinding lists all matching ClusterResourceSets with the cluster it belongs to.  Deprecated: This type will be removed in one of the next releases.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ClusterResourceSetBindingSpec defines the desired state of ClusterResourceSetBinding.",
				MarkdownDescription: "ClusterResourceSetBindingSpec defines the desired state of ClusterResourceSetBinding.",
				Attributes: map[string]schema.Attribute{
					"bindings": schema.ListNestedAttribute{
						Description:         "Bindings is a list of ClusterResourceSets and their resources.",
						MarkdownDescription: "Bindings is a list of ClusterResourceSets and their resources.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cluster_resource_set_name": schema.StringAttribute{
									Description:         "ClusterResourceSetName is the name of the ClusterResourceSet that is applied to the owner cluster of the binding.",
									MarkdownDescription: "ClusterResourceSetName is the name of the ClusterResourceSet that is applied to the owner cluster of the binding.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"resources": schema.ListNestedAttribute{
									Description:         "Resources is a list of resources that the ClusterResourceSet has.",
									MarkdownDescription: "Resources is a list of resources that the ClusterResourceSet has.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"applied": schema.BoolAttribute{
												Description:         "Applied is to track if a resource is applied to the cluster or not.",
												MarkdownDescription: "Applied is to track if a resource is applied to the cluster or not.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"hash": schema.StringAttribute{
												Description:         "Hash is the hash of a resource's data. This can be used to decide if a resource is changed. For 'ApplyOnce' ClusterResourceSet.spec.strategy, this is no-op as that strategy does not act on change.",
												MarkdownDescription: "Hash is the hash of a resource's data. This can be used to decide if a resource is changed. For 'ApplyOnce' ClusterResourceSet.spec.strategy, this is no-op as that strategy does not act on change.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the resource. Supported kinds are: Secrets and ConfigMaps.",
												MarkdownDescription: "Kind of the resource. Supported kinds are: Secrets and ConfigMaps.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"last_applied_time": schema.StringAttribute{
												Description:         "LastAppliedTime identifies when this resource was last applied to the cluster.",
												MarkdownDescription: "LastAppliedTime identifies when this resource was last applied to the cluster.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the resource that is in the same namespace with ClusterResourceSet object.",
												MarkdownDescription: "Name of the resource that is in the same namespace with ClusterResourceSet object.",
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

func (r *AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_addons_cluster_x_k8s_io_cluster_resource_set_binding_v1alpha4")

	var data AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "addons.cluster.x-k8s.io", Version: "v1alpha4", Resource: "clusterresourcesetbindings"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AddonsClusterXK8SIoClusterResourceSetBindingV1Alpha4DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("addons.cluster.x-k8s.io/v1alpha4")
	data.Kind = pointer.String("ClusterResourceSetBinding")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
