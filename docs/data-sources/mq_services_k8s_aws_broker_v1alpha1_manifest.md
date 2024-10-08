---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_mq_services_k8s_aws_broker_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "mq.services.k8s.aws"
description: |-
  Broker is the Schema for the Brokers API
---

# k8s_mq_services_k8s_aws_broker_v1alpha1_manifest (Data Source)

Broker is the Schema for the Brokers API

## Example Usage

```terraform
data "k8s_mq_services_k8s_aws_broker_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) BrokerSpec defines the desired state of Broker. (see [below for nested schema](#nestedatt--spec))

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

- `auto_minor_version_upgrade` (Boolean)
- `deployment_mode` (String)
- `engine_type` (String)
- `engine_version` (String)
- `host_instance_type` (String)
- `name` (String)
- `publicly_accessible` (Boolean)
- `users` (Attributes List) (see [below for nested schema](#nestedatt--spec--users))

Optional:

- `authentication_strategy` (String)
- `configuration` (Attributes) A list of information about the configuration. Does not apply to RabbitMQ brokers. (see [below for nested schema](#nestedatt--spec--configuration))
- `creator_request_id` (String)
- `encryption_options` (Attributes) Does not apply to RabbitMQ brokers. Encryption options for the broker. (see [below for nested schema](#nestedatt--spec--encryption_options))
- `ldap_server_metadata` (Attributes) Optional. The metadata of the LDAP server used to authenticate and authorize connections to the broker. Does not apply to RabbitMQ brokers. (see [below for nested schema](#nestedatt--spec--ldap_server_metadata))
- `logs` (Attributes) The list of information about logs to be enabled for the specified broker. (see [below for nested schema](#nestedatt--spec--logs))
- `maintenance_window_start_time` (Attributes) The scheduled time period relative to UTC during which Amazon MQ begins to apply pending updates or patches to the broker. (see [below for nested schema](#nestedatt--spec--maintenance_window_start_time))
- `security_group_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--security_group_refs))
- `security_groups` (List of String)
- `storage_type` (String)
- `subnet_i_ds` (List of String)
- `subnet_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--subnet_refs))
- `tags` (Map of String)

<a id="nestedatt--spec--users"></a>
### Nested Schema for `spec.users`

Optional:

- `console_access` (Boolean)
- `groups` (List of String)
- `password` (Attributes) SecretKeyReference combines a k8s corev1.SecretReference with a specific key within the referred-to Secret (see [below for nested schema](#nestedatt--spec--users--password))
- `username` (String)

<a id="nestedatt--spec--users--password"></a>
### Nested Schema for `spec.users.password`

Required:

- `key` (String) Key is the key within the secret

Optional:

- `name` (String) name is unique within a namespace to reference a secret resource.
- `namespace` (String) namespace defines the space within which the secret name must be unique.



<a id="nestedatt--spec--configuration"></a>
### Nested Schema for `spec.configuration`

Optional:

- `id` (String)
- `revision` (Number)


<a id="nestedatt--spec--encryption_options"></a>
### Nested Schema for `spec.encryption_options`

Optional:

- `kms_key_id` (String)
- `use_aws_owned_key` (Boolean)


<a id="nestedatt--spec--ldap_server_metadata"></a>
### Nested Schema for `spec.ldap_server_metadata`

Optional:

- `hosts` (List of String)
- `role_base` (String)
- `role_name` (String)
- `role_search_matching` (String)
- `role_search_subtree` (Boolean)
- `service_account_password` (String)
- `service_account_username` (String)
- `user_base` (String)
- `user_role_name` (String)
- `user_search_matching` (String)
- `user_search_subtree` (Boolean)


<a id="nestedatt--spec--logs"></a>
### Nested Schema for `spec.logs`

Optional:

- `audit` (Boolean)
- `general` (Boolean)


<a id="nestedatt--spec--maintenance_window_start_time"></a>
### Nested Schema for `spec.maintenance_window_start_time`

Optional:

- `day_of_week` (String)
- `time_of_day` (String)
- `time_zone` (String)


<a id="nestedatt--spec--security_group_refs"></a>
### Nested Schema for `spec.security_group_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--security_group_refs--from))

<a id="nestedatt--spec--security_group_refs--from"></a>
### Nested Schema for `spec.security_group_refs.from`

Optional:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--subnet_refs"></a>
### Nested Schema for `spec.subnet_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--subnet_refs--from))

<a id="nestedatt--spec--subnet_refs--from"></a>
### Nested Schema for `spec.subnet_refs.from`

Optional:

- `name` (String)
- `namespace` (String)
