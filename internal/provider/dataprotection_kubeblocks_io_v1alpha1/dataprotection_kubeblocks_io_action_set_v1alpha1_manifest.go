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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &DataprotectionKubeblocksIoActionSetV1Alpha1Manifest{}
)

func NewDataprotectionKubeblocksIoActionSetV1Alpha1Manifest() datasource.DataSource {
	return &DataprotectionKubeblocksIoActionSetV1Alpha1Manifest{}
}

type DataprotectionKubeblocksIoActionSetV1Alpha1Manifest struct{}

type DataprotectionKubeblocksIoActionSetV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Backup *struct {
			BackupData *struct {
				Command            *[]string `tfsdk:"command" json:"command,omitempty"`
				Image              *string   `tfsdk:"image" json:"image,omitempty"`
				OnError            *string   `tfsdk:"on_error" json:"onError,omitempty"`
				RunOnTargetPodNode *bool     `tfsdk:"run_on_target_pod_node" json:"runOnTargetPodNode,omitempty"`
				SyncProgress       *struct {
					Enabled         *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					IntervalSeconds *int64 `tfsdk:"interval_seconds" json:"intervalSeconds,omitempty"`
				} `tfsdk:"sync_progress" json:"syncProgress,omitempty"`
			} `tfsdk:"backup_data" json:"backupData,omitempty"`
			PostBackup *[]struct {
				Exec *struct {
					Command   *[]string `tfsdk:"command" json:"command,omitempty"`
					Container *string   `tfsdk:"container" json:"container,omitempty"`
					OnError   *string   `tfsdk:"on_error" json:"onError,omitempty"`
					Timeout   *string   `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				Job *struct {
					Command            *[]string `tfsdk:"command" json:"command,omitempty"`
					Image              *string   `tfsdk:"image" json:"image,omitempty"`
					OnError            *string   `tfsdk:"on_error" json:"onError,omitempty"`
					RunOnTargetPodNode *bool     `tfsdk:"run_on_target_pod_node" json:"runOnTargetPodNode,omitempty"`
				} `tfsdk:"job" json:"job,omitempty"`
			} `tfsdk:"post_backup" json:"postBackup,omitempty"`
			PreBackup *[]struct {
				Exec *struct {
					Command   *[]string `tfsdk:"command" json:"command,omitempty"`
					Container *string   `tfsdk:"container" json:"container,omitempty"`
					OnError   *string   `tfsdk:"on_error" json:"onError,omitempty"`
					Timeout   *string   `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				Job *struct {
					Command            *[]string `tfsdk:"command" json:"command,omitempty"`
					Image              *string   `tfsdk:"image" json:"image,omitempty"`
					OnError            *string   `tfsdk:"on_error" json:"onError,omitempty"`
					RunOnTargetPodNode *bool     `tfsdk:"run_on_target_pod_node" json:"runOnTargetPodNode,omitempty"`
				} `tfsdk:"job" json:"job,omitempty"`
			} `tfsdk:"pre_backup" json:"preBackup,omitempty"`
			PreDelete *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Image   *string   `tfsdk:"image" json:"image,omitempty"`
			} `tfsdk:"pre_delete" json:"preDelete,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		BackupType *string            `tfsdk:"backup_type" json:"backupType,omitempty"`
		Env        *map[string]string `tfsdk:"env" json:"env,omitempty"`
		EnvFrom    *map[string]string `tfsdk:"env_from" json:"envFrom,omitempty"`
		Restore    *struct {
			PostReady *[]struct {
				Exec *struct {
					Command   *[]string `tfsdk:"command" json:"command,omitempty"`
					Container *string   `tfsdk:"container" json:"container,omitempty"`
					OnError   *string   `tfsdk:"on_error" json:"onError,omitempty"`
					Timeout   *string   `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				Job *struct {
					Command            *[]string `tfsdk:"command" json:"command,omitempty"`
					Image              *string   `tfsdk:"image" json:"image,omitempty"`
					OnError            *string   `tfsdk:"on_error" json:"onError,omitempty"`
					RunOnTargetPodNode *bool     `tfsdk:"run_on_target_pod_node" json:"runOnTargetPodNode,omitempty"`
				} `tfsdk:"job" json:"job,omitempty"`
			} `tfsdk:"post_ready" json:"postReady,omitempty"`
			PrepareData *struct {
				Command            *[]string `tfsdk:"command" json:"command,omitempty"`
				Image              *string   `tfsdk:"image" json:"image,omitempty"`
				OnError            *string   `tfsdk:"on_error" json:"onError,omitempty"`
				RunOnTargetPodNode *bool     `tfsdk:"run_on_target_pod_node" json:"runOnTargetPodNode,omitempty"`
			} `tfsdk:"prepare_data" json:"prepareData,omitempty"`
		} `tfsdk:"restore" json:"restore,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataprotectionKubeblocksIoActionSetV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dataprotection_kubeblocks_io_action_set_v1alpha1_manifest"
}

func (r *DataprotectionKubeblocksIoActionSetV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ActionSet is the Schema for the actionsets API",
		MarkdownDescription: "ActionSet is the Schema for the actionsets API",
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
				Description:         "ActionSetSpec defines the desired state of ActionSet",
				MarkdownDescription: "ActionSetSpec defines the desired state of ActionSet",
				Attributes: map[string]schema.Attribute{
					"backup": schema.SingleNestedAttribute{
						Description:         "Specifies the backup action.",
						MarkdownDescription: "Specifies the backup action.",
						Attributes: map[string]schema.Attribute{
							"backup_data": schema.SingleNestedAttribute{
								Description:         "Represents the action to be performed for backing up data.",
								MarkdownDescription: "Represents the action to be performed for backing up data.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
										Description:         "Defines the commands to back up the volume data.",
										MarkdownDescription: "Defines the commands to back up the volume data.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "Specifies the image of the backup container.",
										MarkdownDescription: "Specifies the image of the backup container.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"on_error": schema.StringAttribute{
										Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
										MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Continue", "Fail"),
										},
									},

									"run_on_target_pod_node": schema.BoolAttribute{
										Description:         "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
										MarkdownDescription: "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sync_progress": schema.SingleNestedAttribute{
										Description:         "Determines if the backup progress should be synchronized and the interval for synchronization in seconds.",
										MarkdownDescription: "Determines if the backup progress should be synchronized and the interval for synchronization in seconds.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Determines if the backup progress should be synchronized. If set to true, a sidecar container will be instantiated to synchronize the backup progress with the Backup Custom Resource (CR) status.",
												MarkdownDescription: "Determines if the backup progress should be synchronized. If set to true, a sidecar container will be instantiated to synchronize the backup progress with the Backup Custom Resource (CR) status.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"interval_seconds": schema.Int64Attribute{
												Description:         "Defines the interval in seconds for synchronizing the backup progress.",
												MarkdownDescription: "Defines the interval in seconds for synchronizing the backup progress.",
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

							"post_backup": schema.ListNestedAttribute{
								Description:         "Represents a set of actions that should be executed after the backup process has completed.",
								MarkdownDescription: "Represents a set of actions that should be executed after the backup process has completed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "Specifies that the action should be executed using the pod's exec API within a container.",
											MarkdownDescription: "Specifies that the action should be executed using the pod's exec API within a container.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Defines the command and arguments to be executed.",
													MarkdownDescription: "Defines the command and arguments to be executed.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"container": schema.StringAttribute{
													Description:         "Specifies the container within the pod where the command should be executed. If not specified, the first container in the pod is used by default.",
													MarkdownDescription: "Specifies the container within the pod where the command should be executed. If not specified, the first container in the pod is used by default.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_error": schema.StringAttribute{
													Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
													MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Continue", "Fail"),
													},
												},

												"timeout": schema.StringAttribute{
													Description:         "Specifies the maximum duration to wait for the hook to complete before considering the execution a failure.",
													MarkdownDescription: "Specifies the maximum duration to wait for the hook to complete before considering the execution a failure.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"job": schema.SingleNestedAttribute{
											Description:         "Specifies that the action should be executed by a Kubernetes Job.",
											MarkdownDescription: "Specifies that the action should be executed by a Kubernetes Job.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Defines the commands to back up the volume data.",
													MarkdownDescription: "Defines the commands to back up the volume data.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"image": schema.StringAttribute{
													Description:         "Specifies the image of the backup container.",
													MarkdownDescription: "Specifies the image of the backup container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"on_error": schema.StringAttribute{
													Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
													MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Continue", "Fail"),
													},
												},

												"run_on_target_pod_node": schema.BoolAttribute{
													Description:         "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
													MarkdownDescription: "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pre_backup": schema.ListNestedAttribute{
								Description:         "Represents a set of actions that should be executed before the backup process begins.",
								MarkdownDescription: "Represents a set of actions that should be executed before the backup process begins.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "Specifies that the action should be executed using the pod's exec API within a container.",
											MarkdownDescription: "Specifies that the action should be executed using the pod's exec API within a container.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Defines the command and arguments to be executed.",
													MarkdownDescription: "Defines the command and arguments to be executed.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"container": schema.StringAttribute{
													Description:         "Specifies the container within the pod where the command should be executed. If not specified, the first container in the pod is used by default.",
													MarkdownDescription: "Specifies the container within the pod where the command should be executed. If not specified, the first container in the pod is used by default.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_error": schema.StringAttribute{
													Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
													MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Continue", "Fail"),
													},
												},

												"timeout": schema.StringAttribute{
													Description:         "Specifies the maximum duration to wait for the hook to complete before considering the execution a failure.",
													MarkdownDescription: "Specifies the maximum duration to wait for the hook to complete before considering the execution a failure.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"job": schema.SingleNestedAttribute{
											Description:         "Specifies that the action should be executed by a Kubernetes Job.",
											MarkdownDescription: "Specifies that the action should be executed by a Kubernetes Job.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Defines the commands to back up the volume data.",
													MarkdownDescription: "Defines the commands to back up the volume data.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"image": schema.StringAttribute{
													Description:         "Specifies the image of the backup container.",
													MarkdownDescription: "Specifies the image of the backup container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"on_error": schema.StringAttribute{
													Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
													MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Continue", "Fail"),
													},
												},

												"run_on_target_pod_node": schema.BoolAttribute{
													Description:         "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
													MarkdownDescription: "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pre_delete": schema.SingleNestedAttribute{
								Description:         "Represents a custom deletion action that can be executed before the built-in deletion action. Note: The preDelete action job will ignore the env/envFrom.",
								MarkdownDescription: "Represents a custom deletion action that can be executed before the built-in deletion action. Note: The preDelete action job will ignore the env/envFrom.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
										Description:         "Defines the commands to back up the volume data.",
										MarkdownDescription: "Defines the commands to back up the volume data.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "Specifies the image of the backup container.",
										MarkdownDescription: "Specifies the image of the backup container.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup_type": schema.StringAttribute{
						Description:         "Specifies the backup type. Supported values include:  - 'Full' for a full backup. - 'Incremental' back up data that have changed since the last backup (either full or incremental). - 'Differential' back up data that has changed since the last full backup. - 'Continuous' back up transaction logs continuously, such as MySQL binlog, PostgreSQL WAL, etc.  Continuous backup is essential for implementing Point-in-Time Recovery (PITR).",
						MarkdownDescription: "Specifies the backup type. Supported values include:  - 'Full' for a full backup. - 'Incremental' back up data that have changed since the last backup (either full or incremental). - 'Differential' back up data that has changed since the last full backup. - 'Continuous' back up transaction logs continuously, such as MySQL binlog, PostgreSQL WAL, etc.  Continuous backup is essential for implementing Point-in-Time Recovery (PITR).",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"env": schema.MapAttribute{
						Description:         "Specifies a list of environment variables to be set in the container.",
						MarkdownDescription: "Specifies a list of environment variables to be set in the container.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env_from": schema.MapAttribute{
						Description:         "Specifies a list of sources to populate environment variables in the container. The keys within a source must be a C_IDENTIFIER. Any invalid keys will be reported as an event when the container starts. If a key exists in multiple sources, the value from the last source will take precedence. Any values defined by an Env with a duplicate key will take precedence.  This field cannot be updated.",
						MarkdownDescription: "Specifies a list of sources to populate environment variables in the container. The keys within a source must be a C_IDENTIFIER. Any invalid keys will be reported as an event when the container starts. If a key exists in multiple sources, the value from the last source will take precedence. Any values defined by an Env with a duplicate key will take precedence.  This field cannot be updated.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"restore": schema.SingleNestedAttribute{
						Description:         "Specifies the restore action.",
						MarkdownDescription: "Specifies the restore action.",
						Attributes: map[string]schema.Attribute{
							"post_ready": schema.ListNestedAttribute{
								Description:         "Specifies the actions that should be executed after the data has been prepared and is ready for restoration.",
								MarkdownDescription: "Specifies the actions that should be executed after the data has been prepared and is ready for restoration.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "Specifies that the action should be executed using the pod's exec API within a container.",
											MarkdownDescription: "Specifies that the action should be executed using the pod's exec API within a container.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Defines the command and arguments to be executed.",
													MarkdownDescription: "Defines the command and arguments to be executed.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"container": schema.StringAttribute{
													Description:         "Specifies the container within the pod where the command should be executed. If not specified, the first container in the pod is used by default.",
													MarkdownDescription: "Specifies the container within the pod where the command should be executed. If not specified, the first container in the pod is used by default.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_error": schema.StringAttribute{
													Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
													MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Continue", "Fail"),
													},
												},

												"timeout": schema.StringAttribute{
													Description:         "Specifies the maximum duration to wait for the hook to complete before considering the execution a failure.",
													MarkdownDescription: "Specifies the maximum duration to wait for the hook to complete before considering the execution a failure.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"job": schema.SingleNestedAttribute{
											Description:         "Specifies that the action should be executed by a Kubernetes Job.",
											MarkdownDescription: "Specifies that the action should be executed by a Kubernetes Job.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Defines the commands to back up the volume data.",
													MarkdownDescription: "Defines the commands to back up the volume data.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"image": schema.StringAttribute{
													Description:         "Specifies the image of the backup container.",
													MarkdownDescription: "Specifies the image of the backup container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"on_error": schema.StringAttribute{
													Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
													MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Continue", "Fail"),
													},
												},

												"run_on_target_pod_node": schema.BoolAttribute{
													Description:         "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
													MarkdownDescription: "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"prepare_data": schema.SingleNestedAttribute{
								Description:         "Specifies the action required to prepare data for restoration.",
								MarkdownDescription: "Specifies the action required to prepare data for restoration.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
										Description:         "Defines the commands to back up the volume data.",
										MarkdownDescription: "Defines the commands to back up the volume data.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "Specifies the image of the backup container.",
										MarkdownDescription: "Specifies the image of the backup container.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"on_error": schema.StringAttribute{
										Description:         "Indicates how to behave if an error is encountered during the execution of this action.",
										MarkdownDescription: "Indicates how to behave if an error is encountered during the execution of this action.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Continue", "Fail"),
										},
									},

									"run_on_target_pod_node": schema.BoolAttribute{
										Description:         "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
										MarkdownDescription: "Determines whether to run the job workload on the target pod node. If the backup container needs to mount the target pod's volumes, this field should be set to true. Otherwise, the target pod's volumes will be ignored.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *DataprotectionKubeblocksIoActionSetV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dataprotection_kubeblocks_io_action_set_v1alpha1_manifest")

	var model DataprotectionKubeblocksIoActionSetV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("dataprotection.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("ActionSet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
