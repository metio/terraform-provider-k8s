---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_apps_gitlab_com_git_lab_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "apps.gitlab.com"
description: |-
  GitLab is a complete DevOps platform, delivered in a single application.
---

# k8s_apps_gitlab_com_git_lab_v1beta1_manifest (Data Source)

GitLab is a complete DevOps platform, delivered in a single application.

## Example Usage

```terraform
data "k8s_apps_gitlab_com_git_lab_v1beta1_manifest" "example" {
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

- `spec` (Attributes) Specification of the desired behavior of a GitLab instance. (see [below for nested schema](#nestedatt--spec))

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

- `chart` (Attributes) The specification of GitLab Chart that is used to deploy the instance. (see [below for nested schema](#nestedatt--spec--chart))

<a id="nestedatt--spec--chart"></a>
### Nested Schema for `spec.chart`

Optional:

- `values` (Map of String) ChartValues is the set of Helm values that is used to render the GitLab Chart.
- `version` (String) ChartVersion is the semantic version of the GitLab Chart.