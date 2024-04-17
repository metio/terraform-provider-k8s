---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ceph_rook_io_ceph_object_store_user_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "ceph.rook.io"
description: |-
  CephObjectStoreUser represents a Ceph Object Store Gateway User
---

# k8s_ceph_rook_io_ceph_object_store_user_v1_manifest (Data Source)

CephObjectStoreUser represents a Ceph Object Store Gateway User

## Example Usage

```terraform
data "k8s_ceph_rook_io_ceph_object_store_user_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {}
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) ObjectStoreUserSpec represent the spec of an Objectstoreuser (see [below for nested schema](#nestedatt--spec))

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

- `capabilities` (Attributes) Additional admin-level capabilities for the Ceph object store user (see [below for nested schema](#nestedatt--spec--capabilities))
- `cluster_namespace` (String) The namespace where the parent CephCluster and CephObjectStore are found
- `display_name` (String) The display name for the ceph users
- `quotas` (Attributes) ObjectUserQuotaSpec can be used to set quotas for the object store user to limit their usage. See the [Ceph docs](https://docs.ceph.com/en/latest/radosgw/admin/?#quota-management) for more (see [below for nested schema](#nestedatt--spec--quotas))
- `store` (String) The store the user will be created in

<a id="nestedatt--spec--capabilities"></a>
### Nested Schema for `spec.capabilities`

Optional:

- `amz_cache` (String) Add capabilities for user to send request to RGW Cache API header. Documented in https://docs.ceph.com/en/quincy/radosgw/rgw-cache/#cache-api
- `bilog` (String) Add capabilities for user to change bucket index logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `bucket` (String) Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `buckets` (String) Admin capabilities to read/write Ceph object store buckets. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `datalog` (String) Add capabilities for user to change data logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `info` (String) Admin capabilities to read/write information about the user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `mdlog` (String) Add capabilities for user to change metadata logging. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `metadata` (String) Admin capabilities to read/write Ceph object store metadata. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `oidc_provider` (String) Add capabilities for user to change oidc provider. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `ratelimit` (String) Add capabilities for user to set rate limiter for user and bucket. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `roles` (String) Admin capabilities to read/write roles for user. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `usage` (String) Admin capabilities to read/write Ceph object store usage. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `user` (String) Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `user_policy` (String) Add capabilities for user to change user policies. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `users` (String) Admin capabilities to read/write Ceph object store users. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities
- `zone` (String) Admin capabilities to read/write Ceph object store zones. Documented in https://docs.ceph.com/en/latest/radosgw/admin/?#add-remove-admin-capabilities


<a id="nestedatt--spec--quotas"></a>
### Nested Schema for `spec.quotas`

Optional:

- `max_buckets` (Number) Maximum bucket limit for the ceph user
- `max_objects` (Number) Maximum number of objects across all the user's buckets
- `max_size` (String) Maximum size limit of all objects across all the user's buckets See https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity for more info.