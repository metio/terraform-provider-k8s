/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta2

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2Manifest{}
}

type InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2Manifest struct{}

type InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2ManifestData struct {
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
		Bucket          *string `tfsdk:"bucket" json:"bucket,omitempty"`
		ClusterName     *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		DeletePolicy    *string `tfsdk:"delete_policy" json:"deletePolicy,omitempty"`
		Object          *string `tfsdk:"object" json:"object,omitempty"`
		Region          *string `tfsdk:"region" json:"region,omitempty"`
		ServiceInstance *struct {
			Id    *string `tfsdk:"id" json:"id,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Regex *string `tfsdk:"regex" json:"regex,omitempty"`
		} `tfsdk:"service_instance" json:"serviceInstance,omitempty"`
		ServiceInstanceID *string `tfsdk:"service_instance_id" json:"serviceInstanceID,omitempty"`
		StorageType       *string `tfsdk:"storage_type" json:"storageType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta2_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMPowerVSImage is the Schema for the ibmpowervsimages API.",
		MarkdownDescription: "IBMPowerVSImage is the Schema for the ibmpowervsimages API.",
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
				Description:         "IBMPowerVSImageSpec defines the desired state of IBMPowerVSImage.",
				MarkdownDescription: "IBMPowerVSImageSpec defines the desired state of IBMPowerVSImage.",
				Attributes: map[string]schema.Attribute{
					"bucket": schema.StringAttribute{
						Description:         "Cloud Object Storage bucket name; bucket-name[/optional/folder]",
						MarkdownDescription: "Cloud Object Storage bucket name; bucket-name[/optional/folder]",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the name of the Cluster this object belongs to.",
						MarkdownDescription: "ClusterName is the name of the Cluster this object belongs to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"delete_policy": schema.StringAttribute{
						Description:         "DeletePolicy defines the policy used to identify images to be preserved beyond the lifecycle of associated cluster.",
						MarkdownDescription: "DeletePolicy defines the policy used to identify images to be preserved beyond the lifecycle of associated cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("delete", "retain"),
						},
					},

					"object": schema.StringAttribute{
						Description:         "Cloud Object Storage image filename.",
						MarkdownDescription: "Cloud Object Storage image filename.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"region": schema.StringAttribute{
						Description:         "Cloud Object Storage region.",
						MarkdownDescription: "Cloud Object Storage region.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"service_instance": schema.SingleNestedAttribute{
						Description:         "serviceInstance is the reference to the Power VS workspace on which the server instance(VM) will be created. Power VS workspace is a container for all Power VS instances at a specific geographic region. serviceInstance can be created via IBM Cloud catalog or CLI. supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli. More detail about Power VS service instance. https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server when omitted system will dynamically create the service instance",
						MarkdownDescription: "serviceInstance is the reference to the Power VS workspace on which the server instance(VM) will be created. Power VS workspace is a container for all Power VS instances at a specific geographic region. serviceInstance can be created via IBM Cloud catalog or CLI. supported serviceInstance identifier in PowerVSResource are Name and ID and that can be obtained from IBM Cloud UI or IBM Cloud cli. More detail about Power VS service instance. https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server when omitted system will dynamically create the service instance",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "ID of resource",
								MarkdownDescription: "ID of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name of resource",
								MarkdownDescription: "Name of resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"regex": schema.StringAttribute{
								Description:         "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
								MarkdownDescription: "Regular expression to match resource, In case of multiple resources matches the provided regular expression the first matched resource will be selected",
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

					"service_instance_id": schema.StringAttribute{
						Description:         "ServiceInstanceID is the id of the power cloud instance where the image will get imported. Deprecated: use ServiceInstance instead",
						MarkdownDescription: "ServiceInstanceID is the id of the power cloud instance where the image will get imported. Deprecated: use ServiceInstance instead",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"storage_type": schema.StringAttribute{
						Description:         "Type of storage, storage pool with the most available space will be selected.",
						MarkdownDescription: "Type of storage, storage pool with the most available space will be selected.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("tier1", "tier3"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_image_v1beta2_manifest")

	var model InfrastructureClusterXK8SIoIbmpowerVsimageV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IBMPowerVSImage")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
