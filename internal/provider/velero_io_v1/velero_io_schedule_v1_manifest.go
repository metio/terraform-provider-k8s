/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package velero_io_v1

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
	_ datasource.DataSource = &VeleroIoScheduleV1Manifest{}
)

func NewVeleroIoScheduleV1Manifest() datasource.DataSource {
	return &VeleroIoScheduleV1Manifest{}
}

type VeleroIoScheduleV1Manifest struct{}

type VeleroIoScheduleV1ManifestData struct {
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
		Paused          *bool   `tfsdk:"paused" json:"paused,omitempty"`
		Schedule        *string `tfsdk:"schedule" json:"schedule,omitempty"`
		SkipImmediately *bool   `tfsdk:"skip_immediately" json:"skipImmediately,omitempty"`
		Template        *struct {
			CsiSnapshotTimeout               *string   `tfsdk:"csi_snapshot_timeout" json:"csiSnapshotTimeout,omitempty"`
			Datamover                        *string   `tfsdk:"datamover" json:"datamover,omitempty"`
			DefaultVolumesToFsBackup         *bool     `tfsdk:"default_volumes_to_fs_backup" json:"defaultVolumesToFsBackup,omitempty"`
			DefaultVolumesToRestic           *bool     `tfsdk:"default_volumes_to_restic" json:"defaultVolumesToRestic,omitempty"`
			ExcludedClusterScopedResources   *[]string `tfsdk:"excluded_cluster_scoped_resources" json:"excludedClusterScopedResources,omitempty"`
			ExcludedNamespaceScopedResources *[]string `tfsdk:"excluded_namespace_scoped_resources" json:"excludedNamespaceScopedResources,omitempty"`
			ExcludedNamespaces               *[]string `tfsdk:"excluded_namespaces" json:"excludedNamespaces,omitempty"`
			ExcludedResources                *[]string `tfsdk:"excluded_resources" json:"excludedResources,omitempty"`
			Hooks                            *struct {
				Resources *[]struct {
					ExcludedNamespaces *[]string `tfsdk:"excluded_namespaces" json:"excludedNamespaces,omitempty"`
					ExcludedResources  *[]string `tfsdk:"excluded_resources" json:"excludedResources,omitempty"`
					IncludedNamespaces *[]string `tfsdk:"included_namespaces" json:"includedNamespaces,omitempty"`
					IncludedResources  *[]string `tfsdk:"included_resources" json:"includedResources,omitempty"`
					LabelSelector      *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Post *[]struct {
						Exec *struct {
							Command   *[]string `tfsdk:"command" json:"command,omitempty"`
							Container *string   `tfsdk:"container" json:"container,omitempty"`
							OnError   *string   `tfsdk:"on_error" json:"onError,omitempty"`
							Timeout   *string   `tfsdk:"timeout" json:"timeout,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
					} `tfsdk:"post" json:"post,omitempty"`
					Pre *[]struct {
						Exec *struct {
							Command   *[]string `tfsdk:"command" json:"command,omitempty"`
							Container *string   `tfsdk:"container" json:"container,omitempty"`
							OnError   *string   `tfsdk:"on_error" json:"onError,omitempty"`
							Timeout   *string   `tfsdk:"timeout" json:"timeout,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
					} `tfsdk:"pre" json:"pre,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"hooks" json:"hooks,omitempty"`
			IncludeClusterResources          *bool     `tfsdk:"include_cluster_resources" json:"includeClusterResources,omitempty"`
			IncludedClusterScopedResources   *[]string `tfsdk:"included_cluster_scoped_resources" json:"includedClusterScopedResources,omitempty"`
			IncludedNamespaceScopedResources *[]string `tfsdk:"included_namespace_scoped_resources" json:"includedNamespaceScopedResources,omitempty"`
			IncludedNamespaces               *[]string `tfsdk:"included_namespaces" json:"includedNamespaces,omitempty"`
			IncludedResources                *[]string `tfsdk:"included_resources" json:"includedResources,omitempty"`
			ItemOperationTimeout             *string   `tfsdk:"item_operation_timeout" json:"itemOperationTimeout,omitempty"`
			LabelSelector                    *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Metadata *struct {
				Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			OrLabelSelectors *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"or_label_selectors" json:"orLabelSelectors,omitempty"`
			OrderedResources *map[string]string `tfsdk:"ordered_resources" json:"orderedResources,omitempty"`
			ResourcePolicy   *struct {
				ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
				Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resource_policy" json:"resourcePolicy,omitempty"`
			SnapshotMoveData *bool   `tfsdk:"snapshot_move_data" json:"snapshotMoveData,omitempty"`
			SnapshotVolumes  *bool   `tfsdk:"snapshot_volumes" json:"snapshotVolumes,omitempty"`
			StorageLocation  *string `tfsdk:"storage_location" json:"storageLocation,omitempty"`
			Ttl              *string `tfsdk:"ttl" json:"ttl,omitempty"`
			UploaderConfig   *struct {
				ParallelFilesUpload *int64 `tfsdk:"parallel_files_upload" json:"parallelFilesUpload,omitempty"`
			} `tfsdk:"uploader_config" json:"uploaderConfig,omitempty"`
			VolumeSnapshotLocations *[]string `tfsdk:"volume_snapshot_locations" json:"volumeSnapshotLocations,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		UseOwnerReferencesInBackup *bool `tfsdk:"use_owner_references_in_backup" json:"useOwnerReferencesInBackup,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *VeleroIoScheduleV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_velero_io_schedule_v1_manifest"
}

func (r *VeleroIoScheduleV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Schedule is a Velero resource that represents a pre-scheduled or periodic Backup that should be run.",
		MarkdownDescription: "Schedule is a Velero resource that represents a pre-scheduled or periodic Backup that should be run.",
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
				Description:         "ScheduleSpec defines the specification for a Velero schedule",
				MarkdownDescription: "ScheduleSpec defines the specification for a Velero schedule",
				Attributes: map[string]schema.Attribute{
					"paused": schema.BoolAttribute{
						Description:         "Paused specifies whether the schedule is paused or not",
						MarkdownDescription: "Paused specifies whether the schedule is paused or not",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"schedule": schema.StringAttribute{
						Description:         "Schedule is a Cron expression defining when to run the Backup.",
						MarkdownDescription: "Schedule is a Cron expression defining when to run the Backup.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"skip_immediately": schema.BoolAttribute{
						Description:         "SkipImmediately specifies whether to skip backup if schedule is due immediately from 'schedule.status.lastBackup' timestamp when schedule is unpaused or if schedule is new. If true, backup will be skipped immediately when schedule is unpaused if it is due based on .Status.LastBackupTimestamp or schedule is new, and will run at next schedule time. If false, backup will not be skipped immediately when schedule is unpaused, but will run at next schedule time. If empty, will follow server configuration (default: false).",
						MarkdownDescription: "SkipImmediately specifies whether to skip backup if schedule is due immediately from 'schedule.status.lastBackup' timestamp when schedule is unpaused or if schedule is new. If true, backup will be skipped immediately when schedule is unpaused if it is due based on .Status.LastBackupTimestamp or schedule is new, and will run at next schedule time. If false, backup will not be skipped immediately when schedule is unpaused, but will run at next schedule time. If empty, will follow server configuration (default: false).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template is the definition of the Backup to be run on the provided schedule",
						MarkdownDescription: "Template is the definition of the Backup to be run on the provided schedule",
						Attributes: map[string]schema.Attribute{
							"csi_snapshot_timeout": schema.StringAttribute{
								Description:         "CSISnapshotTimeout specifies the time used to wait for CSI VolumeSnapshot status turns to ReadyToUse during creation, before returning error as timeout. The default value is 10 minute.",
								MarkdownDescription: "CSISnapshotTimeout specifies the time used to wait for CSI VolumeSnapshot status turns to ReadyToUse during creation, before returning error as timeout. The default value is 10 minute.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"datamover": schema.StringAttribute{
								Description:         "DataMover specifies the data mover to be used by the backup. If DataMover is '' or 'velero', the built-in data mover will be used.",
								MarkdownDescription: "DataMover specifies the data mover to be used by the backup. If DataMover is '' or 'velero', the built-in data mover will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_volumes_to_fs_backup": schema.BoolAttribute{
								Description:         "DefaultVolumesToFsBackup specifies whether pod volume file system backup should be used for all volumes by default.",
								MarkdownDescription: "DefaultVolumesToFsBackup specifies whether pod volume file system backup should be used for all volumes by default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_volumes_to_restic": schema.BoolAttribute{
								Description:         "DefaultVolumesToRestic specifies whether restic should be used to take a backup of all pod volumes by default.  Deprecated: this field is no longer used and will be removed entirely in future. Use DefaultVolumesToFsBackup instead.",
								MarkdownDescription: "DefaultVolumesToRestic specifies whether restic should be used to take a backup of all pod volumes by default.  Deprecated: this field is no longer used and will be removed entirely in future. Use DefaultVolumesToFsBackup instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"excluded_cluster_scoped_resources": schema.ListAttribute{
								Description:         "ExcludedClusterScopedResources is a slice of cluster-scoped resource type names to exclude from the backup. If set to '*', all cluster-scoped resource types are excluded. The default value is empty.",
								MarkdownDescription: "ExcludedClusterScopedResources is a slice of cluster-scoped resource type names to exclude from the backup. If set to '*', all cluster-scoped resource types are excluded. The default value is empty.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"excluded_namespace_scoped_resources": schema.ListAttribute{
								Description:         "ExcludedNamespaceScopedResources is a slice of namespace-scoped resource type names to exclude from the backup. If set to '*', all namespace-scoped resource types are excluded. The default value is empty.",
								MarkdownDescription: "ExcludedNamespaceScopedResources is a slice of namespace-scoped resource type names to exclude from the backup. If set to '*', all namespace-scoped resource types are excluded. The default value is empty.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"excluded_namespaces": schema.ListAttribute{
								Description:         "ExcludedNamespaces contains a list of namespaces that are not included in the backup.",
								MarkdownDescription: "ExcludedNamespaces contains a list of namespaces that are not included in the backup.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"excluded_resources": schema.ListAttribute{
								Description:         "ExcludedResources is a slice of resource names that are not included in the backup.",
								MarkdownDescription: "ExcludedResources is a slice of resource names that are not included in the backup.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hooks": schema.SingleNestedAttribute{
								Description:         "Hooks represent custom behaviors that should be executed at different phases of the backup.",
								MarkdownDescription: "Hooks represent custom behaviors that should be executed at different phases of the backup.",
								Attributes: map[string]schema.Attribute{
									"resources": schema.ListNestedAttribute{
										Description:         "Resources are hooks that should be executed when backing up individual instances of a resource.",
										MarkdownDescription: "Resources are hooks that should be executed when backing up individual instances of a resource.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"excluded_namespaces": schema.ListAttribute{
													Description:         "ExcludedNamespaces specifies the namespaces to which this hook spec does not apply.",
													MarkdownDescription: "ExcludedNamespaces specifies the namespaces to which this hook spec does not apply.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"excluded_resources": schema.ListAttribute{
													Description:         "ExcludedResources specifies the resources to which this hook spec does not apply.",
													MarkdownDescription: "ExcludedResources specifies the resources to which this hook spec does not apply.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"included_namespaces": schema.ListAttribute{
													Description:         "IncludedNamespaces specifies the namespaces to which this hook spec applies. If empty, it applies to all namespaces.",
													MarkdownDescription: "IncludedNamespaces specifies the namespaces to which this hook spec applies. If empty, it applies to all namespaces.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"included_resources": schema.ListAttribute{
													Description:         "IncludedResources specifies the resources to which this hook spec applies. If empty, it applies to all resources.",
													MarkdownDescription: "IncludedResources specifies the resources to which this hook spec applies. If empty, it applies to all resources.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label_selector": schema.SingleNestedAttribute{
													Description:         "LabelSelector, if specified, filters the resources to which this hook spec applies.",
													MarkdownDescription: "LabelSelector, if specified, filters the resources to which this hook spec applies.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"name": schema.StringAttribute{
													Description:         "Name is the name of this hook.",
													MarkdownDescription: "Name is the name of this hook.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"post": schema.ListNestedAttribute{
													Description:         "PostHooks is a list of BackupResourceHooks to execute after storing the item in the backup. These are executed after all 'additional items' from item actions are processed.",
													MarkdownDescription: "PostHooks is a list of BackupResourceHooks to execute after storing the item in the backup. These are executed after all 'additional items' from item actions are processed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"exec": schema.SingleNestedAttribute{
																Description:         "Exec defines an exec hook.",
																MarkdownDescription: "Exec defines an exec hook.",
																Attributes: map[string]schema.Attribute{
																	"command": schema.ListAttribute{
																		Description:         "Command is the command and arguments to execute.",
																		MarkdownDescription: "Command is the command and arguments to execute.",
																		ElementType:         types.StringType,
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"container": schema.StringAttribute{
																		Description:         "Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.",
																		MarkdownDescription: "Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"on_error": schema.StringAttribute{
																		Description:         "OnError specifies how Velero should behave if it encounters an error executing this hook.",
																		MarkdownDescription: "OnError specifies how Velero should behave if it encounters an error executing this hook.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Continue", "Fail"),
																		},
																	},

																	"timeout": schema.StringAttribute{
																		Description:         "Timeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
																		MarkdownDescription: "Timeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"pre": schema.ListNestedAttribute{
													Description:         "PreHooks is a list of BackupResourceHooks to execute prior to storing the item in the backup. These are executed before any 'additional items' from item actions are processed.",
													MarkdownDescription: "PreHooks is a list of BackupResourceHooks to execute prior to storing the item in the backup. These are executed before any 'additional items' from item actions are processed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"exec": schema.SingleNestedAttribute{
																Description:         "Exec defines an exec hook.",
																MarkdownDescription: "Exec defines an exec hook.",
																Attributes: map[string]schema.Attribute{
																	"command": schema.ListAttribute{
																		Description:         "Command is the command and arguments to execute.",
																		MarkdownDescription: "Command is the command and arguments to execute.",
																		ElementType:         types.StringType,
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"container": schema.StringAttribute{
																		Description:         "Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.",
																		MarkdownDescription: "Container is the container in the pod where the command should be executed. If not specified, the pod's first container is used.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"on_error": schema.StringAttribute{
																		Description:         "OnError specifies how Velero should behave if it encounters an error executing this hook.",
																		MarkdownDescription: "OnError specifies how Velero should behave if it encounters an error executing this hook.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Continue", "Fail"),
																		},
																	},

																	"timeout": schema.StringAttribute{
																		Description:         "Timeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
																		MarkdownDescription: "Timeout defines the maximum amount of time Velero should wait for the hook to complete before considering the execution a failure.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"include_cluster_resources": schema.BoolAttribute{
								Description:         "IncludeClusterResources specifies whether cluster-scoped resources should be included for consideration in the backup.",
								MarkdownDescription: "IncludeClusterResources specifies whether cluster-scoped resources should be included for consideration in the backup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"included_cluster_scoped_resources": schema.ListAttribute{
								Description:         "IncludedClusterScopedResources is a slice of cluster-scoped resource type names to include in the backup. If set to '*', all cluster-scoped resource types are included. The default value is empty, which means only related cluster-scoped resources are included.",
								MarkdownDescription: "IncludedClusterScopedResources is a slice of cluster-scoped resource type names to include in the backup. If set to '*', all cluster-scoped resource types are included. The default value is empty, which means only related cluster-scoped resources are included.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"included_namespace_scoped_resources": schema.ListAttribute{
								Description:         "IncludedNamespaceScopedResources is a slice of namespace-scoped resource type names to include in the backup. The default value is '*'.",
								MarkdownDescription: "IncludedNamespaceScopedResources is a slice of namespace-scoped resource type names to include in the backup. The default value is '*'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"included_namespaces": schema.ListAttribute{
								Description:         "IncludedNamespaces is a slice of namespace names to include objects from. If empty, all namespaces are included.",
								MarkdownDescription: "IncludedNamespaces is a slice of namespace names to include objects from. If empty, all namespaces are included.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"included_resources": schema.ListAttribute{
								Description:         "IncludedResources is a slice of resource names to include in the backup. If empty, all resources are included.",
								MarkdownDescription: "IncludedResources is a slice of resource names to include in the backup. If empty, all resources are included.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"item_operation_timeout": schema.StringAttribute{
								Description:         "ItemOperationTimeout specifies the time used to wait for asynchronous BackupItemAction operations The default value is 4 hour.",
								MarkdownDescription: "ItemOperationTimeout specifies the time used to wait for asynchronous BackupItemAction operations The default value is 4 hour.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label_selector": schema.SingleNestedAttribute{
								Description:         "LabelSelector is a metav1.LabelSelector to filter with when adding individual objects to the backup. If empty or nil, all objects are included. Optional.",
								MarkdownDescription: "LabelSelector is a metav1.LabelSelector to filter with when adding individual objects to the backup. If empty or nil, all objects are included. Optional.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"metadata": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
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

							"or_label_selectors": schema.ListNestedAttribute{
								Description:         "OrLabelSelectors is list of metav1.LabelSelector to filter with when adding individual objects to the backup. If multiple provided they will be joined by the OR operator. LabelSelector as well as OrLabelSelectors cannot co-exist in backup request, only one of them can be used.",
								MarkdownDescription: "OrLabelSelectors is list of metav1.LabelSelector to filter with when adding individual objects to the backup. If multiple provided they will be joined by the OR operator. LabelSelector as well as OrLabelSelectors cannot co-exist in backup request, only one of them can be used.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"match_labels": schema.MapAttribute{
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ordered_resources": schema.MapAttribute{
								Description:         "OrderedResources specifies the backup order of resources of specific Kind. The map key is the resource name and value is a list of object names separated by commas. Each resource name has format 'namespace/objectname'.  For cluster resources, simply use 'objectname'.",
								MarkdownDescription: "OrderedResources specifies the backup order of resources of specific Kind. The map key is the resource name and value is a list of object names separated by commas. Each resource name has format 'namespace/objectname'.  For cluster resources, simply use 'objectname'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_policy": schema.SingleNestedAttribute{
								Description:         "ResourcePolicy specifies the referenced resource policies that backup should follow",
								MarkdownDescription: "ResourcePolicy specifies the referenced resource policies that backup should follow",
								Attributes: map[string]schema.Attribute{
									"api_group": schema.StringAttribute{
										Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
										MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind is the type of resource being referenced",
										MarkdownDescription: "Kind is the type of resource being referenced",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of resource being referenced",
										MarkdownDescription: "Name is the name of resource being referenced",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"snapshot_move_data": schema.BoolAttribute{
								Description:         "SnapshotMoveData specifies whether snapshot data should be moved",
								MarkdownDescription: "SnapshotMoveData specifies whether snapshot data should be moved",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"snapshot_volumes": schema.BoolAttribute{
								Description:         "SnapshotVolumes specifies whether to take snapshots of any PV's referenced in the set of objects included in the Backup.",
								MarkdownDescription: "SnapshotVolumes specifies whether to take snapshots of any PV's referenced in the set of objects included in the Backup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_location": schema.StringAttribute{
								Description:         "StorageLocation is a string containing the name of a BackupStorageLocation where the backup should be stored.",
								MarkdownDescription: "StorageLocation is a string containing the name of a BackupStorageLocation where the backup should be stored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ttl": schema.StringAttribute{
								Description:         "TTL is a time.Duration-parseable string describing how long the Backup should be retained for.",
								MarkdownDescription: "TTL is a time.Duration-parseable string describing how long the Backup should be retained for.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uploader_config": schema.SingleNestedAttribute{
								Description:         "UploaderConfig specifies the configuration for the uploader.",
								MarkdownDescription: "UploaderConfig specifies the configuration for the uploader.",
								Attributes: map[string]schema.Attribute{
									"parallel_files_upload": schema.Int64Attribute{
										Description:         "ParallelFilesUpload is the number of files parallel uploads to perform when using the uploader.",
										MarkdownDescription: "ParallelFilesUpload is the number of files parallel uploads to perform when using the uploader.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_snapshot_locations": schema.ListAttribute{
								Description:         "VolumeSnapshotLocations is a list containing names of VolumeSnapshotLocations associated with this backup.",
								MarkdownDescription: "VolumeSnapshotLocations is a list containing names of VolumeSnapshotLocations associated with this backup.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"use_owner_references_in_backup": schema.BoolAttribute{
						Description:         "UseOwnerReferencesBackup specifies whether to use OwnerReferences on backups created by this Schedule.",
						MarkdownDescription: "UseOwnerReferencesBackup specifies whether to use OwnerReferences on backups created by this Schedule.",
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

func (r *VeleroIoScheduleV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_velero_io_schedule_v1_manifest")

	var model VeleroIoScheduleV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("velero.io/v1")
	model.Kind = pointer.String("Schedule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
