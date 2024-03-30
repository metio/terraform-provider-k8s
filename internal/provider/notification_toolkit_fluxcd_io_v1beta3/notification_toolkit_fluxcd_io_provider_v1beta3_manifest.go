/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package notification_toolkit_fluxcd_io_v1beta3

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
	_ datasource.DataSource = &NotificationToolkitFluxcdIoProviderV1Beta3Manifest{}
)

func NewNotificationToolkitFluxcdIoProviderV1Beta3Manifest() datasource.DataSource {
	return &NotificationToolkitFluxcdIoProviderV1Beta3Manifest{}
}

type NotificationToolkitFluxcdIoProviderV1Beta3Manifest struct{}

type NotificationToolkitFluxcdIoProviderV1Beta3ManifestData struct {
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
		Address       *string `tfsdk:"address" json:"address,omitempty"`
		CertSecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
		Channel   *string `tfsdk:"channel" json:"channel,omitempty"`
		Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
		Proxy     *string `tfsdk:"proxy" json:"proxy,omitempty"`
		SecretRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		Suspend  *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
		Type     *string `tfsdk:"type" json:"type,omitempty"`
		Username *string `tfsdk:"username" json:"username,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NotificationToolkitFluxcdIoProviderV1Beta3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_notification_toolkit_fluxcd_io_provider_v1beta3_manifest"
}

func (r *NotificationToolkitFluxcdIoProviderV1Beta3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Provider is the Schema for the providers API",
		MarkdownDescription: "Provider is the Schema for the providers API",
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
				Description:         "ProviderSpec defines the desired state of the Provider.",
				MarkdownDescription: "ProviderSpec defines the desired state of the Provider.",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "Address specifies the endpoint, in a generic sense, to where alerts are sent.What kind of endpoint depends on the specific Provider type being used.For the generic Provider, for example, this is an HTTP/S address.For other Provider types this could be a project ID or a namespace.",
						MarkdownDescription: "Address specifies the endpoint, in a generic sense, to where alerts are sent.What kind of endpoint depends on the specific Provider type being used.For the generic Provider, for example, this is an HTTP/S address.For other Provider types this could be a project ID or a namespace.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(2048),
						},
					},

					"cert_secret_ref": schema.SingleNestedAttribute{
						Description:         "CertSecretRef specifies the Secret containinga PEM-encoded CA certificate (in the 'ca.crt' key).Note: Support for the 'caFile' key hasbeen deprecated.",
						MarkdownDescription: "CertSecretRef specifies the Secret containinga PEM-encoded CA certificate (in the 'ca.crt' key).Note: Support for the 'caFile' key hasbeen deprecated.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"channel": schema.StringAttribute{
						Description:         "Channel specifies the destination channel where events should be posted.",
						MarkdownDescription: "Channel specifies the destination channel where events should be posted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(2048),
						},
					},

					"interval": schema.StringAttribute{
						Description:         "Interval at which to reconcile the Provider with its Secret references.Deprecated and not used in v1beta3.",
						MarkdownDescription: "Interval at which to reconcile the Provider with its Secret references.Deprecated and not used in v1beta3.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m|h))+$`), ""),
						},
					},

					"proxy": schema.StringAttribute{
						Description:         "Proxy the HTTP/S address of the proxy server.",
						MarkdownDescription: "Proxy the HTTP/S address of the proxy server.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(2048),
							stringvalidator.RegexMatches(regexp.MustCompile(`^(http|https)://.*$`), ""),
						},
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretRef specifies the Secret containing the authenticationcredentials for this Provider.",
						MarkdownDescription: "SecretRef specifies the Secret containing the authenticationcredentials for this Provider.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend subsequentevents handling for this Provider.",
						MarkdownDescription: "Suspend tells the controller to suspend subsequentevents handling for this Provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout for sending alerts to the Provider.",
						MarkdownDescription: "Timeout for sending alerts to the Provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
						},
					},

					"type": schema.StringAttribute{
						Description:         "Type specifies which Provider implementation to use.",
						MarkdownDescription: "Type specifies which Provider implementation to use.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("slack", "discord", "msteams", "rocket", "generic", "generic-hmac", "github", "gitlab", "gitea", "bitbucketserver", "bitbucket", "azuredevops", "googlechat", "googlepubsub", "webex", "sentry", "azureeventhub", "telegram", "lark", "matrix", "opsgenie", "alertmanager", "grafana", "githubdispatch", "pagerduty", "datadog", "nats"),
						},
					},

					"username": schema.StringAttribute{
						Description:         "Username specifies the name under which events are posted.",
						MarkdownDescription: "Username specifies the name under which events are posted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(2048),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *NotificationToolkitFluxcdIoProviderV1Beta3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_notification_toolkit_fluxcd_io_provider_v1beta3_manifest")

	var model NotificationToolkitFluxcdIoProviderV1Beta3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("notification.toolkit.fluxcd.io/v1beta3")
	model.Kind = pointer.String("Provider")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
