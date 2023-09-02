/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_karmada_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &NetworkingKarmadaIoMultiClusterServiceV1Alpha1Manifest{}
)

func NewNetworkingKarmadaIoMultiClusterServiceV1Alpha1Manifest() datasource.DataSource {
	return &NetworkingKarmadaIoMultiClusterServiceV1Alpha1Manifest{}
}

type NetworkingKarmadaIoMultiClusterServiceV1Alpha1Manifest struct{}

type NetworkingKarmadaIoMultiClusterServiceV1Alpha1ManifestData struct {
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
		Ports *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"ports" json:"ports,omitempty"`
		Range *struct {
			ClusterNames *[]string `tfsdk:"cluster_names" json:"clusterNames,omitempty"`
		} `tfsdk:"range" json:"range,omitempty"`
		Types *[]string `tfsdk:"types" json:"types,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingKarmadaIoMultiClusterServiceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_karmada_io_multi_cluster_service_v1alpha1_manifest"
}

func (r *NetworkingKarmadaIoMultiClusterServiceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MultiClusterService is a named abstraction of multi-cluster software service. The name field of MultiClusterService is the same as that of Service name. Services with the same name in different clusters are regarded as the same service and are associated with the same MultiClusterService. MultiClusterService can control the exposure of services to outside multiple clusters, and also enable service discovery between clusters.",
		MarkdownDescription: "MultiClusterService is a named abstraction of multi-cluster software service. The name field of MultiClusterService is the same as that of Service name. Services with the same name in different clusters are regarded as the same service and are associated with the same MultiClusterService. MultiClusterService can control the exposure of services to outside multiple clusters, and also enable service discovery between clusters.",
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
				Description:         "Spec is the desired state of the MultiClusterService.",
				MarkdownDescription: "Spec is the desired state of the MultiClusterService.",
				Attributes: map[string]schema.Attribute{
					"ports": schema.ListNestedAttribute{
						Description:         "Ports is the list of ports that are exposed by this MultiClusterService. No specified port will be filtered out during the service exposure and discovery process. All ports in the referencing service will be exposed by default.",
						MarkdownDescription: "Ports is the list of ports that are exposed by this MultiClusterService. No specified port will be filtered out during the service exposure and discovery process. All ports in the referencing service will be exposed by default.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name of the port that needs to be exposed within the service. The port name must be the same as that defined in the service.",
									MarkdownDescription: "Name is the name of the port that needs to be exposed within the service. The port name must be the same as that defined in the service.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "Port specifies the exposed service port.",
									MarkdownDescription: "Port specifies the exposed service port.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"range": schema.SingleNestedAttribute{
						Description:         "Range specifies the ranges where the referencing service should be exposed. Only valid and optional in case of Types contains CrossCluster. If not set and Types contains CrossCluster, all clusters will be selected, that means the referencing service will be exposed across all registered clusters.",
						MarkdownDescription: "Range specifies the ranges where the referencing service should be exposed. Only valid and optional in case of Types contains CrossCluster. If not set and Types contains CrossCluster, all clusters will be selected, that means the referencing service will be exposed across all registered clusters.",
						Attributes: map[string]schema.Attribute{
							"cluster_names": schema.ListAttribute{
								Description:         "ClusterNames is the list of clusters to be selected.",
								MarkdownDescription: "ClusterNames is the list of clusters to be selected.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"types": schema.ListAttribute{
						Description:         "Types specifies how to expose the service referencing by this MultiClusterService.",
						MarkdownDescription: "Types specifies how to expose the service referencing by this MultiClusterService.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *NetworkingKarmadaIoMultiClusterServiceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_karmada_io_multi_cluster_service_v1alpha1_manifest")

	var model NetworkingKarmadaIoMultiClusterServiceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("networking.karmada.io/v1alpha1")
	model.Kind = pointer.String("MultiClusterService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
