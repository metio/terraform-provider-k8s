---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "imaging-ingestion.alvearie.org"
description: |-
  Provides a bidirectional proxied DIMSE Application Entity (AE) in the cluster
---

# k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest (Data Source)

Provides a bidirectional proxied DIMSE Application Entity (AE) in the cluster

## Example Usage

```terraform
data "k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) DimseProxySpec defines the desired state of DimseProxy (see [below for nested schema](#nestedatt--spec))

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

- `application_entity_title` (String) Application Entity Title
- `image_pull_policy` (String) Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
- `image_pull_secrets` (Attributes List) Image Pull Secrets (see [below for nested schema](#nestedatt--spec--image_pull_secrets))
- `nats_secure` (Boolean) Make NATS Connection Secure
- `nats_subject_channel` (String) NATS Subject Channel to use
- `nats_subject_root` (String) NATS Subject Root
- `nats_token_secret` (String) NATS Token Secret Name
- `nats_url` (String) NATS URL
- `proxy` (Attributes) DIMSE Proxy Spec (see [below for nested schema](#nestedatt--spec--proxy))
- `target_dimse_host` (String) Target Dimse Host
- `target_dimse_port` (Number) Target Dimse Port

<a id="nestedatt--spec--image_pull_secrets"></a>
### Nested Schema for `spec.image_pull_secrets`

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?


<a id="nestedatt--spec--proxy"></a>
### Nested Schema for `spec.proxy`

Optional:

- `image` (String) Image