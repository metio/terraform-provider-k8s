---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kafka_banzaicloud_io_kafka_user_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "kafka.banzaicloud.io"
description: |-
  KafkaUser is the Schema for the kafka users API
---

# k8s_kafka_banzaicloud_io_kafka_user_v1alpha1_manifest (Data Source)

KafkaUser is the Schema for the kafka users API

## Example Usage

```terraform
data "k8s_kafka_banzaicloud_io_kafka_user_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) KafkaUserSpec defines the desired state of KafkaUser (see [below for nested schema](#nestedatt--spec))

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

- `cluster_ref` (Attributes) ClusterReference states a reference to a cluster for topic/user provisioning (see [below for nested schema](#nestedatt--spec--cluster_ref))
- `secret_name` (String) secretName is used as the name of the K8S secret that contains the certificate of the KafkaUser. SecretName should be unique inside the namespace where KafkaUser is located.

Optional:

- `annotations` (Map of String) Annotations defines the annotations placed on the certificate or certificate signing request object
- `create_cert` (Boolean)
- `dns_names` (List of String)
- `expiration_seconds` (Number) expirationSeconds is the requested duration of validity of the issued certificate. The minimum valid value for expirationSeconds is 3600 i.e. 1h. When it is not specified the default validation duration is 90 days
- `include_jks` (Boolean)
- `pki_backend_spec` (Attributes) (see [below for nested schema](#nestedatt--spec--pki_backend_spec))
- `topic_grants` (Attributes List) (see [below for nested schema](#nestedatt--spec--topic_grants))

<a id="nestedatt--spec--cluster_ref"></a>
### Nested Schema for `spec.cluster_ref`

Required:

- `name` (String)

Optional:

- `namespace` (String)


<a id="nestedatt--spec--pki_backend_spec"></a>
### Nested Schema for `spec.pki_backend_spec`

Required:

- `pki_backend` (String)

Optional:

- `issuer_ref` (Attributes) ObjectReference is a reference to an object with a given name, kind and group. (see [below for nested schema](#nestedatt--spec--pki_backend_spec--issuer_ref))
- `signer_name` (String) SignerName indicates requested signer, and is a qualified name.

<a id="nestedatt--spec--pki_backend_spec--issuer_ref"></a>
### Nested Schema for `spec.pki_backend_spec.issuer_ref`

Required:

- `name` (String) Name of the resource being referred to.

Optional:

- `group` (String) Group of the resource being referred to.
- `kind` (String) Kind of the resource being referred to.



<a id="nestedatt--spec--topic_grants"></a>
### Nested Schema for `spec.topic_grants`

Required:

- `access_type` (String) KafkaAccessType hold info about Kafka ACL
- `topic_name` (String)

Optional:

- `pattern_type` (String) KafkaPatternType hold the Resource Pattern Type of kafka ACL
