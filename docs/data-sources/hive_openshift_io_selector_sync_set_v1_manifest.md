---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_hive_openshift_io_selector_sync_set_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "hive.openshift.io"
description: |-
  SelectorSyncSet is the Schema for the SelectorSyncSet API
---

# k8s_hive_openshift_io_selector_sync_set_v1_manifest (Data Source)

SelectorSyncSet is the Schema for the SelectorSyncSet API

## Example Usage

```terraform
data "k8s_hive_openshift_io_selector_sync_set_v1_manifest" "example" {
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

- `spec` (Attributes) SelectorSyncSetSpec defines the SyncSetCommonSpec resources and patches to sync along with a ClusterDeploymentSelector indicating which clusters the SelectorSyncSet applies to in any namespace. (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `apply_behavior` (String) ApplyBehavior indicates how resources in this syncset will be applied to the target cluster. The default value of 'Apply' indicates that resources should be applied using the 'oc apply' command. If no value is set, 'Apply' is assumed. A value of 'CreateOnly' indicates that the resource will only be created if it does not already exist in the target cluster. Otherwise, it will be left alone. A value of 'CreateOrUpdate' indicates that the resource will be created/updated without the use of the 'oc apply' command, allowing larger resources to be synced, but losing some functionality of the 'oc apply' command such as the ability to remove annotations, labels, and other map entries in general.
- `cluster_deployment_selector` (Attributes) ClusterDeploymentSelector is a LabelSelector indicating which clusters the SelectorSyncSet applies to in any namespace. (see [below for nested schema](#nestedatt--spec--cluster_deployment_selector))
- `enable_resource_templates` (Boolean) EnableResourceTemplates, if True, causes hive to honor golang text/templates in Resources. While the standard syntax is supported, it won't do you a whole lot of good as the parser does not pass a data object (i.e. there is no 'dot' for you to use). This currently exists to expose a single function: {{ fromCDLabel 'some.label/key' }} will be substituted with the string value of ClusterDeployment.Labels['some.label/key']. The empty string is interpolated if there are no labels, or if the indicated key does not exist. Note that this only works in values (not e.g. map keys) that are of type string.
- `patches` (Attributes List) Patches is the list of patches to apply. (see [below for nested schema](#nestedatt--spec--patches))
- `resource_apply_mode` (String) ResourceApplyMode indicates if the Resource apply mode is 'Upsert' (default) or 'Sync'. ApplyMode 'Upsert' indicates create and update. ApplyMode 'Sync' indicates create, update and delete.
- `resources` (List of Map of String) Resources is the list of objects to sync from RawExtension definitions.
- `secret_mappings` (Attributes List) Secrets is the list of secrets to sync along with their respective destinations. (see [below for nested schema](#nestedatt--spec--secret_mappings))

<a id="nestedatt--spec--cluster_deployment_selector"></a>
### Nested Schema for `spec.cluster_deployment_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--cluster_deployment_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--cluster_deployment_selector--match_expressions"></a>
### Nested Schema for `spec.cluster_deployment_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--patches"></a>
### Nested Schema for `spec.patches`

Required:

- `api_version` (String) APIVersion is the Group and Version of the object to be patched.
- `kind` (String) Kind is the Kind of the object to be patched.
- `name` (String) Name is the name of the object to be patched.
- `patch` (String) Patch is the patch to apply.

Optional:

- `namespace` (String) Namespace is the Namespace in which the object to patch exists. Defaults to the SyncSet's Namespace.
- `patch_type` (String) PatchType indicates the PatchType as 'strategic' (default), 'json', or 'merge'.


<a id="nestedatt--spec--secret_mappings"></a>
### Nested Schema for `spec.secret_mappings`

Required:

- `source_ref` (Attributes) SourceRef specifies the name and namespace of a secret on the management cluster (see [below for nested schema](#nestedatt--spec--secret_mappings--source_ref))
- `target_ref` (Attributes) TargetRef specifies the target name and namespace of the secret on the target cluster (see [below for nested schema](#nestedatt--spec--secret_mappings--target_ref))

<a id="nestedatt--spec--secret_mappings--source_ref"></a>
### Nested Schema for `spec.secret_mappings.source_ref`

Required:

- `name` (String) Name is the name of the secret

Optional:

- `namespace` (String) Namespace is the namespace where the secret lives. If not present for the source secret reference, it is assumed to be the same namespace as the syncset with the reference.


<a id="nestedatt--spec--secret_mappings--target_ref"></a>
### Nested Schema for `spec.secret_mappings.target_ref`

Required:

- `name` (String) Name is the name of the secret

Optional:

- `namespace` (String) Namespace is the namespace where the secret lives. If not present for the source secret reference, it is assumed to be the same namespace as the syncset with the reference.
