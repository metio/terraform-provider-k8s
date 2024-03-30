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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ConfigOpenshiftIoAuthenticationV1Manifest{}
)

func NewConfigOpenshiftIoAuthenticationV1Manifest() datasource.DataSource {
	return &ConfigOpenshiftIoAuthenticationV1Manifest{}
}

type ConfigOpenshiftIoAuthenticationV1Manifest struct{}

type ConfigOpenshiftIoAuthenticationV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		OauthMetadata *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"oauth_metadata" json:"oauthMetadata,omitempty"`
		ServiceAccountIssuer      *string `tfsdk:"service_account_issuer" json:"serviceAccountIssuer,omitempty"`
		Type                      *string `tfsdk:"type" json:"type,omitempty"`
		WebhookTokenAuthenticator *struct {
			KubeConfig *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"kube_config" json:"kubeConfig,omitempty"`
		} `tfsdk:"webhook_token_authenticator" json:"webhookTokenAuthenticator,omitempty"`
		WebhookTokenAuthenticators *[]struct {
			KubeConfig *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"kube_config" json:"kubeConfig,omitempty"`
		} `tfsdk:"webhook_token_authenticators" json:"webhookTokenAuthenticators,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ConfigOpenshiftIoAuthenticationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_config_openshift_io_authentication_v1_manifest"
}

func (r *ConfigOpenshiftIoAuthenticationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Authentication specifies cluster-wide settings for authentication (like OAuth and webhook token authenticators). The canonical name of an instance is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "Authentication specifies cluster-wide settings for authentication (like OAuth and webhook token authenticators). The canonical name of an instance is 'cluster'.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
					"oauth_metadata": schema.SingleNestedAttribute{
						Description:         "oauthMetadata contains the discovery endpoint data for OAuth 2.0 Authorization Server Metadata for an external OAuth server. This discovery document can be viewed from its served location: oc get --raw '/.well-known/oauth-authorization-server' For further details, see the IETF Draft: https://tools.ietf.org/html/draft-ietf-oauth-discovery-04#section-2 If oauthMetadata.name is non-empty, this value has precedence over any metadata reference stored in status. The key 'oauthMetadata' is used to locate the data. If specified and the config map or expected key is not found, no metadata is served. If the specified metadata is not valid, no metadata is served. The namespace for this config map is openshift-config.",
						MarkdownDescription: "oauthMetadata contains the discovery endpoint data for OAuth 2.0 Authorization Server Metadata for an external OAuth server. This discovery document can be viewed from its served location: oc get --raw '/.well-known/oauth-authorization-server' For further details, see the IETF Draft: https://tools.ietf.org/html/draft-ietf-oauth-discovery-04#section-2 If oauthMetadata.name is non-empty, this value has precedence over any metadata reference stored in status. The key 'oauthMetadata' is used to locate the data. If specified and the config map or expected key is not found, no metadata is served. If the specified metadata is not valid, no metadata is served. The namespace for this config map is openshift-config.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the metadata.name of the referenced config map",
								MarkdownDescription: "name is the metadata.name of the referenced config map",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_issuer": schema.StringAttribute{
						Description:         "serviceAccountIssuer is the identifier of the bound service account token issuer. The default is https://kubernetes.default.svc WARNING: Updating this field will not result in immediate invalidation of all bound tokens with the previous issuer value. Instead, the tokens issued by previous service account issuer will continue to be trusted for a time period chosen by the platform (currently set to 24h). This time period is subject to change over time. This allows internal components to transition to use new service account issuer without service distruption.",
						MarkdownDescription: "serviceAccountIssuer is the identifier of the bound service account token issuer. The default is https://kubernetes.default.svc WARNING: Updating this field will not result in immediate invalidation of all bound tokens with the previous issuer value. Instead, the tokens issued by previous service account issuer will continue to be trusted for a time period chosen by the platform (currently set to 24h). This time period is subject to change over time. This allows internal components to transition to use new service account issuer without service distruption.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "type identifies the cluster managed, user facing authentication mode in use. Specifically, it manages the component that responds to login attempts. The default is IntegratedOAuth.",
						MarkdownDescription: "type identifies the cluster managed, user facing authentication mode in use. Specifically, it manages the component that responds to login attempts. The default is IntegratedOAuth.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "None", "IntegratedOAuth"),
						},
					},

					"webhook_token_authenticator": schema.SingleNestedAttribute{
						Description:         "webhookTokenAuthenticator configures a remote token reviewer. These remote authentication webhooks can be used to verify bearer tokens via the tokenreviews.authentication.k8s.io REST API. This is required to honor bearer tokens that are provisioned by an external authentication service.  Can only be set if 'Type' is set to 'None'.",
						MarkdownDescription: "webhookTokenAuthenticator configures a remote token reviewer. These remote authentication webhooks can be used to verify bearer tokens via the tokenreviews.authentication.k8s.io REST API. This is required to honor bearer tokens that are provisioned by an external authentication service.  Can only be set if 'Type' is set to 'None'.",
						Attributes: map[string]schema.Attribute{
							"kube_config": schema.SingleNestedAttribute{
								Description:         "kubeConfig references a secret that contains kube config file data which describes how to access the remote webhook service. The namespace for the referenced secret is openshift-config.  For further details, see:  https://kubernetes.io/docs/reference/access-authn-authz/authentication/#webhook-token-authentication  The key 'kubeConfig' is used to locate the data. If the secret or expected key is not found, the webhook is not honored. If the specified kube config data is not valid, the webhook is not honored.",
								MarkdownDescription: "kubeConfig references a secret that contains kube config file data which describes how to access the remote webhook service. The namespace for the referenced secret is openshift-config.  For further details, see:  https://kubernetes.io/docs/reference/access-authn-authz/authentication/#webhook-token-authentication  The key 'kubeConfig' is used to locate the data. If the secret or expected key is not found, the webhook is not honored. If the specified kube config data is not valid, the webhook is not honored.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the metadata.name of the referenced secret",
										MarkdownDescription: "name is the metadata.name of the referenced secret",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"webhook_token_authenticators": schema.ListNestedAttribute{
						Description:         "webhookTokenAuthenticators is DEPRECATED, setting it has no effect.",
						MarkdownDescription: "webhookTokenAuthenticators is DEPRECATED, setting it has no effect.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"kube_config": schema.SingleNestedAttribute{
									Description:         "kubeConfig contains kube config file data which describes how to access the remote webhook service. For further details, see: https://kubernetes.io/docs/reference/access-authn-authz/authentication/#webhook-token-authentication The key 'kubeConfig' is used to locate the data. If the secret or expected key is not found, the webhook is not honored. If the specified kube config data is not valid, the webhook is not honored. The namespace for this secret is determined by the point of use.",
									MarkdownDescription: "kubeConfig contains kube config file data which describes how to access the remote webhook service. For further details, see: https://kubernetes.io/docs/reference/access-authn-authz/authentication/#webhook-token-authentication The key 'kubeConfig' is used to locate the data. If the secret or expected key is not found, the webhook is not honored. If the specified kube config data is not valid, the webhook is not honored. The namespace for this secret is determined by the point of use.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is the metadata.name of the referenced secret",
											MarkdownDescription: "name is the metadata.name of the referenced secret",
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

func (r *ConfigOpenshiftIoAuthenticationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_config_openshift_io_authentication_v1_manifest")

	var model ConfigOpenshiftIoAuthenticationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("config.openshift.io/v1")
	model.Kind = pointer.String("Authentication")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
