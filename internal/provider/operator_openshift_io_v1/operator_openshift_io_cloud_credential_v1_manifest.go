/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_openshift_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OperatorOpenshiftIoCloudCredentialV1Manifest{}
)

func NewOperatorOpenshiftIoCloudCredentialV1Manifest() datasource.DataSource {
	return &OperatorOpenshiftIoCloudCredentialV1Manifest{}
}

type OperatorOpenshiftIoCloudCredentialV1Manifest struct{}

type OperatorOpenshiftIoCloudCredentialV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CredentialsMode            *string            `tfsdk:"credentials_mode" json:"credentialsMode,omitempty"`
		LogLevel                   *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		ManagementState            *string            `tfsdk:"management_state" json:"managementState,omitempty"`
		ObservedConfig             *map[string]string `tfsdk:"observed_config" json:"observedConfig,omitempty"`
		OperatorLogLevel           *string            `tfsdk:"operator_log_level" json:"operatorLogLevel,omitempty"`
		UnsupportedConfigOverrides *map[string]string `tfsdk:"unsupported_config_overrides" json:"unsupportedConfigOverrides,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenshiftIoCloudCredentialV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_openshift_io_cloud_credential_v1_manifest"
}

func (r *OperatorOpenshiftIoCloudCredentialV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CloudCredential provides a means to configure an operator to manage CredentialsRequests.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "CloudCredential provides a means to configure an operator to manage CredentialsRequests.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "CloudCredentialSpec is the specification of the desired behavior of the cloud-credential-operator.",
				MarkdownDescription: "CloudCredentialSpec is the specification of the desired behavior of the cloud-credential-operator.",
				Attributes: map[string]schema.Attribute{
					"credentials_mode": schema.StringAttribute{
						Description:         "CredentialsMode allows informing CCO that it should not attempt to dynamically determine the root cloud credentials capabilities, and it should just run in the specified mode. It also allows putting the operator into 'manual' mode if desired. Leaving the field in default mode runs CCO so that the cluster's cloud credentials will be dynamically probed for capabilities (on supported clouds/platforms). Supported modes: AWS/Azure/GCP: '' (Default), 'Mint', 'Passthrough', 'Manual' Others: Do not set value as other platforms only support running in 'Passthrough'",
						MarkdownDescription: "CredentialsMode allows informing CCO that it should not attempt to dynamically determine the root cloud credentials capabilities, and it should just run in the specified mode. It also allows putting the operator into 'manual' mode if desired. Leaving the field in default mode runs CCO so that the cluster's cloud credentials will be dynamically probed for capabilities (on supported clouds/platforms). Supported modes: AWS/Azure/GCP: '' (Default), 'Mint', 'Passthrough', 'Manual' Others: Do not set value as other platforms only support running in 'Passthrough'",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Manual", "Mint", "Passthrough"),
						},
					},

					"log_level": schema.StringAttribute{
						Description:         "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"management_state": schema.StringAttribute{
						Description:         "managementState indicates whether and how the operator should manage the component",
						MarkdownDescription: "managementState indicates whether and how the operator should manage the component",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(Managed|Unmanaged|Force|Removed)$`), ""),
						},
					},

					"observed_config": schema.MapAttribute{
						Description:         "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						MarkdownDescription: "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operator_log_level": schema.StringAttribute{
						Description:         "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"unsupported_config_overrides": schema.MapAttribute{
						Description:         "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						MarkdownDescription: "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						ElementType:         types.StringType,
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
	}
}

func (r *OperatorOpenshiftIoCloudCredentialV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_openshift_io_cloud_credential_v1_manifest")

	var model OperatorOpenshiftIoCloudCredentialV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.openshift.io/v1")
	model.Kind = pointer.String("CloudCredential")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
