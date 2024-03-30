/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

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
	_ datasource.DataSource = &AnywhereEksAmazonawsComOidcconfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComOidcconfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComOidcconfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComOidcconfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComOidcconfigV1Alpha1ManifestData struct {
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
		ClientId       *string `tfsdk:"client_id" json:"clientId,omitempty"`
		GroupsClaim    *string `tfsdk:"groups_claim" json:"groupsClaim,omitempty"`
		GroupsPrefix   *string `tfsdk:"groups_prefix" json:"groupsPrefix,omitempty"`
		IssuerUrl      *string `tfsdk:"issuer_url" json:"issuerUrl,omitempty"`
		RequiredClaims *[]struct {
			Claim *string `tfsdk:"claim" json:"claim,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"required_claims" json:"requiredClaims,omitempty"`
		UsernameClaim  *string `tfsdk:"username_claim" json:"usernameClaim,omitempty"`
		UsernamePrefix *string `tfsdk:"username_prefix" json:"usernamePrefix,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComOidcconfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_oidc_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComOidcconfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OIDCConfig is the Schema for the oidcconfigs API.",
		MarkdownDescription: "OIDCConfig is the Schema for the oidcconfigs API.",
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
				Description:         "OIDCConfigSpec defines the desired state of OIDCConfig.",
				MarkdownDescription: "OIDCConfigSpec defines the desired state of OIDCConfig.",
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Description:         "ClientId defines the client ID for the OpenID Connect client",
						MarkdownDescription: "ClientId defines the client ID for the OpenID Connect client",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups_claim": schema.StringAttribute{
						Description:         "GroupsClaim defines the name of a custom OpenID Connect claim for specifying user groups",
						MarkdownDescription: "GroupsClaim defines the name of a custom OpenID Connect claim for specifying user groups",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups_prefix": schema.StringAttribute{
						Description:         "GroupsPrefix defines a string to be prefixed to all groups to prevent conflicts with other authentication strategies",
						MarkdownDescription: "GroupsPrefix defines a string to be prefixed to all groups to prevent conflicts with other authentication strategies",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"issuer_url": schema.StringAttribute{
						Description:         "IssuerUrl defines the URL of the OpenID issuer, only HTTPS scheme will be accepted",
						MarkdownDescription: "IssuerUrl defines the URL of the OpenID issuer, only HTTPS scheme will be accepted",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"required_claims": schema.ListNestedAttribute{
						Description:         "RequiredClaims defines a key=value pair that describes a required claim in the ID Token",
						MarkdownDescription: "RequiredClaims defines a key=value pair that describes a required claim in the ID Token",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"claim": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"username_claim": schema.StringAttribute{
						Description:         "UsernameClaim defines the OpenID claim to use as the user name. Note that claims other than the default ('sub') is not guaranteed to be unique and immutable",
						MarkdownDescription: "UsernameClaim defines the OpenID claim to use as the user name. Note that claims other than the default ('sub') is not guaranteed to be unique and immutable",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username_prefix": schema.StringAttribute{
						Description:         "UsernamePrefix defines a string to prefixed to all usernames. If not provided, username claims other than 'email' are prefixed by the issuer URL to avoid clashes. To skip any prefixing, provide the value '-'.",
						MarkdownDescription: "UsernamePrefix defines a string to prefixed to all usernames. If not provided, username claims other than 'email' are prefixed by the issuer URL to avoid clashes. To skip any prefixing, provide the value '-'.",
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

func (r *AnywhereEksAmazonawsComOidcconfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_oidc_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComOidcconfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("OIDCConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
