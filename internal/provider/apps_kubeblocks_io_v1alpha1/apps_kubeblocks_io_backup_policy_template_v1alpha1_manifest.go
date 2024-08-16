/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest{}
}

type AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest struct{}

type AppsKubeblocksIoBackupPolicyTemplateV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BackupPolicies *[]struct {
			BackoffLimit  *int64 `tfsdk:"backoff_limit" json:"backoffLimit,omitempty"`
			BackupMethods *[]struct {
				ActionSetName *string `tfsdk:"action_set_name" json:"actionSetName,omitempty"`
				Env           *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
							Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				EnvMapping *[]struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					ValueFrom *struct {
						ClusterVersionRef *[]struct {
							MappingValue *string   `tfsdk:"mapping_value" json:"mappingValue,omitempty"`
							Names        *[]string `tfsdk:"names" json:"names,omitempty"`
						} `tfsdk:"cluster_version_ref" json:"clusterVersionRef,omitempty"`
						ComponentDef *[]struct {
							MappingValue *string   `tfsdk:"mapping_value" json:"mappingValue,omitempty"`
							Names        *[]string `tfsdk:"names" json:"names,omitempty"`
						} `tfsdk:"component_def" json:"componentDef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env_mapping" json:"envMapping,omitempty"`
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
					Account              *string `tfsdk:"account" json:"account,omitempty"`
					ConnectionCredential *struct {
						HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
						PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
						PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
						SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
					} `tfsdk:"connection_credential" json:"connectionCredential,omitempty"`
					ContainerPort *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						PortName      *string `tfsdk:"port_name" json:"portName,omitempty"`
					} `tfsdk:"container_port" json:"containerPort,omitempty"`
					FallbackRole *string `tfsdk:"fallback_role" json:"fallbackRole,omitempty"`
					Name         *string `tfsdk:"name" json:"name,omitempty"`
					PodSelector  *struct {
						FallbackLabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"fallback_label_selector" json:"fallbackLabelSelector,omitempty"`
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						Strategy    *string            `tfsdk:"strategy" json:"strategy,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Resources *struct {
						Excluded *[]string `tfsdk:"excluded" json:"excluded,omitempty"`
						Included *[]string `tfsdk:"included" json:"included,omitempty"`
						Selector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Role               *string `tfsdk:"role" json:"role,omitempty"`
					ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
					Strategy           *string `tfsdk:"strategy" json:"strategy,omitempty"`
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
				Targets *[]struct {
					ConnectionCredential *struct {
						HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
						PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
						PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
						SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
					} `tfsdk:"connection_credential" json:"connectionCredential,omitempty"`
					ContainerPort *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						PortName      *string `tfsdk:"port_name" json:"portName,omitempty"`
					} `tfsdk:"container_port" json:"containerPort,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					PodSelector *struct {
						FallbackLabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"fallback_label_selector" json:"fallbackLabelSelector,omitempty"`
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						Strategy    *string            `tfsdk:"strategy" json:"strategy,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Resources *struct {
						Excluded *[]string `tfsdk:"excluded" json:"excluded,omitempty"`
						Included *[]string `tfsdk:"included" json:"included,omitempty"`
						Selector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
				} `tfsdk:"targets" json:"targets,omitempty"`
			} `tfsdk:"backup_methods" json:"backupMethods,omitempty"`
			ComponentDefs *[]string `tfsdk:"component_defs" json:"componentDefs,omitempty"`
			Schedules     *[]struct {
				BackupMethod    *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
				CronExpression  *string `tfsdk:"cron_expression" json:"cronExpression,omitempty"`
				Enabled         *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				RetentionPeriod *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
			} `tfsdk:"schedules" json:"schedules,omitempty"`
			Target *struct {
				Account       *string `tfsdk:"account" json:"account,omitempty"`
				ContainerPort *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					PortName      *string `tfsdk:"port_name" json:"portName,omitempty"`
				} `tfsdk:"container_port" json:"containerPort,omitempty"`
				FallbackRole *string `tfsdk:"fallback_role" json:"fallbackRole,omitempty"`
				Role         *string `tfsdk:"role" json:"role,omitempty"`
				Strategy     *string `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"backup_policies" json:"backupPolicies,omitempty"`
		Identifier *string `tfsdk:"identifier" json:"identifier,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupPolicyTemplate should be provided by addon developers and is linked to a ClusterDefinitionand its associated ComponentDefinitions.It is responsible for generating BackupPolicies for Components that require backup operations,also determining the suitable backup methods and strategies.This template is automatically selected based on the specified ClusterDefinition and ComponentDefinitionswhen a Cluster is created.",
		MarkdownDescription: "BackupPolicyTemplate should be provided by addon developers and is linked to a ClusterDefinitionand its associated ComponentDefinitions.It is responsible for generating BackupPolicies for Components that require backup operations,also determining the suitable backup methods and strategies.This template is automatically selected based on the specified ClusterDefinition and ComponentDefinitionswhen a Cluster is created.",
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
					"backup_policies": schema.ListNestedAttribute{
						Description:         "Represents an array of BackupPolicy templates, with each template corresponding to a specified ComponentDefinitionor to a group of ComponentDefinitions that are different versions of definitions of the same component.",
						MarkdownDescription: "Represents an array of BackupPolicy templates, with each template corresponding to a specified ComponentDefinitionor to a group of ComponentDefinitions that are different versions of definitions of the same component.",
						NestedObject: schema.NestedAttributeObject{
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
												Description:         "Refers to the ActionSet object that defines the backup actions.For volume snapshot backup, the actionSet is not required, the controllerwill use the CSI volume snapshotter to create the snapshot.",
												MarkdownDescription: "Refers to the ActionSet object that defines the backup actions.For volume snapshot backup, the actionSet is not required, the controllerwill use the CSI volume snapshotter to create the snapshot.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"env": schema.ListNestedAttribute{
												Description:         "Specifies the environment variables for the backup workload.",
												MarkdownDescription: "Specifies the environment variables for the backup workload.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
															MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
															MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value_from": schema.SingleNestedAttribute{
															Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
															MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
															Attributes: map[string]schema.Attribute{
																"config_map_key_ref": schema.SingleNestedAttribute{
																	Description:         "Selects a key of a ConfigMap.",
																	MarkdownDescription: "Selects a key of a ConfigMap.",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "The key to select.",
																			MarkdownDescription: "The key to select.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
																			Description:         "Specify whether the ConfigMap or its key must be defined",
																			MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"field_ref": schema.SingleNestedAttribute{
																	Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																	MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																	Attributes: map[string]schema.Attribute{
																		"api_version": schema.StringAttribute{
																			Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																			MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"field_path": schema.StringAttribute{
																			Description:         "Path of the field to select in the specified API version.",
																			MarkdownDescription: "Path of the field to select in the specified API version.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"resource_field_ref": schema.SingleNestedAttribute{
																	Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																	MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																	Attributes: map[string]schema.Attribute{
																		"container_name": schema.StringAttribute{
																			Description:         "Container name: required for volumes, optional for env vars",
																			MarkdownDescription: "Container name: required for volumes, optional for env vars",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"divisor": schema.StringAttribute{
																			Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																			MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"resource": schema.StringAttribute{
																			Description:         "Required: resource to select",
																			MarkdownDescription: "Required: resource to select",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"secret_key_ref": schema.SingleNestedAttribute{
																	Description:         "Selects a key of a secret in the pod's namespace",
																	MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "The key of the secret to select from.  Must be a valid secret key.",
																			MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
																			Description:         "Specify whether the Secret or its key must be defined",
																			MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"env_mapping": schema.ListNestedAttribute{
												Description:         "Specifies a mapping of an environment variable key to the appropriate version of the tool imagerequired for backups, as determined by ClusterVersion and ComponentDefinition.The environment variable is then injected into the container executing the backup task.",
												MarkdownDescription: "Specifies a mapping of an environment variable key to the appropriate version of the tool imagerequired for backups, as determined by ClusterVersion and ComponentDefinition.The environment variable is then injected into the container executing the backup task.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Specifies the environment variable key in the mapping.",
															MarkdownDescription: "Specifies the environment variable key in the mapping.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value_from": schema.SingleNestedAttribute{
															Description:         "Specifies the source used to derive the value of the environment variable,which typically represents the tool image required for backup operation.",
															MarkdownDescription: "Specifies the source used to derive the value of the environment variable,which typically represents the tool image required for backup operation.",
															Attributes: map[string]schema.Attribute{
																"cluster_version_ref": schema.ListNestedAttribute{
																	Description:         "Determine the appropriate version of the backup tool image from ClusterVersion.Deprecated since v0.9, since ClusterVersion is deprecated.",
																	MarkdownDescription: "Determine the appropriate version of the backup tool image from ClusterVersion.Deprecated since v0.9, since ClusterVersion is deprecated.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mapping_value": schema.StringAttribute{
																				Description:         "Specifies the appropriate version of the backup tool image.",
																				MarkdownDescription: "Specifies the appropriate version of the backup tool image.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"names": schema.ListAttribute{
																				Description:         "Represents an array of names of ComponentDefinition that can be mapped to the appropriate version of the backup tool image.This mapping allows different versions of component images to correspond to specific versions of backup tool images.",
																				MarkdownDescription: "Represents an array of names of ComponentDefinition that can be mapped to the appropriate version of the backup tool image.This mapping allows different versions of component images to correspond to specific versions of backup tool images.",
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

																"component_def": schema.ListNestedAttribute{
																	Description:         "Determine the appropriate version of the backup tool image from ComponentDefinition.",
																	MarkdownDescription: "Determine the appropriate version of the backup tool image from ComponentDefinition.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mapping_value": schema.StringAttribute{
																				Description:         "Specifies the appropriate version of the backup tool image.",
																				MarkdownDescription: "Specifies the appropriate version of the backup tool image.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"names": schema.ListAttribute{
																				Description:         "Represents an array of names of ComponentDefinition that can be mapped to the appropriate version of the backup tool image.This mapping allows different versions of component images to correspond to specific versions of backup tool images.",
																				MarkdownDescription: "Represents an array of names of ComponentDefinition that can be mapped to the appropriate version of the backup tool image.This mapping allows different versions of component images to correspond to specific versions of backup tool images.",
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
														Description:         "Specifies the resource required by container.More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Specifies the resource required by container.More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														Attributes: map[string]schema.Attribute{
															"claims": schema.ListNestedAttribute{
																Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																			MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"requests": schema.MapAttribute{
																Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
												Description:         "Specifies whether to take snapshots of persistent volumes. If true,the ActionSetName is not required, the controller will use the CSI volumesnapshotter to create the snapshot.",
												MarkdownDescription: "Specifies whether to take snapshots of persistent volumes. If true,the ActionSetName is not required, the controller will use the CSI volumesnapshotter to create the snapshot.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target": schema.SingleNestedAttribute{
												Description:         "Specifies the target information to back up, it will override the target in backup policy.",
												MarkdownDescription: "Specifies the target information to back up, it will override the target in backup policy.",
												Attributes: map[string]schema.Attribute{
													"account": schema.StringAttribute{
														Description:         "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name.This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'.The corresponding secret created by this account is used to connect to the database.",
														MarkdownDescription: "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name.This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'.The corresponding secret created by this account is used to connect to the database.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"connection_credential": schema.SingleNestedAttribute{
														Description:         "Specifies the connection credential to connect to the target database cluster.",
														MarkdownDescription: "Specifies the connection credential to connect to the target database cluster.",
														Attributes: map[string]schema.Attribute{
															"host_key": schema.StringAttribute{
																Description:         "Specifies the map key of the host in the connection credential secret.",
																MarkdownDescription: "Specifies the map key of the host in the connection credential secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"password_key": schema.StringAttribute{
																Description:         "Specifies the map key of the password in the connection credential secret.This password will be saved in the backup annotation for full backup.You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
																MarkdownDescription: "Specifies the map key of the password in the connection credential secret.This password will be saved in the backup annotation for full backup.You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"port_key": schema.StringAttribute{
																Description:         "Specifies the map key of the port in the connection credential secret.",
																MarkdownDescription: "Specifies the map key of the port in the connection credential secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
																Description:         "Refers to the Secret object that contains the connection credential.",
																MarkdownDescription: "Refers to the Secret object that contains the connection credential.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
																},
															},

															"username_key": schema.StringAttribute{
																Description:         "Specifies the map key of the user in the connection credential secret.",
																MarkdownDescription: "Specifies the map key of the user in the connection credential secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"container_port": schema.SingleNestedAttribute{
														Description:         "Specifies the container port in the target pod.If not specified, the first container and its first port will be used.",
														MarkdownDescription: "Specifies the container port in the target pod.If not specified, the first container and its first port will be used.",
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
														Description:         "Specifies the fallback role to select one replica for backup, this only takes effect when the'strategy' field below is set to 'Any'.",
														MarkdownDescription: "Specifies the fallback role to select one replica for backup, this only takes effect when the'strategy' field below is set to 'Any'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Specifies a mandatory and unique identifier for each target when using the 'targets' field.The backup data for the current target is stored in a uniquely named subdirectory.",
														MarkdownDescription: "Specifies a mandatory and unique identifier for each target when using the 'targets' field.The backup data for the current target is stored in a uniquely named subdirectory.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Used to find the target pod. The volumes of the target pod will be backed up.",
														MarkdownDescription: "Used to find the target pod. The volumes of the target pod will be backed up.",
														Attributes: map[string]schema.Attribute{
															"fallback_label_selector": schema.SingleNestedAttribute{
																Description:         "fallbackLabelSelector is used to filter available pods when the labelSelector fails.This only takes effect when the 'strategy' field below is set to 'Any'.",
																MarkdownDescription: "fallbackLabelSelector is used to filter available pods when the labelSelector fails.This only takes effect when the 'strategy' field below is set to 'Any'.",
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
																					Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																			Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																			MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"values": schema.ListAttribute{
																			Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																			MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"strategy": schema.StringAttribute{
																Description:         "Specifies the strategy to select the target pod when multiple pods are selected.Valid values are:- 'Any': select any one pod that match the labelsSelector.- 'All': select all pods that match the labelsSelector. The backup data for the current podwill be stored in a subdirectory named after the pod.",
																MarkdownDescription: "Specifies the strategy to select the target pod when multiple pods are selected.Valid values are:- 'Any': select any one pod that match the labelsSelector.- 'All': select all pods that match the labelsSelector. The backup data for the current podwill be stored in a subdirectory named after the pod.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Any", "All"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "Specifies the kubernetes resources to back up.",
														MarkdownDescription: "Specifies the kubernetes resources to back up.",
														Attributes: map[string]schema.Attribute{
															"excluded": schema.ListAttribute{
																Description:         "excluded is a slice of namespaced-scoped resource type names to exclude inthe kubernetes resources.The default value is empty.",
																MarkdownDescription: "excluded is a slice of namespaced-scoped resource type names to exclude inthe kubernetes resources.The default value is empty.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"included": schema.ListAttribute{
																Description:         "included is a slice of namespaced-scoped resource type names to include inthe kubernetes resources.The default value is empty.",
																MarkdownDescription: "included is a slice of namespaced-scoped resource type names to include inthe kubernetes resources.The default value is empty.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selector": schema.SingleNestedAttribute{
																Description:         "A metav1.LabelSelector to filter the target kubernetes resources that needto be backed up. If not set, will do not back up any kubernetes resources.",
																MarkdownDescription: "A metav1.LabelSelector to filter the target kubernetes resources that needto be backed up. If not set, will do not back up any kubernetes resources.",
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
																					Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

													"role": schema.StringAttribute{
														Description:         "Specifies the role to select one or more replicas for backup.- If no replica with the specified role exists, the backup task will fail.  Special case: If there is only one replica in the cluster, it will be used for backup,  even if its role differs from the specified one.  For example, if you specify backing up on a secondary replica, but the cluster is single-node  with only one primary replica, the primary will be used for backup.  Future versions will address this special case using role priorities.- If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to  the 'strategy' field below.",
														MarkdownDescription: "Specifies the role to select one or more replicas for backup.- If no replica with the specified role exists, the backup task will fail.  Special case: If there is only one replica in the cluster, it will be used for backup,  even if its role differs from the specified one.  For example, if you specify backing up on a secondary replica, but the cluster is single-node  with only one primary replica, the primary will be used for backup.  Future versions will address this special case using role priorities.- If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to  the 'strategy' field below.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"service_account_name": schema.StringAttribute{
														Description:         "Specifies the service account to run the backup workload.",
														MarkdownDescription: "Specifies the service account to run the backup workload.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"strategy": schema.StringAttribute{
														Description:         "Specifies the PodSelectionStrategy to use when multiple pods areselected for the backup target.Valid values are:- Any: Selects any one pod that matches the labelsSelector.- All: Selects all pods that match the labelsSelector.",
														MarkdownDescription: "Specifies the PodSelectionStrategy to use when multiple pods areselected for the backup target.Valid values are:- Any: Selects any one pod that matches the labelsSelector.- All: Selects all pods that match the labelsSelector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Any", "All"),
														},
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
																	Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
																	MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"mount_propagation": schema.StringAttribute{
																	Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																	MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
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
																	Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																	MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sub_path": schema.StringAttribute{
																	Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																	MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sub_path_expr": schema.StringAttribute{
																	Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																	MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
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
														Description:         "Specifies the list of volumes of targeted application that should be mountedon the backup workload.",
														MarkdownDescription: "Specifies the list of volumes of targeted application that should be mountedon the backup workload.",
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

											"targets": schema.ListNestedAttribute{
												Description:         "Specifies multiple target information for backup operations. This includes detailssuch as the target pod and cluster connection credentials. All specified targetswill be backed up collectively.",
												MarkdownDescription: "Specifies multiple target information for backup operations. This includes detailssuch as the target pod and cluster connection credentials. All specified targetswill be backed up collectively.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"connection_credential": schema.SingleNestedAttribute{
															Description:         "Specifies the connection credential to connect to the target database cluster.",
															MarkdownDescription: "Specifies the connection credential to connect to the target database cluster.",
															Attributes: map[string]schema.Attribute{
																"host_key": schema.StringAttribute{
																	Description:         "Specifies the map key of the host in the connection credential secret.",
																	MarkdownDescription: "Specifies the map key of the host in the connection credential secret.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"password_key": schema.StringAttribute{
																	Description:         "Specifies the map key of the password in the connection credential secret.This password will be saved in the backup annotation for full backup.You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
																	MarkdownDescription: "Specifies the map key of the password in the connection credential secret.This password will be saved in the backup annotation for full backup.You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port_key": schema.StringAttribute{
																	Description:         "Specifies the map key of the port in the connection credential secret.",
																	MarkdownDescription: "Specifies the map key of the port in the connection credential secret.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_name": schema.StringAttribute{
																	Description:         "Refers to the Secret object that contains the connection credential.",
																	MarkdownDescription: "Refers to the Secret object that contains the connection credential.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
																	},
																},

																"username_key": schema.StringAttribute{
																	Description:         "Specifies the map key of the user in the connection credential secret.",
																	MarkdownDescription: "Specifies the map key of the user in the connection credential secret.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"container_port": schema.SingleNestedAttribute{
															Description:         "Specifies the container port in the target pod.If not specified, the first container and its first port will be used.",
															MarkdownDescription: "Specifies the container port in the target pod.If not specified, the first container and its first port will be used.",
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

														"name": schema.StringAttribute{
															Description:         "Specifies a mandatory and unique identifier for each target when using the 'targets' field.The backup data for the current target is stored in a uniquely named subdirectory.",
															MarkdownDescription: "Specifies a mandatory and unique identifier for each target when using the 'targets' field.The backup data for the current target is stored in a uniquely named subdirectory.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_selector": schema.SingleNestedAttribute{
															Description:         "Used to find the target pod. The volumes of the target pod will be backed up.",
															MarkdownDescription: "Used to find the target pod. The volumes of the target pod will be backed up.",
															Attributes: map[string]schema.Attribute{
																"fallback_label_selector": schema.SingleNestedAttribute{
																	Description:         "fallbackLabelSelector is used to filter available pods when the labelSelector fails.This only takes effect when the 'strategy' field below is set to 'Any'.",
																	MarkdownDescription: "fallbackLabelSelector is used to filter available pods when the labelSelector fails.This only takes effect when the 'strategy' field below is set to 'Any'.",
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
																						Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"strategy": schema.StringAttribute{
																	Description:         "Specifies the strategy to select the target pod when multiple pods are selected.Valid values are:- 'Any': select any one pod that match the labelsSelector.- 'All': select all pods that match the labelsSelector. The backup data for the current podwill be stored in a subdirectory named after the pod.",
																	MarkdownDescription: "Specifies the strategy to select the target pod when multiple pods are selected.Valid values are:- 'Any': select any one pod that match the labelsSelector.- 'All': select all pods that match the labelsSelector. The backup data for the current podwill be stored in a subdirectory named after the pod.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Any", "All"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Specifies the kubernetes resources to back up.",
															MarkdownDescription: "Specifies the kubernetes resources to back up.",
															Attributes: map[string]schema.Attribute{
																"excluded": schema.ListAttribute{
																	Description:         "excluded is a slice of namespaced-scoped resource type names to exclude inthe kubernetes resources.The default value is empty.",
																	MarkdownDescription: "excluded is a slice of namespaced-scoped resource type names to exclude inthe kubernetes resources.The default value is empty.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"included": schema.ListAttribute{
																	Description:         "included is a slice of namespaced-scoped resource type names to include inthe kubernetes resources.The default value is empty.",
																	MarkdownDescription: "included is a slice of namespaced-scoped resource type names to include inthe kubernetes resources.The default value is empty.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "A metav1.LabelSelector to filter the target kubernetes resources that needto be backed up. If not set, will do not back up any kubernetes resources.",
																	MarkdownDescription: "A metav1.LabelSelector to filter the target kubernetes resources that needto be backed up. If not set, will do not back up any kubernetes resources.",
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
																						Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"service_account_name": schema.StringAttribute{
															Description:         "Specifies the service account to run the backup workload.",
															MarkdownDescription: "Specifies the service account to run the backup workload.",
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
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"component_defs": schema.ListAttribute{
									Description:         "Specifies a list of names of ComponentDefinitions that the specified ClusterDefinition references.They should be different versions of definitions of the same component,thus allowing them to share a single BackupPolicy.Each name must adhere to the IANA Service Naming rule.",
									MarkdownDescription: "Specifies a list of names of ComponentDefinitions that the specified ClusterDefinition references.They should be different versions of definitions of the same component,thus allowing them to share a single BackupPolicy.Each name must adhere to the IANA Service Naming rule.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"schedules": schema.ListNestedAttribute{
									Description:         "Defines the execution plans for backup tasks, specifying when and how backups should occur,and the retention period of backup files.",
									MarkdownDescription: "Defines the execution plans for backup tasks, specifying when and how backups should occur,and the retention period of backup files.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"backup_method": schema.StringAttribute{
												Description:         "Defines the backup method name that is defined in backupPolicy.",
												MarkdownDescription: "Defines the backup method name that is defined in backupPolicy.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"cron_expression": schema.StringAttribute{
												Description:         "Represents the cron expression for schedule, with the timezone set in UTC.Refer to https://en.wikipedia.org/wiki/Cron for more details.",
												MarkdownDescription: "Represents the cron expression for schedule, with the timezone set in UTC.Refer to https://en.wikipedia.org/wiki/Cron for more details.",
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

											"retention_period": schema.StringAttribute{
												Description:         "Determines the duration for which the backup should be retained.The controller will remove all backups that are older than the RetentionPeriod.For instance, a RetentionPeriod of '30d' will retain only the backups from the last 30 days.Sample duration format:- years: 	2y- months: 	6mo- days: 		30d- hours: 	12h- minutes: 	30mThese durations can also be combined, for example: 30d12h30m.",
												MarkdownDescription: "Determines the duration for which the backup should be retained.The controller will remove all backups that are older than the RetentionPeriod.For instance, a RetentionPeriod of '30d' will retain only the backups from the last 30 days.Sample duration format:- years: 	2y- months: 	6mo- days: 		30d- hours: 	12h- minutes: 	30mThese durations can also be combined, for example: 30d12h30m.",
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

								"target": schema.SingleNestedAttribute{
									Description:         "Defines the selection criteria of instance to be backed up, and the connection credential to be usedduring the backup process.",
									MarkdownDescription: "Defines the selection criteria of instance to be backed up, and the connection credential to be usedduring the backup process.",
									Attributes: map[string]schema.Attribute{
										"account": schema.StringAttribute{
											Description:         "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name.This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'.The corresponding secret created by this account is used to connect to the database.",
											MarkdownDescription: "If 'backupPolicy.componentDefs' is set, this field is required to specify the system account name.This account must match one listed in 'componentDefinition.spec.systemAccounts[*].name'.The corresponding secret created by this account is used to connect to the database.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"container_port": schema.SingleNestedAttribute{
											Description:         "Specifies the container port in the target pod.If not specified, the first container and its first port will be used.",
											MarkdownDescription: "Specifies the container port in the target pod.If not specified, the first container and its first port will be used.",
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
											Description:         "Specifies the fallback role to select one replica for backup, this only takes effect when the'strategy' field below is set to 'Any'.",
											MarkdownDescription: "Specifies the fallback role to select one replica for backup, this only takes effect when the'strategy' field below is set to 'Any'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"role": schema.StringAttribute{
											Description:         "Specifies the role to select one or more replicas for backup.- If no replica with the specified role exists, the backup task will fail.  Special case: If there is only one replica in the cluster, it will be used for backup,  even if its role differs from the specified one.  For example, if you specify backing up on a secondary replica, but the cluster is single-node  with only one primary replica, the primary will be used for backup.  Future versions will address this special case using role priorities.- If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to  the 'strategy' field below.",
											MarkdownDescription: "Specifies the role to select one or more replicas for backup.- If no replica with the specified role exists, the backup task will fail.  Special case: If there is only one replica in the cluster, it will be used for backup,  even if its role differs from the specified one.  For example, if you specify backing up on a secondary replica, but the cluster is single-node  with only one primary replica, the primary will be used for backup.  Future versions will address this special case using role priorities.- If multiple replicas satisfy the specified role, the choice ('Any' or 'All') will be made according to  the 'strategy' field below.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strategy": schema.StringAttribute{
											Description:         "Specifies the PodSelectionStrategy to use when multiple pods areselected for the backup target.Valid values are:- Any: Selects any one pod that matches the labelsSelector.- All: Selects all pods that match the labelsSelector.",
											MarkdownDescription: "Specifies the PodSelectionStrategy to use when multiple pods areselected for the backup target.Valid values are:- Any: Selects any one pod that matches the labelsSelector.- All: Selects all pods that match the labelsSelector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Any", "All"),
											},
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

					"identifier": schema.StringAttribute{
						Description:         "Specifies a unique identifier for the BackupPolicyTemplate.This identifier will be used as the suffix of the name of automatically generated BackupPolicy.This prevents unintended overwriting of BackupPolicies due to name conflicts when multiple BackupPolicyTemplatesare present.For instance, using 'backup-policy' for regular backups and 'backup-policy-hscale' for horizontal-scale opscan differentiate the policies.",
						MarkdownDescription: "Specifies a unique identifier for the BackupPolicyTemplate.This identifier will be used as the suffix of the name of automatically generated BackupPolicy.This prevents unintended overwriting of BackupPolicies due to name conflicts when multiple BackupPolicyTemplatesare present.For instance, using 'backup-policy' for regular backups and 'backup-policy-hscale' for horizontal-scale opscan differentiate the policies.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(20),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest")

	var model AppsKubeblocksIoBackupPolicyTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("BackupPolicyTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
