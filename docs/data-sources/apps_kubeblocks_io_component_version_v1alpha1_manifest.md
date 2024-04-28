---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_apps_kubeblocks_io_component_version_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "apps.kubeblocks.io"
description: |-
  ComponentVersion is the Schema for the componentversions API
---

# k8s_apps_kubeblocks_io_component_version_v1alpha1_manifest (Data Source)

ComponentVersion is the Schema for the componentversions API

## Example Usage

```terraform
data "k8s_apps_kubeblocks_io_component_version_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) ComponentVersionSpec defines the desired state of ComponentVersion (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `compatibility_rules` (Attributes List) CompatibilityRules defines compatibility rules between sets of component definitions and releases. (see [below for nested schema](#nestedatt--spec--compatibility_rules))
- `releases` (Attributes List) Releases represents different releases of component instances within this ComponentVersion. (see [below for nested schema](#nestedatt--spec--releases))

<a id="nestedatt--spec--compatibility_rules"></a>
### Nested Schema for `spec.compatibility_rules`

Required:

- `comp_defs` (List of String) CompDefs specifies names for the component definitions associated with this ComponentVersion. Each name in the list can represent an exact name, or a name prefix.  For example:  - 'mysql-8.0.30-v1alpha1': Matches the exact name 'mysql-8.0.30-v1alpha1' - 'mysql-8.0.30': Matches all names starting with 'mysql-8.0.30'
- `releases` (List of String) Releases is a list of identifiers for the releases.


<a id="nestedatt--spec--releases"></a>
### Nested Schema for `spec.releases`

Required:

- `images` (Map of String) Images define the new images for different containers within the release.
- `name` (String) Name is a unique identifier for this release. Cannot be updated.
- `service_version` (String) ServiceVersion defines the version of the well-known service that the component provides. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If the release is used, it will serve as the service version for component instances, overriding the one defined in the component definition. Cannot be updated.

Optional:

- `changes` (String) Changes provides information about the changes made in this release.