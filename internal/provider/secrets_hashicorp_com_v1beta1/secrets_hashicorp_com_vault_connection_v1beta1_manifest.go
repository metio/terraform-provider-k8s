/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_hashicorp_com_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &SecretsHashicorpComVaultConnectionV1Beta1Manifest{}
)

func NewSecretsHashicorpComVaultConnectionV1Beta1Manifest() datasource.DataSource {
	return &SecretsHashicorpComVaultConnectionV1Beta1Manifest{}
}

type SecretsHashicorpComVaultConnectionV1Beta1Manifest struct{}

type SecretsHashicorpComVaultConnectionV1Beta1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Address         *string            `tfsdk:"address" json:"address,omitempty"`
		CaCertSecretRef *string            `tfsdk:"ca_cert_secret_ref" json:"caCertSecretRef,omitempty"`
		Headers         *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
		SkipTLSVerify   *bool              `tfsdk:"skip_tls_verify" json:"skipTLSVerify,omitempty"`
		TlsServerName   *string            `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsHashicorpComVaultConnectionV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_hashicorp_com_vault_connection_v1beta1_manifest"
}

func (r *SecretsHashicorpComVaultConnectionV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VaultConnection is the Schema for the vaultconnections API",
		MarkdownDescription: "VaultConnection is the Schema for the vaultconnections API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "VaultConnectionSpec defines the desired state of VaultConnection",
				MarkdownDescription: "VaultConnectionSpec defines the desired state of VaultConnection",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "Address of the Vault server",
						MarkdownDescription: "Address of the Vault server",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ca_cert_secret_ref": schema.StringAttribute{
						Description:         "CACertSecretRef is the name of a Kubernetes secret containing the trusted PEM encoded CA certificate chain as 'ca.crt'.",
						MarkdownDescription: "CACertSecretRef is the name of a Kubernetes secret containing the trusted PEM encoded CA certificate chain as 'ca.crt'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"headers": schema.MapAttribute{
						Description:         "Headers to be included in all Vault requests.",
						MarkdownDescription: "Headers to be included in all Vault requests.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skip_tls_verify": schema.BoolAttribute{
						Description:         "SkipTLSVerify for TLS connections.",
						MarkdownDescription: "SkipTLSVerify for TLS connections.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tls_server_name": schema.StringAttribute{
						Description:         "TLSServerName to use as the SNI host for TLS connections.",
						MarkdownDescription: "TLSServerName to use as the SNI host for TLS connections.",
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

func (r *SecretsHashicorpComVaultConnectionV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_hashicorp_com_vault_connection_v1beta1_manifest")

	var model SecretsHashicorpComVaultConnectionV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("secrets.hashicorp.com/v1beta1")
	model.Kind = pointer.String("VaultConnection")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
