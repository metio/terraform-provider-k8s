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
	_ datasource.DataSource = &WorkKarmadaIoWorkV1Alpha1Manifest{}
)

func NewWorkKarmadaIoWorkV1Alpha1Manifest() datasource.DataSource {
	return &WorkKarmadaIoWorkV1Alpha1Manifest{}
}

type WorkKarmadaIoWorkV1Alpha1Manifest struct{}

type WorkKarmadaIoWorkV1Alpha1ManifestData struct {
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
		PreserveResourcesOnDeletion *bool `tfsdk:"preserve_resources_on_deletion" json:"preserveResourcesOnDeletion,omitempty"`
		SuspendDispatching          *bool `tfsdk:"suspend_dispatching" json:"suspendDispatching,omitempty"`
		Workload                    *struct {
			Manifests *[]map[string]string `tfsdk:"manifests" json:"manifests,omitempty"`
		} `tfsdk:"workload" json:"workload,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *WorkKarmadaIoWorkV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_work_karmada_io_work_v1alpha1_manifest"
}

func (r *WorkKarmadaIoWorkV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Work defines a list of resources to be deployed on the member cluster.",
		MarkdownDescription: "Work defines a list of resources to be deployed on the member cluster.",
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
				Description:         "Spec represents the desired behavior of Work.",
				MarkdownDescription: "Spec represents the desired behavior of Work.",
				Attributes: map[string]schema.Attribute{
					"preserve_resources_on_deletion": schema.BoolAttribute{
						Description:         "PreserveResourcesOnDeletion controls whether resources should be preserved on the member cluster when the Work object is deleted. If set to true, resources will be preserved on the member cluster. Default is false, which means resources will be deleted along with the Work object.",
						MarkdownDescription: "PreserveResourcesOnDeletion controls whether resources should be preserved on the member cluster when the Work object is deleted. If set to true, resources will be preserved on the member cluster. Default is false, which means resources will be deleted along with the Work object.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suspend_dispatching": schema.BoolAttribute{
						Description:         "SuspendDispatching controls whether dispatching should be suspended, nil means not suspend. Note: true means stop propagating to the corresponding member cluster, and does not prevent status collection.",
						MarkdownDescription: "SuspendDispatching controls whether dispatching should be suspended, nil means not suspend. Note: true means stop propagating to the corresponding member cluster, and does not prevent status collection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"workload": schema.SingleNestedAttribute{
						Description:         "Workload represents the manifest workload to be deployed on managed cluster.",
						MarkdownDescription: "Workload represents the manifest workload to be deployed on managed cluster.",
						Attributes: map[string]schema.Attribute{
							"manifests": schema.ListAttribute{
								Description:         "Manifests represents a list of Kubernetes resources to be deployed on the managed cluster.",
								MarkdownDescription: "Manifests represents a list of Kubernetes resources to be deployed on the managed cluster.",
								ElementType:         types.MapType{ElemType: types.StringType},
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *WorkKarmadaIoWorkV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_work_karmada_io_work_v1alpha1_manifest")

	var model WorkKarmadaIoWorkV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("work.karmada.io/v1alpha1")
	model.Kind = pointer.String("Work")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
