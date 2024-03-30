/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package dex_gpu_ninja_com_v1alpha1

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
	_ datasource.DataSource = &DexGpuNinjaComDexOauth2ClientV1Alpha1Manifest{}
)

func NewDexGpuNinjaComDexOauth2ClientV1Alpha1Manifest() datasource.DataSource {
	return &DexGpuNinjaComDexOauth2ClientV1Alpha1Manifest{}
}

type DexGpuNinjaComDexOauth2ClientV1Alpha1Manifest struct{}

type DexGpuNinjaComDexOauth2ClientV1Alpha1ManifestData struct {
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
		IdentityProviderRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"identity_provider_ref" json:"identityProviderRef,omitempty"`
		LogoURL      *string   `tfsdk:"logo_url" json:"logoURL,omitempty"`
		Name         *string   `tfsdk:"name" json:"name,omitempty"`
		Public       *bool     `tfsdk:"public" json:"public,omitempty"`
		RedirectURIs *[]string `tfsdk:"redirect_ur_is" json:"redirectURIs,omitempty"`
		SecretName   *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
		TrustedPeers *[]string `tfsdk:"trusted_peers" json:"trustedPeers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DexGpuNinjaComDexOauth2ClientV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dex_gpu_ninja_com_dex_o_auth2_client_v1alpha1_manifest"
}

func (r *DexGpuNinjaComDexOauth2ClientV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DexOAuth2Client is an OAuth2 client registered with Dex.",
		MarkdownDescription: "DexOAuth2Client is an OAuth2 client registered with Dex.",
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
				Description:         "DexOAuth2ClientSpec defines the desired state of the OAuth2 client.",
				MarkdownDescription: "DexOAuth2ClientSpec defines the desired state of the OAuth2 client.",
				Attributes: map[string]schema.Attribute{
					"identity_provider_ref": schema.SingleNestedAttribute{
						Description:         "IdentityProviderRef is a reference to the identity provider which this client is associated with.",
						MarkdownDescription: "IdentityProviderRef is a reference to the identity provider which this client is associated with.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referenced DexIdentityProvider.",
								MarkdownDescription: "Name of the referenced DexIdentityProvider.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the optional namespace of the referenced DexIdentityProvider.",
								MarkdownDescription: "Namespace is the optional namespace of the referenced DexIdentityProvider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"logo_url": schema.StringAttribute{
						Description:         "LogoURL is the URL to a logo for the client.",
						MarkdownDescription: "LogoURL is the URL to a logo for the client.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is the human-readable name of the client.",
						MarkdownDescription: "Name is the human-readable name of the client.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"public": schema.BoolAttribute{
						Description:         "Public indicates that this client is a public client, such as a mobile app. Public clients must use either use a redirectURL 127.0.0.1:X or 'urn:ietf:wg:oauth:2.0:oob'.",
						MarkdownDescription: "Public indicates that this client is a public client, such as a mobile app. Public clients must use either use a redirectURL 127.0.0.1:X or 'urn:ietf:wg:oauth:2.0:oob'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redirect_ur_is": schema.ListAttribute{
						Description:         "RedirectURIs is a list of allowed redirect URLs for the client.",
						MarkdownDescription: "RedirectURIs is a list of allowed redirect URLs for the client.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName is the name of the secret that will be created to store the OAuth2 client id and client secret.",
						MarkdownDescription: "SecretName is the name of the secret that will be created to store the OAuth2 client id and client secret.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"trusted_peers": schema.ListAttribute{
						Description:         "TrustedPeers are a list of peers which can issue tokens on this client's behalf using the dynamic 'oauth2:server:client_id:(client_id)' scope. If a peer makes such a request, this client's ID will appear as the ID Token's audience.",
						MarkdownDescription: "TrustedPeers are a list of peers which can issue tokens on this client's behalf using the dynamic 'oauth2:server:client_id:(client_id)' scope. If a peer makes such a request, this client's ID will appear as the ID Token's audience.",
						ElementType:         types.StringType,
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

func (r *DexGpuNinjaComDexOauth2ClientV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dex_gpu_ninja_com_dex_o_auth2_client_v1alpha1_manifest")

	var model DexGpuNinjaComDexOauth2ClientV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("dex.gpu-ninja.com/v1alpha1")
	model.Kind = pointer.String("DexOAuth2Client")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
