---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_opentelemetry_io_instrumentation_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "opentelemetry.io"
description: |-
  Instrumentation is the spec for OpenTelemetry instrumentation.
---

# k8s_opentelemetry_io_instrumentation_v1alpha1 (Resource)

Instrumentation is the spec for OpenTelemetry instrumentation.

## Example Usage

```terraform
resource "k8s_opentelemetry_io_instrumentation_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) InstrumentationSpec defines the desired state of OpenTelemetry SDK and instrumentation. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `dotnet` (Attributes) DotNet defines configuration for DotNet auto-instrumentation. (see [below for nested schema](#nestedatt--spec--dotnet))
- `env` (Attributes List) Env defines common env vars. There are four layers for env vars' definitions and the precedence order is: 'original container env vars' > 'language specific env vars' > 'common env vars' > 'instrument spec configs' vars'. If the former var had been defined, then the other vars would be ignored. (see [below for nested schema](#nestedatt--spec--env))
- `exporter` (Attributes) Exporter defines exporter configuration. (see [below for nested schema](#nestedatt--spec--exporter))
- `java` (Attributes) Java defines configuration for java auto-instrumentation. (see [below for nested schema](#nestedatt--spec--java))
- `nodejs` (Attributes) NodeJS defines configuration for nodejs auto-instrumentation. (see [below for nested schema](#nestedatt--spec--nodejs))
- `propagators` (List of String) Propagators defines inter-process context propagation configuration.
- `python` (Attributes) Python defines configuration for python auto-instrumentation. (see [below for nested schema](#nestedatt--spec--python))
- `resource` (Attributes) Resource defines the configuration for the resource attributes, as defined by the OpenTelemetry specification. (see [below for nested schema](#nestedatt--spec--resource))
- `sampler` (Attributes) Sampler defines sampling configuration. (see [below for nested schema](#nestedatt--spec--sampler))

<a id="nestedatt--spec--dotnet"></a>
### Nested Schema for `spec.dotnet`

Optional:

- `env` (Attributes List) Env defines DotNet specific env vars. There are four layers for env vars' definitions and the precedence order is: 'original container env vars' > 'language specific env vars' > 'common env vars' > 'instrument spec configs' vars'. If the former var had been defined, then the other vars would be ignored. (see [below for nested schema](#nestedatt--spec--dotnet--env))
- `image` (String) Image is a container image with DotNet SDK and auto-instrumentation.

<a id="nestedatt--spec--dotnet--env"></a>
### Nested Schema for `spec.dotnet.env`

Required:

- `name` (String) Name of the environment variable. Must be a C_IDENTIFIER.

Optional:

- `value` (String) Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.
- `value_from` (Attributes) Source for the environment variable's value. Cannot be used if value is not empty. (see [below for nested schema](#nestedatt--spec--dotnet--env--value_from))

<a id="nestedatt--spec--dotnet--env--value_from"></a>
### Nested Schema for `spec.dotnet.env.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--dotnet--env--value_from--config_map_key_ref))
- `field_ref` (Attributes) Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. (see [below for nested schema](#nestedatt--spec--dotnet--env--value_from--field_ref))
- `resource_field_ref` (Attributes) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedatt--spec--dotnet--env--value_from--resource_field_ref))
- `secret_key_ref` (Attributes) Selects a key of a secret in the pod's namespace (see [below for nested schema](#nestedatt--spec--dotnet--env--value_from--secret_key_ref))

<a id="nestedatt--spec--dotnet--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.dotnet.env.value_from.config_map_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--dotnet--env--value_from--field_ref"></a>
### Nested Schema for `spec.dotnet.env.value_from.field_ref`

Required:

- `field_path` (String) Path of the field to select in the specified API version.

Optional:

- `api_version` (String) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.


<a id="nestedatt--spec--dotnet--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.dotnet.env.value_from.resource_field_ref`

Required:

- `resource` (String) Required: resource to select

Optional:

- `container_name` (String) Container name: required for volumes, optional for env vars
- `divisor` (Dynamic) Specifies the output format of the exposed resources, defaults to '1'


<a id="nestedatt--spec--dotnet--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.dotnet.env.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined





<a id="nestedatt--spec--env"></a>
### Nested Schema for `spec.env`

Required:

- `name` (String) Name of the environment variable. Must be a C_IDENTIFIER.

Optional:

- `value` (String) Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.
- `value_from` (Attributes) Source for the environment variable's value. Cannot be used if value is not empty. (see [below for nested schema](#nestedatt--spec--env--value_from))

<a id="nestedatt--spec--env--value_from"></a>
### Nested Schema for `spec.env.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--env--value_from--config_map_key_ref))
- `field_ref` (Attributes) Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. (see [below for nested schema](#nestedatt--spec--env--value_from--field_ref))
- `resource_field_ref` (Attributes) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedatt--spec--env--value_from--resource_field_ref))
- `secret_key_ref` (Attributes) Selects a key of a secret in the pod's namespace (see [below for nested schema](#nestedatt--spec--env--value_from--secret_key_ref))

<a id="nestedatt--spec--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.env.value_from.secret_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--env--value_from--field_ref"></a>
### Nested Schema for `spec.env.value_from.secret_key_ref`

Required:

- `field_path` (String) Path of the field to select in the specified API version.

Optional:

- `api_version` (String) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.


<a id="nestedatt--spec--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.env.value_from.secret_key_ref`

Required:

- `resource` (String) Required: resource to select

Optional:

- `container_name` (String) Container name: required for volumes, optional for env vars
- `divisor` (Dynamic) Specifies the output format of the exposed resources, defaults to '1'


<a id="nestedatt--spec--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.env.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined




<a id="nestedatt--spec--exporter"></a>
### Nested Schema for `spec.exporter`

Optional:

- `endpoint` (String) Endpoint is address of the collector with OTLP endpoint.


<a id="nestedatt--spec--java"></a>
### Nested Schema for `spec.java`

Optional:

- `env` (Attributes List) Env defines java specific env vars. There are four layers for env vars' definitions and the precedence order is: 'original container env vars' > 'language specific env vars' > 'common env vars' > 'instrument spec configs' vars'. If the former var had been defined, then the other vars would be ignored. (see [below for nested schema](#nestedatt--spec--java--env))
- `image` (String) Image is a container image with javaagent auto-instrumentation JAR.

<a id="nestedatt--spec--java--env"></a>
### Nested Schema for `spec.java.env`

Required:

- `name` (String) Name of the environment variable. Must be a C_IDENTIFIER.

Optional:

- `value` (String) Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.
- `value_from` (Attributes) Source for the environment variable's value. Cannot be used if value is not empty. (see [below for nested schema](#nestedatt--spec--java--env--value_from))

<a id="nestedatt--spec--java--env--value_from"></a>
### Nested Schema for `spec.java.env.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--java--env--value_from--config_map_key_ref))
- `field_ref` (Attributes) Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. (see [below for nested schema](#nestedatt--spec--java--env--value_from--field_ref))
- `resource_field_ref` (Attributes) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedatt--spec--java--env--value_from--resource_field_ref))
- `secret_key_ref` (Attributes) Selects a key of a secret in the pod's namespace (see [below for nested schema](#nestedatt--spec--java--env--value_from--secret_key_ref))

<a id="nestedatt--spec--java--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.java.env.value_from.config_map_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--java--env--value_from--field_ref"></a>
### Nested Schema for `spec.java.env.value_from.field_ref`

Required:

- `field_path` (String) Path of the field to select in the specified API version.

Optional:

- `api_version` (String) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.


<a id="nestedatt--spec--java--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.java.env.value_from.resource_field_ref`

Required:

- `resource` (String) Required: resource to select

Optional:

- `container_name` (String) Container name: required for volumes, optional for env vars
- `divisor` (Dynamic) Specifies the output format of the exposed resources, defaults to '1'


<a id="nestedatt--spec--java--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.java.env.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined





<a id="nestedatt--spec--nodejs"></a>
### Nested Schema for `spec.nodejs`

Optional:

- `env` (Attributes List) Env defines nodejs specific env vars. There are four layers for env vars' definitions and the precedence order is: 'original container env vars' > 'language specific env vars' > 'common env vars' > 'instrument spec configs' vars'. If the former var had been defined, then the other vars would be ignored. (see [below for nested schema](#nestedatt--spec--nodejs--env))
- `image` (String) Image is a container image with NodeJS SDK and auto-instrumentation.

<a id="nestedatt--spec--nodejs--env"></a>
### Nested Schema for `spec.nodejs.env`

Required:

- `name` (String) Name of the environment variable. Must be a C_IDENTIFIER.

Optional:

- `value` (String) Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.
- `value_from` (Attributes) Source for the environment variable's value. Cannot be used if value is not empty. (see [below for nested schema](#nestedatt--spec--nodejs--env--value_from))

<a id="nestedatt--spec--nodejs--env--value_from"></a>
### Nested Schema for `spec.nodejs.env.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--nodejs--env--value_from--config_map_key_ref))
- `field_ref` (Attributes) Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. (see [below for nested schema](#nestedatt--spec--nodejs--env--value_from--field_ref))
- `resource_field_ref` (Attributes) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedatt--spec--nodejs--env--value_from--resource_field_ref))
- `secret_key_ref` (Attributes) Selects a key of a secret in the pod's namespace (see [below for nested schema](#nestedatt--spec--nodejs--env--value_from--secret_key_ref))

<a id="nestedatt--spec--nodejs--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.nodejs.env.value_from.config_map_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--nodejs--env--value_from--field_ref"></a>
### Nested Schema for `spec.nodejs.env.value_from.field_ref`

Required:

- `field_path` (String) Path of the field to select in the specified API version.

Optional:

- `api_version` (String) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.


<a id="nestedatt--spec--nodejs--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.nodejs.env.value_from.resource_field_ref`

Required:

- `resource` (String) Required: resource to select

Optional:

- `container_name` (String) Container name: required for volumes, optional for env vars
- `divisor` (Dynamic) Specifies the output format of the exposed resources, defaults to '1'


<a id="nestedatt--spec--nodejs--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.nodejs.env.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined





<a id="nestedatt--spec--python"></a>
### Nested Schema for `spec.python`

Optional:

- `env` (Attributes List) Env defines python specific env vars. There are four layers for env vars' definitions and the precedence order is: 'original container env vars' > 'language specific env vars' > 'common env vars' > 'instrument spec configs' vars'. If the former var had been defined, then the other vars would be ignored. (see [below for nested schema](#nestedatt--spec--python--env))
- `image` (String) Image is a container image with Python SDK and auto-instrumentation.

<a id="nestedatt--spec--python--env"></a>
### Nested Schema for `spec.python.env`

Required:

- `name` (String) Name of the environment variable. Must be a C_IDENTIFIER.

Optional:

- `value` (String) Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.
- `value_from` (Attributes) Source for the environment variable's value. Cannot be used if value is not empty. (see [below for nested schema](#nestedatt--spec--python--env--value_from))

<a id="nestedatt--spec--python--env--value_from"></a>
### Nested Schema for `spec.python.env.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--python--env--value_from--config_map_key_ref))
- `field_ref` (Attributes) Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. (see [below for nested schema](#nestedatt--spec--python--env--value_from--field_ref))
- `resource_field_ref` (Attributes) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedatt--spec--python--env--value_from--resource_field_ref))
- `secret_key_ref` (Attributes) Selects a key of a secret in the pod's namespace (see [below for nested schema](#nestedatt--spec--python--env--value_from--secret_key_ref))

<a id="nestedatt--spec--python--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.python.env.value_from.config_map_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--python--env--value_from--field_ref"></a>
### Nested Schema for `spec.python.env.value_from.field_ref`

Required:

- `field_path` (String) Path of the field to select in the specified API version.

Optional:

- `api_version` (String) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.


<a id="nestedatt--spec--python--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.python.env.value_from.resource_field_ref`

Required:

- `resource` (String) Required: resource to select

Optional:

- `container_name` (String) Container name: required for volumes, optional for env vars
- `divisor` (Dynamic) Specifies the output format of the exposed resources, defaults to '1'


<a id="nestedatt--spec--python--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.python.env.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined





<a id="nestedatt--spec--resource"></a>
### Nested Schema for `spec.resource`

Optional:

- `add_k8s_uid_attributes` (Boolean) AddK8sUIDAttributes defines whether K8s UID attributes should be collected (e.g. k8s.deployment.uid).
- `resource_attributes` (Map of String) Attributes defines attributes that are added to the resource. For example environment: dev


<a id="nestedatt--spec--sampler"></a>
### Nested Schema for `spec.sampler`

Optional:

- `argument` (String) Argument defines sampler argument. The value depends on the sampler type. For instance for parentbased_traceidratio sampler type it is a number in range [0..1] e.g. 0.25.
- `type` (String) Type defines sampler type. The value can be for instance parentbased_always_on, parentbased_always_off, parentbased_traceidratio...


