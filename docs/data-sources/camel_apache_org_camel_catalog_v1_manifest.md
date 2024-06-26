---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_camel_apache_org_camel_catalog_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "camel.apache.org"
description: |-
  CamelCatalog represents the languages, components, data formats and capabilities enabled on a given runtime provider. The catalog may be statically generated.
---

# k8s_camel_apache_org_camel_catalog_v1_manifest (Data Source)

CamelCatalog represents the languages, components, data formats and capabilities enabled on a given runtime provider. The catalog may be statically generated.

## Example Usage

```terraform
data "k8s_camel_apache_org_camel_catalog_v1_manifest" "example" {
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

- `spec` (Attributes) the desired state of the catalog (see [below for nested schema](#nestedatt--spec))

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

Required:

- `artifacts` (Attributes) artifacts required by this catalog (see [below for nested schema](#nestedatt--spec--artifacts))
- `loaders` (Attributes) loaders required by this catalog (see [below for nested schema](#nestedatt--spec--loaders))
- `runtime` (Attributes) the runtime targeted for the catalog (see [below for nested schema](#nestedatt--spec--runtime))

<a id="nestedatt--spec--artifacts"></a>
### Nested Schema for `spec.artifacts`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `dataformats` (List of String) accepted data formats
- `dependencies` (Attributes List) required dependencies (see [below for nested schema](#nestedatt--spec--artifacts--dependencies))
- `exclusions` (Attributes List) provide a list of artifacts to exclude for this dependency (see [below for nested schema](#nestedatt--spec--artifacts--exclusions))
- `java_types` (List of String) the Java types used by the artifact feature (ie, component, data format, ...)
- `languages` (List of String) accepted languages
- `schemes` (Attributes List) accepted URI schemes (see [below for nested schema](#nestedatt--spec--artifacts--schemes))
- `type` (String) Maven Type
- `version` (String) Maven Version

<a id="nestedatt--spec--artifacts--dependencies"></a>
### Nested Schema for `spec.artifacts.dependencies`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `exclusions` (Attributes List) provide a list of artifacts to exclude for this dependency (see [below for nested schema](#nestedatt--spec--artifacts--dependencies--exclusions))
- `type` (String) Maven Type
- `version` (String) Maven Version

<a id="nestedatt--spec--artifacts--dependencies--exclusions"></a>
### Nested Schema for `spec.artifacts.dependencies.exclusions`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group



<a id="nestedatt--spec--artifacts--exclusions"></a>
### Nested Schema for `spec.artifacts.exclusions`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group


<a id="nestedatt--spec--artifacts--schemes"></a>
### Nested Schema for `spec.artifacts.schemes`

Required:

- `http` (Boolean) is a HTTP based scheme
- `id` (String) the ID (ie, timer in a timer:xyz URI)
- `passive` (Boolean) is a passive scheme

Optional:

- `consumer` (Attributes) required scope for consumer (see [below for nested schema](#nestedatt--spec--artifacts--schemes--consumer))
- `producer` (Attributes) required scope for producers (see [below for nested schema](#nestedatt--spec--artifacts--schemes--producer))

<a id="nestedatt--spec--artifacts--schemes--consumer"></a>
### Nested Schema for `spec.artifacts.schemes.consumer`

Optional:

- `dependencies` (Attributes List) list of dependencies needed for this scope (see [below for nested schema](#nestedatt--spec--artifacts--schemes--consumer--dependencies))

<a id="nestedatt--spec--artifacts--schemes--consumer--dependencies"></a>
### Nested Schema for `spec.artifacts.schemes.consumer.dependencies`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `exclusions` (Attributes List) provide a list of artifacts to exclude for this dependency (see [below for nested schema](#nestedatt--spec--artifacts--schemes--consumer--dependencies--exclusions))
- `type` (String) Maven Type
- `version` (String) Maven Version

<a id="nestedatt--spec--artifacts--schemes--consumer--dependencies--exclusions"></a>
### Nested Schema for `spec.artifacts.schemes.consumer.dependencies.exclusions`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group




<a id="nestedatt--spec--artifacts--schemes--producer"></a>
### Nested Schema for `spec.artifacts.schemes.producer`

Optional:

- `dependencies` (Attributes List) list of dependencies needed for this scope (see [below for nested schema](#nestedatt--spec--artifacts--schemes--producer--dependencies))

<a id="nestedatt--spec--artifacts--schemes--producer--dependencies"></a>
### Nested Schema for `spec.artifacts.schemes.producer.dependencies`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `exclusions` (Attributes List) provide a list of artifacts to exclude for this dependency (see [below for nested schema](#nestedatt--spec--artifacts--schemes--producer--dependencies--exclusions))
- `type` (String) Maven Type
- `version` (String) Maven Version

<a id="nestedatt--spec--artifacts--schemes--producer--dependencies--exclusions"></a>
### Nested Schema for `spec.artifacts.schemes.producer.dependencies.exclusions`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group






<a id="nestedatt--spec--loaders"></a>
### Nested Schema for `spec.loaders`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `dependencies` (Attributes List) a list of additional dependencies required beside the base one (see [below for nested schema](#nestedatt--spec--loaders--dependencies))
- `languages` (List of String) a list of DSLs supported
- `metadata` (Map of String) the metadata of the loader
- `type` (String) Maven Type
- `version` (String) Maven Version

<a id="nestedatt--spec--loaders--dependencies"></a>
### Nested Schema for `spec.loaders.dependencies`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `type` (String) Maven Type
- `version` (String) Maven Version



<a id="nestedatt--spec--runtime"></a>
### Nested Schema for `spec.runtime`

Required:

- `application_class` (String) application entry point (main) to be executed
- `dependencies` (Attributes List) list of dependencies needed to run the application (see [below for nested schema](#nestedatt--spec--runtime--dependencies))
- `provider` (String) Camel main application provider, ie, Camel Quarkus
- `version` (String) Camel K Runtime version

Optional:

- `capabilities` (Attributes) features offered by this runtime (see [below for nested schema](#nestedatt--spec--runtime--capabilities))
- `metadata` (Map of String) set of metadata

<a id="nestedatt--spec--runtime--dependencies"></a>
### Nested Schema for `spec.runtime.dependencies`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `type` (String) Maven Type
- `version` (String) Maven Version


<a id="nestedatt--spec--runtime--capabilities"></a>
### Nested Schema for `spec.runtime.capabilities`

Optional:

- `build_time_properties` (Attributes List) Set of required Camel build time properties (see [below for nested schema](#nestedatt--spec--runtime--capabilities--build_time_properties))
- `dependencies` (Attributes List) List of required Maven dependencies (see [below for nested schema](#nestedatt--spec--runtime--capabilities--dependencies))
- `metadata` (Map of String) Set of generic metadata
- `runtime_properties` (Attributes List) Set of required Camel runtime properties (see [below for nested schema](#nestedatt--spec--runtime--capabilities--runtime_properties))

<a id="nestedatt--spec--runtime--capabilities--build_time_properties"></a>
### Nested Schema for `spec.runtime.capabilities.build_time_properties`

Required:

- `key` (String)

Optional:

- `value` (String)


<a id="nestedatt--spec--runtime--capabilities--dependencies"></a>
### Nested Schema for `spec.runtime.capabilities.dependencies`

Required:

- `artifact_id` (String) Maven Artifact
- `group_id` (String) Maven Group

Optional:

- `classifier` (String) Maven Classifier
- `type` (String) Maven Type
- `version` (String) Maven Version


<a id="nestedatt--spec--runtime--capabilities--runtime_properties"></a>
### Nested Schema for `spec.runtime.capabilities.runtime_properties`

Required:

- `key` (String)

Optional:

- `value` (String)
