/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

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
	_ datasource.DataSource = &CamelApacheOrgIntegrationKitV1Manifest{}
)

func NewCamelApacheOrgIntegrationKitV1Manifest() datasource.DataSource {
	return &CamelApacheOrgIntegrationKitV1Manifest{}
}

type CamelApacheOrgIntegrationKitV1Manifest struct{}

type CamelApacheOrgIntegrationKitV1ManifestData struct {
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
		Capabilities  *[]string `tfsdk:"capabilities" json:"capabilities,omitempty"`
		Configuration *[]struct {
			Type  *string `tfsdk:"type" json:"type,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"configuration" json:"configuration,omitempty"`
		Dependencies *[]string `tfsdk:"dependencies" json:"dependencies,omitempty"`
		Image        *string   `tfsdk:"image" json:"image,omitempty"`
		Profile      *string   `tfsdk:"profile" json:"profile,omitempty"`
		Repositories *[]string `tfsdk:"repositories" json:"repositories,omitempty"`
		Sources      *[]struct {
			Compression    *bool     `tfsdk:"compression" json:"compression,omitempty"`
			Content        *string   `tfsdk:"content" json:"content,omitempty"`
			ContentKey     *string   `tfsdk:"content_key" json:"contentKey,omitempty"`
			ContentRef     *string   `tfsdk:"content_ref" json:"contentRef,omitempty"`
			ContentType    *string   `tfsdk:"content_type" json:"contentType,omitempty"`
			From_kamelet   *bool     `tfsdk:"from_kamelet" json:"from-kamelet,omitempty"`
			Interceptors   *[]string `tfsdk:"interceptors" json:"interceptors,omitempty"`
			Language       *string   `tfsdk:"language" json:"language,omitempty"`
			Loader         *string   `tfsdk:"loader" json:"loader,omitempty"`
			Name           *string   `tfsdk:"name" json:"name,omitempty"`
			Path           *string   `tfsdk:"path" json:"path,omitempty"`
			Property_names *[]string `tfsdk:"property_names" json:"property-names,omitempty"`
			RawContent     *string   `tfsdk:"raw_content" json:"rawContent,omitempty"`
			Type           *string   `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"sources" json:"sources,omitempty"`
		Traits *struct {
			Addons  *map[string]string `tfsdk:"addons" json:"addons,omitempty"`
			Builder *struct {
				Annotations           *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				BaseImage             *string            `tfsdk:"base_image" json:"baseImage,omitempty"`
				Configuration         *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled               *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				IncrementalImageBuild *bool              `tfsdk:"incremental_image_build" json:"incrementalImageBuild,omitempty"`
				LimitCPU              *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
				LimitMemory           *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
				MavenProfiles         *[]string          `tfsdk:"maven_profiles" json:"mavenProfiles,omitempty"`
				NodeSelector          *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				OrderStrategy         *string            `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
				Platforms             *[]string          `tfsdk:"platforms" json:"platforms,omitempty"`
				Properties            *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RequestCPU            *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
				RequestMemory         *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
				Strategy              *string            `tfsdk:"strategy" json:"strategy,omitempty"`
				Tasks                 *[]string          `tfsdk:"tasks" json:"tasks,omitempty"`
				TasksFilter           *string            `tfsdk:"tasks_filter" json:"tasksFilter,omitempty"`
				TasksLimitCPU         *[]string          `tfsdk:"tasks_limit_cpu" json:"tasksLimitCPU,omitempty"`
				TasksLimitMemory      *[]string          `tfsdk:"tasks_limit_memory" json:"tasksLimitMemory,omitempty"`
				TasksRequestCPU       *[]string          `tfsdk:"tasks_request_cpu" json:"tasksRequestCPU,omitempty"`
				TasksRequestMemory    *[]string          `tfsdk:"tasks_request_memory" json:"tasksRequestMemory,omitempty"`
				Verbose               *bool              `tfsdk:"verbose" json:"verbose,omitempty"`
			} `tfsdk:"builder" json:"builder,omitempty"`
			Camel *struct {
				Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Properties     *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RuntimeVersion *string            `tfsdk:"runtime_version" json:"runtimeVersion,omitempty"`
			} `tfsdk:"camel" json:"camel,omitempty"`
			Quarkus *struct {
				BuildMode          *[]string          `tfsdk:"build_mode" json:"buildMode,omitempty"`
				Configuration      *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled            *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				NativeBaseImage    *string            `tfsdk:"native_base_image" json:"nativeBaseImage,omitempty"`
				NativeBuilderImage *string            `tfsdk:"native_builder_image" json:"nativeBuilderImage,omitempty"`
				PackageTypes       *[]string          `tfsdk:"package_types" json:"packageTypes,omitempty"`
			} `tfsdk:"quarkus" json:"quarkus,omitempty"`
			Registry *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"registry" json:"registry,omitempty"`
		} `tfsdk:"traits" json:"traits,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgIntegrationKitV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_integration_kit_v1_manifest"
}

func (r *CamelApacheOrgIntegrationKitV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.",
		MarkdownDescription: "IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.",
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
				Description:         "the desired configuration",
				MarkdownDescription: "the desired configuration",
				Attributes: map[string]schema.Attribute{
					"capabilities": schema.ListAttribute{
						Description:         "features offered by the IntegrationKit",
						MarkdownDescription: "features offered by the IntegrationKit",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"configuration": schema.ListNestedAttribute{
						Description:         "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes configuration used by the kit",
						MarkdownDescription: "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes configuration used by the kit",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
									MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
									MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",
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

					"dependencies": schema.ListAttribute{
						Description:         "a list of Camel dependecies used by this kit",
						MarkdownDescription: "a list of Camel dependecies used by this kit",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "the container image as identified in the container registry",
						MarkdownDescription: "the container image as identified in the container registry",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"profile": schema.StringAttribute{
						Description:         "the profile which is expected by this kit",
						MarkdownDescription: "the profile which is expected by this kit",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"repositories": schema.ListAttribute{
						Description:         "Maven repositories that can be used by the kit",
						MarkdownDescription: "Maven repositories that can be used by the kit",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sources": schema.ListNestedAttribute{
						Description:         "the sources to add at build time",
						MarkdownDescription: "the sources to add at build time",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"compression": schema.BoolAttribute{
									Description:         "if the content is compressed (base64 encrypted)",
									MarkdownDescription: "if the content is compressed (base64 encrypted)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"content": schema.StringAttribute{
									Description:         "the source code (plain text)",
									MarkdownDescription: "the source code (plain text)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"content_key": schema.StringAttribute{
									Description:         "the confimap key holding the source content",
									MarkdownDescription: "the confimap key holding the source content",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"content_ref": schema.StringAttribute{
									Description:         "the confimap reference holding the source content",
									MarkdownDescription: "the confimap reference holding the source content",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"content_type": schema.StringAttribute{
									Description:         "the content type (tipically text or binary)",
									MarkdownDescription: "the content type (tipically text or binary)",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"from_kamelet": schema.BoolAttribute{
									Description:         "True if the spec is generated from a Kamelet",
									MarkdownDescription: "True if the spec is generated from a Kamelet",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interceptors": schema.ListAttribute{
									Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"language": schema.StringAttribute{
									Description:         "specify which is the language (Camel DSL) used to interpret this source code",
									MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"loader": schema.StringAttribute{
									Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "the name of the specification",
									MarkdownDescription: "the name of the specification",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "the path where the file is stored",
									MarkdownDescription: "the path where the file is stored",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"property_names": schema.ListAttribute{
									Description:         "List of property names defined in the source (e.g. if type is 'template')",
									MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"raw_content": schema.StringAttribute{
									Description:         "the source code (binary)",
									MarkdownDescription: "the source code (binary)",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.Base64Validator(),
									},
								},

								"type": schema.StringAttribute{
									Description:         "Type defines the kind of source described by this object",
									MarkdownDescription: "Type defines the kind of source described by this object",
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

					"traits": schema.SingleNestedAttribute{
						Description:         "traits that the kit will execute",
						MarkdownDescription: "traits that the kit will execute",
						Attributes: map[string]schema.Attribute{
							"addons": schema.MapAttribute{
								Description:         "The collection of addon trait configurations",
								MarkdownDescription: "The collection of addon trait configurations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"builder": schema.SingleNestedAttribute{
								Description:         "The builder trait is internally used to determine the best strategy to build and configure IntegrationKits.",
								MarkdownDescription: "The builder trait is internally used to determine the best strategy to build and configure IntegrationKits.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "When using 'pod' strategy, annotation to use for the builder pod.",
										MarkdownDescription: "When using 'pod' strategy, annotation to use for the builder pod.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"base_image": schema.StringAttribute{
										Description:         "Specify a base image",
										MarkdownDescription: "Specify a base image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Deprecated: no longer in use.",
										MarkdownDescription: "Deprecated: no longer in use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"incremental_image_build": schema.BoolAttribute{
										Description:         "Use the incremental image build option, to reuse existing containers (default 'true')",
										MarkdownDescription: "Use the incremental image build option, to reuse existing containers (default 'true')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"limit_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"limit_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maven_profiles": schema.ListAttribute{
										Description:         "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										MarkdownDescription: "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_selector": schema.MapAttribute{
										Description:         "Defines a set of nodes the builder pod is eligible to be scheduled on, based on labels on the node.",
										MarkdownDescription: "Defines a set of nodes the builder pod is eligible to be scheduled on, based on labels on the node.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"order_strategy": schema.StringAttribute{
										Description:         "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default 'sequential')",
										MarkdownDescription: "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default 'sequential')",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("dependencies", "fifo", "sequential"),
										},
									},

									"platforms": schema.ListAttribute{
										Description:         "The list of manifest platforms to use to build a container image (default 'linux/amd64').",
										MarkdownDescription: "The list of manifest platforms to use to build a container image (default 'linux/amd64').",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the build task",
										MarkdownDescription: "A list of properties to be provided to the build task",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"strategy": schema.StringAttribute{
										Description:         "The strategy to use, either 'pod' or 'routine' (default 'routine')",
										MarkdownDescription: "The strategy to use, either 'pod' or 'routine' (default 'routine')",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("pod", "routine"),
										},
									},

									"tasks": schema.ListAttribute{
										Description:         "A list of tasks to be executed (available only when using 'pod' strategy) with format '<name>;<container-image>;<container-command>'.",
										MarkdownDescription: "A list of tasks to be executed (available only when using 'pod' strategy) with format '<name>;<container-image>;<container-command>'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tasks_filter": schema.StringAttribute{
										Description:         "A list of tasks sorted by the order of execution in a csv format, ie, '<taskName1>,<taskName2>,...'. Mind that you must include also the operator tasks ('builder', 'quarkus-native', 'package', 'jib', 's2i') if you need to execute them. Useful only with 'pod' strategy.",
										MarkdownDescription: "A list of tasks sorted by the order of execution in a csv format, ie, '<taskName1>,<taskName2>,...'. Mind that you must include also the operator tasks ('builder', 'quarkus-native', 'package', 'jib', 's2i') if you need to execute them. Useful only with 'pod' strategy.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tasks_limit_cpu": schema.ListAttribute{
										Description:         "A list of limit cpu configuration for the specific task with format '<task-name>:<limit-cpu-conf>'.",
										MarkdownDescription: "A list of limit cpu configuration for the specific task with format '<task-name>:<limit-cpu-conf>'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tasks_limit_memory": schema.ListAttribute{
										Description:         "A list of limit memory configuration for the specific task with format '<task-name>:<limit-memory-conf>'.",
										MarkdownDescription: "A list of limit memory configuration for the specific task with format '<task-name>:<limit-memory-conf>'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tasks_request_cpu": schema.ListAttribute{
										Description:         "A list of request cpu configuration for the specific task with format '<task-name>:<request-cpu-conf>'.",
										MarkdownDescription: "A list of request cpu configuration for the specific task with format '<task-name>:<request-cpu-conf>'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tasks_request_memory": schema.ListAttribute{
										Description:         "A list of request memory configuration for the specific task with format '<task-name>:<request-memory-conf>'.",
										MarkdownDescription: "A list of request memory configuration for the specific task with format '<task-name>:<request-memory-conf>'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"verbose": schema.BoolAttribute{
										Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod). Deprecated no longer in use",
										MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod). Deprecated no longer in use",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"camel": schema.SingleNestedAttribute{
								Description:         "The Camel trait sets up Camel configuration.",
								MarkdownDescription: "The Camel trait sets up Camel configuration.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Deprecated: no longer in use.",
										MarkdownDescription: "Deprecated: no longer in use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the Integration runtime",
										MarkdownDescription: "A list of properties to be provided to the Integration runtime",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"runtime_version": schema.StringAttribute{
										Description:         "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform. You can use a fixed version (for example '3.2.3') or a semantic version (for example '3.x') which will try to resolve to the best matching Catalog existing on the cluster.",
										MarkdownDescription: "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform. You can use a fixed version (for example '3.2.3') or a semantic version (for example '3.x') which will try to resolve to the best matching Catalog existing on the cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"quarkus": schema.SingleNestedAttribute{
								Description:         "The Quarkus trait configures the Quarkus runtime. It's enabled by default. NOTE: Compiling to a native executable, requires at least 4GiB of memory, so the Pod running the native build must have enough memory available.",
								MarkdownDescription: "The Quarkus trait configures the Quarkus runtime. It's enabled by default. NOTE: Compiling to a native executable, requires at least 4GiB of memory, so the Pod running the native build must have enough memory available.",
								Attributes: map[string]schema.Attribute{
									"build_mode": schema.ListAttribute{
										Description:         "The Quarkus mode to run: either 'jvm' or 'native' (default 'jvm'). In case both 'jvm' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'jvm' one once ready.",
										MarkdownDescription: "The Quarkus mode to run: either 'jvm' or 'native' (default 'jvm'). In case both 'jvm' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'jvm' one once ready.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Deprecated: no longer in use.",
										MarkdownDescription: "Deprecated: no longer in use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"native_base_image": schema.StringAttribute{
										Description:         "The base image to use when running a native build (default 'quay.io/quarkus/quarkus-micro-image:2.0')",
										MarkdownDescription: "The base image to use when running a native build (default 'quay.io/quarkus/quarkus-micro-image:2.0')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"native_builder_image": schema.StringAttribute{
										Description:         "The image containing the tooling required for a native build (by default it will use the one provided in the runtime catalog)",
										MarkdownDescription: "The image containing the tooling required for a native build (by default it will use the one provided in the runtime catalog)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"package_types": schema.ListAttribute{
										Description:         "The Quarkus package types, 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the native kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists. Deprecated: use 'build-mode' instead.",
										MarkdownDescription: "The Quarkus package types, 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the native kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists. Deprecated: use 'build-mode' instead.",
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

							"registry": schema.SingleNestedAttribute{
								Description:         "The Registry trait sets up Maven to use the Image registry as a Maven repository. Deprecated: use jvm trait or read documentation.",
								MarkdownDescription: "The Registry trait sets up Maven to use the Image registry as a Maven repository. Deprecated: use jvm trait or read documentation.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
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

func (r *CamelApacheOrgIntegrationKitV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_integration_kit_v1_manifest")

	var model CamelApacheOrgIntegrationKitV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("camel.apache.org/v1")
	model.Kind = pointer.String("IntegrationKit")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
