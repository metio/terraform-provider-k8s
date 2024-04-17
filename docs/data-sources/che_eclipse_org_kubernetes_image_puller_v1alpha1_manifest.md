---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "che.eclipse.org"
description: |-
  KubernetesImagePuller is the Schema for the kubernetesimagepullers API
---

# k8s_che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest (Data Source)

KubernetesImagePuller is the Schema for the kubernetesimagepullers API

## Example Usage

```terraform
data "k8s_che_eclipse_org_kubernetes_image_puller_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) KubernetesImagePullerSpec defines the desired state of KubernetesImagePuller (see [below for nested schema](#nestedatt--spec))

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

- `affinity` (String)
- `caching_cpu_limit` (String)
- `caching_cpu_request` (String)
- `caching_interval_hours` (String)
- `caching_memory_limit` (String)
- `caching_memory_request` (String)
- `config_map_name` (String)
- `daemonset_name` (String)
- `deployment_name` (String)
- `image_pull_secrets` (String)
- `image_puller_image` (String)
- `images` (String)
- `node_selector` (String)