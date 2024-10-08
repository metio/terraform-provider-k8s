---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_app_kiegroup_org_kogito_infra_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "app.kiegroup.org"
description: |-
  KogitoInfra is the resource to bind a Custom Resource (CR) not managed by Kogito Operator to a given deployed Kogito service. It holds the reference of a CR managed by another operator such as Strimzi. For example: one can create a Kafka CR via Strimzi and link this resource using KogitoInfra to a given Kogito service (custom or supporting, such as Data Index). Please refer to the Kogito Operator documentation (https://docs.jboss.org/kogito/release/latest/html_single/) for more information.
---

# k8s_app_kiegroup_org_kogito_infra_v1beta1_manifest (Data Source)

KogitoInfra is the resource to bind a Custom Resource (CR) not managed by Kogito Operator to a given deployed Kogito service. It holds the reference of a CR managed by another operator such as Strimzi. For example: one can create a Kafka CR via Strimzi and link this resource using KogitoInfra to a given Kogito service (custom or supporting, such as Data Index). Please refer to the Kogito Operator documentation (https://docs.jboss.org/kogito/release/latest/html_single/) for more information.

## Example Usage

```terraform
data "k8s_app_kiegroup_org_kogito_infra_v1beta1_manifest" "example" {
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

- `spec` (Attributes) KogitoInfraSpec defines the desired state of KogitoInfra. (see [below for nested schema](#nestedatt--spec))

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

- `config_map_env_from_references` (List of String) List of secret that should be mounted to the services as envs
- `config_map_volume_references` (Attributes List) List of configmap that should be added to the services bound to this infra instance (see [below for nested schema](#nestedatt--spec--config_map_volume_references))
- `envs` (Attributes List) Environment variables to be added to the runtime container. Keys must be a C_IDENTIFIER. (see [below for nested schema](#nestedatt--spec--envs))
- `infra_properties` (Map of String) Optional properties which would be needed to setup correct runtime/service configuration, based on the resource type. For example, MongoDB will require 'username' and 'database' as properties for a correct setup, else it will fail
- `resource` (Attributes) Resource for the service. Example: Infinispan/Kafka/Keycloak. (see [below for nested schema](#nestedatt--spec--resource))
- `secret_env_from_references` (List of String) List of secret that should be mounted to the services as envs
- `secret_volume_references` (Attributes List) List of secret that should be munted to the services bound to this infra instance (see [below for nested schema](#nestedatt--spec--secret_volume_references))

<a id="nestedatt--spec--config_map_volume_references"></a>
### Nested Schema for `spec.config_map_volume_references`

Required:

- `name` (String) This must match the Name of a ConfigMap.

Optional:

- `file_mode` (Number) Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.
- `mount_path` (String) Path within the container at which the volume should be mounted. Must not contain ':'. Default mount path is /home/kogito/config
- `optional` (Boolean) Specify whether the Secret or its keys must be defined


<a id="nestedatt--spec--envs"></a>
### Nested Schema for `spec.envs`

Required:

- `name` (String) Name of the environment variable. Must be a C_IDENTIFIER.

Optional:

- `value` (String) Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.
- `value_from` (Attributes) Source for the environment variable's value. Cannot be used if value is not empty. (see [below for nested schema](#nestedatt--spec--envs--value_from))

<a id="nestedatt--spec--envs--value_from"></a>
### Nested Schema for `spec.envs.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--envs--value_from--config_map_key_ref))
- `field_ref` (Attributes) Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. (see [below for nested schema](#nestedatt--spec--envs--value_from--field_ref))
- `resource_field_ref` (Attributes) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedatt--spec--envs--value_from--resource_field_ref))
- `secret_key_ref` (Attributes) Selects a key of a secret in the pod's namespace (see [below for nested schema](#nestedatt--spec--envs--value_from--secret_key_ref))

<a id="nestedatt--spec--envs--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.envs.value_from.config_map_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--envs--value_from--field_ref"></a>
### Nested Schema for `spec.envs.value_from.field_ref`

Required:

- `field_path` (String) Path of the field to select in the specified API version.

Optional:

- `api_version` (String) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.


<a id="nestedatt--spec--envs--value_from--resource_field_ref"></a>
### Nested Schema for `spec.envs.value_from.resource_field_ref`

Required:

- `resource` (String) Required: resource to select

Optional:

- `container_name` (String) Container name: required for volumes, optional for env vars
- `divisor` (String) Specifies the output format of the exposed resources, defaults to '1'


<a id="nestedatt--spec--envs--value_from--secret_key_ref"></a>
### Nested Schema for `spec.envs.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from. Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined




<a id="nestedatt--spec--resource"></a>
### Nested Schema for `spec.resource`

Required:

- `api_version` (String) APIVersion describes the API Version of referred Kubernetes resource for example, infinispan.org/v1
- `kind` (String) Kind describes the kind of referred Kubernetes resource for example, Infinispan
- `name` (String) Name of referred resource.

Optional:

- `namespace` (String) Namespace where referred resource exists.


<a id="nestedatt--spec--secret_volume_references"></a>
### Nested Schema for `spec.secret_volume_references`

Required:

- `name` (String) This must match the Name of a ConfigMap.

Optional:

- `file_mode` (Number) Permission on the file mounted as volume on deployment. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.
- `mount_path` (String) Path within the container at which the volume should be mounted. Must not contain ':'. Default mount path is /home/kogito/config
- `optional` (Boolean) Specify whether the Secret or its keys must be defined
