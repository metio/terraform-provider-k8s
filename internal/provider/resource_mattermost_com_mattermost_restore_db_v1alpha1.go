/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type MattermostComMattermostRestoreDBV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*MattermostComMattermostRestoreDBV1Alpha1Resource)(nil)
)

type MattermostComMattermostRestoreDBV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type MattermostComMattermostRestoreDBV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		InitBucketURL *string `tfsdk:"init_bucket_url" yaml:"initBucketURL,omitempty"`

		MattermostClusterName *string `tfsdk:"mattermost_cluster_name" yaml:"mattermostClusterName,omitempty"`

		MattermostDBName *string `tfsdk:"mattermost_db_name" yaml:"mattermostDBName,omitempty"`

		MattermostDBPassword *string `tfsdk:"mattermost_db_password" yaml:"mattermostDBPassword,omitempty"`

		MattermostDBUser *string `tfsdk:"mattermost_db_user" yaml:"mattermostDBUser,omitempty"`

		RestoreSecret *string `tfsdk:"restore_secret" yaml:"restoreSecret,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewMattermostComMattermostRestoreDBV1Alpha1Resource() resource.Resource {
	return &MattermostComMattermostRestoreDBV1Alpha1Resource{}
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mattermost_com_mattermost_restore_db_v1alpha1"
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "MattermostRestoreDB is the Schema for the mattermostrestoredbs API",
		MarkdownDescription: "MattermostRestoreDB is the Schema for the mattermostrestoredbs API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "MattermostRestoreDBSpec defines the desired state of MattermostRestoreDB",
				MarkdownDescription: "MattermostRestoreDBSpec defines the desired state of MattermostRestoreDB",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"init_bucket_url": {
						Description:         "InitBucketURL defines where the DB backup file is located.",
						MarkdownDescription: "InitBucketURL defines where the DB backup file is located.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mattermost_cluster_name": {
						Description:         "MattermostClusterName defines the ClusterInstallation name.",
						MarkdownDescription: "MattermostClusterName defines the ClusterInstallation name.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mattermost_db_name": {
						Description:         "MattermostDBName defines the database name. Need to set if different from 'mattermost'.",
						MarkdownDescription: "MattermostDBName defines the database name. Need to set if different from 'mattermost'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mattermost_db_password": {
						Description:         "MattermostDBPassword defines the user password to access the database. Need to set if the user is different from the one created by the operator.",
						MarkdownDescription: "MattermostDBPassword defines the user password to access the database. Need to set if the user is different from the one created by the operator.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mattermost_db_user": {
						Description:         "MattermostDBUser defines the user to access the database. Need to set if the user is different from 'mmuser'.",
						MarkdownDescription: "MattermostDBUser defines the user to access the database. Need to set if the user is different from 'mmuser'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"restore_secret": {
						Description:         "RestoreSecret defines the secret that holds the credentials to MySQL Operator be able to download the DB backup file",
						MarkdownDescription: "RestoreSecret defines the secret that holds the credentials to MySQL Operator be able to download the DB backup file",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_mattermost_com_mattermost_restore_db_v1alpha1")

	var state MattermostComMattermostRestoreDBV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MattermostComMattermostRestoreDBV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("mattermost.com/v1alpha1")
	goModel.Kind = utilities.Ptr("MattermostRestoreDB")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mattermost_com_mattermost_restore_db_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_mattermost_com_mattermost_restore_db_v1alpha1")

	var state MattermostComMattermostRestoreDBV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MattermostComMattermostRestoreDBV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("mattermost.com/v1alpha1")
	goModel.Kind = utilities.Ptr("MattermostRestoreDB")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *MattermostComMattermostRestoreDBV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_mattermost_com_mattermost_restore_db_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
