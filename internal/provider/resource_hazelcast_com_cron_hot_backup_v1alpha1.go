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

type HazelcastComCronHotBackupV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*HazelcastComCronHotBackupV1Alpha1Resource)(nil)
)

type HazelcastComCronHotBackupV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type HazelcastComCronHotBackupV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		FailedHotBackupsHistoryLimit *int64 `tfsdk:"failed_hot_backups_history_limit" yaml:"failedHotBackupsHistoryLimit,omitempty"`

		HotBackupTemplate *struct {
			Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Spec *struct {
				BucketURI *string `tfsdk:"bucket_uri" yaml:"bucketURI,omitempty"`

				HazelcastResourceName *string `tfsdk:"hazelcast_resource_name" yaml:"hazelcastResourceName,omitempty"`

				Secret *string `tfsdk:"secret" yaml:"secret,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"hot_backup_template" yaml:"hotBackupTemplate,omitempty"`

		Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`

		SuccessfulHotBackupsHistoryLimit *int64 `tfsdk:"successful_hot_backups_history_limit" yaml:"successfulHotBackupsHistoryLimit,omitempty"`

		Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewHazelcastComCronHotBackupV1Alpha1Resource() resource.Resource {
	return &HazelcastComCronHotBackupV1Alpha1Resource{}
}

func (r *HazelcastComCronHotBackupV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hazelcast_com_cron_hot_backup_v1alpha1"
}

func (r *HazelcastComCronHotBackupV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CronHotBackup is the Schema for the cronhotbackups API",
		MarkdownDescription: "CronHotBackup is the Schema for the cronhotbackups API",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "CronHotBackupSpec defines the desired state of CronHotBackup",
				MarkdownDescription: "CronHotBackupSpec defines the desired state of CronHotBackup",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"failed_hot_backups_history_limit": {
						Description:         "The number of failed finished hot backups to retain. This is a pointer to distinguish between explicit zero and not specified.",
						MarkdownDescription: "The number of failed finished hot backups to retain. This is a pointer to distinguish between explicit zero and not specified.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hot_backup_template": {
						Description:         "Specifies the hot backup that will be created when executing a CronHotBackup.",
						MarkdownDescription: "Specifies the hot backup that will be created when executing a CronHotBackup.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"metadata": {
								Description:         "Standard object's metadata of the hot backups created from this template.",
								MarkdownDescription: "Standard object's metadata of the hot backups created from this template.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": {
								Description:         "Specification of the desired behavior of the hot backup.",
								MarkdownDescription: "Specification of the desired behavior of the hot backup.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"bucket_uri": {
										Description:         "URL of the bucket to download HotBackup folders.",
										MarkdownDescription: "URL of the bucket to download HotBackup folders.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"hazelcast_resource_name": {
										Description:         "HazelcastResourceName defines the name of the Hazelcast resource",
										MarkdownDescription: "HazelcastResourceName defines the name of the Hazelcast resource",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret": {
										Description:         "Name of the secret with credentials for cloud providers.",
										MarkdownDescription: "Name of the secret with credentials for cloud providers.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"schedule": {
						Description:         "Schedule contains a crontab-like expression that defines the schedule in which HotBackup will be started. If the Schedule is empty the HotBackup will start only once when applied. --- Several pre-defined schedules in place of a cron expression can be used. 	Entry                  | Description                                | Equivalent To 	-----                  | -----------                                | ------------- 	@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 * 	@monthly               | Run once a month, midnight, first of month | 0 0 1 * * 	@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0 	@daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * * 	@hourly                | Run once an hour, beginning of hour        | 0 * * * *",
						MarkdownDescription: "Schedule contains a crontab-like expression that defines the schedule in which HotBackup will be started. If the Schedule is empty the HotBackup will start only once when applied. --- Several pre-defined schedules in place of a cron expression can be used. 	Entry                  | Description                                | Equivalent To 	-----                  | -----------                                | ------------- 	@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 * 	@monthly               | Run once a month, midnight, first of month | 0 0 1 * * 	@weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0 	@daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * * 	@hourly                | Run once an hour, beginning of hour        | 0 * * * *",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"successful_hot_backups_history_limit": {
						Description:         "The number of successful finished hot backups to retain. This is a pointer to distinguish between explicit zero and not specified.",
						MarkdownDescription: "The number of successful finished hot backups to retain. This is a pointer to distinguish between explicit zero and not specified.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"suspend": {
						Description:         "When true, CronHotBackup will stop creating HotBackup CRs until it is disabled",
						MarkdownDescription: "When true, CronHotBackup will stop creating HotBackup CRs until it is disabled",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *HazelcastComCronHotBackupV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_hazelcast_com_cron_hot_backup_v1alpha1")

	var state HazelcastComCronHotBackupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HazelcastComCronHotBackupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hazelcast.com/v1alpha1")
	goModel.Kind = utilities.Ptr("CronHotBackup")

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

func (r *HazelcastComCronHotBackupV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hazelcast_com_cron_hot_backup_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *HazelcastComCronHotBackupV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_hazelcast_com_cron_hot_backup_v1alpha1")

	var state HazelcastComCronHotBackupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel HazelcastComCronHotBackupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("hazelcast.com/v1alpha1")
	goModel.Kind = utilities.Ptr("CronHotBackup")

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

func (r *HazelcastComCronHotBackupV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_hazelcast_com_cron_hot_backup_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
