/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package dataprotection_kubeblocks_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &DataprotectionKubeblocksIoBackupRepoV1Alpha1Manifest{}
)

func NewDataprotectionKubeblocksIoBackupRepoV1Alpha1Manifest() datasource.DataSource {
	return &DataprotectionKubeblocksIoBackupRepoV1Alpha1Manifest{}
}

type DataprotectionKubeblocksIoBackupRepoV1Alpha1Manifest struct{}

type DataprotectionKubeblocksIoBackupRepoV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AccessMethod *string            `tfsdk:"access_method" json:"accessMethod,omitempty"`
		Config       *map[string]string `tfsdk:"config" json:"config,omitempty"`
		Credential   *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"credential" json:"credential,omitempty"`
		PathPrefix         *string `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
		PvReclaimPolicy    *string `tfsdk:"pv_reclaim_policy" json:"pvReclaimPolicy,omitempty"`
		StorageProviderRef *string `tfsdk:"storage_provider_ref" json:"storageProviderRef,omitempty"`
		VolumeCapacity     *string `tfsdk:"volume_capacity" json:"volumeCapacity,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataprotectionKubeblocksIoBackupRepoV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dataprotection_kubeblocks_io_backup_repo_v1alpha1_manifest"
}

func (r *DataprotectionKubeblocksIoBackupRepoV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupRepo is a repository for storing backup data.",
		MarkdownDescription: "BackupRepo is a repository for storing backup data.",
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
				Description:         "BackupRepoSpec defines the desired state of 'BackupRepo'.",
				MarkdownDescription: "BackupRepoSpec defines the desired state of 'BackupRepo'.",
				Attributes: map[string]schema.Attribute{
					"access_method": schema.StringAttribute{
						Description:         "Specifies the access method of the backup repository.",
						MarkdownDescription: "Specifies the access method of the backup repository.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Mount", "Tool"),
						},
					},

					"config": schema.MapAttribute{
						Description:         "Stores the non-secret configuration parameters for the 'StorageProvider'.",
						MarkdownDescription: "Stores the non-secret configuration parameters for the 'StorageProvider'.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"credential": schema.SingleNestedAttribute{
						Description:         "References to the secret that holds the credentials for the 'StorageProvider'.",
						MarkdownDescription: "References to the secret that holds the credentials for the 'StorageProvider'.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"path_prefix": schema.StringAttribute{
						Description:         "Specifies the prefix of the path for storing backup data.",
						MarkdownDescription: "Specifies the prefix of the path for storing backup data.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9-_]+/?)*$`), ""),
						},
					},

					"pv_reclaim_policy": schema.StringAttribute{
						Description:         "Specifies reclaim policy of the PV created by this backup repository.",
						MarkdownDescription: "Specifies reclaim policy of the PV created by this backup repository.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Delete", "Retain"),
						},
					},

					"storage_provider_ref": schema.StringAttribute{
						Description:         "Specifies the name of the 'StorageProvider' used by this backup repository.",
						MarkdownDescription: "Specifies the name of the 'StorageProvider' used by this backup repository.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"volume_capacity": schema.StringAttribute{
						Description:         "Specifies the capacity of the PVC created by this backup repository.",
						MarkdownDescription: "Specifies the capacity of the PVC created by this backup repository.",
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

func (r *DataprotectionKubeblocksIoBackupRepoV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dataprotection_kubeblocks_io_backup_repo_v1alpha1_manifest")

	var model DataprotectionKubeblocksIoBackupRepoV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("dataprotection.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("BackupRepo")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
