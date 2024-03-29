/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package awx_ansible_com_v1beta1

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
	_ datasource.DataSource = &AwxAnsibleComAwxbackupV1Beta1Manifest{}
)

func NewAwxAnsibleComAwxbackupV1Beta1Manifest() datasource.DataSource {
	return &AwxAnsibleComAwxbackupV1Beta1Manifest{}
}

type AwxAnsibleComAwxbackupV1Beta1Manifest struct{}

type AwxAnsibleComAwxbackupV1Beta1ManifestData struct {
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
		Additional_labels            *[]string `tfsdk:"additional_labels" json:"additional_labels,omitempty"`
		Backup_pvc                   *string   `tfsdk:"backup_pvc" json:"backup_pvc,omitempty"`
		Backup_pvc_namespace         *string   `tfsdk:"backup_pvc_namespace" json:"backup_pvc_namespace,omitempty"`
		Backup_resource_requirements *struct {
			Limits *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"backup_resource_requirements" json:"backup_resource_requirements,omitempty"`
		Backup_storage_class            *string `tfsdk:"backup_storage_class" json:"backup_storage_class,omitempty"`
		Backup_storage_requirements     *string `tfsdk:"backup_storage_requirements" json:"backup_storage_requirements,omitempty"`
		Clean_backup_on_delete          *bool   `tfsdk:"clean_backup_on_delete" json:"clean_backup_on_delete,omitempty"`
		Db_management_pod_node_selector *string `tfsdk:"db_management_pod_node_selector" json:"db_management_pod_node_selector,omitempty"`
		Deployment_name                 *string `tfsdk:"deployment_name" json:"deployment_name,omitempty"`
		Image_pull_policy               *string `tfsdk:"image_pull_policy" json:"image_pull_policy,omitempty"`
		No_log                          *bool   `tfsdk:"no_log" json:"no_log,omitempty"`
		Pg_dump_suffix                  *string `tfsdk:"pg_dump_suffix" json:"pg_dump_suffix,omitempty"`
		Postgres_image                  *string `tfsdk:"postgres_image" json:"postgres_image,omitempty"`
		Postgres_image_version          *string `tfsdk:"postgres_image_version" json:"postgres_image_version,omitempty"`
		Postgres_label_selector         *string `tfsdk:"postgres_label_selector" json:"postgres_label_selector,omitempty"`
		Precreate_partition_hours       *int64  `tfsdk:"precreate_partition_hours" json:"precreate_partition_hours,omitempty"`
		Set_self_labels                 *bool   `tfsdk:"set_self_labels" json:"set_self_labels,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AwxAnsibleComAwxbackupV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_awx_ansible_com_awx_backup_v1beta1_manifest"
}

func (r *AwxAnsibleComAwxbackupV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Schema validation for the AWXBackup CRD",
		MarkdownDescription: "Schema validation for the AWXBackup CRD",
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

					"backup_pvc": schema.StringAttribute{
						Description:         "Name of the backup PVC",
						MarkdownDescription: "Name of the backup PVC",
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

					"backup_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the management pod used to create a backup",
						MarkdownDescription: "Resource requirements for the management pod used to create a backup",
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

					"backup_storage_class": schema.StringAttribute{
						Description:         "Storage class to use when creating PVC for backup",
						MarkdownDescription: "Storage class to use when creating PVC for backup",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_storage_requirements": schema.StringAttribute{
						Description:         "Storage requirements for backup PVC (may be similar to existing postgres PVC backing up from)",
						MarkdownDescription: "Storage requirements for backup PVC (may be similar to existing postgres PVC backing up from)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clean_backup_on_delete": schema.BoolAttribute{
						Description:         "Flag to indicate if backup should be deleted on PVC if AWXBackup object is deleted",
						MarkdownDescription: "Flag to indicate if backup should be deleted on PVC if AWXBackup object is deleted",
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
						Description:         "Name of the deployment to be backed up",
						MarkdownDescription: "Name of the deployment to be backed up",
						Required:            true,
						Optional:            false,
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

					"pg_dump_suffix": schema.StringAttribute{
						Description:         "Additional parameters for the pg_dump command",
						MarkdownDescription: "Additional parameters for the pg_dump command",
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

					"precreate_partition_hours": schema.Int64Attribute{
						Description:         "Number of hours worth of events table partitions to precreate before backup to avoid pg_dump locks.",
						MarkdownDescription: "Number of hours worth of events table partitions to precreate before backup to avoid pg_dump locks.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"set_self_labels": schema.BoolAttribute{
						Description:         "Maintain some of the recommended 'app.kubernetes.io/*' labels on the resource (self)",
						MarkdownDescription: "Maintain some of the recommended 'app.kubernetes.io/*' labels on the resource (self)",
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

func (r *AwxAnsibleComAwxbackupV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_awx_ansible_com_awx_backup_v1beta1_manifest")

	var model AwxAnsibleComAwxbackupV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("awx.ansible.com/v1beta1")
	model.Kind = pointer.String("AWXBackup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
