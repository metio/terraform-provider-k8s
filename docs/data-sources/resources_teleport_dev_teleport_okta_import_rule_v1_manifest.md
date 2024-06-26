---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_resources_teleport_dev_teleport_okta_import_rule_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "resources.teleport.dev"
description: |-
  OktaImportRule is the Schema for the oktaimportrules API
---

# k8s_resources_teleport_dev_teleport_okta_import_rule_v1_manifest (Data Source)

OktaImportRule is the Schema for the oktaimportrules API

## Example Usage

```terraform
data "k8s_resources_teleport_dev_teleport_okta_import_rule_v1_manifest" "example" {
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

- `spec` (Attributes) OktaImportRule resource definition v1 from Teleport (see [below for nested schema](#nestedatt--spec))

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

- `mappings` (Attributes List) Mappings is a list of matches that will map match conditions to labels. (see [below for nested schema](#nestedatt--spec--mappings))
- `priority` (Number) Priority represents the priority of the rule application. Lower numbered rules will be applied first.

<a id="nestedatt--spec--mappings"></a>
### Nested Schema for `spec.mappings`

Optional:

- `add_labels` (Attributes) AddLabels specifies which labels to add if any of the previous matches match. (see [below for nested schema](#nestedatt--spec--mappings--add_labels))
- `match` (Attributes List) Match is a set of matching rules for this mapping. If any of these match, then the mapping will be applied. (see [below for nested schema](#nestedatt--spec--mappings--match))

<a id="nestedatt--spec--mappings--add_labels"></a>
### Nested Schema for `spec.mappings.add_labels`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--spec--mappings--match"></a>
### Nested Schema for `spec.mappings.match`

Optional:

- `app_ids` (List of String) AppIDs is a list of app IDs to match against.
- `app_name_regexes` (List of String) AppNameRegexes is a list of regexes to match against app names.
- `group_ids` (List of String) GroupIDs is a list of group IDs to match against.
- `group_name_regexes` (List of String) GroupNameRegexes is a list of regexes to match against group names.
