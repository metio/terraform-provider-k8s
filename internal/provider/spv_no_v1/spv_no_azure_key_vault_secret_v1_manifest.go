/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package spv_no_v1

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
	_ datasource.DataSource = &SpvNoAzureKeyVaultSecretV1Manifest{}
)

func NewSpvNoAzureKeyVaultSecretV1Manifest() datasource.DataSource {
	return &SpvNoAzureKeyVaultSecretV1Manifest{}
}

type SpvNoAzureKeyVaultSecretV1Manifest struct{}

type SpvNoAzureKeyVaultSecretV1ManifestData struct {
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
		Output *struct {
			Secret *struct {
				ChainOrder *string `tfsdk:"chain_order" json:"chainOrder,omitempty"`
				DataKey    *string `tfsdk:"data_key" json:"dataKey,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Type       *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			Transform *[]string `tfsdk:"transform" json:"transform,omitempty"`
		} `tfsdk:"output" json:"output,omitempty"`
		Vault *struct {
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Object *struct {
				ContentType *string `tfsdk:"content_type" json:"contentType,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Type        *string `tfsdk:"type" json:"type,omitempty"`
				Version     *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"object" json:"object,omitempty"`
		} `tfsdk:"vault" json:"vault,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SpvNoAzureKeyVaultSecretV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_spv_no_azure_key_vault_secret_v1_manifest"
}

func (r *SpvNoAzureKeyVaultSecretV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AzureKeyVaultSecret is a specification for a AzureKeyVaultSecret resource",
		MarkdownDescription: "AzureKeyVaultSecret is a specification for a AzureKeyVaultSecret resource",
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
				Description:         "AzureKeyVaultSecretSpec is the spec for a AzureKeyVaultSecret resource",
				MarkdownDescription: "AzureKeyVaultSecretSpec is the spec for a AzureKeyVaultSecret resource",
				Attributes: map[string]schema.Attribute{
					"output": schema.SingleNestedAttribute{
						Description:         "AzureKeyVaultOutput defines output sources, currently only support Secret",
						MarkdownDescription: "AzureKeyVaultOutput defines output sources, currently only support Secret",
						Attributes: map[string]schema.Attribute{
							"secret": schema.SingleNestedAttribute{
								Description:         "AzureKeyVaultOutputSecret has information needed to output a secret from Azure Key Vault to Kubernetes as a Secret resource",
								MarkdownDescription: "AzureKeyVaultOutputSecret has information needed to output a secret from Azure Key Vault to Kubernetes as a Secret resource",
								Attributes: map[string]schema.Attribute{
									"chain_order": schema.StringAttribute{
										Description:         "By setting chainOrder to ensureserverfirst the server certificate will be moved first in the chain",
										MarkdownDescription: "By setting chainOrder to ensureserverfirst the server certificate will be moved first in the chain",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ensureserverfirst"),
										},
									},

									"data_key": schema.StringAttribute{
										Description:         "The key to use in Kubernetes secret when setting the value from Azure Key Vault object data",
										MarkdownDescription: "The key to use in Kubernetes secret when setting the value from Azure Key Vault object data",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name for Kubernetes secret",
										MarkdownDescription: "Name for Kubernetes secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type of Secret in Kubernetes",
										MarkdownDescription: "Type of Secret in Kubernetes",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"transform": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"vault": schema.SingleNestedAttribute{
						Description:         "AzureKeyVault contains information needed to get the Azure Key Vault secret from Azure Key Vault",
						MarkdownDescription: "AzureKeyVault contains information needed to get the Azure Key Vault secret from Azure Key Vault",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the Azure Key Vault",
								MarkdownDescription: "Name of the Azure Key Vault",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"object": schema.SingleNestedAttribute{
								Description:         "AzureKeyVaultObject has information about the Azure Key Vault object to get from Azure Key Vault",
								MarkdownDescription: "AzureKeyVaultObject has information about the Azure Key Vault object to get from Azure Key Vault",
								Attributes: map[string]schema.Attribute{
									"content_type": schema.StringAttribute{
										Description:         "AzureKeyVaultObjectContentType defines what content type a secret contains, only used when type is multi-key-value-secret",
										MarkdownDescription: "AzureKeyVaultObjectContentType defines what content type a secret contains, only used when type is multi-key-value-secret",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("application/x-json", "application/x-yaml"),
										},
									},

									"name": schema.StringAttribute{
										Description:         "The object name in Azure Key Vault",
										MarkdownDescription: "The object name in Azure Key Vault",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "AzureKeyVaultObjectType defines which Object type to get from Azure Key Vault",
										MarkdownDescription: "AzureKeyVaultObjectType defines which Object type to get from Azure Key Vault",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("secret", "certificate", "key", "multi-key-value-secret"),
										},
									},

									"version": schema.StringAttribute{
										Description:         "The object version in Azure Key Vault",
										MarkdownDescription: "The object version in Azure Key Vault",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
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

func (r *SpvNoAzureKeyVaultSecretV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_spv_no_azure_key_vault_secret_v1_manifest")

	var model SpvNoAzureKeyVaultSecretV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("spv.no/v1")
	model.Kind = pointer.String("AzureKeyVaultSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
