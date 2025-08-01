/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package v2_edp_epam_com_v1

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
	_ datasource.DataSource = &V2EdpEpamComGitServerV1Manifest{}
)

func NewV2EdpEpamComGitServerV1Manifest() datasource.DataSource {
	return &V2EdpEpamComGitServerV1Manifest{}
}

type V2EdpEpamComGitServerV1Manifest struct{}

type V2EdpEpamComGitServerV1ManifestData struct {
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
		GitHost                    *string `tfsdk:"git_host" json:"gitHost,omitempty"`
		GitProvider                *string `tfsdk:"git_provider" json:"gitProvider,omitempty"`
		GitUser                    *string `tfsdk:"git_user" json:"gitUser,omitempty"`
		HttpsPort                  *int64  `tfsdk:"https_port" json:"httpsPort,omitempty"`
		NameSshKeySecret           *string `tfsdk:"name_ssh_key_secret" json:"nameSshKeySecret,omitempty"`
		SkipWebhookSSLVerification *bool   `tfsdk:"skip_webhook_ssl_verification" json:"skipWebhookSSLVerification,omitempty"`
		SshPort                    *int64  `tfsdk:"ssh_port" json:"sshPort,omitempty"`
		WebhookUrl                 *string `tfsdk:"webhook_url" json:"webhookUrl,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *V2EdpEpamComGitServerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_v2_edp_epam_com_git_server_v1_manifest"
}

func (r *V2EdpEpamComGitServerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GitServer is the Schema for the gitservers API.",
		MarkdownDescription: "GitServer is the Schema for the gitservers API.",
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
				Description:         "GitServerSpec defines the desired state of GitServer.",
				MarkdownDescription: "GitServerSpec defines the desired state of GitServer.",
				Attributes: map[string]schema.Attribute{
					"git_host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"git_provider": schema.StringAttribute{
						Description:         "GitProvider is a git provider type. It can be gerrit, github or gitlab. Default value is gerrit.",
						MarkdownDescription: "GitProvider is a git provider type. It can be gerrit, github or gitlab. Default value is gerrit.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("gerrit", "gitlab", "github", "bitbucket"),
						},
					},

					"git_user": schema.StringAttribute{
						Description:         "GitUser is a user name for git server.",
						MarkdownDescription: "GitUser is a user name for git server.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"https_port": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name_ssh_key_secret": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"skip_webhook_ssl_verification": schema.BoolAttribute{
						Description:         "SkipWebhookSSLVerification is a flag to skip webhook tls verification.",
						MarkdownDescription: "SkipWebhookSSLVerification is a flag to skip webhook tls verification.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ssh_port": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"webhook_url": schema.StringAttribute{
						Description:         "WebhookUrl is a URL for webhook that will be created in the git provider. If not set, a new EventListener and Ingress will be created and used for webhooks.",
						MarkdownDescription: "WebhookUrl is a URL for webhook that will be created in the git provider. If not set, a new EventListener and Ingress will be created and used for webhooks.",
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

func (r *V2EdpEpamComGitServerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_v2_edp_epam_com_git_server_v1_manifest")

	var model V2EdpEpamComGitServerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("v2.edp.epam.com/v1")
	model.Kind = pointer.String("GitServer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
