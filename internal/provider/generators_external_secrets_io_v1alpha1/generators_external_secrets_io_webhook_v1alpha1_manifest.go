/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package generators_external_secrets_io_v1alpha1

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
	_ datasource.DataSource = &GeneratorsExternalSecretsIoWebhookV1Alpha1Manifest{}
)

func NewGeneratorsExternalSecretsIoWebhookV1Alpha1Manifest() datasource.DataSource {
	return &GeneratorsExternalSecretsIoWebhookV1Alpha1Manifest{}
}

type GeneratorsExternalSecretsIoWebhookV1Alpha1Manifest struct{}

type GeneratorsExternalSecretsIoWebhookV1Alpha1ManifestData struct {
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
		Body       *string `tfsdk:"body" json:"body,omitempty"`
		CaBundle   *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
		CaProvider *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Type      *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"ca_provider" json:"caProvider,omitempty"`
		Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
		Method  *string            `tfsdk:"method" json:"method,omitempty"`
		Result  *struct {
			JsonPath *string `tfsdk:"json_path" json:"jsonPath,omitempty"`
		} `tfsdk:"result" json:"result,omitempty"`
		Secrets *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			SecretRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"secrets" json:"secrets,omitempty"`
		Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		Url     *string `tfsdk:"url" json:"url,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GeneratorsExternalSecretsIoWebhookV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_generators_external_secrets_io_webhook_v1alpha1_manifest"
}

func (r *GeneratorsExternalSecretsIoWebhookV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Webhook connects to a third party API server to handle the secrets generation configuration parameters in spec. You can specify the server, the token, and additional body parameters. See documentation for the full API specification for requests and responses.",
		MarkdownDescription: "Webhook connects to a third party API server to handle the secrets generation configuration parameters in spec. You can specify the server, the token, and additional body parameters. See documentation for the full API specification for requests and responses.",
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
				Description:         "WebhookSpec controls the behavior of the external generator. Any body parameters should be passed to the server through the parameters field.",
				MarkdownDescription: "WebhookSpec controls the behavior of the external generator. Any body parameters should be passed to the server through the parameters field.",
				Attributes: map[string]schema.Attribute{
					"body": schema.StringAttribute{
						Description:         "Body",
						MarkdownDescription: "Body",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ca_bundle": schema.StringAttribute{
						Description:         "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
						MarkdownDescription: "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							validators.Base64Validator(),
						},
					},

					"ca_provider": schema.SingleNestedAttribute{
						Description:         "The provider for the CA bundle to use to validate webhook server certificate.",
						MarkdownDescription: "The provider for the CA bundle to use to validate webhook server certificate.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key where the CA certificate can be found in the Secret or ConfigMap.",
								MarkdownDescription: "The key where the CA certificate can be found in the Secret or ConfigMap.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "The name of the object located at the provider type.",
								MarkdownDescription: "The name of the object located at the provider type.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "The namespace the Provider type is in.",
								MarkdownDescription: "The namespace the Provider type is in.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},

							"type": schema.StringAttribute{
								Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
								MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Secret", "ConfigMap"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"headers": schema.MapAttribute{
						Description:         "Headers",
						MarkdownDescription: "Headers",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"method": schema.StringAttribute{
						Description:         "Webhook Method",
						MarkdownDescription: "Webhook Method",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"result": schema.SingleNestedAttribute{
						Description:         "Result formatting",
						MarkdownDescription: "Result formatting",
						Attributes: map[string]schema.Attribute{
							"json_path": schema.StringAttribute{
								Description:         "Json path of return value",
								MarkdownDescription: "Json path of return value",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"secrets": schema.ListNestedAttribute{
						Description:         "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
						MarkdownDescription: "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of this secret in templates",
									MarkdownDescription: "Name of this secret in templates",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "Secret ref to fill in credentials",
									MarkdownDescription: "Secret ref to fill in credentials",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key where the token is found.",
											MarkdownDescription: "The key where the token is found.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
											},
										},

										"name": schema.StringAttribute{
											Description:         "The name of the Secret resource being referred to.",
											MarkdownDescription: "The name of the Secret resource being referred to.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout",
						MarkdownDescription: "Timeout",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"url": schema.StringAttribute{
						Description:         "Webhook url to call",
						MarkdownDescription: "Webhook url to call",
						Required:            true,
						Optional:            false,
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

func (r *GeneratorsExternalSecretsIoWebhookV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_generators_external_secrets_io_webhook_v1alpha1_manifest")

	var model GeneratorsExternalSecretsIoWebhookV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("generators.external-secrets.io/v1alpha1")
	model.Kind = pointer.String("Webhook")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
