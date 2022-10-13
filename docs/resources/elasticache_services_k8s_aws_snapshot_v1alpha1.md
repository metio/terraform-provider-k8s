---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_elasticache_services_k8s_aws_snapshot_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "elasticache.services.k8s.aws/v1alpha1"
description: |-
  Snapshot is the Schema for the Snapshots API
---

# k8s_elasticache_services_k8s_aws_snapshot_v1alpha1 (Resource)

Snapshot is the Schema for the Snapshots API

## Example Usage

```terraform
resource "k8s_elasticache_services_k8s_aws_snapshot_v1alpha1" "minimal" {
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

- `spec` (Attributes) SnapshotSpec defines the desired state of Snapshot.  Represents a copy of an entire Redis cluster as of the time when the snapshot was taken. (see [below for nested schema](#nestedatt--spec))

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

Required:

- `snapshot_name` (String) A name for the snapshot being created.

Optional:

- `cache_cluster_id` (String) The identifier of an existing cluster. The snapshot is created from this cluster.
- `kms_key_id` (String) The ID of the KMS key used to encrypt the snapshot.
- `replication_group_id` (String) The identifier of an existing replication group. The snapshot is created from this replication group.
- `source_snapshot_name` (String) The name of an existing snapshot from which to make a copy.
- `tags` (Attributes List) A list of tags to be added to this resource. A tag is a key-value pair. A tag key must be accompanied by a tag value, although null is accepted. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)

