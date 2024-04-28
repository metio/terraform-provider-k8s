---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_k8s_mariadb_com_user_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "k8s.mariadb.com"
description: |-
  User is the Schema for the users API.  It is used to define grants as if you were running a 'CREATE USER' statement.
---

# k8s_k8s_mariadb_com_user_v1alpha1_manifest (Data Source)

User is the Schema for the users API.  It is used to define grants as if you were running a 'CREATE USER' statement.

## Example Usage

```terraform
data "k8s_k8s_mariadb_com_user_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) UserSpec defines the desired state of User (see [below for nested schema](#nestedatt--spec))

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

- `maria_db_ref` (Attributes) MariaDBRef is a reference to a MariaDB object. (see [below for nested schema](#nestedatt--spec--maria_db_ref))
- `password_secret_key_ref` (Attributes) PasswordSecretKeyRef is a reference to the password to be used by the User. (see [below for nested schema](#nestedatt--spec--password_secret_key_ref))

Optional:

- `host` (String) Host related to the User.
- `max_user_connections` (Number) MaxUserConnections defines the maximum number of connections that the User can establish.
- `name` (String) Name overrides the default name provided by metadata.name.
- `requeue_interval` (String) RequeueInterval is used to perform requeue reconcilizations.
- `retry_interval` (String) RetryInterval is the interval used to perform retries.

<a id="nestedatt--spec--maria_db_ref"></a>
### Nested Schema for `spec.maria_db_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids
- `wait_for_it` (Boolean) WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.


<a id="nestedatt--spec--password_secret_key_ref"></a>
### Nested Schema for `spec.password_secret_key_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined