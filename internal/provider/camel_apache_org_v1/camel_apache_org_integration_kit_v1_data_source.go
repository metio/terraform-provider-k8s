/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &CamelApacheOrgIntegrationKitV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CamelApacheOrgIntegrationKitV1DataSource{}
)

func NewCamelApacheOrgIntegrationKitV1DataSource() datasource.DataSource {
	return &CamelApacheOrgIntegrationKitV1DataSource{}
}

type CamelApacheOrgIntegrationKitV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CamelApacheOrgIntegrationKitV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
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
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				LimitCPU      *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
				LimitMemory   *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
				MavenProfiles *[]string          `tfsdk:"maven_profiles" json:"mavenProfiles,omitempty"`
				OrderStrategy *string            `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
				Properties    *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RequestCPU    *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
				RequestMemory *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
				Strategy      *string            `tfsdk:"strategy" json:"strategy,omitempty"`
				Tasks         *[]string          `tfsdk:"tasks" json:"tasks,omitempty"`
				Verbose       *bool              `tfsdk:"verbose" json:"verbose,omitempty"`
			} `tfsdk:"builder" json:"builder,omitempty"`
			Camel *struct {
				Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Properties     *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RuntimeVersion *string            `tfsdk:"runtime_version" json:"runtimeVersion,omitempty"`
			} `tfsdk:"camel" json:"camel,omitempty"`
			Quarkus *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				PackageTypes  *[]string          `tfsdk:"package_types" json:"packageTypes,omitempty"`
			} `tfsdk:"quarkus" json:"quarkus,omitempty"`
			Registry *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"registry" json:"registry,omitempty"`
		} `tfsdk:"traits" json:"traits,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgIntegrationKitV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_integration_kit_v1"
}

func (r *CamelApacheOrgIntegrationKitV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.",
		MarkdownDescription: "IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "the desired configuration",
				MarkdownDescription: "the desired configuration",
				Attributes: map[string]schema.Attribute{
					"configuration": schema.ListNestedAttribute{
						Description:         "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes configuration used by the kit",
						MarkdownDescription: "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes configuration used by the kit",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
									MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
									MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"dependencies": schema.ListAttribute{
						Description:         "a list of Camel dependecies used by this kit",
						MarkdownDescription: "a list of Camel dependecies used by this kit",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image": schema.StringAttribute{
						Description:         "the container image as identified in the container registry",
						MarkdownDescription: "the container image as identified in the container registry",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"profile": schema.StringAttribute{
						Description:         "the profile which is expected by this kit",
						MarkdownDescription: "the profile which is expected by this kit",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"repositories": schema.ListAttribute{
						Description:         "Maven repositories that can be used by the kit",
						MarkdownDescription: "Maven repositories that can be used by the kit",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
									Optional:            false,
									Computed:            true,
								},

								"content": schema.StringAttribute{
									Description:         "the source code (plain text)",
									MarkdownDescription: "the source code (plain text)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_key": schema.StringAttribute{
									Description:         "the confimap key holding the source content",
									MarkdownDescription: "the confimap key holding the source content",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_ref": schema.StringAttribute{
									Description:         "the confimap reference holding the source content",
									MarkdownDescription: "the confimap reference holding the source content",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_type": schema.StringAttribute{
									Description:         "the content type (tipically text or binary)",
									MarkdownDescription: "the content type (tipically text or binary)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"interceptors": schema.ListAttribute{
									Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"language": schema.StringAttribute{
									Description:         "specify which is the language (Camel DSL) used to interpret this source code",
									MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"loader": schema.StringAttribute{
									Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "the name of the specification",
									MarkdownDescription: "the name of the specification",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"path": schema.StringAttribute{
									Description:         "the path where the file is stored",
									MarkdownDescription: "the path where the file is stored",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"property_names": schema.ListAttribute{
									Description:         "List of property names defined in the source (e.g. if type is 'template')",
									MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"raw_content": schema.StringAttribute{
									Description:         "the source code (binary)",
									MarkdownDescription: "the source code (binary)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "Type defines the kind of source described by this object",
									MarkdownDescription: "Type defines the kind of source described by this object",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
								Optional:            false,
								Computed:            true,
							},

							"builder": schema.SingleNestedAttribute{
								Description:         "The builder trait is internally used to determine the best strategy to build and configure IntegrationKits.",
								MarkdownDescription: "The builder trait is internally used to determine the best strategy to build and configure IntegrationKits.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of CPU required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of CPU required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of memory required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of memory required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"maven_profiles": schema.ListAttribute{
										Description:         "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										MarkdownDescription: "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"order_strategy": schema.StringAttribute{
										Description:         "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default sequential)",
										MarkdownDescription: "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default sequential)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the build task",
										MarkdownDescription: "A list of properties to be provided to the build task",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of CPU required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of CPU required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of memory required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of memory required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"strategy": schema.StringAttribute{
										Description:         "The strategy to use, either 'pod' or 'routine' (default routine)",
										MarkdownDescription: "The strategy to use, either 'pod' or 'routine' (default routine)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tasks": schema.ListAttribute{
										Description:         "A list of tasks to be executed (available only when using 'pod' strategy) with format <name>;<container-image>;<container-command>",
										MarkdownDescription: "A list of tasks to be executed (available only when using 'pod' strategy) with format <name>;<container-image>;<container-command>",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"verbose": schema.BoolAttribute{
										Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the Integration runtime",
										MarkdownDescription: "A list of properties to be provided to the Integration runtime",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"runtime_version": schema.StringAttribute{
										Description:         "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										MarkdownDescription: "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"quarkus": schema.SingleNestedAttribute{
								Description:         "The Quarkus trait configures the Quarkus runtime. It's enabled by default. NOTE: Compiling to a native executable, i.e. when using 'package-type=native', is only supported for kamelets, as well as YAML and XML integrations. It also requires at least 4GiB of memory, so the Pod running the native build, that is either the operator Pod, or the build Pod (depending on the build strategy configured for the platform), must have enough memory available.",
								MarkdownDescription: "The Quarkus trait configures the Quarkus runtime. It's enabled by default. NOTE: Compiling to a native executable, i.e. when using 'package-type=native', is only supported for kamelets, as well as YAML and XML integrations. It also requires at least 4GiB of memory, so the Pod running the native build, that is either the operator Pod, or the build Pod (depending on the build strategy configured for the platform), must have enough memory available.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"package_types": schema.ListAttribute{
										Description:         "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										MarkdownDescription: "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"registry": schema.SingleNestedAttribute{
								Description:         "The Registry trait sets up Maven to use the Image registry as a Maven repository.",
								MarkdownDescription: "The Registry trait sets up Maven to use the Image registry as a Maven repository.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CamelApacheOrgIntegrationKitV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CamelApacheOrgIntegrationKitV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_camel_apache_org_integration_kit_v1")

	var data CamelApacheOrgIntegrationKitV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "integrationkits"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CamelApacheOrgIntegrationKitV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("camel.apache.org/v1")
	data.Kind = pointer.String("IntegrationKit")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
