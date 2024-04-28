---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "apacheweb.arsenal.dev"
description: |-
  
---

# k8s_apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest (Data Source)



## Example Usage

```terraform
data "k8s_apacheweb_arsenal_dev_apacheweb_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

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

- `server_name` (String)
- `size` (Number)
- `type` (String)

Optional:

- `load_balancer` (Attributes) (see [below for nested schema](#nestedatt--spec--load_balancer))
- `web_server` (Attributes) (see [below for nested schema](#nestedatt--spec--web_server))

<a id="nestedatt--spec--load_balancer"></a>
### Nested Schema for `spec.load_balancer`

Required:

- `server_port` (Number)

Optional:

- `back_end_service` (String)
- `end_points_list` (Attributes List) (see [below for nested schema](#nestedatt--spec--load_balancer--end_points_list))
- `path` (String)
- `proto` (String)
- `proxy_paths` (Attributes List) (see [below for nested schema](#nestedatt--spec--load_balancer--proxy_paths))

<a id="nestedatt--spec--load_balancer--end_points_list"></a>
### Nested Schema for `spec.load_balancer.end_points_list`

Required:

- `ip_address` (String)
- `port` (Number)
- `proto` (String)

Optional:

- `status` (Boolean)


<a id="nestedatt--spec--load_balancer--proxy_paths"></a>
### Nested Schema for `spec.load_balancer.proxy_paths`

Required:

- `path` (String)

Optional:

- `end_points_list` (Attributes List) (see [below for nested schema](#nestedatt--spec--load_balancer--proxy_paths--end_points_list))

<a id="nestedatt--spec--load_balancer--proxy_paths--end_points_list"></a>
### Nested Schema for `spec.load_balancer.proxy_paths.end_points_list`

Required:

- `ip_address` (String)
- `port` (Number)
- `proto` (String)

Optional:

- `status` (Boolean)




<a id="nestedatt--spec--web_server"></a>
### Nested Schema for `spec.web_server`

Required:

- `document_root` (String)
- `server_admin` (String)
- `server_port` (Number)