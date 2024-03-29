/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package repo_manager_pulpproject_org_v1beta2

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
	_ datasource.DataSource = &RepoManagerPulpprojectOrgPulpRestoreV1Beta2Manifest{}
)

func NewRepoManagerPulpprojectOrgPulpRestoreV1Beta2Manifest() datasource.DataSource {
	return &RepoManagerPulpprojectOrgPulpRestoreV1Beta2Manifest{}
}

type RepoManagerPulpprojectOrgPulpRestoreV1Beta2Manifest struct{}

type RepoManagerPulpprojectOrgPulpRestoreV1Beta2ManifestData struct {
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
		Backup_dir      *string `tfsdk:"backup_dir" json:"backup_dir,omitempty"`
		Backup_name     *string `tfsdk:"backup_name" json:"backup_name,omitempty"`
		Backup_pvc      *string `tfsdk:"backup_pvc" json:"backup_pvc,omitempty"`
		Deployment_name *string `tfsdk:"deployment_name" json:"deployment_name,omitempty"`
		Deployment_type *string `tfsdk:"deployment_type" json:"deployment_type,omitempty"`
		Keep_replicas   *bool   `tfsdk:"keep_replicas" json:"keep_replicas,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RepoManagerPulpprojectOrgPulpRestoreV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest"
}

func (r *RepoManagerPulpprojectOrgPulpRestoreV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PulpRestore is the Schema for the pulprestores API",
		MarkdownDescription: "PulpRestore is the Schema for the pulprestores API",
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
				Description:         "PulpRestoreSpec defines the desired state of PulpRestore",
				MarkdownDescription: "PulpRestoreSpec defines the desired state of PulpRestore",
				Attributes: map[string]schema.Attribute{
					"backup_dir": schema.StringAttribute{
						Description:         "Backup directory name, set as a status found on the backup object (backupDirectory)",
						MarkdownDescription: "Backup directory name, set as a status found on the backup object (backupDirectory)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_name": schema.StringAttribute{
						Description:         "Name of the backup custom resource",
						MarkdownDescription: "Name of the backup custom resource",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"backup_pvc": schema.StringAttribute{
						Description:         "Name of the PVC to be restored from, set as a status found on the backup object (backupClaim)",
						MarkdownDescription: "Name of the PVC to be restored from, set as a status found on the backup object (backupClaim)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment_name": schema.StringAttribute{
						Description:         "Name of the deployment to be restored to",
						MarkdownDescription: "Name of the deployment to be restored to",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment_type": schema.StringAttribute{
						Description:         "Name of the deployment type. Can be one of {galaxy,pulp}.",
						MarkdownDescription: "Name of the deployment type. Can be one of {galaxy,pulp}.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("galaxy", "pulp"),
						},
					},

					"keep_replicas": schema.BoolAttribute{
						Description:         "KeepBackupReplicasCount allows to define if the restore controller should restore the components with the same number of replicas from backup or restore only a single replica each.",
						MarkdownDescription: "KeepBackupReplicasCount allows to define if the restore controller should restore the components with the same number of replicas from backup or restore only a single replica each.",
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

func (r *RepoManagerPulpprojectOrgPulpRestoreV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_repo_manager_pulpproject_org_pulp_restore_v1beta2_manifest")

	var model RepoManagerPulpprojectOrgPulpRestoreV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("repo-manager.pulpproject.org/v1beta2")
	model.Kind = pointer.String("PulpRestore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
