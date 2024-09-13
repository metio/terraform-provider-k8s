/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package bmc_tinkerbell_org_v1alpha1

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
	_ datasource.DataSource = &BmcTinkerbellOrgJobV1Alpha1Manifest{}
)

func NewBmcTinkerbellOrgJobV1Alpha1Manifest() datasource.DataSource {
	return &BmcTinkerbellOrgJobV1Alpha1Manifest{}
}

type BmcTinkerbellOrgJobV1Alpha1Manifest struct{}

type BmcTinkerbellOrgJobV1Alpha1ManifestData struct {
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
		MachineRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"machine_ref" json:"machineRef,omitempty"`
		Tasks *[]struct {
			OneTimeBootDeviceAction *struct {
				Device  *[]string `tfsdk:"device" json:"device,omitempty"`
				EfiBoot *bool     `tfsdk:"efi_boot" json:"efiBoot,omitempty"`
			} `tfsdk:"one_time_boot_device_action" json:"oneTimeBootDeviceAction,omitempty"`
			PowerAction        *string `tfsdk:"power_action" json:"powerAction,omitempty"`
			VirtualMediaAction *struct {
				Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
				MediaURL *string `tfsdk:"media_url" json:"mediaURL,omitempty"`
			} `tfsdk:"virtual_media_action" json:"virtualMediaAction,omitempty"`
		} `tfsdk:"tasks" json:"tasks,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BmcTinkerbellOrgJobV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_bmc_tinkerbell_org_job_v1alpha1_manifest"
}

func (r *BmcTinkerbellOrgJobV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Job is the Schema for the bmcjobs API.",
		MarkdownDescription: "Job is the Schema for the bmcjobs API.",
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
				Description:         "JobSpec defines the desired state of Job.",
				MarkdownDescription: "JobSpec defines the desired state of Job.",
				Attributes: map[string]schema.Attribute{
					"machine_ref": schema.SingleNestedAttribute{
						Description:         "MachineRef represents the Machine resource to execute the job. All the tasks in the job are executed for the same Machine.",
						MarkdownDescription: "MachineRef represents the Machine resource to execute the job. All the tasks in the job are executed for the same Machine.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the Machine.",
								MarkdownDescription: "Name of the Machine.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace the Machine resides in.",
								MarkdownDescription: "Namespace the Machine resides in.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tasks": schema.ListNestedAttribute{
						Description:         "Tasks represents a list of baseboard management actions to be executed. The tasks are executed sequentially. Controller waits for one task to complete before executing the next. If a single task fails, job execution stops and sets condition Failed. Condition Completed is set only if all the tasks were successful.",
						MarkdownDescription: "Tasks represents a list of baseboard management actions to be executed. The tasks are executed sequentially. Controller waits for one task to complete before executing the next. If a single task fails, job execution stops and sets condition Failed. Condition Completed is set only if all the tasks were successful.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"one_time_boot_device_action": schema.SingleNestedAttribute{
									Description:         "OneTimeBootDeviceAction represents a baseboard management one time set boot device operation.",
									MarkdownDescription: "OneTimeBootDeviceAction represents a baseboard management one time set boot device operation.",
									Attributes: map[string]schema.Attribute{
										"device": schema.ListAttribute{
											Description:         "Devices represents the boot devices, in order for setting one time boot. Currently only the first device in the slice is used to set one time boot.",
											MarkdownDescription: "Devices represents the boot devices, in order for setting one time boot. Currently only the first device in the slice is used to set one time boot.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"efi_boot": schema.BoolAttribute{
											Description:         "EFIBoot instructs the machine to use EFI boot.",
											MarkdownDescription: "EFIBoot instructs the machine to use EFI boot.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"power_action": schema.StringAttribute{
									Description:         "PowerAction represents a baseboard management power operation.",
									MarkdownDescription: "PowerAction represents a baseboard management power operation.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("on", "off", "soft", "status", "cycle", "reset"),
									},
								},

								"virtual_media_action": schema.SingleNestedAttribute{
									Description:         "VirtualMediaAction represents a baseboard management virtual media insert/eject.",
									MarkdownDescription: "VirtualMediaAction represents a baseboard management virtual media insert/eject.",
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"media_url": schema.StringAttribute{
											Description:         "mediaURL represents the URL of the image to be inserted into the virtual media, or empty to eject media.",
											MarkdownDescription: "mediaURL represents the URL of the image to be inserted into the virtual media, or empty to eject media.",
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

func (r *BmcTinkerbellOrgJobV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_bmc_tinkerbell_org_job_v1alpha1_manifest")

	var model BmcTinkerbellOrgJobV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("bmc.tinkerbell.org/v1alpha1")
	model.Kind = pointer.String("Job")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
