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
	_ datasource.DataSource = &VeleroIoBackupRepositoryV1Manifest{}
)

func NewVeleroIoBackupRepositoryV1Manifest() datasource.DataSource {
	return &VeleroIoBackupRepositoryV1Manifest{}
}

type VeleroIoBackupRepositoryV1Manifest struct{}

type VeleroIoBackupRepositoryV1ManifestData struct {
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
		BackupStorageLocation *string `tfsdk:"backup_storage_location" json:"backupStorageLocation,omitempty"`
		MaintenanceFrequency  *string `tfsdk:"maintenance_frequency" json:"maintenanceFrequency,omitempty"`
		RepositoryType        *string `tfsdk:"repository_type" json:"repositoryType,omitempty"`
		ResticIdentifier      *string `tfsdk:"restic_identifier" json:"resticIdentifier,omitempty"`
		VolumeNamespace       *string `tfsdk:"volume_namespace" json:"volumeNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *VeleroIoBackupRepositoryV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_velero_io_backup_repository_v1_manifest"
}

func (r *VeleroIoBackupRepositoryV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "BackupRepositorySpec is the specification for a BackupRepository.",
				MarkdownDescription: "BackupRepositorySpec is the specification for a BackupRepository.",
				Attributes: map[string]schema.Attribute{
					"backup_storage_location": schema.StringAttribute{
						Description:         "BackupStorageLocation is the name of the BackupStorageLocationthat should contain this repository.",
						MarkdownDescription: "BackupStorageLocation is the name of the BackupStorageLocationthat should contain this repository.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"maintenance_frequency": schema.StringAttribute{
						Description:         "MaintenanceFrequency is how often maintenance should be run.",
						MarkdownDescription: "MaintenanceFrequency is how often maintenance should be run.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"repository_type": schema.StringAttribute{
						Description:         "RepositoryType indicates the type of the backend repository",
						MarkdownDescription: "RepositoryType indicates the type of the backend repository",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("kopia", "restic", ""),
						},
					},

					"restic_identifier": schema.StringAttribute{
						Description:         "ResticIdentifier is the full restic-compatible string for identifyingthis repository.",
						MarkdownDescription: "ResticIdentifier is the full restic-compatible string for identifyingthis repository.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"volume_namespace": schema.StringAttribute{
						Description:         "VolumeNamespace is the namespace this backup repository containspod volume backups for.",
						MarkdownDescription: "VolumeNamespace is the namespace this backup repository containspod volume backups for.",
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

func (r *VeleroIoBackupRepositoryV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_velero_io_backup_repository_v1_manifest")

	var model VeleroIoBackupRepositoryV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("velero.io/v1")
	model.Kind = pointer.String("BackupRepository")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
