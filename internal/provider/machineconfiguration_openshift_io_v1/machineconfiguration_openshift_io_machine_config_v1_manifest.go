/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package machineconfiguration_openshift_io_v1

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
	_ datasource.DataSource = &MachineconfigurationOpenshiftIoMachineConfigV1Manifest{}
)

func NewMachineconfigurationOpenshiftIoMachineConfigV1Manifest() datasource.DataSource {
	return &MachineconfigurationOpenshiftIoMachineConfigV1Manifest{}
}

type MachineconfigurationOpenshiftIoMachineConfigV1Manifest struct{}

type MachineconfigurationOpenshiftIoMachineConfigV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BaseOSExtensionsContainerImage *string            `tfsdk:"base_os_extensions_container_image" json:"baseOSExtensionsContainerImage,omitempty"`
		Config                         *map[string]string `tfsdk:"config" json:"config,omitempty"`
		Extensions                     *[]string          `tfsdk:"extensions" json:"extensions,omitempty"`
		Fips                           *bool              `tfsdk:"fips" json:"fips,omitempty"`
		KernelArguments                *[]string          `tfsdk:"kernel_arguments" json:"kernelArguments,omitempty"`
		KernelType                     *string            `tfsdk:"kernel_type" json:"kernelType,omitempty"`
		OsImageURL                     *string            `tfsdk:"os_image_url" json:"osImageURL,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MachineconfigurationOpenshiftIoMachineConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_machineconfiguration_openshift_io_machine_config_v1_manifest"
}

func (r *MachineconfigurationOpenshiftIoMachineConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MachineConfig defines the configuration for a machine  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "MachineConfig defines the configuration for a machine  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "MachineConfigSpec is the spec for MachineConfig",
				MarkdownDescription: "MachineConfigSpec is the spec for MachineConfig",
				Attributes: map[string]schema.Attribute{
					"base_os_extensions_container_image": schema.StringAttribute{
						Description:         "BaseOSExtensionsContainerImage specifies the remote location that will be used to fetch the extensions container matching a new-format OS image",
						MarkdownDescription: "BaseOSExtensionsContainerImage specifies the remote location that will be used to fetch the extensions container matching a new-format OS image",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config": schema.MapAttribute{
						Description:         "Config is a Ignition Config object.",
						MarkdownDescription: "Config is a Ignition Config object.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extensions": schema.ListAttribute{
						Description:         "extensions contains a list of additional features that can be enabled on host",
						MarkdownDescription: "extensions contains a list of additional features that can be enabled on host",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"fips": schema.BoolAttribute{
						Description:         "fips controls FIPS mode",
						MarkdownDescription: "fips controls FIPS mode",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kernel_arguments": schema.ListAttribute{
						Description:         "kernelArguments contains a list of kernel arguments to be added",
						MarkdownDescription: "kernelArguments contains a list of kernel arguments to be added",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kernel_type": schema.StringAttribute{
						Description:         "kernelType contains which kernel we want to be running like default (traditional), realtime, 64k-pages (aarch64 only).",
						MarkdownDescription: "kernelType contains which kernel we want to be running like default (traditional), realtime, 64k-pages (aarch64 only).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"os_image_url": schema.StringAttribute{
						Description:         "OSImageURL specifies the remote location that will be used to fetch the OS.",
						MarkdownDescription: "OSImageURL specifies the remote location that will be used to fetch the OS.",
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

func (r *MachineconfigurationOpenshiftIoMachineConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_machineconfiguration_openshift_io_machine_config_v1_manifest")

	var model MachineconfigurationOpenshiftIoMachineConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("machineconfiguration.openshift.io/v1")
	model.Kind = pointer.String("MachineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
