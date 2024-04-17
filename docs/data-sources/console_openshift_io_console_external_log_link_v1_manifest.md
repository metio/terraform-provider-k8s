---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_console_openshift_io_console_external_log_link_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "console.openshift.io"
description: |-
  ConsoleExternalLogLink is an extension for customizing OpenShift web console log links.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).
---

# k8s_console_openshift_io_console_external_log_link_v1_manifest (Data Source)

ConsoleExternalLogLink is an extension for customizing OpenShift web console log links.  Compatibility level 2: Stable within a major release for a minimum of 9 months or 3 minor releases (whichever is longer).

## Example Usage

```terraform
data "k8s_console_openshift_io_console_external_log_link_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) ConsoleExternalLogLinkSpec is the desired log link configuration. The log link will appear on the logs tab of the pod details page. (see [below for nested schema](#nestedatt--spec))

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

- `href_template` (String) hrefTemplate is an absolute secure URL (must use https) for the log link including variables to be replaced. Variables are specified in the URL with the format ${variableName}, for instance, ${containerName} and will be replaced with the corresponding values from the resource. Resource is a pod. Supported variables are: - ${resourceName} - name of the resource which containes the logs - ${resourceUID} - UID of the resource which contains the logs - e.g. '11111111-2222-3333-4444-555555555555' - ${containerName} - name of the resource's container that contains the logs - ${resourceNamespace} - namespace of the resource that contains the logs - ${resourceNamespaceUID} - namespace UID of the resource that contains the logs - ${podLabels} - JSON representation of labels matching the pod with the logs - e.g. '{'key1':'value1','key2':'value2'}'  e.g., https://example.com/logs?resourceName=${resourceName}&containerName=${containerName}&resourceNamespace=${resourceNamespace}&podLabels=${podLabels}
- `text` (String) text is the display text for the link

Optional:

- `namespace_filter` (String) namespaceFilter is a regular expression used to restrict a log link to a matching set of namespaces (e.g., '^openshift-'). The string is converted into a regular expression using the JavaScript RegExp constructor. If not specified, links will be displayed for all the namespaces.