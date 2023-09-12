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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest{}
)

func NewInfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest{}
}

type InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest struct{}

type InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1ManifestData struct {
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

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest"
}

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the tag that represents this failure domain",
								MarkdownDescription: "Name is the name of the tag that represents this failure domain",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tag_category": schema.StringAttribute{
								Description:         "TagCategory is the category used for the tag",
								MarkdownDescription: "TagCategory is the category used for the tag",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								MarkdownDescription: "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Datacenter", "ComputeCluster", "HostGroup"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"topology": schema.SingleNestedAttribute{
						Description:         "Topology describes a given failure domain using vSphere constructs",
						MarkdownDescription: "Topology describes a given failure domain using vSphere constructs",
						Attributes: map[string]schema.Attribute{
							"compute_cluster": schema.StringAttribute{
								Description:         "ComputeCluster as the failure domain",
								MarkdownDescription: "ComputeCluster as the failure domain",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"datacenter": schema.StringAttribute{
								Description:         "Datacenter as the failure domain.",
								MarkdownDescription: "Datacenter as the failure domain.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"datastore": schema.StringAttribute{
								Description:         "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								MarkdownDescription: "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hosts": schema.SingleNestedAttribute{
								Description:         "Hosts has information required for placement of machines on VSphere hosts.",
								MarkdownDescription: "Hosts has information required for placement of machines on VSphere hosts.",
								Attributes: map[string]schema.Attribute{
									"host_group_name": schema.StringAttribute{
										Description:         "HostGroupName is the name of the Host group",
										MarkdownDescription: "HostGroupName is the name of the Host group",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"vm_group_name": schema.StringAttribute{
										Description:         "VMGroupName is the name of the VM group",
										MarkdownDescription: "VMGroupName is the name of the VM group",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"networks": schema.ListAttribute{
								Description:         "Networks is the list of networks within this failure domain",
								MarkdownDescription: "Networks is the list of networks within this failure domain",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"zone": schema.SingleNestedAttribute{
						Description:         "Zone defines the name and type of a zone",
						MarkdownDescription: "Zone defines the name and type of a zone",
						Attributes: map[string]schema.Attribute{
							"auto_configure": schema.BoolAttribute{
								Description:         "AutoConfigure tags the Type which is specified in the Topology  Deprecated: This field is going to be removed in a future release.",
								MarkdownDescription: "AutoConfigure tags the Type which is specified in the Topology  Deprecated: This field is going to be removed in a future release.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the tag that represents this failure domain",
								MarkdownDescription: "Name is the name of the tag that represents this failure domain",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tag_category": schema.StringAttribute{
								Description:         "TagCategory is the category used for the tag",
								MarkdownDescription: "TagCategory is the category used for the tag",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								MarkdownDescription: "Type is the type of failure domain, the current values are 'Datacenter', 'ComputeCluster' and 'HostGroup'",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Datacenter", "ComputeCluster", "HostGroup"),
								},
							},
						},
						Required: true,
						Optional: false,
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

func (r *InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_failure_domain_v1beta1_manifest")

	var model InfrastructureClusterXK8SIoVsphereFailureDomainV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("VSphereFailureDomain")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
