/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package capabilities_3scale_net_v1alpha1

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
	_ datasource.DataSource = &Capabilities3ScaleNetTenantV1Alpha1Manifest{}
)

func NewCapabilities3ScaleNetTenantV1Alpha1Manifest() datasource.DataSource {
	return &Capabilities3ScaleNetTenantV1Alpha1Manifest{}
}

type Capabilities3ScaleNetTenantV1Alpha1Manifest struct{}

type Capabilities3ScaleNetTenantV1Alpha1ManifestData struct {
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
		Email                *string `tfsdk:"email" json:"email,omitempty"`
		FinanceSupportEmail  *string `tfsdk:"finance_support_email" json:"financeSupportEmail,omitempty"`
		FromEmail            *string `tfsdk:"from_email" json:"fromEmail,omitempty"`
		MasterCredentialsRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"master_credentials_ref" json:"masterCredentialsRef,omitempty"`
		OrganizationName       *string `tfsdk:"organization_name" json:"organizationName,omitempty"`
		PasswordCredentialsRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"password_credentials_ref" json:"passwordCredentialsRef,omitempty"`
		SiteAccessCode  *string `tfsdk:"site_access_code" json:"siteAccessCode,omitempty"`
		SupportEmail    *string `tfsdk:"support_email" json:"supportEmail,omitempty"`
		SystemMasterUrl *string `tfsdk:"system_master_url" json:"systemMasterUrl,omitempty"`
		TenantSecretRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"tenant_secret_ref" json:"tenantSecretRef,omitempty"`
		Username *string `tfsdk:"username" json:"username,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Capabilities3ScaleNetTenantV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_capabilities_3scale_net_tenant_v1alpha1_manifest"
}

func (r *Capabilities3ScaleNetTenantV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Tenant is the Schema for the tenants API",
		MarkdownDescription: "Tenant is the Schema for the tenants API",
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
				Description:         "TenantSpec defines the desired state of Tenant",
				MarkdownDescription: "TenantSpec defines the desired state of Tenant",
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"finance_support_email": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"from_email": schema.StringAttribute{
						Description:         "additional parameters, used for Update, as in master portal Api Docs",
						MarkdownDescription: "additional parameters, used for Update, as in master portal Api Docs",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"master_credentials_ref": schema.SingleNestedAttribute{
						Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"organization_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"password_credentials_ref": schema.SingleNestedAttribute{
						Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"site_access_code": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"support_email": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"system_master_url": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tenant_secret_ref": schema.SingleNestedAttribute{
						Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"username": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

func (r *Capabilities3ScaleNetTenantV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_capabilities_3scale_net_tenant_v1alpha1_manifest")

	var model Capabilities3ScaleNetTenantV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("capabilities.3scale.net/v1alpha1")
	model.Kind = pointer.String("Tenant")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
