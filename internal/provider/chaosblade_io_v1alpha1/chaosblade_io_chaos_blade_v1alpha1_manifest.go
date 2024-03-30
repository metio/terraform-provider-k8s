/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaosblade_io_v1alpha1

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
	_ datasource.DataSource = &ChaosbladeIoChaosBladeV1Alpha1Manifest{}
)

func NewChaosbladeIoChaosBladeV1Alpha1Manifest() datasource.DataSource {
	return &ChaosbladeIoChaosBladeV1Alpha1Manifest{}
}

type ChaosbladeIoChaosBladeV1Alpha1Manifest struct{}

type ChaosbladeIoChaosBladeV1Alpha1ManifestData struct {
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
		Experiments *[]struct {
			Action   *string `tfsdk:"action" json:"action,omitempty"`
			Desc     *string `tfsdk:"desc" json:"desc,omitempty"`
			Matchers *[]struct {
				Name  *string   `tfsdk:"name" json:"name,omitempty"`
				Value *[]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"matchers" json:"matchers,omitempty"`
			Scope  *string `tfsdk:"scope" json:"scope,omitempty"`
			Target *string `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"experiments" json:"experiments,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosbladeIoChaosBladeV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaosblade_io_chaos_blade_v1alpha1_manifest"
}

func (r *ChaosbladeIoChaosBladeV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ChaosBlade is the Schema for the chaosblades API",
		MarkdownDescription: "ChaosBlade is the Schema for the chaosblades API",
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
				Description:         "ChaosBladeSpec defines the desired state of ChaosBlade",
				MarkdownDescription: "ChaosBladeSpec defines the desired state of ChaosBlade",
				Attributes: map[string]schema.Attribute{
					"experiments": schema.ListNestedAttribute{
						Description:         "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'operator-sdk generate k8s' to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
						MarkdownDescription: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'operator-sdk generate k8s' to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action is the experiment scenario of the target, such as delay, load",
									MarkdownDescription: "Action is the experiment scenario of the target, such as delay, load",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"desc": schema.StringAttribute{
									Description:         "Desc is the experiment description",
									MarkdownDescription: "Desc is the experiment description",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"matchers": schema.ListNestedAttribute{
									Description:         "Matchers is the experiment rules",
									MarkdownDescription: "Matchers is the experiment rules",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name is the name of flag",
												MarkdownDescription: "Name is the name of flag",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"value": schema.ListAttribute{
												Description:         "TODO: Temporarily defined as an array for all flags Value is the value of flag",
												MarkdownDescription: "TODO: Temporarily defined as an array for all flags Value is the value of flag",
												ElementType:         types.StringType,
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

								"scope": schema.StringAttribute{
									Description:         "Scope is the area of the experiments, currently support node, pod and container",
									MarkdownDescription: "Scope is the area of the experiments, currently support node, pod and container",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"target": schema.StringAttribute{
									Description:         "Target is the experiment target, such as cpu, network",
									MarkdownDescription: "Target is the experiment target, such as cpu, network",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ChaosbladeIoChaosBladeV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaosblade_io_chaos_blade_v1alpha1_manifest")

	var model ChaosbladeIoChaosBladeV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("chaosblade.io/v1alpha1")
	model.Kind = pointer.String("ChaosBlade")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
