/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package dataprotection_kubeblocks_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest{}
)

func NewDataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest() datasource.DataSource {
	return &DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest{}
}

type DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest struct{}

type DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BackoffLimit  *int64 `tfsdk:"backoff_limit" json:"backoffLimit,omitempty"`
		BackupMethods *[]struct {
			ActionSetName    *string `tfsdk:"action_set_name" json:"actionSetName,omitempty"`
			CompatibleMethod *string `tfsdk:"compatible_method" json:"compatibleMethod,omitempty"`
			Env              *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					VersionMapping *[]struct {
						MappedValue     *string   `tfsdk:"mapped_value" json:"mappedValue,omitempty"`
						ServiceVersions *[]string `tfsdk:"service_versions" json:"serviceVersions,omitempty"`
					} `tfsdk:"version_mapping" json:"versionMapping,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"env" json:"env,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			RuntimeSettings *struct {
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"runtime_settings" json:"runtimeSettings,omitempty"`
			SnapshotVolumes *bool `tfsdk:"snapshot_volumes" json:"snapshotVolumes,omitempty"`
			Target          *struct {
				Account       *string `tfsdk:"account" json:"account,omitempty"`
				ContainerPort *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					PortName      *string `tfsdk:"port_name" json:"portName,omitempty"`
				} `tfsdk:"container_port" json:"containerPort,omitempty"`
				FallbackRole          *string `tfsdk:"fallback_role" json:"fallbackRole,omitempty"`
				Role                  *string `tfsdk:"role" json:"role,omitempty"`
				Strategy              *string `tfsdk:"strategy" json:"strategy,omitempty"`
				UseParentSelectedPods *bool   `tfsdk:"use_parent_selected_pods" json:"useParentSelectedPods,omitempty"`
			} `tfsdk:"target" json:"target,omitempty"`
			TargetVolumes *struct {
				VolumeMounts *[]struct {
					MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
					Name             *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
				Volumes *[]string `tfsdk:"volumes" json:"volumes,omitempty"`
			} `tfsdk:"target_volumes" json:"targetVolumes,omitempty"`
		} `tfsdk:"backup_methods" json:"backupMethods,omitempty"`
		CompDefs        *[]string `tfsdk:"comp_defs" json:"compDefs,omitempty"`
		RetentionPolicy *string   `tfsdk:"retention_policy" json:"retentionPolicy,omitempty"`
		Schedules       *[]struct {
			BackupMethod   *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
			CronExpression *string `tfsdk:"cron_expression" json:"cronExpression,omitempty"`
			Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Parameters     *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
			RetentionPeriod *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		} `tfsdk:"schedules" json:"schedules,omitempty"`
		ServiceKind *string `tfsdk:"service_kind" json:"serviceKind,omitempty"`
		Target      *struct {
			Account       *string `tfsdk:"account" json:"account,omitempty"`
			ContainerPort *struct {
				ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
				PortName      *string `tfsdk:"port_name" json:"portName,omitempty"`
			} `tfsdk:"container_port" json:"containerPort,omitempty"`
			FallbackRole          *string `tfsdk:"fallback_role" json:"fallbackRole,omitempty"`
			Role                  *string `tfsdk:"role" json:"role,omitempty"`
			Strategy              *string `tfsdk:"strategy" json:"strategy,omitempty"`
			UseParentSelectedPods *bool   `tfsdk:"use_parent_selected_pods" json:"useParentSelectedPods,omitempty"`
		} `tfsdk:"target" json:"target,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dataprotection_kubeblocks_io_backup_policy_template_v1alpha1_manifest"
}

func (r *DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupPolicyTemplate should be provided by addon developers. It is responsible for generating BackupPolicies for the addon that requires backup operations, also determining the suitable backup methods and strategies.",
		MarkdownDescription: "BackupPolicyTemplate should be provided by addon developers. It is responsible for generating BackupPolicies for the addon that requires backup operations, also determining the suitable backup methods and strategies.",
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
				Description:         "Defines the desired state of the BackupPolicyTemplate.",
				MarkdownDescription: "Defines the desired state of the BackupPolicyTemplate.",
				Attributes: map[string]schema.Attribute{
					"backoff_limit": schema.Int64Attribute{
						Description:         "Specifies the maximum number of retry attempts for a backup before it is considered a failure.",
						MarkdownDescription: "Specifies the maximum number of retry attempts for a backup before it is considered a failure.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(10),
						},
					},

					"backup_methods": schema.ListNestedAttribute{
						Description:         "Defines an array of BackupMethods to be used.",
						MarkdownDescription: "Defines an array of BackupMethods to be used.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action_set_name": schema.StringAttribute{
									Description:         "Refers to the ActionSet object that defines the backup actions. For volume snapshot backup, the actionSet is not required, the controller will use the CSI volume snapshotter to create the snapshot.",
									MarkdownDescription: "Refers to the ActionSet object that defines the backup actions. For volume snapshot backup, the actionSet is not required, the controller will use the CSI volume snapshotter to create the snapshot.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"compatible_method": schema.StringAttribute{
									Description:         "The name of the compatible full backup method, used by incremental backups.",
									MarkdownDescription: "The name of the compatible full backup method, used by incremental backups.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"env": schema.ListNestedAttribute{
									Description:         "Specifies the environment variables for the backup workload.",
									MarkdownDescription: "Specifies the environment variables for the backup workload.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Specifies the environment variable key.",
												MarkdownDescription: "Specifies the environment variable key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"value": schema.StringAttribute{
												Description:         "Specifies the environment variable value.",
												MarkdownDescription: "Specifies the environment variable value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "Specifies the source used to determine the value of the environment variable. Cannot be used if value is not empty.",
												MarkdownDescription: "Specifies the source used to determine the value of the environment variable. Cannot be used if value is not empty.",
												Attributes: map[string]schema.Attribute{
													"version_mapping": schema.ListNestedAttribute{
														Description:         "Determine the appropriate version of the backup tool image from service version.",
														MarkdownDescription: "Determine the appropriate version of the backup tool image from service version.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"mapped_value": schema.StringAttribute{
																	Description:         "Specifies a mapping value based on service version. Typically used to set up the tools image required for backup operations.",
																	MarkdownDescription: "Specifies a mapping value based on service version. Typically used to set up the tools image required for backup operations.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service_versions": schema.ListAttribute{
																	Description:         "Represents an array of the service version that can be mapped to the appropriate value. Each name in the list can represent an exact name, a name prefix, or a regular expression pattern. For example: - '8.0.33': Matches the exact name '8.0.33' - '8.0': Matches all names starting with '8.0' - '^8.0.d{1,2}$': Matches all names starting with '8.0.' followed by one or two digits.",
																	MarkdownDescription: "Represents an array of the service version that can be mapped to the appropriate value. Each name in the list can represent an exact name, a name prefix, or a regular expression pattern. For example: - '8.0.33': Matches the exact name '8.0.33' - '8.0': Matches all names starting with '8.0' - '^8.0.d{1,2}$': Matches all names starting with '8.0.' followed by one or two digits.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
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
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "The name of backup method.",
									MarkdownDescription: "The name of backup method.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"runtime_settings": schema.SingleNestedAttribute{
									Description:         "Specifies runtime settings for the backup workload container.",
									MarkdownDescription: "Specifies runtime settings for the backup workload container.",
									Attributes: map[string]schema.Attribute{
										"resources": schema.SingleNestedAttribute{
											Description:         "Specifies the resource required by container. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
											MarkdownDescription: "Specifies the resource required by container. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
													MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"limits": schema.MapAttribute{
													Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"snapshot_volumes": schema.BoolAttribute{
									Description:         "Specifies whether to take snapshots of persistent volumes. If true, the ActionSetName is not required, the controller will use the CSI volume snapshotter to create the snapshot.",
									MarkdownDescription: "Specifies whether to take snapshots of persistent volumes. If true, the ActionSetName is not required, the controller will use the CSI volume snapshotter to create the snapshot.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target": schema.SingleNestedAttribute{
									Description:         "If set, specifies the method for selecting the replica to be backed up using the criteria defined here. If this field is not set, the selection method specified in 'backupPolicy.target' is used. This field provides a way to override the global 'backupPolicy.target' setting for specific BackupMethod.",
									MarkdownDescription: "If set, specifies the method for selecting the replica to be backed up using the criteria defined here. If this field is not set, the selection method specified in 'backupPolicy.target' is used. This field provides a way to override the global 'backupPolicy.target' setting for specific BackupMethod.",
									Attributes: map[string]schema.Attribute{
										"account": schema.StringAttribute{
											Description:         "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name. This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'. The corresponding secret created by this account is used to connect to the database.",
											MarkdownDescription: "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name. This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'. The corresponding secret created by this account is used to connect to the database.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"container_port": schema.SingleNestedAttribute{
											Description:         "Specifies the container port in the target pod. If not specified, the first container and its first port will be used.",
											MarkdownDescription: "Specifies the container port in the target pod. If not specified, the first container and its first port will be used.",
											Attributes: map[string]schema.Attribute{
												"container_name": schema.StringAttribute{
													Description:         "Specifies the name of container with the port.",
													MarkdownDescription: "Specifies the name of container with the port.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port_name": schema.StringAttribute{
													Description:         "Specifies the port name.",
													MarkdownDescription: "Specifies the port name.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"fallback_role": schema.StringAttribute{
											Description:         "Specifies the fallback role to select one replica for backup, this only takes effect when the 'strategy' field below is set to 'Any'.",
											MarkdownDescription: "Specifies the fallback role to select one replica for backup, this only takes effect when the 'strategy' field below is set to 'Any'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"role": schema.StringAttribute{
											Description:         "Specifies the role to select one or more replicas for backup. - If no replica with the specified role exists, the backup task will fail. Special case: If there is only one replica in the cluster, it will be used for backup, even if its role differs from the specified one. For example, if you specify backing up on a secondary replica, but the cluster is single-node with only one primary replica, the primary will be used for backup. Future versions will address this special case using role priorities. - If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to the 'strategy' field below.",
											MarkdownDescription: "Specifies the role to select one or more replicas for backup. - If no replica with the specified role exists, the backup task will fail. Special case: If there is only one replica in the cluster, it will be used for backup, even if its role differs from the specified one. For example, if you specify backing up on a secondary replica, but the cluster is single-node with only one primary replica, the primary will be used for backup. Future versions will address this special case using role priorities. - If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to the 'strategy' field below.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strategy": schema.StringAttribute{
											Description:         "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are: - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
											MarkdownDescription: "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are: - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Any", "All"),
											},
										},

										"use_parent_selected_pods": schema.BoolAttribute{
											Description:         "UseParentSelectedPods indicates whether to use the pods selected by the parent for backup. If set to true, the backup will use the same pods selected by the parent. And only takes effect when the 'strategy' is set to 'Any'.",
											MarkdownDescription: "UseParentSelectedPods indicates whether to use the pods selected by the parent for backup. If set to true, the backup will use the same pods selected by the parent. And only takes effect when the 'strategy' is set to 'Any'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"target_volumes": schema.SingleNestedAttribute{
									Description:         "Specifies which volumes from the target should be mounted in the backup workload.",
									MarkdownDescription: "Specifies which volumes from the target should be mounted in the backup workload.",
									Attributes: map[string]schema.Attribute{
										"volume_mounts": schema.ListNestedAttribute{
											Description:         "Specifies the mount for the volumes specified in 'volumes' section.",
											MarkdownDescription: "Specifies the mount for the volumes specified in 'volumes' section.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "Path within the container at which the volume should be mounted. Must not contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted. Must not contain ':'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"mount_propagation": schema.StringAttribute{
														Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "This must match the Name of a Volume.",
														MarkdownDescription: "This must match the Name of a Volume.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"read_only": schema.BoolAttribute{
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sub_path": schema.StringAttribute{
														Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sub_path_expr": schema.StringAttribute{
														Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

										"volumes": schema.ListAttribute{
											Description:         "Specifies the list of volumes of targeted application that should be mounted on the backup workload.",
											MarkdownDescription: "Specifies the list of volumes of targeted application that should be mounted on the backup workload.",
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"comp_defs": schema.ListAttribute{
						Description:         "CompDefs specifies names for the component definitions associated with this BackupPolicyTemplate. Each name in the list can represent an exact name, a name prefix, or a regular expression pattern. For example: - 'mysql-8.0.30-v1alpha1': Matches the exact name 'mysql-8.0.30-v1alpha1' - 'mysql-8.0.30': Matches all names starting with 'mysql-8.0.30' - '^mysql-8.0.d{1,2}$': Matches all names starting with 'mysql-8.0.' followed by one or two digits.",
						MarkdownDescription: "CompDefs specifies names for the component definitions associated with this BackupPolicyTemplate. Each name in the list can represent an exact name, a name prefix, or a regular expression pattern. For example: - 'mysql-8.0.30-v1alpha1': Matches the exact name 'mysql-8.0.30-v1alpha1' - 'mysql-8.0.30': Matches all names starting with 'mysql-8.0.30' - '^mysql-8.0.d{1,2}$': Matches all names starting with 'mysql-8.0.' followed by one or two digits.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retention_policy": schema.StringAttribute{
						Description:         "Defines the backup retention policy to be used.",
						MarkdownDescription: "Defines the backup retention policy to be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("retainLatestBackup", "none"),
						},
					},

					"schedules": schema.ListNestedAttribute{
						Description:         "Defines the execution plans for backup tasks, specifying when and how backups should occur, and the retention period of backup files.",
						MarkdownDescription: "Defines the execution plans for backup tasks, specifying when and how backups should occur, and the retention period of backup files.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backup_method": schema.StringAttribute{
									Description:         "Specifies the backup method name that is defined in backupPolicy.",
									MarkdownDescription: "Specifies the backup method name that is defined in backupPolicy.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cron_expression": schema.StringAttribute{
									Description:         "Specifies the cron expression for the schedule. The timezone is in UTC. see https://en.wikipedia.org/wiki/Cron.",
									MarkdownDescription: "Specifies the cron expression for the schedule. The timezone is in UTC. see https://en.wikipedia.org/wiki/Cron.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"enabled": schema.BoolAttribute{
									Description:         "Specifies whether the backup schedule is enabled or not.",
									MarkdownDescription: "Specifies whether the backup schedule is enabled or not.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Specifies the name of the schedule. Names cannot be duplicated. If the name is empty, it will be considered the same as the value of the backupMethod below.",
									MarkdownDescription: "Specifies the name of the schedule. Names cannot be duplicated. If the name is empty, it will be considered the same as the value of the backupMethod below.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"parameters": schema.ListNestedAttribute{
									Description:         "Specifies a list of name-value pairs representing parameters and their corresponding values. Parameters match the schema specified in the 'actionset.spec.parametersSchema'",
									MarkdownDescription: "Specifies a list of name-value pairs representing parameters and their corresponding values. Parameters match the schema specified in the 'actionset.spec.parametersSchema'",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Represents the name of the parameter.",
												MarkdownDescription: "Represents the name of the parameter.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"value": schema.StringAttribute{
												Description:         "Represents the parameter values.",
												MarkdownDescription: "Represents the parameter values.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"retention_period": schema.StringAttribute{
									Description:         "Determines the duration for which the backup should be kept. KubeBlocks will remove all backups that are older than the RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format: - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m You can also combine the above durations. For example: 30d12h30m",
									MarkdownDescription: "Determines the duration for which the backup should be kept. KubeBlocks will remove all backups that are older than the RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format: - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m You can also combine the above durations. For example: 30d12h30m",
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

					"service_kind": schema.StringAttribute{
						Description:         "Defines the type of well-known service protocol that the BackupPolicyTemplate provides, and it is optional. Some examples of well-known service protocols include: - 'MySQL': Indicates that the Component provides a MySQL database service. - 'PostgreSQL': Indicates that the Component offers a PostgreSQL database service. - 'Redis': Signifies that the Component functions as a Redis key-value store. - 'ETCD': Denotes that the Component serves as an ETCD distributed key-value store",
						MarkdownDescription: "Defines the type of well-known service protocol that the BackupPolicyTemplate provides, and it is optional. Some examples of well-known service protocols include: - 'MySQL': Indicates that the Component provides a MySQL database service. - 'PostgreSQL': Indicates that the Component offers a PostgreSQL database service. - 'Redis': Signifies that the Component functions as a Redis key-value store. - 'ETCD': Denotes that the Component serves as an ETCD distributed key-value store",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(32),
						},
					},

					"target": schema.SingleNestedAttribute{
						Description:         "Defines the selection criteria of instance to be backed up, and the connection credential to be used during the backup process.",
						MarkdownDescription: "Defines the selection criteria of instance to be backed up, and the connection credential to be used during the backup process.",
						Attributes: map[string]schema.Attribute{
							"account": schema.StringAttribute{
								Description:         "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name. This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'. The corresponding secret created by this account is used to connect to the database.",
								MarkdownDescription: "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name. This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'. The corresponding secret created by this account is used to connect to the database.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"container_port": schema.SingleNestedAttribute{
								Description:         "Specifies the container port in the target pod. If not specified, the first container and its first port will be used.",
								MarkdownDescription: "Specifies the container port in the target pod. If not specified, the first container and its first port will be used.",
								Attributes: map[string]schema.Attribute{
									"container_name": schema.StringAttribute{
										Description:         "Specifies the name of container with the port.",
										MarkdownDescription: "Specifies the name of container with the port.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_name": schema.StringAttribute{
										Description:         "Specifies the port name.",
										MarkdownDescription: "Specifies the port name.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"fallback_role": schema.StringAttribute{
								Description:         "Specifies the fallback role to select one replica for backup, this only takes effect when the 'strategy' field below is set to 'Any'.",
								MarkdownDescription: "Specifies the fallback role to select one replica for backup, this only takes effect when the 'strategy' field below is set to 'Any'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role": schema.StringAttribute{
								Description:         "Specifies the role to select one or more replicas for backup. - If no replica with the specified role exists, the backup task will fail. Special case: If there is only one replica in the cluster, it will be used for backup, even if its role differs from the specified one. For example, if you specify backing up on a secondary replica, but the cluster is single-node with only one primary replica, the primary will be used for backup. Future versions will address this special case using role priorities. - If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to the 'strategy' field below.",
								MarkdownDescription: "Specifies the role to select one or more replicas for backup. - If no replica with the specified role exists, the backup task will fail. Special case: If there is only one replica in the cluster, it will be used for backup, even if its role differs from the specified one. For example, if you specify backing up on a secondary replica, but the cluster is single-node with only one primary replica, the primary will be used for backup. Future versions will address this special case using role priorities. - If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to the 'strategy' field below.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"strategy": schema.StringAttribute{
								Description:         "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are: - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
								MarkdownDescription: "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are: - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Any", "All"),
								},
							},

							"use_parent_selected_pods": schema.BoolAttribute{
								Description:         "UseParentSelectedPods indicates whether to use the pods selected by the parent for backup. If set to true, the backup will use the same pods selected by the parent. And only takes effect when the 'strategy' is set to 'Any'.",
								MarkdownDescription: "UseParentSelectedPods indicates whether to use the pods selected by the parent for backup. If set to true, the backup will use the same pods selected by the parent. And only takes effect when the 'strategy' is set to 'Any'.",
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
	}
}

func (r *DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dataprotection_kubeblocks_io_backup_policy_template_v1alpha1_manifest")

	var model DataprotectionKubeblocksIoBackupPolicyTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("dataprotection.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("BackupPolicyTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
