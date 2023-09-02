/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mattermost_com_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MattermostComMattermostRestoreDBV1Alpha1Manifest{}
)

func NewMattermostComMattermostRestoreDBV1Alpha1Manifest() datasource.DataSource {
	return &MattermostComMattermostRestoreDBV1Alpha1Manifest{}
}

type MattermostComMattermostRestoreDBV1Alpha1Manifest struct{}

type MattermostComMattermostRestoreDBV1Alpha1ManifestData struct {
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
		InitBucketURL         *string `tfsdk:"init_bucket_url" json:"initBucketURL,omitempty"`
		MattermostClusterName *string `tfsdk:"mattermost_cluster_name" json:"mattermostClusterName,omitempty"`
		MattermostDBName      *string `tfsdk:"mattermost_db_name" json:"mattermostDBName,omitempty"`
		MattermostDBPassword  *string `tfsdk:"mattermost_db_password" json:"mattermostDBPassword,omitempty"`
		MattermostDBUser      *string `tfsdk:"mattermost_db_user" json:"mattermostDBUser,omitempty"`
		RestoreSecret         *string `tfsdk:"restore_secret" json:"restoreSecret,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_mattermost_com_mattermost_restore_db_v1alpha1_manifest"
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MattermostRestoreDB is the Schema for the mattermostrestoredbs API",
		MarkdownDescription: "MattermostRestoreDB is the Schema for the mattermostrestoredbs API",
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
				Description:         "MattermostRestoreDBSpec defines the desired state of MattermostRestoreDB",
				MarkdownDescription: "MattermostRestoreDBSpec defines the desired state of MattermostRestoreDB",
				Attributes: map[string]schema.Attribute{
					"init_bucket_url": schema.StringAttribute{
						Description:         "InitBucketURL defines where the DB backup file is located.",
						MarkdownDescription: "InitBucketURL defines where the DB backup file is located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mattermost_cluster_name": schema.StringAttribute{
						Description:         "MattermostClusterName defines the ClusterInstallation name.",
						MarkdownDescription: "MattermostClusterName defines the ClusterInstallation name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mattermost_db_name": schema.StringAttribute{
						Description:         "MattermostDBName defines the database name. Need to set if different from 'mattermost'.",
						MarkdownDescription: "MattermostDBName defines the database name. Need to set if different from 'mattermost'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mattermost_db_password": schema.StringAttribute{
						Description:         "MattermostDBPassword defines the user password to access the database. Need to set if the user is different from the one created by the operator.",
						MarkdownDescription: "MattermostDBPassword defines the user password to access the database. Need to set if the user is different from the one created by the operator.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mattermost_db_user": schema.StringAttribute{
						Description:         "MattermostDBUser defines the user to access the database. Need to set if the user is different from 'mmuser'.",
						MarkdownDescription: "MattermostDBUser defines the user to access the database. Need to set if the user is different from 'mmuser'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"restore_secret": schema.StringAttribute{
						Description:         "RestoreSecret defines the secret that holds the credentials to MySQL Operator be able to download the DB backup file",
						MarkdownDescription: "RestoreSecret defines the secret that holds the credentials to MySQL Operator be able to download the DB backup file",
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

func (r *MattermostComMattermostRestoreDBV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mattermost_com_mattermost_restore_db_v1alpha1_manifest")

	var model MattermostComMattermostRestoreDBV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("mattermost.com/v1alpha1")
	model.Kind = pointer.String("MattermostRestoreDB")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
