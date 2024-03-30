/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_kubeadmiral_io_v1alpha1

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
	_ datasource.DataSource = &CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1Manifest{}
}

type CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1Manifest struct{}

type CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		FilterPath      *string   `tfsdk:"filter_path" json:"filterPath,omitempty"`
		HttpTimeout     *string   `tfsdk:"http_timeout" json:"httpTimeout,omitempty"`
		PayloadVersions *[]string `tfsdk:"payload_versions" json:"payloadVersions,omitempty"`
		ScorePath       *string   `tfsdk:"score_path" json:"scorePath,omitempty"`
		SelectPath      *string   `tfsdk:"select_path" json:"selectPath,omitempty"`
		TlsConfig       *struct {
			CaData     *string `tfsdk:"ca_data" json:"caData,omitempty"`
			CertData   *string `tfsdk:"cert_data" json:"certData,omitempty"`
			Insecure   *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
			KeyData    *string `tfsdk:"key_data" json:"keyData,omitempty"`
			ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
		} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
		UrlPrefix *string `tfsdk:"url_prefix" json:"urlPrefix,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_scheduler_plugin_webhook_configuration_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SchedulerPluginWebhookConfiguration is a webhook that can be used as a scheduler plugin.",
		MarkdownDescription: "SchedulerPluginWebhookConfiguration is a webhook that can be used as a scheduler plugin.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"filter_path": schema.StringAttribute{
						Description:         "Path for the filter call, empty if not supported. This path is appended to the URLPrefix when issuing the filter call to webhook.",
						MarkdownDescription: "Path for the filter call, empty if not supported. This path is appended to the URLPrefix when issuing the filter call to webhook.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"http_timeout": schema.StringAttribute{
						Description:         "HTTPTimeout specifies the timeout duration for a call to the webhook. Timeout fails the scheduling of the workload. Defaults to 5 seconds.",
						MarkdownDescription: "HTTPTimeout specifies the timeout duration for a call to the webhook. Timeout fails the scheduling of the workload. Defaults to 5 seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"payload_versions": schema.ListAttribute{
						Description:         "PayloadVersions is an ordered list of preferred request and response versions the webhook expects. The scheduler will try to use the first version in the list which it supports. If none of the versions specified in this list supported by the scheduler, scheduling will fail for this object.",
						MarkdownDescription: "PayloadVersions is an ordered list of preferred request and response versions the webhook expects. The scheduler will try to use the first version in the list which it supports. If none of the versions specified in this list supported by the scheduler, scheduling will fail for this object.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"score_path": schema.StringAttribute{
						Description:         "Path for the score call, empty if not supported. This verb is appended to the URLPrefix when issuing the score call to webhook.",
						MarkdownDescription: "Path for the score call, empty if not supported. This verb is appended to the URLPrefix when issuing the score call to webhook.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"select_path": schema.StringAttribute{
						Description:         "Path for the select call, empty if not supported. This verb is appended to the URLPrefix when issuing the select call to webhook.",
						MarkdownDescription: "Path for the select call, empty if not supported. This verb is appended to the URLPrefix when issuing the select call to webhook.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls_config": schema.SingleNestedAttribute{
						Description:         "TLSConfig specifies the transport layer security config.",
						MarkdownDescription: "TLSConfig specifies the transport layer security config.",
						Attributes: map[string]schema.Attribute{
							"ca_data": schema.StringAttribute{
								Description:         "CAData holds PEM-encoded bytes (typically read from a root certificates bundle).",
								MarkdownDescription: "CAData holds PEM-encoded bytes (typically read from a root certificates bundle).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"cert_data": schema.StringAttribute{
								Description:         "CertData holds PEM-encoded bytes (typically read from a client certificate file).",
								MarkdownDescription: "CertData holds PEM-encoded bytes (typically read from a client certificate file).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"insecure": schema.BoolAttribute{
								Description:         "Server should be accessed without verifying the TLS certificate. For testing only.",
								MarkdownDescription: "Server should be accessed without verifying the TLS certificate. For testing only.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_data": schema.StringAttribute{
								Description:         "KeyData holds PEM-encoded bytes (typically read from a client certificate key file).",
								MarkdownDescription: "KeyData holds PEM-encoded bytes (typically read from a client certificate key file).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"server_name": schema.StringAttribute{
								Description:         "ServerName is passed to the server for SNI and is used in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used.",
								MarkdownDescription: "ServerName is passed to the server for SNI and is used in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"url_prefix": schema.StringAttribute{
						Description:         "URLPrefix at which the webhook is available",
						MarkdownDescription: "URLPrefix at which the webhook is available",
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

func (r *CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_scheduler_plugin_webhook_configuration_v1alpha1_manifest")

	var model CoreKubeadmiralIoSchedulerPluginWebhookConfigurationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("SchedulerPluginWebhookConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
