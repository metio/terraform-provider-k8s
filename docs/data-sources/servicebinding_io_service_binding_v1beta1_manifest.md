---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_servicebinding_io_service_binding_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "servicebinding.io"
description: |-
  ServiceBinding is the Schema for the servicebindings API
---

# k8s_servicebinding_io_service_binding_v1beta1_manifest (Data Source)

ServiceBinding is the Schema for the servicebindings API

## Example Usage

```terraform
data "k8s_servicebinding_io_service_binding_v1beta1_manifest" "example" {
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

- `spec` (Attributes) ServiceBindingSpec defines the desired state of ServiceBinding (see [below for nested schema](#nestedatt--spec))

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

- `service` (Attributes) Service is a reference to an object that fulfills the ProvisionedService duck type (see [below for nested schema](#nestedatt--spec--service))
- `workload` (Attributes) Workload is a reference to an object (see [below for nested schema](#nestedatt--spec--workload))

Optional:

- `env` (Attributes List) Env is the collection of mappings from Secret entries to environment variables (see [below for nested schema](#nestedatt--spec--env))
- `name` (String) Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.
- `provider` (String) Provider is the provider of the service as projected into the workload container
- `type` (String) Type is the type of the service as projected into the workload container

<a id="nestedatt--spec--service"></a>
### Nested Schema for `spec.service`

Required:

- `api_version` (String) API version of the referent.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names


<a id="nestedatt--spec--workload"></a>
### Nested Schema for `spec.workload`

Required:

- `api_version` (String) API version of the referent.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds

Optional:

- `containers` (List of String) Containers describes which containers in a Pod should be bound to
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `selector` (Attributes) Selector is a query that selects the workload or workloads to bind the service to (see [below for nested schema](#nestedatt--spec--workload--selector))

<a id="nestedatt--spec--workload--selector"></a>
### Nested Schema for `spec.workload.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--workload--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--workload--selector--match_expressions"></a>
### Nested Schema for `spec.workload.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.




<a id="nestedatt--spec--env"></a>
### Nested Schema for `spec.env`

Required:

- `key` (String) Key is the key in the Secret that will be exposed
- `name` (String) Name is the name of the environment variable