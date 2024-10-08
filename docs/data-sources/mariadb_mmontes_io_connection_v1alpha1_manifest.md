---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_mariadb_mmontes_io_connection_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "mariadb.mmontes.io"
description: |-
  Connection is the Schema for the connections API. It is used to configure connection strings for the applications connecting to MariaDB.
---

# k8s_mariadb_mmontes_io_connection_v1alpha1_manifest (Data Source)

Connection is the Schema for the connections API. It is used to configure connection strings for the applications connecting to MariaDB.

## Example Usage

```terraform
data "k8s_mariadb_mmontes_io_connection_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) ConnectionSpec defines the desired state of Connection (see [below for nested schema](#nestedatt--spec))

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

- `password_secret_key_ref` (Attributes) PasswordSecretKeyRef is a reference to the password to use for configuring the Connection. (see [below for nested schema](#nestedatt--spec--password_secret_key_ref))
- `username` (String) Username to use for configuring the Connection.

Optional:

- `database` (String) Database to use when configuring the Connection.
- `health_check` (Attributes) HealthCheck to be used in the Connection. (see [below for nested schema](#nestedatt--spec--health_check))
- `host` (String) Host to connect to. If not provided, it defaults to the MariaDB host or to the MaxScale host.
- `maria_db_ref` (Attributes) MariaDBRef is a reference to the MariaDB to connect to. Either MariaDBRef or MaxScaleRef must be provided. (see [below for nested schema](#nestedatt--spec--maria_db_ref))
- `max_scale_ref` (Attributes) MaxScaleRef is a reference to the MaxScale to connect to. Either MariaDBRef or MaxScaleRef must be provided. (see [below for nested schema](#nestedatt--spec--max_scale_ref))
- `params` (Map of String) Params to be used in the Connection.
- `port` (Number) Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.
- `secret_name` (String) SecretName to be used in the Connection.
- `secret_template` (Attributes) SecretTemplate to be used in the Connection. (see [below for nested schema](#nestedatt--spec--secret_template))
- `service_name` (String) ServiceName to be used in the Connection.

<a id="nestedatt--spec--password_secret_key_ref"></a>
### Nested Schema for `spec.password_secret_key_ref`

Required:

- `key` (String) The key of the secret to select from. Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined


<a id="nestedatt--spec--health_check"></a>
### Nested Schema for `spec.health_check`

Optional:

- `interval` (String) Interval used to perform health checks.
- `retry_interval` (String) RetryInterval is the intervañ used to perform health check retries.


<a id="nestedatt--spec--maria_db_ref"></a>
### Nested Schema for `spec.maria_db_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
- `wait_for_it` (Boolean) WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.


<a id="nestedatt--spec--max_scale_ref"></a>
### Nested Schema for `spec.max_scale_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--spec--secret_template"></a>
### Nested Schema for `spec.secret_template`

Optional:

- `annotations` (Map of String) Annotations to be added to the Secret object.
- `database_key` (String) DatabaseKey to be used in the Secret.
- `format` (String) Format to be used in the Secret.
- `host_key` (String) HostKey to be used in the Secret.
- `key` (String) Key to be used in the Secret.
- `labels` (Map of String) Labels to be added to the Secret object.
- `password_key` (String) PasswordKey to be used in the Secret.
- `port_key` (String) PortKey to be used in the Secret.
- `username_key` (String) UsernameKey to be used in the Secret.
