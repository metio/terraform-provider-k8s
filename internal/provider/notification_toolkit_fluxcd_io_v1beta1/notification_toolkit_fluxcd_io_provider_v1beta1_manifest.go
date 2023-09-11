/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package notification_toolkit_fluxcd_io_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &NotificationToolkitFluxcdIoProviderV1Beta1Manifest{}
)

func NewNotificationToolkitFluxcdIoProviderV1Beta1Manifest() datasource.DataSource {
	return &NotificationToolkitFluxcdIoProviderV1Beta1Manifest{}
}

type NotificationToolkitFluxcdIoProviderV1Beta1Manifest struct{}

type NotificationToolkitFluxcdIoProviderV1Beta1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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

func (r *NotificationToolkitFluxcdIoProviderV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_notification_toolkit_fluxcd_io_provider_v1beta1_manifest"
}

func (r *NotificationToolkitFluxcdIoProviderV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Provider is the Schema for the providers API",
		MarkdownDescription: "Provider is the Schema for the providers API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "ProviderSpec defines the desired state of Provider",
				MarkdownDescription: "ProviderSpec defines the desired state of Provider",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "HTTP/S webhook address of this provider",
						MarkdownDescription: "HTTP/S webhook address of this provider",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(http|https)://`), ""),
						},
					},

					"cert_secret_ref": schema.SingleNestedAttribute{
						Description:         "CertSecretRef can be given the name of a secret containing a PEM-encoded CA certificate ('caFile')",
						MarkdownDescription: "CertSecretRef can be given the name of a secret containing a PEM-encoded CA certificate ('caFile')",
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
						Description:         "Alert channel for this provider",
						MarkdownDescription: "Alert channel for this provider",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proxy": schema.StringAttribute{
						Description:         "HTTP/S address of the proxy",
						MarkdownDescription: "HTTP/S address of the proxy",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(http|https)://`), ""),
						},
					},

					"secret_ref": schema.SingleNestedAttribute{
						Description:         "Secret reference containing the provider webhook URL using 'address' as data key",
						MarkdownDescription: "Secret reference containing the provider webhook URL using 'address' as data key",
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
						Description:         "This flag tells the controller to suspend subsequent events handling. Defaults to false.",
						MarkdownDescription: "This flag tells the controller to suspend subsequent events handling. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout for sending alerts to the provider.",
						MarkdownDescription: "Timeout for sending alerts to the provider.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ms|s|m))+$`), ""),
						},
					},

					"type": schema.StringAttribute{
						Description:         "Type of provider",
						MarkdownDescription: "Type of provider",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("slack", "discord", "msteams", "rocket", "generic", "generic-hmac", "github", "gitlab", "bitbucket", "azuredevops", "googlechat", "webex", "sentry", "azureeventhub", "telegram", "lark", "matrix", "opsgenie", "alertmanager", "grafana", "githubdispatch"),
						},
					},

					"username": schema.StringAttribute{
						Description:         "Bot username for this provider",
						MarkdownDescription: "Bot username for this provider",
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

func (r *NotificationToolkitFluxcdIoProviderV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_notification_toolkit_fluxcd_io_provider_v1beta1_manifest")

	var model NotificationToolkitFluxcdIoProviderV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("notification.toolkit.fluxcd.io/v1beta1")
	model.Kind = pointer.String("Provider")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
