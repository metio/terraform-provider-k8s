---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_charts_flagsmith_com_flagsmith_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "charts.flagsmith.com"
description: |-
  Flagsmith is the Schema for the flagsmiths API
---

# k8s_charts_flagsmith_com_flagsmith_v1alpha1_manifest (Data Source)

Flagsmith is the Schema for the flagsmiths API

## Example Usage

```terraform
data "k8s_charts_flagsmith_com_flagsmith_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) Spec defines the desired state of Flagsmith (see [below for nested schema](#nestedatt--spec))

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

- `api` (Map of String) Configuration how to setup the flagsmith api service.
- `frontend` (Map of String) Configuration how to setup the flagsmith frontend service.

Optional:

- `hooks` (Map of String) Configuration how to setup the flagsmith hooks.
- `influxdb` (Attributes) Configuration how to setup the flagsmith InfluxDB service. (see [below for nested schema](#nestedatt--spec--influxdb))
- `ingress` (Map of String) Configuration how to setup ingress to the flagsmith if flagsmith is using Kubernetes and not OpenShift.
- `metrics` (Map of String) Configuration how to setup the flagsmith metrics.
- `openshift` (Boolean) If flagsmith install on OpenShift set value to true otherwise false.
- `postgresql` (Attributes) Configuration how to setup the flagsmith postgresql service. (see [below for nested schema](#nestedatt--spec--postgresql))
- `service` (Map of String) Configuration how to setup the flagsmith kubernetes service.

<a id="nestedatt--spec--influxdb"></a>
### Nested Schema for `spec.influxdb`

Required:

- `enabled` (Boolean) Set to true if InfluxDB will be installed. If the value is false InfluxDB will not be installed.


<a id="nestedatt--spec--postgresql"></a>
### Nested Schema for `spec.postgresql`

Required:

- `enabled` (Boolean) Set to true if PostgreSQL will be installed. If the value is false PostgreSQL will not be installed.