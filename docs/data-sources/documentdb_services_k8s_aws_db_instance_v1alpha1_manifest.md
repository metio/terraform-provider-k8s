---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_documentdb_services_k8s_aws_db_instance_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "documentdb.services.k8s.aws"
description: |-
  DBInstance is the Schema for the DBInstances API
---

# k8s_documentdb_services_k8s_aws_db_instance_v1alpha1_manifest (Data Source)

DBInstance is the Schema for the DBInstances API

## Example Usage

```terraform
data "k8s_documentdb_services_k8s_aws_db_instance_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) DBInstanceSpec defines the desired state of DBInstance.Detailed information about an instance. (see [below for nested schema](#nestedatt--spec))

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

- `db_cluster_identifier` (String) The identifier of the cluster that the instance will belong to.
- `db_instance_class` (String) The compute and memory capacity of the instance; for example, db.r5.large.
- `db_instance_identifier` (String) The instance identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * The first character must be a letter.   * Cannot end with a hyphen or contain two consecutive hyphens.Example: mydbinstance
- `engine` (String) The name of the database engine to be used for this instance.Valid value: docdb

Optional:

- `auto_minor_version_upgrade` (Boolean) This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB doesnot perform minor version upgrades regardless of the value set.Default: false
- `availability_zone` (String) The Amazon EC2 Availability Zone that the instance is created in.Default: A random, system-chosen Availability Zone in the endpoint's AmazonWeb Services Region.Example: us-east-1d
- `ca_certificate_identifier` (String) The CA certificate identifier to use for the DB instance's server certificate.For more information, see Updating Your Amazon DocumentDB TLS Certificates(https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html)and Encrypting Data in Transit (https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html)in the Amazon DocumentDB Developer Guide.
- `copy_tags_to_snapshot` (Boolean) A value that indicates whether to copy tags from the DB instance to snapshotsof the DB instance. By default, tags are not copied.
- `performance_insights_enabled` (Boolean) A value that indicates whether to enable Performance Insights for the DBInstance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html).
- `performance_insights_kms_key_id` (String) The KMS key identifier for encryption of Performance Insights data.The KMS key identifier is the key ARN, key ID, alias ARN, or alias name forthe KMS key.If you do not specify a value for PerformanceInsightsKMSKeyId, then AmazonDocumentDB uses your default KMS key. There is a default KMS key for yourAmazon Web Services account. Your Amazon Web Services account has a differentdefault KMS key for each Amazon Web Services region.
- `performance_insights_kms_key_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api (see [below for nested schema](#nestedatt--spec--performance_insights_kms_key_ref))
- `preferred_maintenance_window` (String) The time range each week during which system maintenance can occur, in UniversalCoordinated Time (UTC).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid days: Mon, Tue, Wed, Thu, Fri, Sat, SunConstraints: Minimum 30-minute window.
- `promotion_tier` (Number) A value that specifies the order in which an Amazon DocumentDB replica ispromoted to the primary instance after a failure of the existing primaryinstance.Default: 1Valid values: 0-15
- `tags` (Attributes List) The tags to be assigned to the instance. You can assign up to 10 tags toan instance. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--performance_insights_kms_key_ref"></a>
### Nested Schema for `spec.performance_insights_kms_key_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--performance_insights_kms_key_ref--from))

<a id="nestedatt--spec--performance_insights_kms_key_ref--from"></a>
### Nested Schema for `spec.performance_insights_kms_key_ref.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)