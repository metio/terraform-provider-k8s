/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package work_karmada_io_v1alpha1

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
	_ datasource.DataSource = &WorkKarmadaIoResourceBindingV1Alpha1Manifest{}
)

func NewWorkKarmadaIoResourceBindingV1Alpha1Manifest() datasource.DataSource {
	return &WorkKarmadaIoResourceBindingV1Alpha1Manifest{}
}

type WorkKarmadaIoResourceBindingV1Alpha1Manifest struct{}

type WorkKarmadaIoResourceBindingV1Alpha1ManifestData struct {
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
		Clusters *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Replicas *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"clusters" json:"clusters,omitempty"`
		Resource *struct {
			ApiVersion          *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind                *string            `tfsdk:"kind" json:"kind,omitempty"`
			Name                *string            `tfsdk:"name" json:"name,omitempty"`
			Namespace           *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			Replicas            *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			ResourcePerReplicas *map[string]string `tfsdk:"resource_per_replicas" json:"resourcePerReplicas,omitempty"`
			ResourceVersion     *string            `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
		} `tfsdk:"resource" json:"resource,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *WorkKarmadaIoResourceBindingV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_work_karmada_io_resource_binding_v1alpha1_manifest"
}

func (r *WorkKarmadaIoResourceBindingV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceBinding represents a binding of a kubernetes resource with a propagation policy.",
		MarkdownDescription: "ResourceBinding represents a binding of a kubernetes resource with a propagation policy.",
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
				Description:         "Spec represents the desired behavior.",
				MarkdownDescription: "Spec represents the desired behavior.",
				Attributes: map[string]schema.Attribute{
					"clusters": schema.ListNestedAttribute{
						Description:         "Clusters represents target member clusters where the resource to be deployed.",
						MarkdownDescription: "Clusters represents target member clusters where the resource to be deployed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of target cluster.",
									MarkdownDescription: "Name of target cluster.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Replicas in target cluster",
									MarkdownDescription: "Replicas in target cluster",
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

					"resource": schema.SingleNestedAttribute{
						Description:         "Resource represents the Kubernetes resource to be propagated.",
						MarkdownDescription: "Resource represents the Kubernetes resource to be propagated.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion represents the API version of the referent.",
								MarkdownDescription: "APIVersion represents the API version of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind represents the Kind of the referent.",
								MarkdownDescription: "Kind represents the Kind of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name represents the name of the referent.",
								MarkdownDescription: "Name represents the name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace represents the namespace for the referent. For non-namespace scoped resources(e.g. 'ClusterRole')，do not need specify Namespace, and for namespace scoped resources, Namespace is required. If Namespace is not specified, means the resource is non-namespace scoped.",
								MarkdownDescription: "Namespace represents the namespace for the referent. For non-namespace scoped resources(e.g. 'ClusterRole')，do not need specify Namespace, and for namespace scoped resources, Namespace is required. If Namespace is not specified, means the resource is non-namespace scoped.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas represents the replica number of the referencing resource.",
								MarkdownDescription: "Replicas represents the replica number of the referencing resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_per_replicas": schema.MapAttribute{
								Description:         "ReplicaResourceRequirements represents the resources required by each replica.",
								MarkdownDescription: "ReplicaResourceRequirements represents the resources required by each replica.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "ResourceVersion represents the internal version of the referenced object, that can be used by clients to determine when object has changed.",
								MarkdownDescription: "ResourceVersion represents the internal version of the referenced object, that can be used by clients to determine when object has changed.",
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
	}
}

func (r *WorkKarmadaIoResourceBindingV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_work_karmada_io_resource_binding_v1alpha1_manifest")

	var model WorkKarmadaIoResourceBindingV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("work.karmada.io/v1alpha1")
	model.Kind = pointer.String("ResourceBinding")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
