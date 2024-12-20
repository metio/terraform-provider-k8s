---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_camel_apache_org_integration_kit_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "camel.apache.org"
description: |-
  IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.
---

# k8s_camel_apache_org_integration_kit_v1_manifest (Data Source)

IntegrationKit defines a container image and additional configuration needed to run an 'Integration'. An 'IntegrationKit' is a generic image generally built from the requirements of an 'Integration', but agnostic to it, in order to be reused by any other 'Integration' which has the same required set of capabilities. An 'IntegrationKit' may be used for other kits as a base container layer, when the 'incremental' build option is enabled.

## Example Usage

```terraform
data "k8s_camel_apache_org_integration_kit_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) the desired configuration (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `capabilities` (List of String) features offered by the IntegrationKit
- `configuration` (Attributes List) Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes configuration used by the kit (see [below for nested schema](#nestedatt--spec--configuration))
- `dependencies` (List of String) a list of Camel dependecies used by this kit
- `image` (String) the container image as identified in the container registry
- `profile` (String) the profile which is expected by this kit
- `repositories` (List of String) Maven repositories that can be used by the kit
- `sources` (Attributes List) the sources to add at build time (see [below for nested schema](#nestedatt--spec--sources))
- `traits` (Attributes) traits that the kit will execute (see [below for nested schema](#nestedatt--spec--traits))

<a id="nestedatt--spec--configuration"></a>
### Nested Schema for `spec.configuration`

Required:

- `type` (String) represents the type of configuration, ie: property, configmap, secret, ...
- `value` (String) the value to assign to the configuration (syntax may vary depending on the 'Type')


<a id="nestedatt--spec--sources"></a>
### Nested Schema for `spec.sources`

Optional:

- `compression` (Boolean) if the content is compressed (base64 encrypted)
- `content` (String) the source code (plain text)
- `content_key` (String) the confimap key holding the source content
- `content_ref` (String) the confimap reference holding the source content
- `content_type` (String) the content type (tipically text or binary)
- `from_kamelet` (Boolean) True if the spec is generated from a Kamelet
- `interceptors` (List of String) Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources Deprecated: no longer in use.
- `language` (String) specify which is the language (Camel DSL) used to interpret this source code
- `loader` (String) Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime
- `name` (String) the name of the specification
- `path` (String) the path where the file is stored
- `property_names` (List of String) List of property names defined in the source (e.g. if type is 'template')
- `raw_content` (String) the source code (binary)
- `type` (String) Type defines the kind of source described by this object


<a id="nestedatt--spec--traits"></a>
### Nested Schema for `spec.traits`

Optional:

- `addons` (Map of String) The collection of addon trait configurations
- `builder` (Attributes) The builder trait is internally used to determine the best strategy to build and configure IntegrationKits. (see [below for nested schema](#nestedatt--spec--traits--builder))
- `camel` (Attributes) The Camel trait sets up Camel configuration. (see [below for nested schema](#nestedatt--spec--traits--camel))
- `quarkus` (Attributes) The Quarkus trait configures the Quarkus runtime. It's enabled by default. NOTE: Compiling to a native executable, requires at least 4GiB of memory, so the Pod running the native build must have enough memory available. (see [below for nested schema](#nestedatt--spec--traits--quarkus))
- `registry` (Attributes) The Registry trait sets up Maven to use the Image registry as a Maven repository (support removed since version 2.5.0). Deprecated: use jvm trait or read documentation. (see [below for nested schema](#nestedatt--spec--traits--registry))

<a id="nestedatt--spec--traits--builder"></a>
### Nested Schema for `spec.traits.builder`

Optional:

- `annotations` (Map of String) When using 'pod' strategy, annotation to use for the builder pod.
- `base_image` (String) Specify a base image. In order to have the application working properly it must be a container image which has a Java JDK installed and ready to use on path (ie '/usr/bin/java').
- `configuration` (Map of String) Legacy trait configuration parameters. Deprecated: for backward compatibility.
- `enabled` (Boolean) Deprecated: no longer in use.
- `incremental_image_build` (Boolean) Use the incremental image build option, to reuse existing containers (default 'true')
- `limit_cpu` (String) When using 'pod' strategy, the maximum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.
- `limit_memory` (String) When using 'pod' strategy, the maximum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.
- `maven_profiles` (List of String) A list of references pointing to configmaps/secrets that contains a maven profile. This configmap/secret is a resource of the IntegrationKit created, therefore it needs to be present in the namespace where the operator is going to create the IntegrationKit. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).
- `node_selector` (Map of String) Defines a set of nodes the builder pod is eligible to be scheduled on, based on labels on the node.
- `order_strategy` (String) The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default is the platform default)
- `platforms` (List of String) The list of manifest platforms to use to build a container image (default 'linux/amd64').
- `properties` (List of String) A list of properties to be provided to the build task
- `request_cpu` (String) When using 'pod' strategy, the minimum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.
- `request_memory` (String) When using 'pod' strategy, the minimum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.
- `strategy` (String) The strategy to use, either 'pod' or 'routine' (default 'routine')
- `tasks` (List of String) A list of tasks to be executed (available only when using 'pod' strategy) with format '<name>;<container-image>;<container-command>'.
- `tasks_filter` (String) A list of tasks sorted by the order of execution in a csv format, ie, '<taskName1>,<taskName2>,...'. Mind that you must include also the operator tasks ('builder', 'quarkus-native', 'package', 'jib', 's2i') if you need to execute them. Useful only with 'pod' strategy.
- `tasks_limit_cpu` (List of String) A list of limit cpu configuration for the specific task with format '<task-name>:<limit-cpu-conf>'.
- `tasks_limit_memory` (List of String) A list of limit memory configuration for the specific task with format '<task-name>:<limit-memory-conf>'.
- `tasks_request_cpu` (List of String) A list of request cpu configuration for the specific task with format '<task-name>:<request-cpu-conf>'.
- `tasks_request_memory` (List of String) A list of request memory configuration for the specific task with format '<task-name>:<request-memory-conf>'.
- `verbose` (Boolean) Enable verbose logging on build components that support it (e.g. Kaniko build pod). Deprecated no longer in use


<a id="nestedatt--spec--traits--camel"></a>
### Nested Schema for `spec.traits.camel`

Optional:

- `configuration` (Map of String) Legacy trait configuration parameters. Deprecated: for backward compatibility.
- `enabled` (Boolean) Deprecated: no longer in use.
- `properties` (List of String) A list of properties to be provided to the Integration runtime
- `runtime_version` (String) The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform. You can use a fixed version (for example '3.2.3') or a semantic version (for example '3.x') which will try to resolve to the best matching Catalog existing on the cluster.


<a id="nestedatt--spec--traits--quarkus"></a>
### Nested Schema for `spec.traits.quarkus`

Optional:

- `build_mode` (List of String) The Quarkus mode to run: either 'jvm' or 'native' (default 'jvm'). In case both 'jvm' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'jvm' one once ready.
- `configuration` (Map of String) Legacy trait configuration parameters. Deprecated: for backward compatibility.
- `enabled` (Boolean) Deprecated: no longer in use.
- `native_base_image` (String) The base image to use when running a native build (default 'quay.io/quarkus/quarkus-micro-image:2.0')
- `native_builder_image` (String) The image containing the tooling required for a native build (by default it will use the one provided in the runtime catalog)
- `package_types` (List of String) The Quarkus package types, 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the native kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists. Deprecated: use 'build-mode' instead.


<a id="nestedatt--spec--traits--registry"></a>
### Nested Schema for `spec.traits.registry`

Optional:

- `configuration` (Map of String) Legacy trait configuration parameters. Deprecated: for backward compatibility.
- `enabled` (Boolean) Can be used to enable or disable a trait. All traits share this common property.
