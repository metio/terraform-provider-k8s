/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1beta1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppsKubeblocksIoConfigConstraintV1Beta1Manifest{}
)

func NewAppsKubeblocksIoConfigConstraintV1Beta1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoConfigConstraintV1Beta1Manifest{}
}

type AppsKubeblocksIoConfigConstraintV1Beta1Manifest struct{}

type AppsKubeblocksIoConfigConstraintV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ConfigSchema *struct {
			Cue          *string            `tfsdk:"cue" json:"cue,omitempty"`
			SchemaInJSON *map[string]string `tfsdk:"schema_in_json" json:"schemaInJSON,omitempty"`
		} `tfsdk:"config_schema" json:"configSchema,omitempty"`
		ConfigSchemaTopLevelKey *string `tfsdk:"config_schema_top_level_key" json:"configSchemaTopLevelKey,omitempty"`
		DownwardActions         *[]struct {
			Command *[]string `tfsdk:"command" json:"command,omitempty"`
			Items   *[]struct {
				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
				Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
				Path             *string `tfsdk:"path" json:"path,omitempty"`
				ResourceFieldRef *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
					Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
				} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
			} `tfsdk:"items" json:"items,omitempty"`
			MountPoint *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"downward_actions" json:"downwardActions,omitempty"`
		DynamicActionCanBeMerged       *bool     `tfsdk:"dynamic_action_can_be_merged" json:"dynamicActionCanBeMerged,omitempty"`
		DynamicParameterSelectedPolicy *string   `tfsdk:"dynamic_parameter_selected_policy" json:"dynamicParameterSelectedPolicy,omitempty"`
		DynamicParameters              *[]string `tfsdk:"dynamic_parameters" json:"dynamicParameters,omitempty"`
		DynamicReloadAction            *struct {
			AutoTrigger *struct {
				ProcessName *string `tfsdk:"process_name" json:"processName,omitempty"`
			} `tfsdk:"auto_trigger" json:"autoTrigger,omitempty"`
			ShellTrigger *struct {
				BatchParametersTemplate *string   `tfsdk:"batch_parameters_template" json:"batchParametersTemplate,omitempty"`
				BatchReload             *bool     `tfsdk:"batch_reload" json:"batchReload,omitempty"`
				Command                 *[]string `tfsdk:"command" json:"command,omitempty"`
				Sync                    *bool     `tfsdk:"sync" json:"sync,omitempty"`
			} `tfsdk:"shell_trigger" json:"shellTrigger,omitempty"`
			TplScriptTrigger *struct {
				Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ScriptConfigMapRef *string `tfsdk:"script_config_map_ref" json:"scriptConfigMapRef,omitempty"`
				Sync               *bool   `tfsdk:"sync" json:"sync,omitempty"`
			} `tfsdk:"tpl_script_trigger" json:"tplScriptTrigger,omitempty"`
			UnixSignalTrigger *struct {
				ProcessName *string `tfsdk:"process_name" json:"processName,omitempty"`
				Signal      *string `tfsdk:"signal" json:"signal,omitempty"`
			} `tfsdk:"unix_signal_trigger" json:"unixSignalTrigger,omitempty"`
		} `tfsdk:"dynamic_reload_action" json:"dynamicReloadAction,omitempty"`
		DynamicReloadSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"dynamic_reload_selector" json:"dynamicReloadSelector,omitempty"`
		FormatterConfig *struct {
			Format    *string `tfsdk:"format" json:"format,omitempty"`
			IniConfig *struct {
				SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
			} `tfsdk:"ini_config" json:"iniConfig,omitempty"`
		} `tfsdk:"formatter_config" json:"formatterConfig,omitempty"`
		ImmutableParameters *[]string `tfsdk:"immutable_parameters" json:"immutableParameters,omitempty"`
		ReloadToolsImage    *struct {
			MountPoint  *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
			ToolConfigs *[]struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Image   *string   `tfsdk:"image" json:"image,omitempty"`
				Name    *string   `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"tool_configs" json:"toolConfigs,omitempty"`
		} `tfsdk:"reload_tools_image" json:"reloadToolsImage,omitempty"`
		ScriptConfigs *[]struct {
			Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ScriptConfigMapRef *string `tfsdk:"script_config_map_ref" json:"scriptConfigMapRef,omitempty"`
		} `tfsdk:"script_configs" json:"scriptConfigs,omitempty"`
		StaticParameters *[]string `tfsdk:"static_parameters" json:"staticParameters,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoConfigConstraintV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_config_constraint_v1beta1_manifest"
}

func (r *AppsKubeblocksIoConfigConstraintV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConfigConstraint manages the parameters across multiple configuration files contained in a single configure template. These configuration files should have the same format (e.g. ini, xml, properties, json).  It provides the following functionalities:  1. **Parameter Value Validation**: Validates and ensures compliance of parameter values with defined constraints. 2. **Dynamic Reload on Modification**: Monitors parameter changes and triggers dynamic reloads to apply updates. 3. **Parameter Rendering in Templates**: Injects parameters into templates to generate up-to-date configuration files.",
		MarkdownDescription: "ConfigConstraint manages the parameters across multiple configuration files contained in a single configure template. These configuration files should have the same format (e.g. ini, xml, properties, json).  It provides the following functionalities:  1. **Parameter Value Validation**: Validates and ensures compliance of parameter values with defined constraints. 2. **Dynamic Reload on Modification**: Monitors parameter changes and triggers dynamic reloads to apply updates. 3. **Parameter Rendering in Templates**: Injects parameters into templates to generate up-to-date configuration files.",
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
				Description:         "ConfigConstraintSpec defines the desired state of ConfigConstraint",
				MarkdownDescription: "ConfigConstraintSpec defines the desired state of ConfigConstraint",
				Attributes: map[string]schema.Attribute{
					"config_schema": schema.SingleNestedAttribute{
						Description:         "Defines a list of parameters including their names, default values, descriptions, types, and constraints (permissible values or the range of valid values).",
						MarkdownDescription: "Defines a list of parameters including their names, default values, descriptions, types, and constraints (permissible values or the range of valid values).",
						Attributes: map[string]schema.Attribute{
							"cue": schema.StringAttribute{
								Description:         "Hold a string that contains a script written in CUE language that defines a list of configuration items. Each item is detailed with its name, default value, description, type (e.g. string, integer, float), and constraints (permissible values or the valid range of values).  CUE (Configure, Unify, Execute) is a declarative language designed for defining and validating complex data configurations. It is particularly useful in environments like K8s where complex configurations and validation rules are common.  This script functions as a validator for user-provided configurations, ensuring compliance with the established specifications and constraints.",
								MarkdownDescription: "Hold a string that contains a script written in CUE language that defines a list of configuration items. Each item is detailed with its name, default value, description, type (e.g. string, integer, float), and constraints (permissible values or the valid range of values).  CUE (Configure, Unify, Execute) is a declarative language designed for defining and validating complex data configurations. It is particularly useful in environments like K8s where complex configurations and validation rules are common.  This script functions as a validator for user-provided configurations, ensuring compliance with the established specifications and constraints.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"schema_in_json": schema.MapAttribute{
								Description:         "Generated from the 'cue' field and transformed into a JSON format.",
								MarkdownDescription: "Generated from the 'cue' field and transformed into a JSON format.",
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

					"config_schema_top_level_key": schema.StringAttribute{
						Description:         "Specifies the top-level key in the 'configSchema.cue' that organizes the validation rules for parameters. This key must exist within the CUE script defined in 'configSchema.cue'.",
						MarkdownDescription: "Specifies the top-level key in the 'configSchema.cue' that organizes the validation rules for parameters. This key must exist within the CUE script defined in 'configSchema.cue'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"downward_actions": schema.ListNestedAttribute{
						Description:         "Specifies a list of actions to execute specified commands based on Pod labels.  It utilizes the K8s Downward API to mount label information as a volume into the pod. The 'config-manager' sidecar container watches for changes in the role label and dynamically invoke registered commands (usually execute some SQL statements) when a change is detected.  It is designed for scenarios where:  - Replicas with different roles have different configurations, such as Redis primary & secondary replicas. - After a role switch (e.g., from secondary to primary), some changes in configuration are needed to reflect the new role.",
						MarkdownDescription: "Specifies a list of actions to execute specified commands based on Pod labels.  It utilizes the K8s Downward API to mount label information as a volume into the pod. The 'config-manager' sidecar container watches for changes in the role label and dynamically invoke registered commands (usually execute some SQL statements) when a change is detected.  It is designed for scenarios where:  - Replicas with different roles have different configurations, such as Redis primary & secondary replicas. - After a role switch (e.g., from secondary to primary), some changes in configuration are needed to reflect the new role.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"command": schema.ListAttribute{
									Description:         "Specifies the command to be triggered when changes are detected in Downward API volume files. It relies on the inotify mechanism in the config-manager sidecar to monitor file changes.",
									MarkdownDescription: "Specifies the command to be triggered when changes are detected in Downward API volume files. It relies on the inotify mechanism in the config-manager sidecar to monitor file changes.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"items": schema.ListNestedAttribute{
									Description:         "Represents a list of files under the Downward API volume.",
									MarkdownDescription: "Represents a list of files under the Downward API volume.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"field_ref": schema.SingleNestedAttribute{
												Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
												MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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

											"mode": schema.Int64Attribute{
												Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
												MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"resource_field_ref": schema.SingleNestedAttribute{
												Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
												MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"mount_point": schema.StringAttribute{
									Description:         "Specifies the mount point of the Downward API volume.",
									MarkdownDescription: "Specifies the mount point of the Downward API volume.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(128),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Specifies the name of the field. It must be a string of maximum length 63. The name should match the regex pattern '^[a-z0-9]([a-z0-9.-]*[a-z0-9])?$'.",
									MarkdownDescription: "Specifies the name of the field. It must be a string of maximum length 63. The name should match the regex pattern '^[a-z0-9]([a-z0-9.-]*[a-z0-9])?$'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"dynamic_action_can_be_merged": schema.BoolAttribute{
						Description:         "Indicates whether to consolidate dynamic reload and restart actions into a single restart.  - If true, updates requiring both actions will result in only a restart, merging the actions. - If false, updates will trigger both actions executed sequentially: first dynamic reload, then restart.  This flag allows for more efficient handling of configuration changes by potentially eliminating an unnecessary reload step.",
						MarkdownDescription: "Indicates whether to consolidate dynamic reload and restart actions into a single restart.  - If true, updates requiring both actions will result in only a restart, merging the actions. - If false, updates will trigger both actions executed sequentially: first dynamic reload, then restart.  This flag allows for more efficient handling of configuration changes by potentially eliminating an unnecessary reload step.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dynamic_parameter_selected_policy": schema.StringAttribute{
						Description:         "Configures whether the dynamic reload specified in 'dynamicReloadAction' applies only to dynamic parameters or to all parameters (including static parameters).  - 'dynamic' (default): Only modifications to the dynamic parameters listed in 'dynamicParameters' will trigger a dynamic reload. - 'all': Modifications to both dynamic parameters listed in 'dynamicParameters' and static parameters listed in 'staticParameters' will trigger a dynamic reload. The 'all' option is for certain engines that require static parameters to be set via SQL statements before they can take effect on restart.",
						MarkdownDescription: "Configures whether the dynamic reload specified in 'dynamicReloadAction' applies only to dynamic parameters or to all parameters (including static parameters).  - 'dynamic' (default): Only modifications to the dynamic parameters listed in 'dynamicParameters' will trigger a dynamic reload. - 'all': Modifications to both dynamic parameters listed in 'dynamicParameters' and static parameters listed in 'staticParameters' will trigger a dynamic reload. The 'all' option is for certain engines that require static parameters to be set via SQL statements before they can take effect on restart.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("all", "dynamic"),
						},
					},

					"dynamic_parameters": schema.ListAttribute{
						Description:         "List dynamic parameters. Modifications to these parameters trigger a configuration reload without requiring a process restart.",
						MarkdownDescription: "List dynamic parameters. Modifications to these parameters trigger a configuration reload without requiring a process restart.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dynamic_reload_action": schema.SingleNestedAttribute{
						Description:         "Specifies the dynamic reload (dynamic reconfiguration) actions supported by the engine. When set, the controller executes the scripts defined in these actions to handle dynamic parameter updates.  Dynamic reloading is triggered only if both of the following conditions are met:  1. The modified parameters are listed in the 'dynamicParameters' field. If 'dynamicParameterSelectedPolicy' is set to 'all', modifications to 'staticParameters' can also trigger a reload. 2. 'dynamicReloadAction' is set.  If 'dynamicReloadAction' is not set or the modified parameters are not listed in 'dynamicParameters', dynamic reloading will not be triggered.  Example: '''yaml reloadOptions: tplScriptTrigger: namespace: kb-system scriptConfigMapRef: mysql-reload-script sync: true '''",
						MarkdownDescription: "Specifies the dynamic reload (dynamic reconfiguration) actions supported by the engine. When set, the controller executes the scripts defined in these actions to handle dynamic parameter updates.  Dynamic reloading is triggered only if both of the following conditions are met:  1. The modified parameters are listed in the 'dynamicParameters' field. If 'dynamicParameterSelectedPolicy' is set to 'all', modifications to 'staticParameters' can also trigger a reload. 2. 'dynamicReloadAction' is set.  If 'dynamicReloadAction' is not set or the modified parameters are not listed in 'dynamicParameters', dynamic reloading will not be triggered.  Example: '''yaml reloadOptions: tplScriptTrigger: namespace: kb-system scriptConfigMapRef: mysql-reload-script sync: true '''",
						Attributes: map[string]schema.Attribute{
							"auto_trigger": schema.SingleNestedAttribute{
								Description:         "Automatically perform the reload when specified conditions are met.",
								MarkdownDescription: "Automatically perform the reload when specified conditions are met.",
								Attributes: map[string]schema.Attribute{
									"process_name": schema.StringAttribute{
										Description:         "The name of the process.",
										MarkdownDescription: "The name of the process.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"shell_trigger": schema.SingleNestedAttribute{
								Description:         "Allows to execute a custom shell script to reload the process.",
								MarkdownDescription: "Allows to execute a custom shell script to reload the process.",
								Attributes: map[string]schema.Attribute{
									"batch_parameters_template": schema.StringAttribute{
										Description:         "BatchParametersTemplate provides an optional Go template string to format the batch input data passed into the STDIN of the script when 'batchReload' is set to 'True'. The template uses the updated parameters' key-value pairs, accessible via the '$' variable. This allows for custom formatting of the input data. Example template:  '''yaml batchParametersTemplate: |- {{- range $pKey, $pValue := $ }} {{ printf '%s:%s' $pKey $pValue }} {{- end }} '''  This example generates batch input data in a key:value format, sorted by keys. ''' key1:value1 key2:value2 key3:value3 '''  If not specified, the default format is key=value, sorted by keys, for each updated parameter. ''' key1=value1 key2=value2 key3=value3 '''",
										MarkdownDescription: "BatchParametersTemplate provides an optional Go template string to format the batch input data passed into the STDIN of the script when 'batchReload' is set to 'True'. The template uses the updated parameters' key-value pairs, accessible via the '$' variable. This allows for custom formatting of the input data. Example template:  '''yaml batchParametersTemplate: |- {{- range $pKey, $pValue := $ }} {{ printf '%s:%s' $pKey $pValue }} {{- end }} '''  This example generates batch input data in a key:value format, sorted by keys. ''' key1:value1 key2:value2 key3:value3 '''  If not specified, the default format is key=value, sorted by keys, for each updated parameter. ''' key1=value1 key2=value2 key3=value3 '''",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"batch_reload": schema.BoolAttribute{
										Description:         "Specifies whether to process dynamic parameter updates individually or collectively in a batch:  - Set to 'True' to execute all parameter changes in one batch reload action. - Set to 'False' to execute a reload action for each individual parameter change. The default behavior, if not specified, is 'False'.",
										MarkdownDescription: "Specifies whether to process dynamic parameter updates individually or collectively in a batch:  - Set to 'True' to execute all parameter changes in one batch reload action. - Set to 'False' to execute a reload action for each individual parameter change. The default behavior, if not specified, is 'False'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"command": schema.ListAttribute{
										Description:         "Specifies the command to execute in order to reload the process. It should be a valid shell command.",
										MarkdownDescription: "Specifies the command to execute in order to reload the process. It should be a valid shell command.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"sync": schema.BoolAttribute{
										Description:         "Determines whether parameter updates should be synchronized with the config manager. Specifies the controller's reload strategy:  - If set to 'True', the controller executes the reload action in synchronous mode, pausing execution until the reload completes. - If set to 'False', the controller executes the reload action in asynchronous mode, updating the ConfigMap without waiting for the reload process to finish.",
										MarkdownDescription: "Determines whether parameter updates should be synchronized with the config manager. Specifies the controller's reload strategy:  - If set to 'True', the controller executes the reload action in synchronous mode, pausing execution until the reload completes. - If set to 'False', the controller executes the reload action in asynchronous mode, updating the ConfigMap without waiting for the reload process to finish.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tpl_script_trigger": schema.SingleNestedAttribute{
								Description:         "Enables reloading process using a Go template script.",
								MarkdownDescription: "Enables reloading process using a Go template script.",
								Attributes: map[string]schema.Attribute{
									"namespace": schema.StringAttribute{
										Description:         "Specifies the namespace where the referenced tpl script ConfigMap in. If left empty, by default in the 'default' namespace.",
										MarkdownDescription: "Specifies the namespace where the referenced tpl script ConfigMap in. If left empty, by default in the 'default' namespace.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtMost(63),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`), ""),
										},
									},

									"script_config_map_ref": schema.StringAttribute{
										Description:         "Specifies the reference to the ConfigMap that contains the script to be executed for reload.",
										MarkdownDescription: "Specifies the reference to the ConfigMap that contains the script to be executed for reload.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"sync": schema.BoolAttribute{
										Description:         "Determines whether parameter updates should be synchronized with the config manager. Specifies the controller's reload strategy:  - If set to 'True', the controller executes the reload action in synchronous mode, pausing execution until the reload completes. - If set to 'False', the controller executes the reload action in asynchronous mode, updating the ConfigMap without waiting for the reload process to finish.",
										MarkdownDescription: "Determines whether parameter updates should be synchronized with the config manager. Specifies the controller's reload strategy:  - If set to 'True', the controller executes the reload action in synchronous mode, pausing execution until the reload completes. - If set to 'False', the controller executes the reload action in asynchronous mode, updating the ConfigMap without waiting for the reload process to finish.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"unix_signal_trigger": schema.SingleNestedAttribute{
								Description:         "Used to trigger a reload by sending a specific Unix signal to the process.",
								MarkdownDescription: "Used to trigger a reload by sending a specific Unix signal to the process.",
								Attributes: map[string]schema.Attribute{
									"process_name": schema.StringAttribute{
										Description:         "Identifies the name of the process to which the Unix signal will be sent.",
										MarkdownDescription: "Identifies the name of the process to which the Unix signal will be sent.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"signal": schema.StringAttribute{
										Description:         "Specifies a valid Unix signal to be sent. For a comprehensive list of all Unix signals, see: ../../pkg/configuration/configmap/handler.go:allUnixSignals",
										MarkdownDescription: "Specifies a valid Unix signal to be sent. For a comprehensive list of all Unix signals, see: ../../pkg/configuration/configmap/handler.go:allUnixSignals",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("SIGHUP", "SIGINT", "SIGQUIT", "SIGILL", "SIGTRAP", "SIGABRT", "SIGBUS", "SIGFPE", "SIGKILL", "SIGUSR1", "SIGSEGV", "SIGUSR2", "SIGPIPE", "SIGALRM", "SIGTERM", "SIGSTKFLT", "SIGCHLD", "SIGCONT", "SIGSTOP", "SIGTSTP", "SIGTTIN", "SIGTTOU", "SIGURG", "SIGXCPU", "SIGXFSZ", "SIGVTALRM", "SIGPROF", "SIGWINCH", "SIGIO", "SIGPWR", "SIGSYS"),
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

					"dynamic_reload_selector": schema.SingleNestedAttribute{
						Description:         "Used to match labels on the pod to determine whether a dynamic reload should be performed.  In some scenarios, only specific pods (e.g., primary replicas) need to undergo a dynamic reload. The 'dynamicReloadSelector' allows you to specify label selectors to target the desired pods for the reload process.  If the 'dynamicReloadSelector' is not specified or is nil, all pods managed by the workload will be considered for the dynamic reload.",
						MarkdownDescription: "Used to match labels on the pod to determine whether a dynamic reload should be performed.  In some scenarios, only specific pods (e.g., primary replicas) need to undergo a dynamic reload. The 'dynamicReloadSelector' allows you to specify label selectors to target the desired pods for the reload process.  If the 'dynamicReloadSelector' is not specified or is nil, all pods managed by the workload will be considered for the dynamic reload.",
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

					"formatter_config": schema.SingleNestedAttribute{
						Description:         "Specifies the format of the configuration file and any associated parameters that are specific to the chosen format. Supported formats include 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties', and 'toml'.  Each format may have its own set of parameters that can be configured. For instance, when using the 'ini' format, you can specify the section name.  Example: ''' formatterConfig: format: ini iniConfig: sectionName: mysqld '''",
						MarkdownDescription: "Specifies the format of the configuration file and any associated parameters that are specific to the chosen format. Supported formats include 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties', and 'toml'.  Each format may have its own set of parameters that can be configured. For instance, when using the 'ini' format, you can specify the section name.  Example: ''' formatterConfig: format: ini iniConfig: sectionName: mysqld '''",
						Attributes: map[string]schema.Attribute{
							"format": schema.StringAttribute{
								Description:         "The config file format. Valid values are 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties' and 'toml'. Each format has its own characteristics and use cases.  - ini: is a text-based content with a structure and syntax comprising key–value pairs for properties, reference wiki: https://en.wikipedia.org/wiki/INI_file - xml: refers to wiki: https://en.wikipedia.org/wiki/XML - yaml: supports for complex data types and structures. - json: refers to wiki: https://en.wikipedia.org/wiki/JSON - hcl: The HashiCorp Configuration Language (HCL) is a configuration language authored by HashiCorp, reference url: https://www.linode.com/docs/guides/introduction-to-hcl/ - dotenv: is a plain text file with simple key–value pairs, reference wiki: https://en.wikipedia.org/wiki/Configuration_file#MS-DOS - properties: a file extension mainly used in Java, reference wiki: https://en.wikipedia.org/wiki/.properties - toml: refers to wiki: https://en.wikipedia.org/wiki/TOML - props-plus: a file extension mainly used in Java, supports CamelCase(e.g: brokerMaxConnectionsPerIp)",
								MarkdownDescription: "The config file format. Valid values are 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties' and 'toml'. Each format has its own characteristics and use cases.  - ini: is a text-based content with a structure and syntax comprising key–value pairs for properties, reference wiki: https://en.wikipedia.org/wiki/INI_file - xml: refers to wiki: https://en.wikipedia.org/wiki/XML - yaml: supports for complex data types and structures. - json: refers to wiki: https://en.wikipedia.org/wiki/JSON - hcl: The HashiCorp Configuration Language (HCL) is a configuration language authored by HashiCorp, reference url: https://www.linode.com/docs/guides/introduction-to-hcl/ - dotenv: is a plain text file with simple key–value pairs, reference wiki: https://en.wikipedia.org/wiki/Configuration_file#MS-DOS - properties: a file extension mainly used in Java, reference wiki: https://en.wikipedia.org/wiki/.properties - toml: refers to wiki: https://en.wikipedia.org/wiki/TOML - props-plus: a file extension mainly used in Java, supports CamelCase(e.g: brokerMaxConnectionsPerIp)",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("xml", "ini", "yaml", "json", "hcl", "dotenv", "toml", "properties", "redis", "props-plus"),
								},
							},

							"ini_config": schema.SingleNestedAttribute{
								Description:         "Holds options specific to the 'ini' file format.",
								MarkdownDescription: "Holds options specific to the 'ini' file format.",
								Attributes: map[string]schema.Attribute{
									"section_name": schema.StringAttribute{
										Description:         "A string that describes the name of the ini section.",
										MarkdownDescription: "A string that describes the name of the ini section.",
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

					"immutable_parameters": schema.ListAttribute{
						Description:         "Lists the parameters that cannot be modified once set. Attempting to change any of these parameters will be ignored.",
						MarkdownDescription: "Lists the parameters that cannot be modified once set. Attempting to change any of these parameters will be ignored.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"reload_tools_image": schema.SingleNestedAttribute{
						Description:         "Specifies the tools container image used by ShellTrigger for dynamic reload. If the dynamic reload action is triggered by a ShellTrigger, this field is required. This image must contain all necessary tools for executing the ShellTrigger scripts.  Usually the specified image is referenced by the init container, which is then responsible for copy the tools from the image to a bin volume. This ensures that the tools are available to the 'config-manager' sidecar.",
						MarkdownDescription: "Specifies the tools container image used by ShellTrigger for dynamic reload. If the dynamic reload action is triggered by a ShellTrigger, this field is required. This image must contain all necessary tools for executing the ShellTrigger scripts.  Usually the specified image is referenced by the init container, which is then responsible for copy the tools from the image to a bin volume. This ensures that the tools are available to the 'config-manager' sidecar.",
						Attributes: map[string]schema.Attribute{
							"mount_point": schema.StringAttribute{
								Description:         "Specifies the directory path in the container where the tools-related files are to be copied. This field is typically used with an emptyDir volume to ensure a temporary, empty directory is provided at pod creation.",
								MarkdownDescription: "Specifies the directory path in the container where the tools-related files are to be copied. This field is typically used with an emptyDir volume to ensure a temporary, empty directory is provided at pod creation.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(128),
								},
							},

							"tool_configs": schema.ListNestedAttribute{
								Description:         "Specifies a list of settings of init containers that prepare tools for dynamic reload.",
								MarkdownDescription: "Specifies a list of settings of init containers that prepare tools for dynamic reload.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"command": schema.ListAttribute{
											Description:         "Specifies the command to be executed by the init container.",
											MarkdownDescription: "Specifies the command to be executed by the init container.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "Specifies the tool container image.",
											MarkdownDescription: "Specifies the tool container image.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the init container.",
											MarkdownDescription: "Specifies the name of the init container.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
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

					"script_configs": schema.ListNestedAttribute{
						Description:         "A list of ScriptConfig Object.  Each ScriptConfig object specifies a ConfigMap that contains script files that should be mounted inside the pod. The scripts are mounted as volumes and can be referenced and executed by the dynamic reload and DownwardAction to perform specific tasks or configurations.",
						MarkdownDescription: "A list of ScriptConfig Object.  Each ScriptConfig object specifies a ConfigMap that contains script files that should be mounted inside the pod. The scripts are mounted as volumes and can be referenced and executed by the dynamic reload and DownwardAction to perform specific tasks or configurations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"namespace": schema.StringAttribute{
									Description:         "Specifies the namespace where the referenced tpl script ConfigMap in. If left empty, by default in the 'default' namespace.",
									MarkdownDescription: "Specifies the namespace where the referenced tpl script ConfigMap in. If left empty, by default in the 'default' namespace.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`), ""),
									},
								},

								"script_config_map_ref": schema.StringAttribute{
									Description:         "Specifies the reference to the ConfigMap that contains the script to be executed for reload.",
									MarkdownDescription: "Specifies the reference to the ConfigMap that contains the script to be executed for reload.",
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

					"static_parameters": schema.ListAttribute{
						Description:         "List static parameters. Modifications to any of these parameters require a restart of the process to take effect.",
						MarkdownDescription: "List static parameters. Modifications to any of these parameters require a restart of the process to take effect.",
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
	}
}

func (r *AppsKubeblocksIoConfigConstraintV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_config_constraint_v1beta1_manifest")

	var model AppsKubeblocksIoConfigConstraintV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1beta1")
	model.Kind = pointer.String("ConfigConstraint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
