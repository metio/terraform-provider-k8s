/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1Manifest{}
)

func NewInfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1Manifest{}
}

type InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1Manifest struct{}

type InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1beta1_manifest"
}

func (r *InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereDeploymentZone is the Schema for the vspheredeploymentzones API.",
		MarkdownDescription: "VSphereDeploymentZone is the Schema for the vspheredeploymentzones API.",
		Attributes: map[string]schema.Attribute{
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
				Description:         "VSphereDeploymentZoneSpec defines the desired state of VSphereDeploymentZone.",
				MarkdownDescription: "VSphereDeploymentZoneSpec defines the desired state of VSphereDeploymentZone.",
				Attributes: map[string]schema.Attribute{
					"control_plane": schema.BoolAttribute{
						Description:         "ControlPlane determines if this failure domain is suitable for use by control plane machines.",
						MarkdownDescription: "ControlPlane determines if this failure domain is suitable for use by control plane machines.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failure_domain": schema.StringAttribute{
						Description:         "FailureDomain is the name of the VSphereFailureDomain used for this VSphereDeploymentZone",
						MarkdownDescription: "FailureDomain is the name of the VSphereFailureDomain used for this VSphereDeploymentZone",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"placement_constraint": schema.SingleNestedAttribute{
						Description:         "PlacementConstraint encapsulates the placement constraintsused within this deployment zone.",
						MarkdownDescription: "PlacementConstraint encapsulates the placement constraintsused within this deployment zone.",
						Attributes: map[string]schema.Attribute{
							"folder": schema.StringAttribute{
								Description:         "Folder is the name or inventory path of the folder in which thevirtual machine is created/located.",
								MarkdownDescription: "Folder is the name or inventory path of the folder in which thevirtual machine is created/located.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_pool": schema.StringAttribute{
								Description:         "ResourcePool is the name or inventory path of the resource pool in whichthe virtual machine is created/located.",
								MarkdownDescription: "ResourcePool is the name or inventory path of the resource pool in whichthe virtual machine is created/located.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"server": schema.StringAttribute{
						Description:         "Server is the address of the vSphere endpoint.",
						MarkdownDescription: "Server is the address of the vSphere endpoint.",
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
	}
}

func (r *InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1beta1_manifest")

	var model InfrastructureClusterXK8SIoVsphereDeploymentZoneV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("VSphereDeploymentZone")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
