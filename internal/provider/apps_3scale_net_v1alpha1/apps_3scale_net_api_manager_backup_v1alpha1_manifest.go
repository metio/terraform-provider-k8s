/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_3scale_net_v1alpha1

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
	_ datasource.DataSource = &Apps3ScaleNetApimanagerBackupV1Alpha1Manifest{}
)

func NewApps3ScaleNetApimanagerBackupV1Alpha1Manifest() datasource.DataSource {
	return &Apps3ScaleNetApimanagerBackupV1Alpha1Manifest{}
}

type Apps3ScaleNetApimanagerBackupV1Alpha1Manifest struct{}

type Apps3ScaleNetApimanagerBackupV1Alpha1ManifestData struct {
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
		BackupDestination *struct {
			PersistentVolumeClaim *struct {
				Resources *struct {
					Requests *string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				VolumeName   *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
		} `tfsdk:"backup_destination" json:"backupDestination,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Apps3ScaleNetApimanagerBackupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_3scale_net_api_manager_backup_v1alpha1_manifest"
}

func (r *Apps3ScaleNetApimanagerBackupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "APIManagerBackup represents an APIManager backup",
		MarkdownDescription: "APIManagerBackup represents an APIManager backup",
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
				Description:         "APIManagerBackupSpec defines the desired state of APIManagerBackup",
				MarkdownDescription: "APIManagerBackupSpec defines the desired state of APIManagerBackup",
				Attributes: map[string]schema.Attribute{
					"backup_destination": schema.SingleNestedAttribute{
						Description:         "Backup data destination configuration",
						MarkdownDescription: "Backup data destination configuration",
						Attributes: map[string]schema.Attribute{
							"persistent_volume_claim": schema.SingleNestedAttribute{
								Description:         "PersistentVolumeClaim as backup data destination configuration",
								MarkdownDescription: "PersistentVolumeClaim as backup data destination configuration",
								Attributes: map[string]schema.Attribute{
									"resources": schema.SingleNestedAttribute{
										Description:         "Resources configuration for the backup data PersistentVolumeClaim. Ignored when VolumeName field is set",
										MarkdownDescription: "Resources configuration for the backup data PersistentVolumeClaim. Ignored when VolumeName field is set",
										Attributes: map[string]schema.Attribute{
											"requests": schema.StringAttribute{
												Description:         "Storage Resource requests to be used on the PersistentVolumeClaim. To learn more about resource requests see: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Storage Resource requests to be used on the PersistentVolumeClaim. To learn more about resource requests see: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_class": schema.StringAttribute{
										Description:         "Storage class to be used by the PersistentVolumeClaim. Ignored when VolumeName field is set",
										MarkdownDescription: "Storage class to be used by the PersistentVolumeClaim. Ignored when VolumeName field is set",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_name": schema.StringAttribute{
										Description:         "Name of an existing PersistentVolume to be bound to the backup data PersistentVolumeClaim",
										MarkdownDescription: "Name of an existing PersistentVolume to be bound to the backup data PersistentVolumeClaim",
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
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Apps3ScaleNetApimanagerBackupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_3scale_net_api_manager_backup_v1alpha1_manifest")

	var model Apps3ScaleNetApimanagerBackupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("apps.3scale.net/v1alpha1")
	model.Kind = pointer.String("APIManagerBackup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
