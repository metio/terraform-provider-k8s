/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha3

import (
	"context"
	"encoding/json"
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
	_ datasource.DataSource              = &InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource{}
	_ datasource.DataSourceWithConfigure = &InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource{}
)

func NewInfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource{}
}

type InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource struct {
	kubernetesClient dynamic.Interface
}

type InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ControlPlane        *bool   `tfsdk:"control_plane" json:"controlPlane,omitempty"`
		FailureDomain       *string `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
		PlacementConstraint *struct {
			Folder       *string `tfsdk:"folder" json:"folder,omitempty"`
			ResourcePool *string `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
		} `tfsdk:"placement_constraint" json:"placementConstraint,omitempty"`
		Server *string `tfsdk:"server" json:"server,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3"
}

func (r *InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereDeploymentZone is the Schema for the vspheredeploymentzones API  Deprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "VSphereDeploymentZone is the Schema for the vspheredeploymentzones API  Deprecated: This type will be removed in one of the next releases.",
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
				Description:         "VSphereDeploymentZoneSpec defines the desired state of VSphereDeploymentZone",
				MarkdownDescription: "VSphereDeploymentZoneSpec defines the desired state of VSphereDeploymentZone",
				Attributes: map[string]schema.Attribute{
					"control_plane": schema.BoolAttribute{
						Description:         "ControlPlane determines if this failure domain is suitable for use by control plane machines.",
						MarkdownDescription: "ControlPlane determines if this failure domain is suitable for use by control plane machines.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"failure_domain": schema.StringAttribute{
						Description:         "failureDomain is the name of the VSphereFailureDomain used for this VSphereDeploymentZone",
						MarkdownDescription: "failureDomain is the name of the VSphereFailureDomain used for this VSphereDeploymentZone",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"placement_constraint": schema.SingleNestedAttribute{
						Description:         "PlacementConstraint encapsulates the placement constraints used within this deployment zone.",
						MarkdownDescription: "PlacementConstraint encapsulates the placement constraints used within this deployment zone.",
						Attributes: map[string]schema.Attribute{
							"folder": schema.StringAttribute{
								Description:         "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
								MarkdownDescription: "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resource_pool": schema.StringAttribute{
								Description:         "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
								MarkdownDescription: "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"server": schema.StringAttribute{
						Description:         "Server is the address of the vSphere endpoint.",
						MarkdownDescription: "Server is the address of the vSphere endpoint.",
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
	}
}

func (r *InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3")

	var data InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha3", Resource: "vspheredeploymentzones"}).
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

	var readResponse InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha3")
	data.Kind = pointer.String("VSphereDeploymentZone")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
