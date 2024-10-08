---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_grafana_integreatly_org_grafana_datasource_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "grafana.integreatly.org"
description: |-
  GrafanaDatasource is the Schema for the grafanadatasources API
---

# k8s_grafana_integreatly_org_grafana_datasource_v1beta1_manifest (Data Source)

GrafanaDatasource is the Schema for the grafanadatasources API

## Example Usage

```terraform
data "k8s_grafana_integreatly_org_grafana_datasource_v1beta1_manifest" "example" {
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

- `spec` (Attributes) GrafanaDatasourceSpec defines the desired state of GrafanaDatasource (see [below for nested schema](#nestedatt--spec))

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

- `datasource` (Attributes) (see [below for nested schema](#nestedatt--spec--datasource))
- `instance_selector` (Attributes) selects Grafana instances for import (see [below for nested schema](#nestedatt--spec--instance_selector))

Optional:

- `allow_cross_namespace_import` (Boolean) allow to import this resources from an operator in a different namespace
- `plugins` (Attributes List) plugins (see [below for nested schema](#nestedatt--spec--plugins))
- `resync_period` (String) how often the datasource is refreshed, defaults to 5m if not set
- `values_from` (Attributes List) environments variables from secrets or config maps (see [below for nested schema](#nestedatt--spec--values_from))

<a id="nestedatt--spec--datasource"></a>
### Nested Schema for `spec.datasource`

Optional:

- `access` (String)
- `basic_auth` (Boolean)
- `basic_auth_user` (String)
- `database` (String)
- `editable` (Boolean) Deprecated field, it has no effect
- `is_default` (Boolean)
- `json_data` (Map of String)
- `name` (String)
- `org_id` (Number) Deprecated field, it has no effect
- `secure_json_data` (Map of String)
- `type` (String)
- `uid` (String)
- `url` (String)
- `user` (String)


<a id="nestedatt--spec--instance_selector"></a>
### Nested Schema for `spec.instance_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--instance_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--instance_selector--match_expressions"></a>
### Nested Schema for `spec.instance_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--plugins"></a>
### Nested Schema for `spec.plugins`

Required:

- `name` (String)
- `version` (String)


<a id="nestedatt--spec--values_from"></a>
### Nested Schema for `spec.values_from`

Required:

- `target_path` (String)
- `value_from` (Attributes) (see [below for nested schema](#nestedatt--spec--values_from--value_from))

<a id="nestedatt--spec--values_from--value_from"></a>
### Nested Schema for `spec.values_from.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--values_from--value_from--config_map_key_ref))
- `secret_key_ref` (Attributes) Selects a key of a Secret. (see [below for nested schema](#nestedatt--spec--values_from--value_from--secret_key_ref))

<a id="nestedatt--spec--values_from--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.values_from.value_from.config_map_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--values_from--value_from--secret_key_ref"></a>
### Nested Schema for `spec.values_from.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from. Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `optional` (Boolean) Specify whether the Secret or its key must be defined
