/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package multicluster_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &MulticlusterXK8SIoAppliedWorkV1Alpha1Manifest{}
)

func NewMulticlusterXK8SIoAppliedWorkV1Alpha1Manifest() datasource.DataSource {
	return &MulticlusterXK8SIoAppliedWorkV1Alpha1Manifest{}
}

type MulticlusterXK8SIoAppliedWorkV1Alpha1Manifest struct{}

type MulticlusterXK8SIoAppliedWorkV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		WorkName      *string `tfsdk:"work_name" json:"workName,omitempty"`
		WorkNamespace *string `tfsdk:"work_namespace" json:"workNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MulticlusterXK8SIoAppliedWorkV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_multicluster_x_k8s_io_applied_work_v1alpha1_manifest"
}

func (r *MulticlusterXK8SIoAppliedWorkV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AppliedWork represents an applied work on managed cluster that is placedon a managed cluster. An appliedwork links to a work on a hub recording resourcesdeployed in the managed cluster.When the agent is removed from managed cluster, cluster-admin on managed clustercan delete appliedwork to remove resources deployed by the agent.The name of the appliedwork must be the same as {work name}The namespace of the appliedwork should be the same as the resource applied onthe managed cluster.",
		MarkdownDescription: "AppliedWork represents an applied work on managed cluster that is placedon a managed cluster. An appliedwork links to a work on a hub recording resourcesdeployed in the managed cluster.When the agent is removed from managed cluster, cluster-admin on managed clustercan delete appliedwork to remove resources deployed by the agent.The name of the appliedwork must be the same as {work name}The namespace of the appliedwork should be the same as the resource applied onthe managed cluster.",
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
				Description:         "Spec represents the desired configuration of AppliedWork.",
				MarkdownDescription: "Spec represents the desired configuration of AppliedWork.",
				Attributes: map[string]schema.Attribute{
					"work_name": schema.StringAttribute{
						Description:         "WorkName represents the name of the related work on the hub.",
						MarkdownDescription: "WorkName represents the name of the related work on the hub.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"work_namespace": schema.StringAttribute{
						Description:         "WorkNamespace represents the namespace of the related work on the hub.",
						MarkdownDescription: "WorkNamespace represents the namespace of the related work on the hub.",
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

func (r *MulticlusterXK8SIoAppliedWorkV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_multicluster_x_k8s_io_applied_work_v1alpha1_manifest")

	var model MulticlusterXK8SIoAppliedWorkV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("multicluster.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("AppliedWork")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
