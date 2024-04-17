---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_keycloak_org_keycloak_backup_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "keycloak.org"
description: |-
  KeycloakBackup is the Schema for the keycloakbackups API.
---

# k8s_keycloak_org_keycloak_backup_v1alpha1_manifest (Data Source)

KeycloakBackup is the Schema for the keycloakbackups API.

## Example Usage

```terraform
data "k8s_keycloak_org_keycloak_backup_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) KeycloakBackupSpec defines the desired state of KeycloakBackup. (see [below for nested schema](#nestedatt--spec))

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

- `aws` (Attributes) If provided, an automatic database backup will be created on AWS S3 instead of a local Persistent Volume. If this property is not provided - a local Persistent Volume backup will be chosen. (see [below for nested schema](#nestedatt--spec--aws))
- `instance_selector` (Attributes) Selector for looking up Keycloak Custom Resources. (see [below for nested schema](#nestedatt--spec--instance_selector))
- `restore` (Boolean) Controls automatic restore behavior. Currently not implemented.  In the future this will be used to trigger automatic restore for a given KeycloakBackup. Each backup will correspond to a single snapshot of the database (stored either in a Persistent Volume or AWS). If a user wants to restore it, all he/she needs to do is to change this flag to true. Potentially, it will be possible to restore a single backup multiple times.
- `storage_class_name` (String) Name of the StorageClass for Postgresql Backup Persistent Volume Claim

<a id="nestedatt--spec--aws"></a>
### Nested Schema for `spec.aws`

Optional:

- `credentials_secret_name` (String) Provides a secret name used for connecting to AWS S3 Service. The secret needs to be in the following form:      apiVersion: v1     kind: Secret     metadata:       name: <Secret name>     type: Opaque     stringData:       AWS_S3_BUCKET_NAME: <S3 Bucket Name>       AWS_ACCESS_KEY_ID: <AWS Access Key ID>       AWS_SECRET_ACCESS_KEY: <AWS Secret Key>  For more information, please refer to the Operator documentation.
- `encryption_key_secret_name` (String) If provided, the database backup will be encrypted. Provides a secret name used for encrypting database data. The secret needs to be in the following form:      apiVersion: v1     kind: Secret     metadata:       name: <Secret name>     type: Opaque     stringData:       GPG_PUBLIC_KEY: <GPG Public Key>       GPG_TRUST_MODEL: <GPG Trust Model>       GPG_RECIPIENT: <GPG Recipient>  For more information, please refer to the Operator documentation.
- `schedule` (String) If specified, it will be used as a schedule for creating a CronJob.


<a id="nestedatt--spec--instance_selector"></a>
### Nested Schema for `spec.instance_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--instance_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--instance_selector--match_expressions"></a>
### Nested Schema for `spec.instance_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.