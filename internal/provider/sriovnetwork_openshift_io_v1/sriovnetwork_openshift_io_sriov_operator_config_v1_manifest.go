/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sriovnetwork_openshift_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &SriovnetworkOpenshiftIoSriovOperatorConfigV1Manifest{}
)

func NewSriovnetworkOpenshiftIoSriovOperatorConfigV1Manifest() datasource.DataSource {
	return &SriovnetworkOpenshiftIoSriovOperatorConfigV1Manifest{}
}

type SriovnetworkOpenshiftIoSriovOperatorConfigV1Manifest struct{}

type SriovnetworkOpenshiftIoSriovOperatorConfigV1ManifestData struct {
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
		ConfigDaemonNodeSelector *map[string]string `tfsdk:"config_daemon_node_selector" json:"configDaemonNodeSelector,omitempty"`
		ConfigurationMode        *string            `tfsdk:"configuration_mode" json:"configurationMode,omitempty"`
		DisableDrain             *bool              `tfsdk:"disable_drain" json:"disableDrain,omitempty"`
		DisablePlugins           *[]string          `tfsdk:"disable_plugins" json:"disablePlugins,omitempty"`
		EnableInjector           *bool              `tfsdk:"enable_injector" json:"enableInjector,omitempty"`
		EnableOperatorWebhook    *bool              `tfsdk:"enable_operator_webhook" json:"enableOperatorWebhook,omitempty"`
		EnableOvsOffload         *bool              `tfsdk:"enable_ovs_offload" json:"enableOvsOffload,omitempty"`
		FeatureGates             *map[string]string `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		LogLevel                 *int64             `tfsdk:"log_level" json:"logLevel,omitempty"`
		UseCDI                   *bool              `tfsdk:"use_cdi" json:"useCDI,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SriovnetworkOpenshiftIoSriovOperatorConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sriovnetwork_openshift_io_sriov_operator_config_v1_manifest"
}

func (r *SriovnetworkOpenshiftIoSriovOperatorConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SriovOperatorConfig is the Schema for the sriovoperatorconfigs API",
		MarkdownDescription: "SriovOperatorConfig is the Schema for the sriovoperatorconfigs API",
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
				Description:         "SriovOperatorConfigSpec defines the desired state of SriovOperatorConfig",
				MarkdownDescription: "SriovOperatorConfigSpec defines the desired state of SriovOperatorConfig",
				Attributes: map[string]schema.Attribute{
					"config_daemon_node_selector": schema.MapAttribute{
						Description:         "NodeSelector selects the nodes to be configured",
						MarkdownDescription: "NodeSelector selects the nodes to be configured",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"configuration_mode": schema.StringAttribute{
						Description:         "Flag to enable the sriov-network-config-daemon to use a systemd service to configure SR-IOV devices on boot Default mode: daemon",
						MarkdownDescription: "Flag to enable the sriov-network-config-daemon to use a systemd service to configure SR-IOV devices on boot Default mode: daemon",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("daemon", "systemd"),
						},
					},

					"disable_drain": schema.BoolAttribute{
						Description:         "Flag to disable nodes drain during debugging",
						MarkdownDescription: "Flag to disable nodes drain during debugging",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_plugins": schema.ListAttribute{
						Description:         "DisablePlugins is a list of sriov-network-config-daemon plugins to disable",
						MarkdownDescription: "DisablePlugins is a list of sriov-network-config-daemon plugins to disable",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_injector": schema.BoolAttribute{
						Description:         "Flag to control whether the network resource injector webhook shall be deployed",
						MarkdownDescription: "Flag to control whether the network resource injector webhook shall be deployed",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_operator_webhook": schema.BoolAttribute{
						Description:         "Flag to control whether the operator admission controller webhook shall be deployed",
						MarkdownDescription: "Flag to control whether the operator admission controller webhook shall be deployed",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_ovs_offload": schema.BoolAttribute{
						Description:         "Flag to enable OVS hardware offload. Set to 'true' to provision switchdev-configuration.service and enable OpenvSwitch hw-offload on nodes.",
						MarkdownDescription: "Flag to enable OVS hardware offload. Set to 'true' to provision switchdev-configuration.service and enable OpenvSwitch hw-offload on nodes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"feature_gates": schema.MapAttribute{
						Description:         "FeatureGates to enable experimental features",
						MarkdownDescription: "FeatureGates to enable experimental features",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.Int64Attribute{
						Description:         "Flag to control the log verbose level of the operator. Set to '0' to show only the basic logs. And set to '2' to show all the available logs.",
						MarkdownDescription: "Flag to control the log verbose level of the operator. Set to '0' to show only the basic logs. And set to '2' to show all the available logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(2),
						},
					},

					"use_cdi": schema.BoolAttribute{
						Description:         "Flag to enable Container Device Interface mode for SR-IOV Network Device Plugin",
						MarkdownDescription: "Flag to enable Container Device Interface mode for SR-IOV Network Device Plugin",
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

func (r *SriovnetworkOpenshiftIoSriovOperatorConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sriovnetwork_openshift_io_sriov_operator_config_v1_manifest")

	var model SriovnetworkOpenshiftIoSriovOperatorConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sriovnetwork.openshift.io/v1")
	model.Kind = pointer.String("SriovOperatorConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
