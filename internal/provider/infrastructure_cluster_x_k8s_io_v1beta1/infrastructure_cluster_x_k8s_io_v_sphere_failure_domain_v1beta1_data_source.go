/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource              = &InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource{}
)

func NewInfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource{}
}

type InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Region *struct {
			AutoConfigure *bool   `tfsdk:"auto_configure" json:"autoConfigure,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			TagCategory   *string `tfsdk:"tag_category" json:"tagCategory,omitempty"`
			Type          *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"region" json:"region,omitempty"`
		Topology *struct {
			ComputeCluster *string `tfsdk:"compute_cluster" json:"computeCluster,omitempty"`
			Datacenter     *string `tfsdk:"datacenter" json:"datacenter,omitempty"`
			Datastore      *string `tfsdk:"datastore" json:"datastore,omitempty"`
			Hosts          *struct {
				HostGroupName *string `tfsdk:"host_group_name" json:"hostGroupName,omitempty"`
				VmGroupName   *string `tfsdk:"vm_group_name" json:"vmGroupName,omitempty"`
			} `tfsdk:"hosts" json:"hosts,omitempty"`
			Networks *[]string `tfsdk:"networks" json:"networks,omitempty"`
		} `tfsdk:"topology" json:"topology,omitempty"`
		Zone *struct {
			AutoConfigure *bool   `tfsdk:"auto_configure" json:"autoConfigure,omitempty"`
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			TagCategory   *string `tfsdk:"tag_category" json:"tagCategory,omitempty"`
			Type          *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1"
}

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereFailureDomain is the Schema for the vspherefailuredomains API",
		MarkdownDescription: "VSphereFailureDomain is the Schema for the vspherefailuredomains API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "VSphereFailureDomainSpec defines the desired state of VSphereFailureDomain",
				MarkdownDescription: "VSphereFailureDomainSpec defines the desired state of VSphereFailureDomain",
				Attributes: map[string]schema.Attribute{
					"region": schema.SingleNestedAttribute{
						Description:         "Region defines the name and type of a region",
						MarkdownDescription: "Region defines the name and type of a region",
						Attributes: map[string]schema.Attribute{
							"auto_configure": schema.BoolAttribute{
								Description:         "AutoConfigure tags the Type which is specified in the Topology  Deprecated: This field is going to be removed in a future release.",
								MarkdownDescription: "AutoConfigure tags the Type which is specified in the Topology  Deprecated: This field is going to be removed in a future release.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the tag that represents this failure domain",
								MarkdownDescription: "Name is the name of the tag that represents this failure domain",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tag_category": schema.StringAttribute{
								Description:         "TagCategory is the category used for the tag",
								MarkdownDescription: "TagCategory is the category used for the tag",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								MarkdownDescription: "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"topology": schema.SingleNestedAttribute{
						Description:         "Topology describes a given failure domain using vSphere constructs",
						MarkdownDescription: "Topology describes a given failure domain using vSphere constructs",
						Attributes: map[string]schema.Attribute{
							"compute_cluster": schema.StringAttribute{
								Description:         "ComputeCluster as the failure domain",
								MarkdownDescription: "ComputeCluster as the failure domain",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"datacenter": schema.StringAttribute{
								Description:         "Datacenter as the failure domain.",
								MarkdownDescription: "Datacenter as the failure domain.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"datastore": schema.StringAttribute{
								Description:         "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								MarkdownDescription: "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"hosts": schema.SingleNestedAttribute{
								Description:         "Hosts has information required for placement of machines on VSphere hosts.",
								MarkdownDescription: "Hosts has information required for placement of machines on VSphere hosts.",
								Attributes: map[string]schema.Attribute{
									"host_group_name": schema.StringAttribute{
										Description:         "HostGroupName is the name of the Host group",
										MarkdownDescription: "HostGroupName is the name of the Host group",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"vm_group_name": schema.StringAttribute{
										Description:         "VMGroupName is the name of the VM group",
										MarkdownDescription: "VMGroupName is the name of the VM group",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"networks": schema.ListAttribute{
								Description:         "Networks is the list of networks within this failure domain",
								MarkdownDescription: "Networks is the list of networks within this failure domain",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"zone": schema.SingleNestedAttribute{
						Description:         "Zone defines the name and type of a zone",
						MarkdownDescription: "Zone defines the name and type of a zone",
						Attributes: map[string]schema.Attribute{
							"auto_configure": schema.BoolAttribute{
								Description:         "AutoConfigure tags the Type which is specified in the Topology  Deprecated: This field is going to be removed in a future release.",
								MarkdownDescription: "AutoConfigure tags the Type which is specified in the Topology  Deprecated: This field is going to be removed in a future release.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the tag that represents this failure domain",
								MarkdownDescription: "Name is the name of the tag that represents this failure domain",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tag_category": schema.StringAttribute{
								Description:         "TagCategory is the category used for the tag",
								MarkdownDescription: "TagCategory is the category used for the tag",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								MarkdownDescription: "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
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

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1")

	var data InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1beta1", Resource: "vspherefailuredomains"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	data.Kind = pointer.String("VSphereFailureDomain")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
