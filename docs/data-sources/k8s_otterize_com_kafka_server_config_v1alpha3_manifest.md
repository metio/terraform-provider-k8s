---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_k8s_otterize_com_kafka_server_config_v1alpha3_manifest Data Source - terraform-provider-k8s"
subcategory: "k8s.otterize.com"
description: |-
  KafkaServerConfig is the Schema for the kafkaserverconfigs API
---

# k8s_k8s_otterize_com_kafka_server_config_v1alpha3_manifest (Data Source)

KafkaServerConfig is the Schema for the kafkaserverconfigs API

## Example Usage

```terraform
data "k8s_k8s_otterize_com_kafka_server_config_v1alpha3_manifest" "example" {
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

- `spec` (Attributes) KafkaServerConfigSpec defines the desired state of KafkaServerConfig (see [below for nested schema](#nestedatt--spec))

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

- `addr` (String)
- `no_auto_create_intents_for_operator` (Boolean) If Intents for network policies are enabled, and there are other Intents to this Kafka server, will automatically create an Intent so that the Intents Operator can connect. Set to true to disable.
- `service` (Attributes) (see [below for nested schema](#nestedatt--spec--service))
- `tls` (Attributes) (see [below for nested schema](#nestedatt--spec--tls))
- `topics` (Attributes List) (see [below for nested schema](#nestedatt--spec--topics))

<a id="nestedatt--spec--service"></a>
### Nested Schema for `spec.service`

Required:

- `name` (String)


<a id="nestedatt--spec--tls"></a>
### Nested Schema for `spec.tls`

Required:

- `cert_file` (String)
- `key_file` (String)
- `root_ca_file` (String)


<a id="nestedatt--spec--topics"></a>
### Nested Schema for `spec.topics`

Required:

- `client_identity_required` (Boolean)
- `intents_required` (Boolean)
- `pattern` (String)
- `topic` (String)
