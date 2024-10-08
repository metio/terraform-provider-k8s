---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_velero_io_server_status_request_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "velero.io"
description: |-
  ServerStatusRequest is a request to access current status information about the Velero server.
---

# k8s_velero_io_server_status_request_v1_manifest (Data Source)

ServerStatusRequest is a request to access current status information about the Velero server.

## Example Usage

```terraform
data "k8s_velero_io_server_status_request_v1_manifest" "example" {
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

- `spec` (Map of String) ServerStatusRequestSpec is the specification for a ServerStatusRequest.

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
