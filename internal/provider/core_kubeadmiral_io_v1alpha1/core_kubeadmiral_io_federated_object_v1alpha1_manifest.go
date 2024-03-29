/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_kubeadmiral_io_v1alpha1

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
	_ datasource.DataSource = &CoreKubeadmiralIoFederatedObjectV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoFederatedObjectV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoFederatedObjectV1Alpha1Manifest{}
}

type CoreKubeadmiralIoFederatedObjectV1Alpha1Manifest struct{}

type CoreKubeadmiralIoFederatedObjectV1Alpha1ManifestData struct {
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
		Follows *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"follows" json:"follows,omitempty"`
		Overrides *[]struct {
			Clusters *[]struct {
				Cluster *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Patches *[]struct {
					Op    *string            `tfsdk:"op" json:"op,omitempty"`
					Path  *string            `tfsdk:"path" json:"path,omitempty"`
					Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"patches" json:"patches,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			Controller *string `tfsdk:"controller" json:"controller,omitempty"`
		} `tfsdk:"overrides" json:"overrides,omitempty"`
		Placements *[]struct {
			Controller *string `tfsdk:"controller" json:"controller,omitempty"`
			Placement  *[]struct {
				Cluster *string `tfsdk:"cluster" json:"cluster,omitempty"`
			} `tfsdk:"placement" json:"placement,omitempty"`
		} `tfsdk:"placements" json:"placements,omitempty"`
		Template *map[string]string `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CoreKubeadmiralIoFederatedObjectV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_federated_object_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoFederatedObjectV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FederatedObject describes a namespace-scoped Kubernetes object and how it should be propagated to different member clusters.",
		MarkdownDescription: "FederatedObject describes a namespace-scoped Kubernetes object and how it should be propagated to different member clusters.",
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
				Description:         "Spec defines the desired behavior of the FederatedObject.",
				MarkdownDescription: "Spec defines the desired behavior of the FederatedObject.",
				Attributes: map[string]schema.Attribute{
					"follows": schema.ListNestedAttribute{
						Description:         "Follows defines other objects, or 'leaders', that the Kubernetes object should follow during propagation, i.e. the Kubernetes object should be propagated to all member clusters that its 'leaders' are placed in.",
						MarkdownDescription: "Follows defines other objects, or 'leaders', that the Kubernetes object should follow during propagation, i.e. the Kubernetes object should be propagated to all member clusters that its 'leaders' are placed in.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
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

					"overrides": schema.ListNestedAttribute{
						Description:         "Overrides describe the overrides that should be applied to the base template of the Kubernetes object before it is propagated to individual member clusters.",
						MarkdownDescription: "Overrides describe the overrides that should be applied to the base template of the Kubernetes object before it is propagated to individual member clusters.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"clusters": schema.ListNestedAttribute{
									Description:         "Override is the list of member clusters and their respective override patches.",
									MarkdownDescription: "Override is the list of member clusters and their respective override patches.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cluster": schema.StringAttribute{
												Description:         "Cluster is the name of the member cluster.",
												MarkdownDescription: "Cluster is the name of the member cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"patches": schema.ListNestedAttribute{
												Description:         "Patches is the list of override patches for the member cluster.",
												MarkdownDescription: "Patches is the list of override patches for the member cluster.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"op": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
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
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"controller": schema.StringAttribute{
									Description:         "Controller identifies the controller responsible for this override.",
									MarkdownDescription: "Controller identifies the controller responsible for this override.",
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

					"placements": schema.ListNestedAttribute{
						Description:         "Placements describe the member clusters that the Kubernetes object will be propagated to, which is a union of all the listed clusters.",
						MarkdownDescription: "Placements describe the member clusters that the Kubernetes object will be propagated to, which is a union of all the listed clusters.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"controller": schema.StringAttribute{
									Description:         "Controller identifies the controller responsible for this placement.",
									MarkdownDescription: "Controller identifies the controller responsible for this placement.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"placement": schema.ListNestedAttribute{
									Description:         "Placement is the list of member clusters that the Kubernetes object should be propagated to.",
									MarkdownDescription: "Placement is the list of member clusters that the Kubernetes object should be propagated to.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cluster": schema.StringAttribute{
												Description:         "Cluster is the name of the member cluster.",
												MarkdownDescription: "Cluster is the name of the member cluster.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"template": schema.MapAttribute{
						Description:         "Template is the base template of the Kubernetes object to be propagated.",
						MarkdownDescription: "Template is the base template of the Kubernetes object to be propagated.",
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

func (r *CoreKubeadmiralIoFederatedObjectV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_federated_object_v1alpha1_manifest")

	var model CoreKubeadmiralIoFederatedObjectV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("FederatedObject")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
