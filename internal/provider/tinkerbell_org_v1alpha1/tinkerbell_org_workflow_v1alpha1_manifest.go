/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tinkerbell_org_v1alpha1

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
	_ datasource.DataSource = &TinkerbellOrgWorkflowV1Alpha1Manifest{}
)

func NewTinkerbellOrgWorkflowV1Alpha1Manifest() datasource.DataSource {
	return &TinkerbellOrgWorkflowV1Alpha1Manifest{}
}

type TinkerbellOrgWorkflowV1Alpha1Manifest struct{}

type TinkerbellOrgWorkflowV1Alpha1ManifestData struct {
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
		BootOptions *struct {
			BootMode           *string `tfsdk:"boot_mode" json:"bootMode,omitempty"`
			IsoURL             *string `tfsdk:"iso_url" json:"isoURL,omitempty"`
			ToggleAllowNetboot *bool   `tfsdk:"toggle_allow_netboot" json:"toggleAllowNetboot,omitempty"`
		} `tfsdk:"boot_options" json:"bootOptions,omitempty"`
		HardwareMap *map[string]string `tfsdk:"hardware_map" json:"hardwareMap,omitempty"`
		HardwareRef *string            `tfsdk:"hardware_ref" json:"hardwareRef,omitempty"`
		TemplateRef *string            `tfsdk:"template_ref" json:"templateRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TinkerbellOrgWorkflowV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tinkerbell_org_workflow_v1alpha1_manifest"
}

func (r *TinkerbellOrgWorkflowV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Workflow is the Schema for the Workflows API.",
		MarkdownDescription: "Workflow is the Schema for the Workflows API.",
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
				Description:         "WorkflowSpec defines the desired state of Workflow.",
				MarkdownDescription: "WorkflowSpec defines the desired state of Workflow.",
				Attributes: map[string]schema.Attribute{
					"boot_options": schema.SingleNestedAttribute{
						Description:         "BootOptions are options that control the booting of Hardware.",
						MarkdownDescription: "BootOptions are options that control the booting of Hardware.",
						Attributes: map[string]schema.Attribute{
							"boot_mode": schema.StringAttribute{
								Description:         "BootMode is the type of booting that will be done.",
								MarkdownDescription: "BootMode is the type of booting that will be done.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("netboot", "iso"),
								},
							},

							"iso_url": schema.StringAttribute{
								Description:         "ISOURL is the URL of the ISO that will be one-time booted. When this field is set, the controller will create a job.bmc.tinkerbell.org object for getting the associated hardware into a CDROM booting state. A HardwareRef that contains a spec.BmcRef must be provided.",
								MarkdownDescription: "ISOURL is the URL of the ISO that will be one-time booted. When this field is set, the controller will create a job.bmc.tinkerbell.org object for getting the associated hardware into a CDROM booting state. A HardwareRef that contains a spec.BmcRef must be provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"toggle_allow_netboot": schema.BoolAttribute{
								Description:         "ToggleAllowNetboot indicates whether the controller should toggle the field in the associated hardware for allowing PXE booting. This will be enabled before a Workflow is executed and disabled after the Workflow has completed successfully. A HardwareRef must be provided.",
								MarkdownDescription: "ToggleAllowNetboot indicates whether the controller should toggle the field in the associated hardware for allowing PXE booting. This will be enabled before a Workflow is executed and disabled after the Workflow has completed successfully. A HardwareRef must be provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"hardware_map": schema.MapAttribute{
						Description:         "A mapping of template devices to hadware mac addresses.",
						MarkdownDescription: "A mapping of template devices to hadware mac addresses.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hardware_ref": schema.StringAttribute{
						Description:         "Name of the Hardware associated with this workflow.",
						MarkdownDescription: "Name of the Hardware associated with this workflow.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template_ref": schema.StringAttribute{
						Description:         "Name of the Template associated with this workflow.",
						MarkdownDescription: "Name of the Template associated with this workflow.",
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

func (r *TinkerbellOrgWorkflowV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tinkerbell_org_workflow_v1alpha1_manifest")

	var model TinkerbellOrgWorkflowV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tinkerbell.org/v1alpha1")
	model.Kind = pointer.String("Workflow")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
