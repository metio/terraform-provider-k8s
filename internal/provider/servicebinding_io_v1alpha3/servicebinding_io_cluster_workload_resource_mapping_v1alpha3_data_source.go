/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package servicebinding_io_v1alpha3

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
	_ datasource.DataSource              = &ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource{}
	_ datasource.DataSourceWithConfigure = &ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource{}
)

func NewServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource() datasource.DataSource {
	return &ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource{}
}

type ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource struct {
	kubernetesClient dynamic.Interface
}

type ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_servicebinding_io_cluster_workload_resource_mapping_v1alpha3"
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
									Optional:            false,
									Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name is a Restricted JSONPath that references the name of the container with the container-like workload resource fragment. If not defined, container name filtering is ignored.",
												MarkdownDescription: "Name is a Restricted JSONPath that references the name of the container with the container-like workload resource fragment. If not defined, container name filtering is ignored.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"path": schema.StringAttribute{
												Description:         "Path is the JSONPath within the workload resource that matches an existing fragment that is container-like.",
												MarkdownDescription: "Path is the JSONPath within the workload resource that matches an existing fragment that is container-like.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"volume_mounts": schema.StringAttribute{
												Description:         "VolumeMounts is a Restricted JSONPath that references the slice of volume mounts for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.volumeMounts'.",
												MarkdownDescription: "VolumeMounts is a Restricted JSONPath that references the slice of volume mounts for the container with the container-like workload resource fragment. The referenced location is created if it does not exist. Defaults to '.volumeMounts'.",
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

								"version": schema.StringAttribute{
									Description:         "Version is the version of the workload resource that this mapping is for.",
									MarkdownDescription: "Version is the version of the workload resource that this mapping is for.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"volumes": schema.StringAttribute{
									Description:         "Volumes is a Restricted JSONPath that references the slice of volumes within the workload resource. Defaults to '.spec.template.spec.volumes'.",
									MarkdownDescription: "Volumes is a Restricted JSONPath that references the slice of volumes within the workload resource. Defaults to '.spec.template.spec.volumes'.",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3")

	var data ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "servicebinding.io", Version: "v1alpha3", Resource: "ClusterWorkloadResourceMapping"}).
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

	var readResponse ServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("servicebinding.io/v1alpha3")
	data.Kind = pointer.String("ClusterWorkloadResourceMapping")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
