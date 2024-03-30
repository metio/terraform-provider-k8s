/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package appmesh_k8s_aws_v1beta2

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
	_ datasource.DataSource = &AppmeshK8SAwsVirtualServiceV1Beta2Manifest{}
)

func NewAppmeshK8SAwsVirtualServiceV1Beta2Manifest() datasource.DataSource {
	return &AppmeshK8SAwsVirtualServiceV1Beta2Manifest{}
}

type AppmeshK8SAwsVirtualServiceV1Beta2Manifest struct{}

type AppmeshK8SAwsVirtualServiceV1Beta2ManifestData struct {
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
		AwsName *string `tfsdk:"aws_name" json:"awsName,omitempty"`
		MeshRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Uid  *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"mesh_ref" json:"meshRef,omitempty"`
		Provider *struct {
			VirtualNode *struct {
				VirtualNodeARN *string `tfsdk:"virtual_node_arn" json:"virtualNodeARN,omitempty"`
				VirtualNodeRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"virtual_node_ref" json:"virtualNodeRef,omitempty"`
			} `tfsdk:"virtual_node" json:"virtualNode,omitempty"`
			VirtualRouter *struct {
				VirtualRouterARN *string `tfsdk:"virtual_router_arn" json:"virtualRouterARN,omitempty"`
				VirtualRouterRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"virtual_router_ref" json:"virtualRouterRef,omitempty"`
			} `tfsdk:"virtual_router" json:"virtualRouter,omitempty"`
		} `tfsdk:"provider" json:"provider,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppmeshK8SAwsVirtualServiceV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_appmesh_k8s_aws_virtual_service_v1beta2_manifest"
}

func (r *AppmeshK8SAwsVirtualServiceV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VirtualService is the Schema for the virtualservices API",
		MarkdownDescription: "VirtualService is the Schema for the virtualservices API",
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
				Description:         "VirtualServiceSpec defines the desired state of VirtualService refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualServiceSpec.html",
				MarkdownDescription: "VirtualServiceSpec defines the desired state of VirtualService refers to https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualServiceSpec.html",
				Attributes: map[string]schema.Attribute{
					"aws_name": schema.StringAttribute{
						Description:         "AWSName is the AppMesh VirtualService object's name. If unspecified or empty, it defaults to be '${name}.${namespace}' of k8s VirtualService",
						MarkdownDescription: "AWSName is the AppMesh VirtualService object's name. If unspecified or empty, it defaults to be '${name}.${namespace}' of k8s VirtualService",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mesh_ref": schema.SingleNestedAttribute{
						Description:         "A reference to k8s Mesh CR that this VirtualService belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						MarkdownDescription: "A reference to k8s Mesh CR that this VirtualService belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field.  Populated by the system. Read-only.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of Mesh CR",
								MarkdownDescription: "Name is the name of Mesh CR",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID is the UID of Mesh CR",
								MarkdownDescription: "UID is the UID of Mesh CR",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "The provider for virtual services. You can specify a single virtual node or virtual router.",
						MarkdownDescription: "The provider for virtual services. You can specify a single virtual node or virtual router.",
						Attributes: map[string]schema.Attribute{
							"virtual_node": schema.SingleNestedAttribute{
								Description:         "The virtual node associated with a virtual service.",
								MarkdownDescription: "The virtual node associated with a virtual service.",
								Attributes: map[string]schema.Attribute{
									"virtual_node_arn": schema.StringAttribute{
										Description:         "Amazon Resource Name to AppMesh VirtualNode object that is acting as a service provider. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
										MarkdownDescription: "Amazon Resource Name to AppMesh VirtualNode object that is acting as a service provider. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"virtual_node_ref": schema.SingleNestedAttribute{
										Description:         "Reference to Kubernetes VirtualNode CR in cluster that is acting as a service provider. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
										MarkdownDescription: "Reference to Kubernetes VirtualNode CR in cluster that is acting as a service provider. Exactly one of 'virtualNodeRef' or 'virtualNodeARN' must be specified.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name is the name of VirtualNode CR",
												MarkdownDescription: "Name is the name of VirtualNode CR",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
												MarkdownDescription: "Namespace is the namespace of VirtualNode CR. If unspecified, defaults to the referencing object's namespace",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"virtual_router": schema.SingleNestedAttribute{
								Description:         "The virtual router associated with a virtual service.",
								MarkdownDescription: "The virtual router associated with a virtual service.",
								Attributes: map[string]schema.Attribute{
									"virtual_router_arn": schema.StringAttribute{
										Description:         "Amazon Resource Name to AppMesh VirtualRouter object that is acting as a service provider. Exactly one of 'virtualRouterRef' or 'virtualRouterARN' must be specified.",
										MarkdownDescription: "Amazon Resource Name to AppMesh VirtualRouter object that is acting as a service provider. Exactly one of 'virtualRouterRef' or 'virtualRouterARN' must be specified.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"virtual_router_ref": schema.SingleNestedAttribute{
										Description:         "Reference to Kubernetes VirtualRouter CR in cluster that is acting as a service provider. Exactly one of 'virtualRouterRef' or 'virtualRouterARN' must be specified.",
										MarkdownDescription: "Reference to Kubernetes VirtualRouter CR in cluster that is acting as a service provider. Exactly one of 'virtualRouterRef' or 'virtualRouterARN' must be specified.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name is the name of VirtualRouter CR",
												MarkdownDescription: "Name is the name of VirtualRouter CR",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace of VirtualRouter CR. If unspecified, defaults to the referencing object's namespace",
												MarkdownDescription: "Namespace is the namespace of VirtualRouter CR. If unspecified, defaults to the referencing object's namespace",
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
								Required: false,
								Optional: true,
								Computed: false,
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
		},
	}
}

func (r *AppmeshK8SAwsVirtualServiceV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_appmesh_k8s_aws_virtual_service_v1beta2_manifest")

	var model AppmeshK8SAwsVirtualServiceV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("appmesh.k8s.aws/v1beta2")
	model.Kind = pointer.String("VirtualService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
