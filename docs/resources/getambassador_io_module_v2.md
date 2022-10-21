---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_getambassador_io_module_v2 Resource - terraform-provider-k8s"
subcategory: "getambassador.io/v2"
description: |-
  A Module defines system-wide configuration.  The type of module is controlled by the .metadata.name; valid names are 'ambassador' or 'tls'.  https://www.getambassador.io/docs/edge-stack/latest/topics/running/ambassador/#the-ambassador-module https://www.getambassador.io/docs/edge-stack/latest/topics/running/tls/#tls-module-deprecated
---

# k8s_getambassador_io_module_v2 (Resource)

A Module defines system-wide configuration.  The type of module is controlled by the .metadata.name; valid names are 'ambassador' or 'tls'.  https://www.getambassador.io/docs/edge-stack/latest/topics/running/ambassador/#the-ambassador-module https://www.getambassador.io/docs/edge-stack/latest/topics/running/tls/#tls-module-deprecated

## Example Usage

```terraform
resource "k8s_getambassador_io_module_v2" "minimal" {
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

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `ambassador_id` (List of String) AmbassadorID declares which Ambassador instances should pay attention to this resource.  May either be a string or a list of strings.  If no value is provided, the default is:  	ambassador_id: 	- 'default'
- `config` (Dynamic) UntypedDict is relatively opaque as a Go type, but it preserves its contents in a roundtrippable way.

