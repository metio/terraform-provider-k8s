/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package config_openshift_io_v1

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
	_ datasource.DataSource = &ConfigOpenshiftIoConsoleV1Manifest{}
)

func NewConfigOpenshiftIoConsoleV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoConsoleV1Manifest{}
}

type ConfigOpenshiftIoConsoleV1Manifest struct{}

type ConfigOpenshiftIoConsoleV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Authentication *struct {
			LogoutRedirect *string `tfsdk:"logout_redirect" json:"logoutRedirect,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoConsoleV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_console_v1_manifest"
}

func (r *ConfigOpenshiftIoConsoleV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Console holds cluster-wide configuration for the web console, including the logout URL, and reports the public URL of the console. The canonical name is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Console holds cluster-wide configuration for the web console, including the logout URL, and reports the public URL of the console. The canonical name is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec holds user settable values for configuration",
				MarkdownDescription: "spec holds user settable values for configuration",
				Attributes: map[string]schema.Attribute{
					"authentication": schema.SingleNestedAttribute{
						Description:         "ConsoleAuthentication defines a list of optional configuration for console authentication.",
						MarkdownDescription: "ConsoleAuthentication defines a list of optional configuration for console authentication.",
						Attributes: map[string]schema.Attribute{
							"logout_redirect": schema.StringAttribute{
								Description:         "An optional, absolute URL to redirect web browsers to after logging out of the console. If not specified, it will redirect to the default login page. This is required when using an identity provider that supports single sign-on (SSO) such as: - OpenID (Keycloak, Azure) - RequestHeader (GSSAPI, SSPI, SAML) - OAuth (GitHub, GitLab, Google) Logging out of the console will destroy the user's token. The logoutRedirect provides the user the option to perform single logout (SLO) through the identity provider to destroy their single sign-on session.",
								MarkdownDescription: "An optional, absolute URL to redirect web browsers to after logging out of the console. If not specified, it will redirect to the default login page. This is required when using an identity provider that supports single sign-on (SSO) such as: - OpenID (Keycloak, Azure) - RequestHeader (GSSAPI, SSPI, SAML) - OAuth (GitHub, GitLab, Google) Logging out of the console will destroy the user's token. The logoutRedirect provides the user the option to perform single logout (SLO) through the identity provider to destroy their single sign-on session.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^((https):\/\/?)[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|\/?))$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ConfigOpenshiftIoConsoleV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_console_v1_manifest")

	var model ConfigOpenshiftIoConsoleV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("Console")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
