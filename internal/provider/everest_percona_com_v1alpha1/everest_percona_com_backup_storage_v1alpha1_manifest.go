/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package everest_percona_com_v1alpha1

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
	_ datasource.DataSource = &EverestPerconaComBackupStorageV1Alpha1Manifest{}
)

func NewEverestPerconaComBackupStorageV1Alpha1Manifest() datasource.DataSource {
	return &EverestPerconaComBackupStorageV1Alpha1Manifest{}
}

type EverestPerconaComBackupStorageV1Alpha1Manifest struct{}

type EverestPerconaComBackupStorageV1Alpha1ManifestData struct {
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
		AllowedNamespaces     *[]string `tfsdk:"allowed_namespaces" json:"allowedNamespaces,omitempty"`
		Bucket                *string   `tfsdk:"bucket" json:"bucket,omitempty"`
		CredentialsSecretName *string   `tfsdk:"credentials_secret_name" json:"credentialsSecretName,omitempty"`
		Description           *string   `tfsdk:"description" json:"description,omitempty"`
		EndpointURL           *string   `tfsdk:"endpoint_url" json:"endpointURL,omitempty"`
		ForcePathStyle        *bool     `tfsdk:"force_path_style" json:"forcePathStyle,omitempty"`
		Region                *string   `tfsdk:"region" json:"region,omitempty"`
		Type                  *string   `tfsdk:"type" json:"type,omitempty"`
		VerifyTLS             *bool     `tfsdk:"verify_tls" json:"verifyTLS,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EverestPerconaComBackupStorageV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_everest_percona_com_backup_storage_v1alpha1_manifest"
}

func (r *EverestPerconaComBackupStorageV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupStorage is the Schema for the backupstorages API.",
		MarkdownDescription: "BackupStorage is the Schema for the backupstorages API.",
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
				Description:         "BackupStorageSpec defines the desired state of BackupStorage.",
				MarkdownDescription: "BackupStorageSpec defines the desired state of BackupStorage.",
				Attributes: map[string]schema.Attribute{
					"allowed_namespaces": schema.ListAttribute{
						Description:         "AllowedNamespaces is the list of namespaces where the operator will copy secrets provided in the CredentialsSecretsName. Deprecated: BackupStorages are now used only in the namespaces where they are created.",
						MarkdownDescription: "AllowedNamespaces is the list of namespaces where the operator will copy secrets provided in the CredentialsSecretsName. Deprecated: BackupStorages are now used only in the namespaces where they are created.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bucket": schema.StringAttribute{
						Description:         "Bucket is a name of bucket.",
						MarkdownDescription: "Bucket is a name of bucket.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"credentials_secret_name": schema.StringAttribute{
						Description:         "CredentialsSecretName is the name of the secret with credentials.",
						MarkdownDescription: "CredentialsSecretName is the name of the secret with credentials.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "Description stores description of a backup storage.",
						MarkdownDescription: "Description stores description of a backup storage.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoint_url": schema.StringAttribute{
						Description:         "EndpointURL is an endpoint URL of backup storage.",
						MarkdownDescription: "EndpointURL is an endpoint URL of backup storage.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"force_path_style": schema.BoolAttribute{
						Description:         "ForcePathStyle is set to use path-style URLs. If unspecified, the default value is false.",
						MarkdownDescription: "ForcePathStyle is set to use path-style URLs. If unspecified, the default value is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"region": schema.StringAttribute{
						Description:         "Region is a region where the bucket is located.",
						MarkdownDescription: "Region is a region where the bucket is located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is a type of backup storage.",
						MarkdownDescription: "Type is a type of backup storage.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("s3", "azure"),
						},
					},

					"verify_tls": schema.BoolAttribute{
						Description:         "VerifyTLS is set to ensure TLS/SSL verification. If unspecified, the default value is true.",
						MarkdownDescription: "VerifyTLS is set to ensure TLS/SSL verification. If unspecified, the default value is true.",
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

func (r *EverestPerconaComBackupStorageV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_everest_percona_com_backup_storage_v1alpha1_manifest")

	var model EverestPerconaComBackupStorageV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("everest.percona.com/v1alpha1")
	model.Kind = pointer.String("BackupStorage")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
