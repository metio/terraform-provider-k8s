/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package velero_io_v1

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
	_ datasource.DataSource = &VeleroIoBackupStorageLocationV1Manifest{}
)

func NewVeleroIoBackupStorageLocationV1Manifest() datasource.DataSource {
	return &VeleroIoBackupStorageLocationV1Manifest{}
}

type VeleroIoBackupStorageLocationV1Manifest struct{}

type VeleroIoBackupStorageLocationV1ManifestData struct {
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
		AccessMode       *string            `tfsdk:"access_mode" json:"accessMode,omitempty"`
		BackupSyncPeriod *string            `tfsdk:"backup_sync_period" json:"backupSyncPeriod,omitempty"`
		Config           *map[string]string `tfsdk:"config" json:"config,omitempty"`
		Credential       *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"credential" json:"credential,omitempty"`
		Default       *bool `tfsdk:"default" json:"default,omitempty"`
		ObjectStorage *struct {
			Bucket *string `tfsdk:"bucket" json:"bucket,omitempty"`
			CaCert *string `tfsdk:"ca_cert" json:"caCert,omitempty"`
			Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
		} `tfsdk:"object_storage" json:"objectStorage,omitempty"`
		Provider            *string `tfsdk:"provider" json:"provider,omitempty"`
		ValidationFrequency *string `tfsdk:"validation_frequency" json:"validationFrequency,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *VeleroIoBackupStorageLocationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_velero_io_backup_storage_location_v1_manifest"
}

func (r *VeleroIoBackupStorageLocationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupStorageLocation is a location where Velero stores backup objects",
		MarkdownDescription: "BackupStorageLocation is a location where Velero stores backup objects",
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
				Description:         "BackupStorageLocationSpec defines the desired state of a Velero BackupStorageLocation",
				MarkdownDescription: "BackupStorageLocationSpec defines the desired state of a Velero BackupStorageLocation",
				Attributes: map[string]schema.Attribute{
					"access_mode": schema.StringAttribute{
						Description:         "AccessMode defines the permissions for the backup storage location.",
						MarkdownDescription: "AccessMode defines the permissions for the backup storage location.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ReadOnly", "ReadWrite"),
						},
					},

					"backup_sync_period": schema.StringAttribute{
						Description:         "BackupSyncPeriod defines how frequently to sync backup API objects from object storage. A value of 0 disables sync.",
						MarkdownDescription: "BackupSyncPeriod defines how frequently to sync backup API objects from object storage. A value of 0 disables sync.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config": schema.MapAttribute{
						Description:         "Config is for provider-specific configuration fields.",
						MarkdownDescription: "Config is for provider-specific configuration fields.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"credential": schema.SingleNestedAttribute{
						Description:         "Credential contains the credential information intended to be used with this location",
						MarkdownDescription: "Credential contains the credential information intended to be used with this location",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key of the secret to select from.  Must be a valid secret key.",
								MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"optional": schema.BoolAttribute{
								Description:         "Specify whether the Secret or its key must be defined",
								MarkdownDescription: "Specify whether the Secret or its key must be defined",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"default": schema.BoolAttribute{
						Description:         "Default indicates this location is the default backup storage location.",
						MarkdownDescription: "Default indicates this location is the default backup storage location.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"object_storage": schema.SingleNestedAttribute{
						Description:         "ObjectStorageLocation specifies the settings necessary to connect to a provider's object storage.",
						MarkdownDescription: "ObjectStorageLocation specifies the settings necessary to connect to a provider's object storage.",
						Attributes: map[string]schema.Attribute{
							"bucket": schema.StringAttribute{
								Description:         "Bucket is the bucket to use for object storage.",
								MarkdownDescription: "Bucket is the bucket to use for object storage.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"ca_cert": schema.StringAttribute{
								Description:         "CACert defines a CA bundle to use when verifying TLS connections to the provider.",
								MarkdownDescription: "CACert defines a CA bundle to use when verifying TLS connections to the provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"prefix": schema.StringAttribute{
								Description:         "Prefix is the path inside a bucket to use for Velero storage. Optional.",
								MarkdownDescription: "Prefix is the path inside a bucket to use for Velero storage. Optional.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"provider": schema.StringAttribute{
						Description:         "Provider is the provider of the backup storage.",
						MarkdownDescription: "Provider is the provider of the backup storage.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"validation_frequency": schema.StringAttribute{
						Description:         "ValidationFrequency defines how frequently to validate the corresponding object storage. A value of 0 disables validation.",
						MarkdownDescription: "ValidationFrequency defines how frequently to validate the corresponding object storage. A value of 0 disables validation.",
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

func (r *VeleroIoBackupStorageLocationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_velero_io_backup_storage_location_v1_manifest")

	var model VeleroIoBackupStorageLocationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("velero.io/v1")
	model.Kind = pointer.String("BackupStorageLocation")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
