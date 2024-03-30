/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &AppsKubeblocksIoConfigConstraintV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoConfigConstraintV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoConfigConstraintV1Alpha1Manifest{}
}

type AppsKubeblocksIoConfigConstraintV1Alpha1Manifest struct{}

type AppsKubeblocksIoConfigConstraintV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CfgSchemaTopLevelName *string `tfsdk:"cfg_schema_top_level_name" json:"cfgSchemaTopLevelName,omitempty"`
		ConfigurationSchema   *struct {
			Cue    *string            `tfsdk:"cue" json:"cue,omitempty"`
			Schema *map[string]string `tfsdk:"schema" json:"schema,omitempty"`
		} `tfsdk:"configuration_schema" json:"configurationSchema,omitempty"`
		DownwardAPIOptions *[]struct {
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
		} `tfsdk:"downward_api_options" json:"downwardAPIOptions,omitempty"`
		DynamicActionCanBeMerged       *bool     `tfsdk:"dynamic_action_can_be_merged" json:"dynamicActionCanBeMerged,omitempty"`
		DynamicParameterSelectedPolicy *string   `tfsdk:"dynamic_parameter_selected_policy" json:"dynamicParameterSelectedPolicy,omitempty"`
		DynamicParameters              *[]string `tfsdk:"dynamic_parameters" json:"dynamicParameters,omitempty"`
		FormatterConfig                *struct {
			Format    *string `tfsdk:"format" json:"format,omitempty"`
			IniConfig *struct {
				SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
			} `tfsdk:"ini_config" json:"iniConfig,omitempty"`
		} `tfsdk:"formatter_config" json:"formatterConfig,omitempty"`
		ImmutableParameters *[]string `tfsdk:"immutable_parameters" json:"immutableParameters,omitempty"`
		ReloadOptions       *struct {
			AutoTrigger *struct {
				ProcessName *string `tfsdk:"process_name" json:"processName,omitempty"`
			} `tfsdk:"auto_trigger" json:"autoTrigger,omitempty"`
			ShellTrigger *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Sync    *bool     `tfsdk:"sync" json:"sync,omitempty"`
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
		} `tfsdk:"reload_options" json:"reloadOptions,omitempty"`
		ScriptConfigs *[]struct {
			Namespace          *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ScriptConfigMapRef *string `tfsdk:"script_config_map_ref" json:"scriptConfigMapRef,omitempty"`
		} `tfsdk:"script_configs" json:"scriptConfigs,omitempty"`
		Selector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		StaticParameters *[]string `tfsdk:"static_parameters" json:"staticParameters,omitempty"`
		ToolsImageSpec   *struct {
			MountPoint  *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
			ToolConfigs *[]struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Image   *string   `tfsdk:"image" json:"image,omitempty"`
				Name    *string   `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"tool_configs" json:"toolConfigs,omitempty"`
		} `tfsdk:"tools_image_spec" json:"toolsImageSpec,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoConfigConstraintV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_config_constraint_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoConfigConstraintV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConfigConstraint is the Schema for the configconstraint API",
		MarkdownDescription: "ConfigConstraint is the Schema for the configconstraint API",
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
					"cfg_schema_top_level_name": schema.StringAttribute{
						Description:         "Top level key used to get the cue rules to validate the config file. It must exist in 'ConfigSchema' TODO (refactored to ConfigSchemaTopLevelKey)",
						MarkdownDescription: "Top level key used to get the cue rules to validate the config file. It must exist in 'ConfigSchema' TODO (refactored to ConfigSchemaTopLevelKey)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"configuration_schema": schema.SingleNestedAttribute{
						Description:         "List constraints rules for each config parameters. TODO (refactored to ConfigSchema)",
						MarkdownDescription: "List constraints rules for each config parameters. TODO (refactored to ConfigSchema)",
						Attributes: map[string]schema.Attribute{
							"cue": schema.StringAttribute{
								Description:         "Enables providers to verify user configurations using the CUE language.",
								MarkdownDescription: "Enables providers to verify user configurations using the CUE language.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"schema": schema.MapAttribute{
								Description:         "Transforms the schema from CUE to json for further OpenAPI validation TODO (refactored to SchemaInJson)",
								MarkdownDescription: "Transforms the schema from CUE to json for further OpenAPI validation TODO (refactored to SchemaInJson)",
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

					"downward_api_options": schema.ListNestedAttribute{
						Description:         "A set of actions for regenerating local configs.  It works when: - different engine roles have different config, such as redis primary & secondary - after a role switch, the local config will be regenerated with the help of DownwardActions TODO (refactored to DownwardActions)",
						MarkdownDescription: "A set of actions for regenerating local configs.  It works when: - different engine roles have different config, such as redis primary & secondary - after a role switch, the local config will be regenerated with the help of DownwardActions TODO (refactored to DownwardActions)",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"command": schema.ListAttribute{
									Description:         "The command used to execute for the downward API.",
									MarkdownDescription: "The command used to execute for the downward API.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"items": schema.ListNestedAttribute{
									Description:         "Represents a list of downward API volume files.",
									MarkdownDescription: "Represents a list of downward API volume files.",
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
									Description:         "Specifies the mount point of the scripts file.",
									MarkdownDescription: "Specifies the mount point of the scripts file.",
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
						Description:         "Indicates the dynamic reload action and restart action can be merged to a restart action.  When a batch of parameters updates incur both restart & dynamic reload, it works as: - set to true, the two actions merged to only one restart action - set to false, the two actions cannot be merged, the actions executed in order [dynamic reload, restart]",
						MarkdownDescription: "Indicates the dynamic reload action and restart action can be merged to a restart action.  When a batch of parameters updates incur both restart & dynamic reload, it works as: - set to true, the two actions merged to only one restart action - set to false, the two actions cannot be merged, the actions executed in order [dynamic reload, restart]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dynamic_parameter_selected_policy": schema.StringAttribute{
						Description:         "Specifies the policy for selecting the parameters of dynamic reload actions.",
						MarkdownDescription: "Specifies the policy for selecting the parameters of dynamic reload actions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("all", "dynamic"),
						},
					},

					"dynamic_parameters": schema.ListAttribute{
						Description:         "A list of DynamicParameter. Modifications of dynamic parameters trigger a reload action without process restart.",
						MarkdownDescription: "A list of DynamicParameter. Modifications of dynamic parameters trigger a reload action without process restart.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"formatter_config": schema.SingleNestedAttribute{
						Description:         "Describes the format of the config file. The controller works as follows: 1. Parse the config file 2. Get the modified parameters 3. Trigger the corresponding action",
						MarkdownDescription: "Describes the format of the config file. The controller works as follows: 1. Parse the config file 2. Get the modified parameters 3. Trigger the corresponding action",
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
								Description:         "A pointer to an IniConfig struct that holds the ini options.",
								MarkdownDescription: "A pointer to an IniConfig struct that holds the ini options.",
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
						Description:         "Describes parameters that are prohibited to do any modifications.",
						MarkdownDescription: "Describes parameters that are prohibited to do any modifications.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"reload_options": schema.SingleNestedAttribute{
						Description:         "Specifies the dynamic reload actions supported by the engine. If set, the controller call the scripts defined in the actions for a dynamic parameter upgrade. The actions are called only when the modified parameter is defined in dynamicParameters part && DynamicReloadActions != nil TODO (refactored to DynamicReloadActions)",
						MarkdownDescription: "Specifies the dynamic reload actions supported by the engine. If set, the controller call the scripts defined in the actions for a dynamic parameter upgrade. The actions are called only when the modified parameter is defined in dynamicParameters part && DynamicReloadActions != nil TODO (refactored to DynamicReloadActions)",
						Attributes: map[string]schema.Attribute{
							"auto_trigger": schema.SingleNestedAttribute{
								Description:         "Used to automatically perform the reload command when conditions are met.",
								MarkdownDescription: "Used to automatically perform the reload command when conditions are met.",
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
								Description:         "Used to perform the reload command in shell script.",
								MarkdownDescription: "Used to perform the reload command in shell script.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
										Description:         "Specifies the list of commands for reload.",
										MarkdownDescription: "Specifies the list of commands for reload.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"sync": schema.BoolAttribute{
										Description:         "Specifies whether to synchronize updates parameters to the config manager. Specifies two ways of controller to reload the parameter: - set to 'True', execute the reload action in sync mode, wait for the completion of reload - set to 'False', execute the reload action in async mode, just update the 'Configmap', no need to wait",
										MarkdownDescription: "Specifies whether to synchronize updates parameters to the config manager. Specifies two ways of controller to reload the parameter: - set to 'True', execute the reload action in sync mode, wait for the completion of reload - set to 'False', execute the reload action in async mode, just update the 'Configmap', no need to wait",
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
								Description:         "Used to perform the reload command by Go template script.",
								MarkdownDescription: "Used to perform the reload command by Go template script.",
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
										Description:         "Specifies whether to synchronize updates parameters to the config manager. Specifies two ways of controller to reload the parameter: - set to 'True', execute the reload action in sync mode, wait for the completion of reload - set to 'False', execute the reload action in async mode, just update the 'Configmap', no need to wait",
										MarkdownDescription: "Specifies whether to synchronize updates parameters to the config manager. Specifies two ways of controller to reload the parameter: - set to 'True', execute the reload action in sync mode, wait for the completion of reload - set to 'False', execute the reload action in async mode, just update the 'Configmap', no need to wait",
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
								Description:         "Used to trigger a reload by sending a Unix signal to the process.",
								MarkdownDescription: "Used to trigger a reload by sending a Unix signal to the process.",
								Attributes: map[string]schema.Attribute{
									"process_name": schema.StringAttribute{
										Description:         "Represents the name of the process that the Unix signal sent to.",
										MarkdownDescription: "Represents the name of the process that the Unix signal sent to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"signal": schema.StringAttribute{
										Description:         "Represents a valid Unix signal. Refer to the following URL for a list of all Unix signals: ../../pkg/configuration/configmap/handler.go:allUnixSignals",
										MarkdownDescription: "Represents a valid Unix signal. Refer to the following URL for a list of all Unix signals: ../../pkg/configuration/configmap/handler.go:allUnixSignals",
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

					"script_configs": schema.ListNestedAttribute{
						Description:         "A list of ScriptConfig used by the actions defined in dynamic reload and downward actions.",
						MarkdownDescription: "A list of ScriptConfig used by the actions defined in dynamic reload and downward actions.",
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

					"selector": schema.SingleNestedAttribute{
						Description:         "Used to match labels on the pod to do a dynamic reload TODO (refactored to DynamicReloadSelector)",
						MarkdownDescription: "Used to match labels on the pod to do a dynamic reload TODO (refactored to DynamicReloadSelector)",
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

					"static_parameters": schema.ListAttribute{
						Description:         "A list of StaticParameter. Modifications of static parameters trigger a process restart.",
						MarkdownDescription: "A list of StaticParameter. Modifications of static parameters trigger a process restart.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tools_image_spec": schema.SingleNestedAttribute{
						Description:         "Tools used by the dynamic reload actions. Usually it is referenced by the 'init container' for 'cp' it to a binary volume. TODO (refactored to ReloadToolsImage)",
						MarkdownDescription: "Tools used by the dynamic reload actions. Usually it is referenced by the 'init container' for 'cp' it to a binary volume. TODO (refactored to ReloadToolsImage)",
						Attributes: map[string]schema.Attribute{
							"mount_point": schema.StringAttribute{
								Description:         "Represents the point where the scripts file will be mounted.",
								MarkdownDescription: "Represents the point where the scripts file will be mounted.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(128),
								},
							},

							"tool_configs": schema.ListNestedAttribute{
								Description:         "Used to configure the initialization container.",
								MarkdownDescription: "Used to configure the initialization container.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"command": schema.ListAttribute{
											Description:         "Commands to be executed when init containers.",
											MarkdownDescription: "Commands to be executed when init containers.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "Represents the url of the tool container image.",
											MarkdownDescription: "Represents the url of the tool container image.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the initContainer.",
											MarkdownDescription: "Specifies the name of the initContainer.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoConfigConstraintV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_config_constraint_v1alpha1_manifest")

	var model AppsKubeblocksIoConfigConstraintV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("ConfigConstraint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
