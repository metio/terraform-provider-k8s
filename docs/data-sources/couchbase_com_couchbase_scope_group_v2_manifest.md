---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_couchbase_com_couchbase_scope_group_v2_manifest Data Source - terraform-provider-k8s"
subcategory: "couchbase.com"
description: |-
  CouchbaseScopeGroup represents a logical unit of data storage that sits between buckets and collections e.g. a bucket may contain multiple scopes, and a scope may contain multiple collections.  At present, scopes are not nested, so provide only a single level of abstraction.  Scopes provide a coarser grained basis for role-based access control (RBAC) and cross-datacenter replication (XDCR) than collections, but finer that buckets. In order to be considered by the Operator, a scope must be referenced by either a 'CouchbaseBucket' or 'CouchbaseEphemeralBucket' resource. Unlike 'CouchbaseScope' resources, scope groups represents multiple scopes, with the same common set of collections, to be expressed as a single resource, minimizing required configuration and Kubernetes API traffic.  It also forms the basis of Couchbase RBAC security boundaries.
---

# k8s_couchbase_com_couchbase_scope_group_v2_manifest (Data Source)

CouchbaseScopeGroup represents a logical unit of data storage that sits between buckets and collections e.g. a bucket may contain multiple scopes, and a scope may contain multiple collections.  At present, scopes are not nested, so provide only a single level of abstraction.  Scopes provide a coarser grained basis for role-based access control (RBAC) and cross-datacenter replication (XDCR) than collections, but finer that buckets. In order to be considered by the Operator, a scope must be referenced by either a 'CouchbaseBucket' or 'CouchbaseEphemeralBucket' resource. Unlike 'CouchbaseScope' resources, scope groups represents multiple scopes, with the same common set of collections, to be expressed as a single resource, minimizing required configuration and Kubernetes API traffic.  It also forms the basis of Couchbase RBAC security boundaries.

## Example Usage

```terraform
data "k8s_couchbase_com_couchbase_scope_group_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    names = ["name-1", "name-2"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Spec defines the desired state of the resource. (see [below for nested schema](#nestedatt--spec))

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

- `names` (List of String) Names specifies the names of the scopes.  Unlike CouchbaseScope, which specifies a single scope, a scope group specifies multiple, and the scope group must specify at least one scope name. Any scope names specified must be unique. Scope names must be 1-251 characters in length, contain only [a-zA-Z0-9_-%] and not start with either _ or %.

Optional:

- `collections` (Attributes) Collections defines how to collate collections included in this scope or scope group. Any of the provided methods may be used to collate a set of collections to manage.  Collated collections must have unique names, otherwise it is considered ambiguous, and an error condition. (see [below for nested schema](#nestedatt--spec--collections))

<a id="nestedatt--spec--collections"></a>
### Nested Schema for `spec.collections`

Optional:

- `managed` (Boolean) Managed indicates whether collections within this scope are managed. If not then you can dynamically create and delete collections with the Couchbase UI or SDKs.
- `preserve_default_collection` (Boolean) PreserveDefaultCollection indicates whether the Operator should manage the default collection within the default scope.  The default collection can be deleted, but can not be recreated by Couchbase Server.  By setting this field to 'true', the Operator will implicitly manage the default collection within the default scope.  The default collection cannot be modified and will have no document time-to-live (TTL).  When set to 'false', the operator will not manage the default collection, which will be deleted and cannot be used or recreated.
- `resources` (Attributes List) Resources is an explicit list of named resources that will be considered for inclusion in this scope or scopes.  If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency. (see [below for nested schema](#nestedatt--spec--collections--resources))
- `selector` (Attributes) Selector allows resources to be implicitly considered for inclusion in this scope or scopes.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta (see [below for nested schema](#nestedatt--spec--collections--selector))

<a id="nestedatt--spec--collections--resources"></a>
### Nested Schema for `spec.collections.resources`

Required:

- `name` (String) Name is the name of the Kubernetes resource name that is being referenced. Legal collection names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.

Optional:

- `kind` (String) Kind indicates the kind of resource that is being referenced.  A scope can only reference 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource kinds.  This field defaults to 'CouchbaseCollection' if not specified.


<a id="nestedatt--spec--collections--selector"></a>
### Nested Schema for `spec.collections.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--collections--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--collections--selector--match_expressions"></a>
### Nested Schema for `spec.collections.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.