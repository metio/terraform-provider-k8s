/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package self_node_remediation_medik8s_io_v1alpha1

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
	_ datasource.DataSource = &SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1Manifest{}
)

func NewSelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1Manifest() datasource.DataSource {
	return &SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1Manifest{}
}

type SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1Manifest struct{}

type SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1ManifestData struct {
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
				RemediationStrategy *string `tfsdk:"remediation_strategy" json:"remediationStrategy,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_self_node_remediation_medik8s_io_self_node_remediation_template_v1alpha1_manifest"
}

func (r *SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SelfNodeRemediationTemplate is the Schema for the selfnoderemediationtemplates API",
		MarkdownDescription: "SelfNodeRemediationTemplate is the Schema for the selfnoderemediationtemplates API",
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
				Description:         "SelfNodeRemediationTemplateSpec defines the desired state of SelfNodeRemediationTemplate",
				MarkdownDescription: "SelfNodeRemediationTemplateSpec defines the desired state of SelfNodeRemediationTemplate",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "Template defines the desired state of SelfNodeRemediationTemplate",
						MarkdownDescription: "Template defines the desired state of SelfNodeRemediationTemplate",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "SelfNodeRemediationSpec defines the desired state of SelfNodeRemediation",
								MarkdownDescription: "SelfNodeRemediationSpec defines the desired state of SelfNodeRemediation",
								Attributes: map[string]schema.Attribute{
									"remediation_strategy": schema.StringAttribute{
										Description:         "RemediationStrategy is the remediation method for unhealthy nodes.Currently, it could be either 'Automatic', 'OutOfServiceTaint' or 'ResourceDeletion'.ResourceDeletion will iterate over all pods and VolumeAttachment related to the unhealthy node and delete them.OutOfServiceTaint will add the out-of-service taint which is a new well-known taint 'node.kubernetes.io/out-of-service'that enables automatic deletion of pv-attached pods on failed nodes, 'out-of-service' taint is only supported on clusters with k8s version 1.26+ or OCP/OKD version 4.13+.Automatic will choose the most appropriate strategy during runtime.",
										MarkdownDescription: "RemediationStrategy is the remediation method for unhealthy nodes.Currently, it could be either 'Automatic', 'OutOfServiceTaint' or 'ResourceDeletion'.ResourceDeletion will iterate over all pods and VolumeAttachment related to the unhealthy node and delete them.OutOfServiceTaint will add the out-of-service taint which is a new well-known taint 'node.kubernetes.io/out-of-service'that enables automatic deletion of pv-attached pods on failed nodes, 'out-of-service' taint is only supported on clusters with k8s version 1.26+ or OCP/OKD version 4.13+.Automatic will choose the most appropriate strategy during runtime.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Automatic", "ResourceDeletion", "OutOfServiceTaint"),
										},
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

func (r *SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_self_node_remediation_medik8s_io_self_node_remediation_template_v1alpha1_manifest")

	var model SelfNodeRemediationMedik8SIoSelfNodeRemediationTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("self-node-remediation.medik8s.io/v1alpha1")
	model.Kind = pointer.String("SelfNodeRemediationTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
