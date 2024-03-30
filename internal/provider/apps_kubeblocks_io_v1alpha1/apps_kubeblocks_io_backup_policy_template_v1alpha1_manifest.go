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
					ConnectionCredentialKey *struct {
						HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
						PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
						PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
						UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
					} `tfsdk:"connection_credential_key" json:"connectionCredentialKey,omitempty"`
					PodSelector *struct {
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
			} `tfsdk:"backup_methods" json:"backupMethods,omitempty"`
			ComponentDefRef *string   `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
			ComponentDefs   *[]string `tfsdk:"component_defs" json:"componentDefs,omitempty"`
			Schedules       *[]struct {
				BackupMethod    *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
				CronExpression  *string `tfsdk:"cron_expression" json:"cronExpression,omitempty"`
				Enabled         *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				RetentionPeriod *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
			} `tfsdk:"schedules" json:"schedules,omitempty"`
			Target *struct {
				Account                 *string `tfsdk:"account" json:"account,omitempty"`
				ConnectionCredentialKey *struct {
					HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
					PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
					PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
					UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
				} `tfsdk:"connection_credential_key" json:"connectionCredentialKey,omitempty"`
				Role     *string `tfsdk:"role" json:"role,omitempty"`
				Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"backup_policies" json:"backupPolicies,omitempty"`
		ClusterDefinitionRef *string `tfsdk:"cluster_definition_ref" json:"clusterDefinitionRef,omitempty"`
		Identifier           *string `tfsdk:"identifier" json:"identifier,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_backup_policy_template_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoBackupPolicyTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupPolicyTemplate is the Schema for the BackupPolicyTemplates API (defined by provider)",
		MarkdownDescription: "BackupPolicyTemplate is the Schema for the BackupPolicyTemplates API (defined by provider)",
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
						Description:         "Represents an array of backup policy templates for the specified ComponentDefinition.",
						MarkdownDescription: "Represents an array of backup policy templates for the specified ComponentDefinition.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backoff_limit": schema.Int64Attribute{
									Description:         "Specifies the number of retries before marking the backup as failed.",
									MarkdownDescription: "Specifies the number of retries before marking the backup as failed.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(10),
									},
								},

								"backup_methods": schema.ListNestedAttribute{
									Description:         "Define the methods to be used for backups.",
									MarkdownDescription: "Define the methods to be used for backups.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"action_set_name": schema.StringAttribute{
												Description:         "Refers to the ActionSet object that defines the backup actions. For volume snapshot backup, the actionSet is not required, the controller will use the CSI volume snapshotter to create the snapshot.",
												MarkdownDescription: "Refers to the ActionSet object that defines the backup actions. For volume snapshot backup, the actionSet is not required, the controller will use the CSI volume snapshotter to create the snapshot.",
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
															Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
															MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
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
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																	MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
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
																	Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																	MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
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
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
												Description:         "Defines the mapping between the environment variables of the cluster and the keys of the environment values.",
												MarkdownDescription: "Defines the mapping between the environment variables of the cluster and the keys of the environment values.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Specifies the environment variable key that requires mapping.",
															MarkdownDescription: "Specifies the environment variable key that requires mapping.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value_from": schema.SingleNestedAttribute{
															Description:         "Defines the source from which the environment variable value is derived.",
															MarkdownDescription: "Defines the source from which the environment variable value is derived.",
															Attributes: map[string]schema.Attribute{
																"cluster_version_ref": schema.ListNestedAttribute{
																	Description:         "Maps to the environment value. This is an optional field.",
																	MarkdownDescription: "Maps to the environment value. This is an optional field.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mapping_value": schema.StringAttribute{
																				Description:         "The value that corresponds to the specified ClusterVersion names.",
																				MarkdownDescription: "The value that corresponds to the specified ClusterVersion names.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"names": schema.ListAttribute{
																				Description:         "Represents an array of ClusterVersion names that can be mapped to an environment variable value.",
																				MarkdownDescription: "Represents an array of ClusterVersion names that can be mapped to an environment variable value.",
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
																	Description:         "Maps to the environment value. This is also an optional field.",
																	MarkdownDescription: "Maps to the environment value. This is also an optional field.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mapping_value": schema.StringAttribute{
																				Description:         "The value that corresponds to the specified ClusterVersion names.",
																				MarkdownDescription: "The value that corresponds to the specified ClusterVersion names.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"names": schema.ListAttribute{
																				Description:         "Represents an array of ClusterVersion names that can be mapped to an environment variable value.",
																				MarkdownDescription: "Represents an array of ClusterVersion names that can be mapped to an environment variable value.",
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
														Description:         "Specifies the resource required by container. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Specifies the resource required by container. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														Attributes: map[string]schema.Attribute{
															"claims": schema.ListNestedAttribute{
																Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
												Description:         "Specifies the target information to back up, it will override the target in backup policy.",
												MarkdownDescription: "Specifies the target information to back up, it will override the target in backup policy.",
												Attributes: map[string]schema.Attribute{
													"account": schema.StringAttribute{
														Description:         "Refers to spec.componentDef.systemAccounts.accounts[*].name in the ClusterDefinition. The secret created by this account will be used to connect to the database. If not set, the secret created by spec.ConnectionCredential of the ClusterDefinition will be used.  It will be transformed into a secret for the BackupPolicy's target secret.",
														MarkdownDescription: "Refers to spec.componentDef.systemAccounts.accounts[*].name in the ClusterDefinition. The secret created by this account will be used to connect to the database. If not set, the secret created by spec.ConnectionCredential of the ClusterDefinition will be used.  It will be transformed into a secret for the BackupPolicy's target secret.",
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
																Description:         "Specifies the map key of the password in the connection credential secret. This password will be saved in the backup annotation for full backup. You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
																MarkdownDescription: "Specifies the map key of the password in the connection credential secret. This password will be saved in the backup annotation for full backup. You can use the environment variable DP_ENCRYPTION_KEY to specify encryption key.",
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

													"connection_credential_key": schema.SingleNestedAttribute{
														Description:         "Defines the connection credential key in the secret created by spec.ConnectionCredential of the ClusterDefinition. It will be ignored when the 'account' is set.",
														MarkdownDescription: "Defines the connection credential key in the secret created by spec.ConnectionCredential of the ClusterDefinition. It will be ignored when the 'account' is set.",
														Attributes: map[string]schema.Attribute{
															"host_key": schema.StringAttribute{
																Description:         "Defines the map key of the host in the connection credential secret.",
																MarkdownDescription: "Defines the map key of the host in the connection credential secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"password_key": schema.StringAttribute{
																Description:         "Represents the key of the password in the ConnectionCredential secret. If not specified, the default key 'password' is used.",
																MarkdownDescription: "Represents the key of the password in the ConnectionCredential secret. If not specified, the default key 'password' is used.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"port_key": schema.StringAttribute{
																Description:         "Indicates the map key of the port in the connection credential secret.",
																MarkdownDescription: "Indicates the map key of the port in the connection credential secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username_key": schema.StringAttribute{
																Description:         "Represents the key of the username in the ConnectionCredential secret. If not specified, the default key 'username' is used.",
																MarkdownDescription: "Represents the key of the username in the ConnectionCredential secret. If not specified, the default key 'username' is used.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Used to find the target pod. The volumes of the target pod will be backed up.",
														MarkdownDescription: "Used to find the target pod. The volumes of the target pod will be backed up.",
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

															"strategy": schema.StringAttribute{
																Description:         "Specifies the strategy to select the target pod when multiple pods are selected. Valid values are: Any: select any one pod that match the labelsSelector.  - 'Any': select any one pod that match the labelsSelector. - 'All': select all pods that match the labelsSelector.",
																MarkdownDescription: "Specifies the strategy to select the target pod when multiple pods are selected. Valid values are: Any: select any one pod that match the labelsSelector.  - 'Any': select any one pod that match the labelsSelector. - 'All': select all pods that match the labelsSelector.",
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
																Description:         "excluded is a slice of namespaced-scoped resource type names to exclude in the kubernetes resources. The default value is empty.",
																MarkdownDescription: "excluded is a slice of namespaced-scoped resource type names to exclude in the kubernetes resources. The default value is empty.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"included": schema.ListAttribute{
																Description:         "included is a slice of namespaced-scoped resource type names to include in the kubernetes resources. The default value is empty.",
																MarkdownDescription: "included is a slice of namespaced-scoped resource type names to include in the kubernetes resources. The default value is empty.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selector": schema.SingleNestedAttribute{
																Description:         "A metav1.LabelSelector to filter the target kubernetes resources that need to be backed up. If not set, will do not back up any kubernetes resources.",
																MarkdownDescription: "A metav1.LabelSelector to filter the target kubernetes resources that need to be backed up. If not set, will do not back up any kubernetes resources.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"role": schema.StringAttribute{
														Description:         "Specifies the instance of the corresponding role for backup. The roles can be:  - Leader, Follower, or Leaner for the Consensus component. - Primary or Secondary for the Replication component.  Invalid roles of the component will be ignored. For example, if the workload type is Replication and the component's replicas is 1, the secondary role is invalid. It will also be ignored when the component is Stateful or Stateless.  The role will be transformed into a role LabelSelector for the BackupPolicy's target attribute.",
														MarkdownDescription: "Specifies the instance of the corresponding role for backup. The roles can be:  - Leader, Follower, or Leaner for the Consensus component. - Primary or Secondary for the Replication component.  Invalid roles of the component will be ignored. For example, if the workload type is Replication and the component's replicas is 1, the secondary role is invalid. It will also be ignored when the component is Stateful or Stateless.  The role will be transformed into a role LabelSelector for the BackupPolicy's target attribute.",
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
														Description:         "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are:  - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
														MarkdownDescription: "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are:  - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
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
																	Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
																	MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
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

								"component_def_ref": schema.StringAttribute{
									Description:         "References a componentDef defined in the ClusterDefinition spec. Must comply with the IANA Service Naming rule.",
									MarkdownDescription: "References a componentDef defined in the ClusterDefinition spec. Must comply with the IANA Service Naming rule.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(22),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
									},
								},

								"component_defs": schema.ListAttribute{
									Description:         "References to componentDefinitions. Must comply with the IANA Service Naming rule.",
									MarkdownDescription: "References to componentDefinitions. Must comply with the IANA Service Naming rule.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"schedules": schema.ListNestedAttribute{
									Description:         "Define the policy for backup scheduling.",
									MarkdownDescription: "Define the policy for backup scheduling.",
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
												Description:         "Represents the cron expression for schedule, with the timezone set in UTC. Refer to https://en.wikipedia.org/wiki/Cron for more details.",
												MarkdownDescription: "Represents the cron expression for schedule, with the timezone set in UTC. Refer to https://en.wikipedia.org/wiki/Cron for more details.",
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
												Description:         "Determines the duration for which the backup should be retained. The controller will remove all backups that are older than the RetentionPeriod. For instance, a RetentionPeriod of '30d' will retain only the backups from the last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  These durations can also be combined, for example: 30d12h30m.",
												MarkdownDescription: "Determines the duration for which the backup should be retained. The controller will remove all backups that are older than the RetentionPeriod. For instance, a RetentionPeriod of '30d' will retain only the backups from the last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  These durations can also be combined, for example: 30d12h30m.",
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
									Description:         "The instance to be backed up.",
									MarkdownDescription: "The instance to be backed up.",
									Attributes: map[string]schema.Attribute{
										"account": schema.StringAttribute{
											Description:         "Refers to spec.componentDef.systemAccounts.accounts[*].name in the ClusterDefinition. The secret created by this account will be used to connect to the database. If not set, the secret created by spec.ConnectionCredential of the ClusterDefinition will be used.  It will be transformed into a secret for the BackupPolicy's target secret.",
											MarkdownDescription: "Refers to spec.componentDef.systemAccounts.accounts[*].name in the ClusterDefinition. The secret created by this account will be used to connect to the database. If not set, the secret created by spec.ConnectionCredential of the ClusterDefinition will be used.  It will be transformed into a secret for the BackupPolicy's target secret.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"connection_credential_key": schema.SingleNestedAttribute{
											Description:         "Defines the connection credential key in the secret created by spec.ConnectionCredential of the ClusterDefinition. It will be ignored when the 'account' is set.",
											MarkdownDescription: "Defines the connection credential key in the secret created by spec.ConnectionCredential of the ClusterDefinition. It will be ignored when the 'account' is set.",
											Attributes: map[string]schema.Attribute{
												"host_key": schema.StringAttribute{
													Description:         "Defines the map key of the host in the connection credential secret.",
													MarkdownDescription: "Defines the map key of the host in the connection credential secret.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"password_key": schema.StringAttribute{
													Description:         "Represents the key of the password in the ConnectionCredential secret. If not specified, the default key 'password' is used.",
													MarkdownDescription: "Represents the key of the password in the ConnectionCredential secret. If not specified, the default key 'password' is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port_key": schema.StringAttribute{
													Description:         "Indicates the map key of the port in the connection credential secret.",
													MarkdownDescription: "Indicates the map key of the port in the connection credential secret.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"username_key": schema.StringAttribute{
													Description:         "Represents the key of the username in the ConnectionCredential secret. If not specified, the default key 'username' is used.",
													MarkdownDescription: "Represents the key of the username in the ConnectionCredential secret. If not specified, the default key 'username' is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"role": schema.StringAttribute{
											Description:         "Specifies the instance of the corresponding role for backup. The roles can be:  - Leader, Follower, or Leaner for the Consensus component. - Primary or Secondary for the Replication component.  Invalid roles of the component will be ignored. For example, if the workload type is Replication and the component's replicas is 1, the secondary role is invalid. It will also be ignored when the component is Stateful or Stateless.  The role will be transformed into a role LabelSelector for the BackupPolicy's target attribute.",
											MarkdownDescription: "Specifies the instance of the corresponding role for backup. The roles can be:  - Leader, Follower, or Leaner for the Consensus component. - Primary or Secondary for the Replication component.  Invalid roles of the component will be ignored. For example, if the workload type is Replication and the component's replicas is 1, the secondary role is invalid. It will also be ignored when the component is Stateful or Stateless.  The role will be transformed into a role LabelSelector for the BackupPolicy's target attribute.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"strategy": schema.StringAttribute{
											Description:         "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are:  - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
											MarkdownDescription: "Specifies the PodSelectionStrategy to use when multiple pods are selected for the backup target. Valid values are:  - Any: Selects any one pod that matches the labelsSelector. - All: Selects all pods that match the labelsSelector.",
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

					"cluster_definition_ref": schema.StringAttribute{
						Description:         "Specifies a reference to the ClusterDefinition name. This is an immutable attribute that cannot be changed after creation.",
						MarkdownDescription: "Specifies a reference to the ClusterDefinition name. This is an immutable attribute that cannot be changed after creation.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"identifier": schema.StringAttribute{
						Description:         "Acts as a unique identifier for this BackupPolicyTemplate. This identifier will be used as a suffix for the automatically generated backupPolicy name. It is required when multiple BackupPolicyTemplates exist to prevent backupPolicy override.",
						MarkdownDescription: "Acts as a unique identifier for this BackupPolicyTemplate. This identifier will be used as a suffix for the automatically generated backupPolicy name. It is required when multiple BackupPolicyTemplates exist to prevent backupPolicy override.",
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
