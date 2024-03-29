/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package console_openshift_io_v1

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
	_ datasource.DataSource = &ConsoleOpenshiftIoConsoleNotificationV1Manifest{}
)

func NewConsoleOpenshiftIoConsoleNotificationV1Manifest() datasource.DataSource {
	return &ConsoleOpenshiftIoConsoleNotificationV1Manifest{}
}

type ConsoleOpenshiftIoConsoleNotificationV1Manifest struct{}

type ConsoleOpenshiftIoConsoleNotificationV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BackgroundColor *string `tfsdk:"background_color" json:"backgroundColor,omitempty"`
		Color           *string `tfsdk:"color" json:"color,omitempty"`
		Link            *struct {
			Href *string `tfsdk:"href" json:"href,omitempty"`
			Text *string `tfsdk:"text" json:"text,omitempty"`
		} `tfsdk:"link" json:"link,omitempty"`
		Location *string `tfsdk:"location" json:"location,omitempty"`
		Text     *string `tfsdk:"text" json:"text,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConsoleOpenshiftIoConsoleNotificationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_console_openshift_io_console_notification_v1_manifest"
}

func (r *ConsoleOpenshiftIoConsoleNotificationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConsoleNotification is the extension for configuring openshift web console notifications.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ConsoleNotification is the extension for configuring openshift web console notifications.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "ConsoleNotificationSpec is the desired console notification configuration.",
				MarkdownDescription: "ConsoleNotificationSpec is the desired console notification configuration.",
				Attributes: map[string]schema.Attribute{
					"background_color": schema.StringAttribute{
						Description:         "backgroundColor is the color of the background for the notification as CSS data type color.",
						MarkdownDescription: "backgroundColor is the color of the background for the notification as CSS data type color.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"color": schema.StringAttribute{
						Description:         "color is the color of the text for the notification as CSS data type color.",
						MarkdownDescription: "color is the color of the text for the notification as CSS data type color.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"link": schema.SingleNestedAttribute{
						Description:         "link is an object that holds notification link details.",
						MarkdownDescription: "link is an object that holds notification link details.",
						Attributes: map[string]schema.Attribute{
							"href": schema.StringAttribute{
								Description:         "href is the absolute secure URL for the link (must use https)",
								MarkdownDescription: "href is the absolute secure URL for the link (must use https)",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^https://`), ""),
								},
							},

							"text": schema.StringAttribute{
								Description:         "text is the display text for the link",
								MarkdownDescription: "text is the display text for the link",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"location": schema.StringAttribute{
						Description:         "location is the location of the notification in the console. Valid values are: 'BannerTop', 'BannerBottom', 'BannerTopBottom'.",
						MarkdownDescription: "location is the location of the notification in the console. Valid values are: 'BannerTop', 'BannerBottom', 'BannerTopBottom'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(BannerTop|BannerBottom|BannerTopBottom)$`), ""),
						},
					},

					"text": schema.StringAttribute{
						Description:         "text is the visible text of the notification.",
						MarkdownDescription: "text is the visible text of the notification.",
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

func (r *ConsoleOpenshiftIoConsoleNotificationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_console_openshift_io_console_notification_v1_manifest")

	var model ConsoleOpenshiftIoConsoleNotificationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("console.openshift.io/v1")
	model.Kind = pointer.String("ConsoleNotification")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
