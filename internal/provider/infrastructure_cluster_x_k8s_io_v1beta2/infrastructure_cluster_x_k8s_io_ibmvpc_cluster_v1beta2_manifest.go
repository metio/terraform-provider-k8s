/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest{}
}

type InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest struct{}

type InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2ManifestData struct {
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
		ControlPlaneEndpoint *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
		ControlPlaneLoadBalancer *struct {
			AdditionalListeners *[]struct {
				Port *int64 `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"additional_listeners" json:"additionalListeners,omitempty"`
			Id     *string `tfsdk:"id" json:"id,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Public *bool   `tfsdk:"public" json:"public,omitempty"`
		} `tfsdk:"control_plane_load_balancer" json:"controlPlaneLoadBalancer,omitempty"`
		Network *struct {
			ControlPlaneSubnets *[]struct {
				Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Zone *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"control_plane_subnets" json:"controlPlaneSubnets,omitempty"`
			ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
			Vpc           *struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"vpc" json:"vpc,omitempty"`
			WorkerSubnets *[]struct {
				Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Zone *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"worker_subnets" json:"workerSubnets,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		Region        *string `tfsdk:"region" json:"region,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		Vpc           *string `tfsdk:"vpc" json:"vpc,omitempty"`
		Zone          *string `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMVPCCluster is the Schema for the ibmvpcclusters API.",
		MarkdownDescription: "IBMVPCCluster is the Schema for the ibmvpcclusters API.",
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
				Description:         "IBMVPCClusterSpec defines the desired state of IBMVPCCluster.",
				MarkdownDescription: "IBMVPCClusterSpec defines the desired state of IBMVPCCluster.",
				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.SingleNestedAttribute{
						Description:         "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						MarkdownDescription: "ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.",
						Attributes: map[string]schema.Attribute{
							"host": schema.StringAttribute{
								Description:         "The hostname on which the API server is serving.",
								MarkdownDescription: "The hostname on which the API server is serving.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "The port on which the API server is serving.",
								MarkdownDescription: "The port on which the API server is serving.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"control_plane_load_balancer": schema.SingleNestedAttribute{
						Description:         "ControlPlaneLoadBalancer is optional configuration for customizing control plane behavior.",
						MarkdownDescription: "ControlPlaneLoadBalancer is optional configuration for customizing control plane behavior.",
						Attributes: map[string]schema.Attribute{
							"additional_listeners": schema.ListNestedAttribute{
								Description:         "AdditionalListeners sets the additional listeners for the control plane load balancer.",
								MarkdownDescription: "AdditionalListeners sets the additional listeners for the control plane load balancer.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Description:         "Port sets the port for the additional listener.",
											MarkdownDescription: "Port sets the port for the additional listener.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(65535),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"id": schema.StringAttribute{
								Description:         "id of the loadbalancer",
								MarkdownDescription: "id of the loadbalancer",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(64),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[-0-9a-z_]+$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name sets the name of the VPC load balancer.",
								MarkdownDescription: "Name sets the name of the VPC load balancer.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
								},
							},

							"public": schema.BoolAttribute{
								Description:         "public indicates that load balancer is public or private",
								MarkdownDescription: "public indicates that load balancer is public or private",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "network represents the VPC network to use for the cluster.",
						MarkdownDescription: "network represents the VPC network to use for the cluster.",
						Attributes: map[string]schema.Attribute{
							"control_plane_subnets": schema.ListNestedAttribute{
								Description:         "controlPlaneSubnets is a set of Subnet's which define the Control Plane subnets.",
								MarkdownDescription: "controlPlaneSubnets is a set of Subnet's which define the Control Plane subnets.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cidr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(64),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[-0-9a-z_]+$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
											},
										},

										"zone": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"resource_group": schema.StringAttribute{
								Description:         "resourceGroup is the name of the Resource Group containing all of the newtork resources.This can be different than the Resource Group containing the remaining cluster resources.",
								MarkdownDescription: "resourceGroup is the name of the Resource Group containing all of the newtork resources.This can be different than the Resource Group containing the remaining cluster resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vpc": schema.SingleNestedAttribute{
								Description:         "vpc defines the IBM Cloud VPC for extended VPC Infrastructure support.",
								MarkdownDescription: "vpc defines the IBM Cloud VPC for extended VPC Infrastructure support.",
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "id of the resource.",
										MarkdownDescription: "id of the resource.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"name": schema.StringAttribute{
										Description:         "name of the resource.",
										MarkdownDescription: "name of the resource.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"worker_subnets": schema.ListNestedAttribute{
								Description:         "workerSubnets is a set of Subnet's which define the Worker subnets.",
								MarkdownDescription: "workerSubnets is a set of Subnet's which define the Worker subnets.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cidr": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(64),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[-0-9a-z_]+$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`), ""),
											},
										},

										"zone": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

					"region": schema.StringAttribute{
						Description:         "The IBM Cloud Region the cluster lives in.",
						MarkdownDescription: "The IBM Cloud Region the cluster lives in.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"resource_group": schema.StringAttribute{
						Description:         "The VPC resources should be created under the resource group.",
						MarkdownDescription: "The VPC resources should be created under the resource group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"vpc": schema.StringAttribute{
						Description:         "The Name of VPC.",
						MarkdownDescription: "The Name of VPC.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"zone": schema.StringAttribute{
						Description:         "The Name of availability zone.",
						MarkdownDescription: "The Name of availability zone.",
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

func (r *InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibmvpc_cluster_v1beta2_manifest")

	var model InfrastructureClusterXK8SIoIbmvpcclusterV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IBMVPCCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
