/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1Manifest{}
)

func NewInfrastructureClusterXK8SIoTinkerbellClusterV1Beta1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1Manifest{}
}

type InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1Manifest struct{}

type InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1ManifestData struct {
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
		ControlPlaneEndpoint *struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
		ImageLookupBaseRegistry *string `tfsdk:"image_lookup_base_registry" json:"imageLookupBaseRegistry,omitempty"`
		ImageLookupFormat       *string `tfsdk:"image_lookup_format" json:"imageLookupFormat,omitempty"`
		ImageLookupOSDistro     *string `tfsdk:"image_lookup_os_distro" json:"imageLookupOSDistro,omitempty"`
		ImageLookupOSVersion    *string `tfsdk:"image_lookup_os_version" json:"imageLookupOSVersion,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest"
}

func (r *InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TinkerbellCluster is the Schema for the tinkerbellclusters API.",
		MarkdownDescription: "TinkerbellCluster is the Schema for the tinkerbellclusters API.",
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
				Description:         "TinkerbellClusterSpec defines the desired state of TinkerbellCluster.",
				MarkdownDescription: "TinkerbellClusterSpec defines the desired state of TinkerbellCluster.",
				Attributes: map[string]schema.Attribute{
					"control_plane_endpoint": schema.SingleNestedAttribute{
						Description:         "ControlPlaneEndpoint is a required field by ClusterAPI v1beta1.See https://cluster-api.sigs.k8s.io/developer/architecture/controllers/cluster.htmlfor more details.",
						MarkdownDescription: "ControlPlaneEndpoint is a required field by ClusterAPI v1beta1.See https://cluster-api.sigs.k8s.io/developer/architecture/controllers/cluster.htmlfor more details.",
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

					"image_lookup_base_registry": schema.StringAttribute{
						Description:         "ImageLookupBaseRegistry is the base Registry URL that is used for pulling images,if not set, the default will be to use ghcr.io/tinkerbell/cluster-api-provider-tinkerbell.",
						MarkdownDescription: "ImageLookupBaseRegistry is the base Registry URL that is used for pulling images,if not set, the default will be to use ghcr.io/tinkerbell/cluster-api-provider-tinkerbell.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_lookup_format": schema.StringAttribute{
						Description:         "ImageLookupFormat is the URL naming format to use for machine images whena machine does not specify. When set, this will be used for all cluster machinesunless a machine specifies a different ImageLookupFormat. Supports substitutionsfor {{.BaseRegistry}}, {{.OSDistro}}, {{.OSVersion}} and {{.KubernetesVersion}} withthe basse URL, OS distribution, OS version, and kubernetes version, respectively.BaseRegistry will be the value in ImageLookupBaseRegistry or ghcr.io/tinkerbell/cluster-api-provider-tinkerbell(the default), OSDistro will be the value in ImageLookupOSDistro or ubuntu (the default),OSVersion will be the value in ImageLookupOSVersion or default based on the OSDistro(if known), and the kubernetes version as defined by the packages produced bykubernetes/release: v1.13.0, v1.12.5-mybuild.1, or v1.17.3. For example, the defaultimage format of {{.BaseRegistry}}/{{.OSDistro}}-{{.OSVersion}}:{{.KubernetesVersion}}.gz willattempt to pull the image from that location. See also: https://golang.org/pkg/text/template/",
						MarkdownDescription: "ImageLookupFormat is the URL naming format to use for machine images whena machine does not specify. When set, this will be used for all cluster machinesunless a machine specifies a different ImageLookupFormat. Supports substitutionsfor {{.BaseRegistry}}, {{.OSDistro}}, {{.OSVersion}} and {{.KubernetesVersion}} withthe basse URL, OS distribution, OS version, and kubernetes version, respectively.BaseRegistry will be the value in ImageLookupBaseRegistry or ghcr.io/tinkerbell/cluster-api-provider-tinkerbell(the default), OSDistro will be the value in ImageLookupOSDistro or ubuntu (the default),OSVersion will be the value in ImageLookupOSVersion or default based on the OSDistro(if known), and the kubernetes version as defined by the packages produced bykubernetes/release: v1.13.0, v1.12.5-mybuild.1, or v1.17.3. For example, the defaultimage format of {{.BaseRegistry}}/{{.OSDistro}}-{{.OSVersion}}:{{.KubernetesVersion}}.gz willattempt to pull the image from that location. See also: https://golang.org/pkg/text/template/",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_lookup_os_distro": schema.StringAttribute{
						Description:         "ImageLookupOSDistro is the name of the OS distro to use when fetching machine images,if not set it will default to ubuntu.",
						MarkdownDescription: "ImageLookupOSDistro is the name of the OS distro to use when fetching machine images,if not set it will default to ubuntu.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_lookup_os_version": schema.StringAttribute{
						Description:         "ImageLookupOSVersion is the version of the OS distribution to use when fetching machineimages. If not set it will default based on ImageLookupOSDistro.",
						MarkdownDescription: "ImageLookupOSVersion is the version of the OS distribution to use when fetching machineimages. If not set it will default based on ImageLookupOSDistro.",
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

func (r *InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest")

	var model InfrastructureClusterXK8SIoTinkerbellClusterV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("TinkerbellCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
