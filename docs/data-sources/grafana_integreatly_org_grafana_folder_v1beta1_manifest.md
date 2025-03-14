---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_grafana_integreatly_org_grafana_folder_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "grafana.integreatly.org"
description: |-
  GrafanaFolder is the Schema for the grafanafolders API
---

# k8s_grafana_integreatly_org_grafana_folder_v1beta1_manifest (Data Source)

GrafanaFolder is the Schema for the grafanafolders API

## Example Usage

```terraform
data "k8s_grafana_integreatly_org_grafana_folder_v1beta1_manifest" "example" {
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

- `spec` (Attributes) GrafanaFolderSpec defines the desired state of GrafanaFolder (see [below for nested schema](#nestedatt--spec))

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

- `instance_selector` (Attributes) Selects Grafanas for import (see [below for nested schema](#nestedatt--spec--instance_selector))

Optional:

- `allow_cross_namespace_import` (Boolean) Enable matching Grafana instances outside the current namespace
- `parent_folder_ref` (String) Reference to an existing GrafanaFolder CR in the same namespace
- `parent_folder_uid` (String) UID of the folder in which the current folder should be created
- `permissions` (String) Raw json with folder permissions, potentially exported from Grafana
- `resync_period` (String) How often the folder is synced, defaults to 5m if not set
- `title` (String) Display name of the folder in Grafana
- `uid` (String) Manually specify the UID the Folder is created with

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
