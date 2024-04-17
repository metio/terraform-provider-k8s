---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_oracle_db_anthosapis_com_export_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "oracle.db.anthosapis.com"
description: |-
  Export is the Schema for the exports API.
---

# k8s_oracle_db_anthosapis_com_export_v1alpha1_manifest (Data Source)

Export is the Schema for the exports API.

## Example Usage

```terraform
data "k8s_oracle_db_anthosapis_com_export_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) ExportSpec defines the desired state of Export (see [below for nested schema](#nestedatt--spec))

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

- `database_name` (String) DatabaseName is the database resource name within Instance to export from.
- `instance` (String) Instance is the resource name within namespace to export from.

Optional:

- `export_object_type` (String) ExportObjectType is the type of objects to export. If omitted, the default of Schemas is assumed. Supported options at this point are: Schemas or Tables.
- `export_objects` (List of String) ExportObjects are objects, schemas or tables, exported by DataPump.
- `flashback_time` (String) FlashbackTime is an optional time. If this time is set, the SCN that most closely matches the time is found, and this SCN is used to enable the Flashback utility. The export operation is performed with data that is consistent up to this SCN.
- `gcs_log_path` (String) GcsLogPath is an optional full path in GCS. If set up ahead of time, export logs can be optionally transferred to set GCS bucket. A user is to ensure proper write access to the bucket from within the Oracle Operator.
- `gcs_path` (String) GcsPath is a full path in GCS bucket to transfer exported files to. A user is to ensure proper write access to the bucket from within the Oracle Operator.
- `type` (String) Type of the Export. If omitted, the default of DataPump is assumed.