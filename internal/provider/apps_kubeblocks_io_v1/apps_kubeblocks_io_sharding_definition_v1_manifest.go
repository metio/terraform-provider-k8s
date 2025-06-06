/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1

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
	_ datasource.DataSource = &AppsKubeblocksIoShardingDefinitionV1Manifest{}
)

func NewAppsKubeblocksIoShardingDefinitionV1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoShardingDefinitionV1Manifest{}
}

type AppsKubeblocksIoShardingDefinitionV1Manifest struct{}

type AppsKubeblocksIoShardingDefinitionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		LifecycleActions *struct {
			PostProvision *struct {
				Exec *struct {
					Args      *[]string `tfsdk:"args" json:"args,omitempty"`
					Command   *[]string `tfsdk:"command" json:"command,omitempty"`
					Container *string   `tfsdk:"container" json:"container,omitempty"`
					Env       *[]struct {
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
					Image             *string `tfsdk:"image" json:"image,omitempty"`
					MatchingKey       *string `tfsdk:"matching_key" json:"matchingKey,omitempty"`
					TargetPodSelector *string `tfsdk:"target_pod_selector" json:"targetPodSelector,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				PreCondition *string `tfsdk:"pre_condition" json:"preCondition,omitempty"`
				RetryPolicy  *struct {
					MaxRetries    *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					RetryInterval *int64 `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
				} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"post_provision" json:"postProvision,omitempty"`
			PreTerminate *struct {
				Exec *struct {
					Args      *[]string `tfsdk:"args" json:"args,omitempty"`
					Command   *[]string `tfsdk:"command" json:"command,omitempty"`
					Container *string   `tfsdk:"container" json:"container,omitempty"`
					Env       *[]struct {
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
					Image             *string `tfsdk:"image" json:"image,omitempty"`
					MatchingKey       *string `tfsdk:"matching_key" json:"matchingKey,omitempty"`
					TargetPodSelector *string `tfsdk:"target_pod_selector" json:"targetPodSelector,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				PreCondition *string `tfsdk:"pre_condition" json:"preCondition,omitempty"`
				RetryPolicy  *struct {
					MaxRetries    *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					RetryInterval *int64 `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
				} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"pre_terminate" json:"preTerminate,omitempty"`
			ShardAdd *struct {
				Exec *struct {
					Args      *[]string `tfsdk:"args" json:"args,omitempty"`
					Command   *[]string `tfsdk:"command" json:"command,omitempty"`
					Container *string   `tfsdk:"container" json:"container,omitempty"`
					Env       *[]struct {
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
					Image             *string `tfsdk:"image" json:"image,omitempty"`
					MatchingKey       *string `tfsdk:"matching_key" json:"matchingKey,omitempty"`
					TargetPodSelector *string `tfsdk:"target_pod_selector" json:"targetPodSelector,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				PreCondition *string `tfsdk:"pre_condition" json:"preCondition,omitempty"`
				RetryPolicy  *struct {
					MaxRetries    *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					RetryInterval *int64 `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
				} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"shard_add" json:"shardAdd,omitempty"`
			ShardRemove *struct {
				Exec *struct {
					Args      *[]string `tfsdk:"args" json:"args,omitempty"`
					Command   *[]string `tfsdk:"command" json:"command,omitempty"`
					Container *string   `tfsdk:"container" json:"container,omitempty"`
					Env       *[]struct {
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
					Image             *string `tfsdk:"image" json:"image,omitempty"`
					MatchingKey       *string `tfsdk:"matching_key" json:"matchingKey,omitempty"`
					TargetPodSelector *string `tfsdk:"target_pod_selector" json:"targetPodSelector,omitempty"`
				} `tfsdk:"exec" json:"exec,omitempty"`
				PreCondition *string `tfsdk:"pre_condition" json:"preCondition,omitempty"`
				RetryPolicy  *struct {
					MaxRetries    *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					RetryInterval *int64 `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
				} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"shard_remove" json:"shardRemove,omitempty"`
		} `tfsdk:"lifecycle_actions" json:"lifecycleActions,omitempty"`
		ProvisionStrategy *string `tfsdk:"provision_strategy" json:"provisionStrategy,omitempty"`
		ShardsLimit       *struct {
			MaxShards *int64 `tfsdk:"max_shards" json:"maxShards,omitempty"`
			MinShards *int64 `tfsdk:"min_shards" json:"minShards,omitempty"`
		} `tfsdk:"shards_limit" json:"shardsLimit,omitempty"`
		SystemAccounts *[]struct {
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Shared *bool   `tfsdk:"shared" json:"shared,omitempty"`
		} `tfsdk:"system_accounts" json:"systemAccounts,omitempty"`
		Template *struct {
			CompDef *string `tfsdk:"comp_def" json:"compDef,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		Tls *struct {
			Shared *bool `tfsdk:"shared" json:"shared,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		UpdateStrategy *string `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoShardingDefinitionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_sharding_definition_v1_manifest"
}

func (r *AppsKubeblocksIoShardingDefinitionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ShardingDefinition is the Schema for the shardingdefinitions API",
		MarkdownDescription: "ShardingDefinition is the Schema for the shardingdefinitions API",
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
				Description:         "ShardingDefinitionSpec defines the desired state of ShardingDefinition",
				MarkdownDescription: "ShardingDefinitionSpec defines the desired state of ShardingDefinition",
				Attributes: map[string]schema.Attribute{
					"lifecycle_actions": schema.SingleNestedAttribute{
						Description:         "Defines a set of hooks and procedures that customize the behavior of a sharding throughout its lifecycle. This field is immutable.",
						MarkdownDescription: "Defines a set of hooks and procedures that customize the behavior of a sharding throughout its lifecycle. This field is immutable.",
						Attributes: map[string]schema.Attribute{
							"post_provision": schema.SingleNestedAttribute{
								Description:         "Specifies the hook to be executed after a sharding's creation. By setting 'postProvision.preCondition', you can determine the specific lifecycle stage at which the action should trigger, available conditions for sharding include: 'Immediately', 'ComponentReady', and 'ClusterReady'. For sharding, the 'ComponentReady' condition means all components of the sharding are ready. With 'ComponentReady' being the default. The PostProvision Action is intended to run only once. Note: This field is immutable once it has been set.",
								MarkdownDescription: "Specifies the hook to be executed after a sharding's creation. By setting 'postProvision.preCondition', you can determine the specific lifecycle stage at which the action should trigger, available conditions for sharding include: 'Immediately', 'ComponentReady', and 'ClusterReady'. For sharding, the 'ComponentReady' condition means all components of the sharding are ready. With 'ComponentReady' being the default. The PostProvision Action is intended to run only once. Note: This field is immutable once it has been set.",
								Attributes: map[string]schema.Attribute{
									"exec": schema.SingleNestedAttribute{
										Description:         "Defines the command to run. This field cannot be updated.",
										MarkdownDescription: "Defines the command to run. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"args": schema.ListAttribute{
												Description:         "Args represents the arguments that are passed to the 'command' for execution.",
												MarkdownDescription: "Args represents the arguments that are passed to the 'command' for execution.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"command": schema.ListAttribute{
												Description:         "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												MarkdownDescription: "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"container": schema.StringAttribute{
												Description:         "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												MarkdownDescription: "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"env": schema.ListNestedAttribute{
												Description:         "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
												MarkdownDescription: "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
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
																			Description:         "The key of the secret to select from. Must be a valid secret key.",
																			MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
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

											"image": schema.StringAttribute{
												Description:         "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												MarkdownDescription: "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"matching_key": schema.StringAttribute{
												Description:         "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												MarkdownDescription: "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_pod_selector": schema.StringAttribute{
												Description:         "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												MarkdownDescription: "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Any", "All", "Role", "Ordinal"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pre_condition": schema.StringAttribute{
										Description:         "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										MarkdownDescription: "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_policy": schema.SingleNestedAttribute{
										Description:         "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										MarkdownDescription: "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"max_retries": schema.Int64Attribute{
												Description:         "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												MarkdownDescription: "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_interval": schema.Int64Attribute{
												Description:         "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												MarkdownDescription: "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
										MarkdownDescription: "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pre_terminate": schema.SingleNestedAttribute{
								Description:         "Specifies the hook to be executed prior to terminating a sharding. The PreTerminate Action is intended to run only once. This action is executed immediately when a terminate operation for the sharding is initiated. The actual termination and cleanup of the sharding and its associated resources will not proceed until the PreTerminate action has completed successfully. Note: This field is immutable once it has been set.",
								MarkdownDescription: "Specifies the hook to be executed prior to terminating a sharding. The PreTerminate Action is intended to run only once. This action is executed immediately when a terminate operation for the sharding is initiated. The actual termination and cleanup of the sharding and its associated resources will not proceed until the PreTerminate action has completed successfully. Note: This field is immutable once it has been set.",
								Attributes: map[string]schema.Attribute{
									"exec": schema.SingleNestedAttribute{
										Description:         "Defines the command to run. This field cannot be updated.",
										MarkdownDescription: "Defines the command to run. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"args": schema.ListAttribute{
												Description:         "Args represents the arguments that are passed to the 'command' for execution.",
												MarkdownDescription: "Args represents the arguments that are passed to the 'command' for execution.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"command": schema.ListAttribute{
												Description:         "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												MarkdownDescription: "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"container": schema.StringAttribute{
												Description:         "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												MarkdownDescription: "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"env": schema.ListNestedAttribute{
												Description:         "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
												MarkdownDescription: "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
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
																			Description:         "The key of the secret to select from. Must be a valid secret key.",
																			MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
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

											"image": schema.StringAttribute{
												Description:         "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												MarkdownDescription: "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"matching_key": schema.StringAttribute{
												Description:         "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												MarkdownDescription: "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_pod_selector": schema.StringAttribute{
												Description:         "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												MarkdownDescription: "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Any", "All", "Role", "Ordinal"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pre_condition": schema.StringAttribute{
										Description:         "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										MarkdownDescription: "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_policy": schema.SingleNestedAttribute{
										Description:         "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										MarkdownDescription: "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"max_retries": schema.Int64Attribute{
												Description:         "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												MarkdownDescription: "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_interval": schema.Int64Attribute{
												Description:         "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												MarkdownDescription: "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
										MarkdownDescription: "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"shard_add": schema.SingleNestedAttribute{
								Description:         "Specifies the hook to be executed after a shard added. Note: This field is immutable once it has been set.",
								MarkdownDescription: "Specifies the hook to be executed after a shard added. Note: This field is immutable once it has been set.",
								Attributes: map[string]schema.Attribute{
									"exec": schema.SingleNestedAttribute{
										Description:         "Defines the command to run. This field cannot be updated.",
										MarkdownDescription: "Defines the command to run. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"args": schema.ListAttribute{
												Description:         "Args represents the arguments that are passed to the 'command' for execution.",
												MarkdownDescription: "Args represents the arguments that are passed to the 'command' for execution.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"command": schema.ListAttribute{
												Description:         "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												MarkdownDescription: "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"container": schema.StringAttribute{
												Description:         "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												MarkdownDescription: "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"env": schema.ListNestedAttribute{
												Description:         "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
												MarkdownDescription: "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
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
																			Description:         "The key of the secret to select from. Must be a valid secret key.",
																			MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
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

											"image": schema.StringAttribute{
												Description:         "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												MarkdownDescription: "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"matching_key": schema.StringAttribute{
												Description:         "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												MarkdownDescription: "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_pod_selector": schema.StringAttribute{
												Description:         "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												MarkdownDescription: "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Any", "All", "Role", "Ordinal"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pre_condition": schema.StringAttribute{
										Description:         "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										MarkdownDescription: "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_policy": schema.SingleNestedAttribute{
										Description:         "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										MarkdownDescription: "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"max_retries": schema.Int64Attribute{
												Description:         "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												MarkdownDescription: "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_interval": schema.Int64Attribute{
												Description:         "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												MarkdownDescription: "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
										MarkdownDescription: "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"shard_remove": schema.SingleNestedAttribute{
								Description:         "Specifies the hook to be executed prior to remove a shard. Note: This field is immutable once it has been set.",
								MarkdownDescription: "Specifies the hook to be executed prior to remove a shard. Note: This field is immutable once it has been set.",
								Attributes: map[string]schema.Attribute{
									"exec": schema.SingleNestedAttribute{
										Description:         "Defines the command to run. This field cannot be updated.",
										MarkdownDescription: "Defines the command to run. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"args": schema.ListAttribute{
												Description:         "Args represents the arguments that are passed to the 'command' for execution.",
												MarkdownDescription: "Args represents the arguments that are passed to the 'command' for execution.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"command": schema.ListAttribute{
												Description:         "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												MarkdownDescription: "Specifies the command to be executed inside the container. The working directory for this command is the container's root directory('/'). Commands are executed directly without a shell environment, meaning shell-specific syntax ('|', etc.) is not supported. If the shell is required, it must be explicitly invoked in the command. A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"container": schema.StringAttribute{
												Description:         "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												MarkdownDescription: "Specifies the name of the container within the same pod whose resources will be shared with the action. This allows the action to utilize the specified container's resources without executing within it. The name must match one of the containers defined in 'componentDefinition.spec.runtime'. The resources that can be shared are included: - volume mounts This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"env": schema.ListNestedAttribute{
												Description:         "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
												MarkdownDescription: "Represents a list of environment variables that will be injected into the container. These variables enable the container to adapt its behavior based on the environment it's running in. This field cannot be updated.",
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
																			Description:         "The key of the secret to select from. Must be a valid secret key.",
																			MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
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

											"image": schema.StringAttribute{
												Description:         "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												MarkdownDescription: "Specifies the container image to be used for running the Action. When specified, a dedicated container will be created using this image to execute the Action. All actions with same image will share the same container. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"matching_key": schema.StringAttribute{
												Description:         "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												MarkdownDescription: "Used in conjunction with the 'targetPodSelector' field to refine the selection of target pod(s) for Action execution. The impact of this field depends on the 'targetPodSelector' value: - When 'targetPodSelector' is set to 'Any' or 'All', this field will be ignored. - When 'targetPodSelector' is set to 'Role', only those replicas whose role matches the 'matchingKey' will be selected for the Action. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_pod_selector": schema.StringAttribute{
												Description:         "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												MarkdownDescription: "Defines the criteria used to select the target Pod(s) for executing the Action. This is useful when there is no default target replica identified. It allows for precise control over which Pod(s) the Action should run in. If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod to be removed or added; or a random pod if the Action is triggered at the component level, such as post-provision or pre-terminate of the component. This field cannot be updated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Any", "All", "Role", "Ordinal"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pre_condition": schema.StringAttribute{
										Description:         "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										MarkdownDescription: "Specifies the state that the cluster must reach before the Action is executed. Currently, this is only applicable to the 'postProvision' action. The conditions are as follows: - 'Immediately': Executed right after the Component object is created. The readiness of the Component and its resources is not guaranteed at this stage. - 'RuntimeReady': The Action is triggered after the Component object has been created and all associated runtime resources (e.g. Pods) are in a ready state. - 'ComponentReady': The Action is triggered after the Component itself is in a ready state. This process does not affect the readiness state of the Component or the Cluster. - 'ClusterReady': The Action is executed after the Cluster is in a ready state. This execution does not alter the Component or the Cluster's state of readiness. This field cannot be updated.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_policy": schema.SingleNestedAttribute{
										Description:         "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										MarkdownDescription: "Defines the strategy to be taken when retrying the Action after a failure. It specifies the conditions under which the Action should be retried and the limits to apply, such as the maximum number of retries and backoff strategy. This field cannot be updated.",
										Attributes: map[string]schema.Attribute{
											"max_retries": schema.Int64Attribute{
												Description:         "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												MarkdownDescription: "Defines the maximum number of retry attempts that should be made for a given Action. This value is set to 0 by default, indicating that no retries will be made.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_interval": schema.Int64Attribute{
												Description:         "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												MarkdownDescription: "Indicates the duration of time to wait between each retry attempt. This value is set to 0 by default, indicating that there will be no delay between retry attempts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
										MarkdownDescription: "Specifies the maximum duration in seconds that the Action is allowed to run. If the Action does not complete within this time frame, it will be terminated. This field cannot be updated.",
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

					"provision_strategy": schema.StringAttribute{
						Description:         "Specifies the strategy for provisioning shards of the sharding. Only 'Serial' and 'Parallel' are supported. This field is immutable.",
						MarkdownDescription: "Specifies the strategy for provisioning shards of the sharding. Only 'Serial' and 'Parallel' are supported. This field is immutable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Serial", "BestEffortParallel", "Parallel"),
						},
					},

					"shards_limit": schema.SingleNestedAttribute{
						Description:         "Defines the upper limit of the number of shards supported by the sharding. This field is immutable.",
						MarkdownDescription: "Defines the upper limit of the number of shards supported by the sharding. This field is immutable.",
						Attributes: map[string]schema.Attribute{
							"max_shards": schema.Int64Attribute{
								Description:         "The maximum limit of shards.",
								MarkdownDescription: "The maximum limit of shards.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"min_shards": schema.Int64Attribute{
								Description:         "The minimum limit of shards.",
								MarkdownDescription: "The minimum limit of shards.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"system_accounts": schema.ListNestedAttribute{
						Description:         "Defines the system accounts for the sharding. This field is immutable.",
						MarkdownDescription: "Defines the system accounts for the sharding. This field is immutable.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "The name of the system account defined in the sharding template. This field is immutable once set.",
									MarkdownDescription: "The name of the system account defined in the sharding template. This field is immutable once set.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"shared": schema.BoolAttribute{
									Description:         "Specifies whether the account is shared across all shards in the sharding.",
									MarkdownDescription: "Specifies whether the account is shared across all shards in the sharding.",
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

					"template": schema.SingleNestedAttribute{
						Description:         "This field is immutable.",
						MarkdownDescription: "This field is immutable.",
						Attributes: map[string]schema.Attribute{
							"comp_def": schema.StringAttribute{
								Description:         "The component definition(s) that the sharding is based on. The component definition can be specified using one of the following: - the full name - the regular expression pattern ('^' will be added to the beginning of the pattern automatically) This field is immutable.",
								MarkdownDescription: "The component definition(s) that the sharding is based on. The component definition can be specified using one of the following: - the full name - the regular expression pattern ('^' will be added to the beginning of the pattern automatically) This field is immutable.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "Defines the TLS for the sharding. This field is immutable.",
						MarkdownDescription: "Defines the TLS for the sharding. This field is immutable.",
						Attributes: map[string]schema.Attribute{
							"shared": schema.BoolAttribute{
								Description:         "Specifies whether the TLS configuration is shared across all shards in the sharding.",
								MarkdownDescription: "Specifies whether the TLS configuration is shared across all shards in the sharding.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"update_strategy": schema.StringAttribute{
						Description:         "Specifies the strategy for updating shards of the sharding. Only 'Serial' and 'Parallel' are supported. This field is immutable.",
						MarkdownDescription: "Specifies the strategy for updating shards of the sharding. Only 'Serial' and 'Parallel' are supported. This field is immutable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Serial", "BestEffortParallel", "Parallel"),
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

func (r *AppsKubeblocksIoShardingDefinitionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_sharding_definition_v1_manifest")

	var model AppsKubeblocksIoShardingDefinitionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1")
	model.Kind = pointer.String("ShardingDefinition")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
