/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package awx_ansible_com_v1beta1

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
	_ datasource.DataSource = &AwxAnsibleComAwxrestoreV1Beta1Manifest{}
)

func NewAwxAnsibleComAwxrestoreV1Beta1Manifest() datasource.DataSource {
	return &AwxAnsibleComAwxrestoreV1Beta1Manifest{}
}

type AwxAnsibleComAwxrestoreV1Beta1Manifest struct{}

type AwxAnsibleComAwxrestoreV1Beta1ManifestData struct {
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
		Additional_labels               *[]string `tfsdk:"additional_labels" json:"additional_labels,omitempty"`
		Backup_dir                      *string   `tfsdk:"backup_dir" json:"backup_dir,omitempty"`
		Backup_name                     *string   `tfsdk:"backup_name" json:"backup_name,omitempty"`
		Backup_pvc                      *string   `tfsdk:"backup_pvc" json:"backup_pvc,omitempty"`
		Backup_pvc_namespace            *string   `tfsdk:"backup_pvc_namespace" json:"backup_pvc_namespace,omitempty"`
		Backup_source                   *string   `tfsdk:"backup_source" json:"backup_source,omitempty"`
		Cluster_name                    *string   `tfsdk:"cluster_name" json:"cluster_name,omitempty"`
		Db_management_pod_node_selector *string   `tfsdk:"db_management_pod_node_selector" json:"db_management_pod_node_selector,omitempty"`
		Deployment_name                 *string   `tfsdk:"deployment_name" json:"deployment_name,omitempty"`
		Force_drop_db                   *bool     `tfsdk:"force_drop_db" json:"force_drop_db,omitempty"`
		Image_pull_policy               *string   `tfsdk:"image_pull_policy" json:"image_pull_policy,omitempty"`
		No_log                          *bool     `tfsdk:"no_log" json:"no_log,omitempty"`
		Postgres_image                  *string   `tfsdk:"postgres_image" json:"postgres_image,omitempty"`
		Postgres_image_version          *string   `tfsdk:"postgres_image_version" json:"postgres_image_version,omitempty"`
		Postgres_label_selector         *string   `tfsdk:"postgres_label_selector" json:"postgres_label_selector,omitempty"`
		Restore_resource_requirements   *struct {
			Limits *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"restore_resource_requirements" json:"restore_resource_requirements,omitempty"`
		Set_self_labels *bool              `tfsdk:"set_self_labels" json:"set_self_labels,omitempty"`
		Spec_overrides  *map[string]string `tfsdk:"spec_overrides" json:"spec_overrides,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AwxAnsibleComAwxrestoreV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_awx_ansible_com_awx_restore_v1beta1_manifest"
}

func (r *AwxAnsibleComAwxrestoreV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Schema validation for the AWXRestore CRD",
		MarkdownDescription: "Schema validation for the AWXRestore CRD",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"additional_labels": schema.ListAttribute{
						Description:         "Additional labels defined on the resource, which should be propagated to child resources",
						MarkdownDescription: "Additional labels defined on the resource, which should be propagated to child resources",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_dir": schema.StringAttribute{
						Description:         "Backup directory name, set as a status found on the awxbackup object (backupDirectory)",
						MarkdownDescription: "Backup directory name, set as a status found on the awxbackup object (backupDirectory)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_name": schema.StringAttribute{
						Description:         "AWXBackup object name",
						MarkdownDescription: "AWXBackup object name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_pvc": schema.StringAttribute{
						Description:         "Name of the PVC to be restored from, set as a status found on the awxbackup object (backupClaim)",
						MarkdownDescription: "Name of the PVC to be restored from, set as a status found on the awxbackup object (backupClaim)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_pvc_namespace": schema.StringAttribute{
						Description:         "(Deprecated) Namespace the PVC is in",
						MarkdownDescription: "(Deprecated) Namespace the PVC is in",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_source": schema.StringAttribute{
						Description:         "Backup source",
						MarkdownDescription: "Backup source",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Backup CR", "PVC"),
						},
					},

					"cluster_name": schema.StringAttribute{
						Description:         "Cluster name",
						MarkdownDescription: "Cluster name",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_management_pod_node_selector": schema.StringAttribute{
						Description:         "nodeSelector for the Postgres pods to backup",
						MarkdownDescription: "nodeSelector for the Postgres pods to backup",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment_name": schema.StringAttribute{
						Description:         "Name of the restored deployment. This should be different from the original deployment name if the original deployment still exists.",
						MarkdownDescription: "Name of the restored deployment. This should be different from the original deployment name if the original deployment still exists.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"force_drop_db": schema.BoolAttribute{
						Description:         "Force drop the database before restoring. USE WITH CAUTION!",
						MarkdownDescription: "Force drop the database before restoring. USE WITH CAUTION!",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "The image pull policy",
						MarkdownDescription: "The image pull policy",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Always", "always", "Never", "never", "IfNotPresent", "ifnotpresent"),
						},
					},

					"no_log": schema.BoolAttribute{
						Description:         "Configure no_log for no_log tasks",
						MarkdownDescription: "Configure no_log for no_log tasks",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_image": schema.StringAttribute{
						Description:         "Registry path to the PostgreSQL container to use",
						MarkdownDescription: "Registry path to the PostgreSQL container to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_image_version": schema.StringAttribute{
						Description:         "PostgreSQL container image version to use",
						MarkdownDescription: "PostgreSQL container image version to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_label_selector": schema.StringAttribute{
						Description:         "Label selector used to identify postgres pod for backing up data",
						MarkdownDescription: "Label selector used to identify postgres pod for backing up data",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"restore_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the management pod that restores AWX from a backup",
						MarkdownDescription: "Resource requirements for the management pod that restores AWX from a backup",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"set_self_labels": schema.BoolAttribute{
						Description:         "Maintain some of the recommended 'app.kubernetes.io/*' labels on the resource (self)",
						MarkdownDescription: "Maintain some of the recommended 'app.kubernetes.io/*' labels on the resource (self)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"spec_overrides": schema.MapAttribute{
						Description:         "Overrides for the AWX spec",
						MarkdownDescription: "Overrides for the AWX spec",
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
		},
	}
}

func (r *AwxAnsibleComAwxrestoreV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_awx_ansible_com_awx_restore_v1beta1_manifest")

	var model AwxAnsibleComAwxrestoreV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("awx.ansible.com/v1beta1")
	model.Kind = pointer.String("AWXRestore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
