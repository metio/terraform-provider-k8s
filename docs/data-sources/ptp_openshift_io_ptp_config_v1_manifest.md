---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_ptp_openshift_io_ptp_config_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "ptp.openshift.io"
description: |-
  PtpConfig is the Schema for the ptpconfigs API
---

# k8s_ptp_openshift_io_ptp_config_v1_manifest (Data Source)

PtpConfig is the Schema for the ptpconfigs API

## Example Usage

```terraform
data "k8s_ptp_openshift_io_ptp_config_v1_manifest" "example" {
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

- `spec` (Attributes) PtpConfigSpec defines the desired state of PtpConfig (see [below for nested schema](#nestedatt--spec))

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

- `profile` (Attributes List) (see [below for nested schema](#nestedatt--spec--profile))
- `recommend` (Attributes List) (see [below for nested schema](#nestedatt--spec--recommend))

<a id="nestedatt--spec--profile"></a>
### Nested Schema for `spec.profile`

Required:

- `name` (String)

Optional:

- `interface` (String)
- `phc2sys_conf` (String)
- `phc2sys_opts` (String)
- `plugins` (Map of String)
- `ptp4l_conf` (String)
- `ptp4l_opts` (String)
- `ptp_clock_threshold` (Attributes) (see [below for nested schema](#nestedatt--spec--profile--ptp_clock_threshold))
- `ptp_scheduling_policy` (String)
- `ptp_scheduling_priority` (Number)
- `ptp_settings` (Map of String)
- `synce4l_conf` (String)
- `synce4l_opts` (String)
- `ts2phc_conf` (String)
- `ts2phc_opts` (String)

<a id="nestedatt--spec--profile--ptp_clock_threshold"></a>
### Nested Schema for `spec.profile.ptp_clock_threshold`

Optional:

- `hold_over_timeout` (Number) clock state to stay in holdover state in secs
- `max_offset_threshold` (Number) max offset in nano secs
- `min_offset_threshold` (Number) min offset in nano secs



<a id="nestedatt--spec--recommend"></a>
### Nested Schema for `spec.recommend`

Required:

- `priority` (Number)
- `profile` (String)

Optional:

- `match` (Attributes List) (see [below for nested schema](#nestedatt--spec--recommend--match))

<a id="nestedatt--spec--recommend--match"></a>
### Nested Schema for `spec.recommend.match`

Optional:

- `node_label` (String)
- `node_name` (String)
