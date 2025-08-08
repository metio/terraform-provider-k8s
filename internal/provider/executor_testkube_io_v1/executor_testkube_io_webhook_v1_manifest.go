/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package executor_testkube_io_v1

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
	_ datasource.DataSource = &ExecutorTestkubeIoWebhookV1Manifest{}
)

func NewExecutorTestkubeIoWebhookV1Manifest() datasource.DataSource {
	return &ExecutorTestkubeIoWebhookV1Manifest{}
}

type ExecutorTestkubeIoWebhookV1Manifest struct{}

type ExecutorTestkubeIoWebhookV1ManifestData struct {
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
		Config *struct {
			Secret *struct {
				Key       *string `tfsdk:"key" json:"key,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
		Disabled      *bool              `tfsdk:"disabled" json:"disabled,omitempty"`
		Events        *[]string          `tfsdk:"events" json:"events,omitempty"`
		Headers       *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
		OnStateChange *bool              `tfsdk:"on_state_change" json:"onStateChange,omitempty"`
		Parameters    *[]struct {
			Default     *string `tfsdk:"default" json:"default,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Example     *string `tfsdk:"example" json:"example,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			Pattern     *string `tfsdk:"pattern" json:"pattern,omitempty"`
			Required    *bool   `tfsdk:"required" json:"required,omitempty"`
		} `tfsdk:"parameters" json:"parameters,omitempty"`
		PayloadObjectField       *string `tfsdk:"payload_object_field" json:"payloadObjectField,omitempty"`
		PayloadTemplate          *string `tfsdk:"payload_template" json:"payloadTemplate,omitempty"`
		PayloadTemplateReference *string `tfsdk:"payload_template_reference" json:"payloadTemplateReference,omitempty"`
		Selector                 *string `tfsdk:"selector" json:"selector,omitempty"`
		Uri                      *string `tfsdk:"uri" json:"uri,omitempty"`
		WebhookTemplateRef       *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"webhook_template_ref" json:"webhookTemplateRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExecutorTestkubeIoWebhookV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_executor_testkube_io_webhook_v1_manifest"
}

func (r *ExecutorTestkubeIoWebhookV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Webhook is the Schema for the webhooks API",
		MarkdownDescription: "Webhook is the Schema for the webhooks API",
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
				Description:         "WebhookSpec defines the desired state of Webhook",
				MarkdownDescription: "WebhookSpec defines the desired state of Webhook",
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Description:         "webhook configuration",
						MarkdownDescription: "webhook configuration",
						Attributes: map[string]schema.Attribute{
							"secret": schema.SingleNestedAttribute{
								Description:         "private value stored in secret to use in webhook template",
								MarkdownDescription: "private value stored in secret to use in webhook template",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "object key",
										MarkdownDescription: "object key",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "object name",
										MarkdownDescription: "object name",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "object kubernetes namespace",
										MarkdownDescription: "object kubernetes namespace",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": schema.StringAttribute{
								Description:         "public value to use in webhook template",
								MarkdownDescription: "public value to use in webhook template",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disabled": schema.BoolAttribute{
						Description:         "Disabled will disable the webhook",
						MarkdownDescription: "Disabled will disable the webhook",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"events": schema.ListAttribute{
						Description:         "Events declare list if events on which webhook should be called",
						MarkdownDescription: "Events declare list if events on which webhook should be called",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"headers": schema.MapAttribute{
						Description:         "webhook headers (golang template supported)",
						MarkdownDescription: "webhook headers (golang template supported)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"on_state_change": schema.BoolAttribute{
						Description:         "OnStateChange will trigger the webhook only when the result of the current execution differs from the previous result of the same test/test suite/workflow Deprecated: field is not used",
						MarkdownDescription: "OnStateChange will trigger the webhook only when the result of the current execution differs from the previous result of the same test/test suite/workflow Deprecated: field is not used",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters": schema.ListNestedAttribute{
						Description:         "webhook parameters",
						MarkdownDescription: "webhook parameters",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"default": schema.StringAttribute{
									Description:         "default parameter value",
									MarkdownDescription: "default parameter value",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "description for the parameter",
									MarkdownDescription: "description for the parameter",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"example": schema.StringAttribute{
									Description:         "example value for the parameter",
									MarkdownDescription: "example value for the parameter",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "unique parameter name",
									MarkdownDescription: "unique parameter name",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"pattern": schema.StringAttribute{
									Description:         "regular expression to match",
									MarkdownDescription: "regular expression to match",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"required": schema.BoolAttribute{
									Description:         "whether parameter is required",
									MarkdownDescription: "whether parameter is required",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"payload_object_field": schema.StringAttribute{
						Description:         "will load the generated payload for notification inside the object",
						MarkdownDescription: "will load the generated payload for notification inside the object",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"payload_template": schema.StringAttribute{
						Description:         "golang based template for notification payload",
						MarkdownDescription: "golang based template for notification payload",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"payload_template_reference": schema.StringAttribute{
						Description:         "name of the template resource",
						MarkdownDescription: "name of the template resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.StringAttribute{
						Description:         "Labels to filter for tests and test suites",
						MarkdownDescription: "Labels to filter for tests and test suites",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uri": schema.StringAttribute{
						Description:         "Uri is address where webhook should be made (golang template supported)",
						MarkdownDescription: "Uri is address where webhook should be made (golang template supported)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"webhook_template_ref": schema.SingleNestedAttribute{
						Description:         "webhook template reference",
						MarkdownDescription: "webhook template reference",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "webhook template name to include",
								MarkdownDescription: "webhook template name to include",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ExecutorTestkubeIoWebhookV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_executor_testkube_io_webhook_v1_manifest")

	var model ExecutorTestkubeIoWebhookV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("executor.testkube.io/v1")
	model.Kind = pointer.String("Webhook")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
