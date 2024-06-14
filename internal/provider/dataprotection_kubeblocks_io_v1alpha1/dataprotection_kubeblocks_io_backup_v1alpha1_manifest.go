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
	_ datasource.DataSource = &DataprotectionKubeblocksIoBackupV1Alpha1Manifest{}
)

func NewDataprotectionKubeblocksIoBackupV1Alpha1Manifest() datasource.DataSource {
	return &DataprotectionKubeblocksIoBackupV1Alpha1Manifest{}
}

type DataprotectionKubeblocksIoBackupV1Alpha1Manifest struct{}

type DataprotectionKubeblocksIoBackupV1Alpha1ManifestData struct {
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
		BackupMethod     *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
		BackupPolicyName *string `tfsdk:"backup_policy_name" json:"backupPolicyName,omitempty"`
		DeletionPolicy   *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
		ParentBackupName *string `tfsdk:"parent_backup_name" json:"parentBackupName,omitempty"`
		RetentionPeriod  *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataprotectionKubeblocksIoBackupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dataprotection_kubeblocks_io_backup_v1alpha1_manifest"
}

func (r *DataprotectionKubeblocksIoBackupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Backup is the Schema for the backups API.",
		MarkdownDescription: "Backup is the Schema for the backups API.",
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
				Description:         "BackupSpec defines the desired state of Backup.",
				MarkdownDescription: "BackupSpec defines the desired state of Backup.",
				Attributes: map[string]schema.Attribute{
					"backup_method": schema.StringAttribute{
						Description:         "Specifies the backup method name that is defined in the backup policy.",
						MarkdownDescription: "Specifies the backup method name that is defined in the backup policy.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"backup_policy_name": schema.StringAttribute{
						Description:         "Specifies the backup policy to be applied for this backup.",
						MarkdownDescription: "Specifies the backup policy to be applied for this backup.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"deletion_policy": schema.StringAttribute{
						Description:         "Determines whether the backup contents stored in the backup repositoryshould be deleted when the backup custom resource(CR) is deleted.Supported values are 'Retain' and 'Delete'.- 'Retain' means that the backup content and its physical snapshot on backup repository are kept.- 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.TODO: for the retain policy, we should support in the future for only deleting  the backup CR but retaining the backup contents in backup repository.  The current implementation only prevent accidental deletion of backup data.",
						MarkdownDescription: "Determines whether the backup contents stored in the backup repositoryshould be deleted when the backup custom resource(CR) is deleted.Supported values are 'Retain' and 'Delete'.- 'Retain' means that the backup content and its physical snapshot on backup repository are kept.- 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.TODO: for the retain policy, we should support in the future for only deleting  the backup CR but retaining the backup contents in backup repository.  The current implementation only prevent accidental deletion of backup data.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parent_backup_name": schema.StringAttribute{
						Description:         "Determines the parent backup name for incremental or differential backup.",
						MarkdownDescription: "Determines the parent backup name for incremental or differential backup.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retention_period": schema.StringAttribute{
						Description:         "Determines a duration up to which the backup should be kept.Controller will remove all backups that are older than the RetentionPeriod.If not set, the backup will be kept forever.For example, RetentionPeriod of '30d' will keep only the backups of last 30 days.Sample duration format:- years: 	2y- months: 	6mo- days: 		30d- hours: 	12h- minutes: 	30mYou can also combine the above durations. For example: 30d12h30m.",
						MarkdownDescription: "Determines a duration up to which the backup should be kept.Controller will remove all backups that are older than the RetentionPeriod.If not set, the backup will be kept forever.For example, RetentionPeriod of '30d' will keep only the backups of last 30 days.Sample duration format:- years: 	2y- months: 	6mo- days: 		30d- hours: 	12h- minutes: 	30mYou can also combine the above durations. For example: 30d12h30m.",
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

func (r *DataprotectionKubeblocksIoBackupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dataprotection_kubeblocks_io_backup_v1alpha1_manifest")

	var model DataprotectionKubeblocksIoBackupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("dataprotection.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Backup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
