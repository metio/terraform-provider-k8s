/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha4

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4Manifest{}
)

func NewInfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4Manifest{}
}

type InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4Manifest struct{}

type InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4ManifestData struct {
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
		Template *struct {
			Spec *struct {
				ControlPlaneEndpoint *struct {
					Host *string `tfsdk:"host" json:"host,omitempty"`
					Port *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"control_plane_endpoint" json:"controlPlaneEndpoint,omitempty"`
				IdentityRef *struct {
					Kind *string `tfsdk:"kind" json:"kind,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"identity_ref" json:"identityRef,omitempty"`
				Server     *string `tfsdk:"server" json:"server,omitempty"`
				Thumbprint *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest"
}

func (r *InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereClusterTemplate is the Schema for the vsphereclustertemplates API Deprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "VSphereClusterTemplate is the Schema for the vsphereclustertemplates API Deprecated: This type will be removed in one of the next releases.",
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
				Description:         "VSphereClusterTemplateSpec defines the desired state of VSphereClusterTemplate",
				MarkdownDescription: "VSphereClusterTemplateSpec defines the desired state of VSphereClusterTemplate",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "VSphereClusterSpec defines the desired state of VSphereCluster",
								MarkdownDescription: "VSphereClusterSpec defines the desired state of VSphereCluster",
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

									"identity_ref": schema.SingleNestedAttribute{
										Description:         "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that contains the identity to use when reconciling the cluster.",
										MarkdownDescription: "IdentityRef is a reference to either a Secret or VSphereClusterIdentity that contains the identity to use when reconciling the cluster.",
										Attributes: map[string]schema.Attribute{
											"kind": schema.StringAttribute{
												Description:         "Kind of the identity. Can either be VSphereClusterIdentity or Secret",
												MarkdownDescription: "Kind of the identity. Can either be VSphereClusterIdentity or Secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("VSphereClusterIdentity", "Secret"),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Name of the identity.",
												MarkdownDescription: "Name of the identity.",
												Required:            true,
												Optional:            false,
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

									"server": schema.StringAttribute{
										Description:         "Server is the address of the vSphere endpoint.",
										MarkdownDescription: "Server is the address of the vSphere endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thumbprint": schema.StringAttribute{
										Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate",
										MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
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

func (r *InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_cluster_template_v1alpha4_manifest")

	var model InfrastructureClusterXK8SIoVsphereClusterTemplateV1Alpha4ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha4")
	model.Kind = pointer.String("VSphereClusterTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
