---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_trino_stackable_tech_trino_catalog_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "trino.stackable.tech"
description: |-
  Auto-generated derived type for TrinoCatalogSpec via 'CustomResource'
---

# k8s_trino_stackable_tech_trino_catalog_v1alpha1_manifest (Data Source)

Auto-generated derived type for TrinoCatalogSpec via 'CustomResource'

## Example Usage

```terraform
data "k8s_trino_stackable_tech_trino_catalog_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    connector = {}
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) The TrinoCatalog resource can be used to define catalogs in Kubernetes objects. Read more about it in the [Trino operator concept docs](https://docs.stackable.tech/home/nightly/trino/concepts) and the [Trino operator usage guide](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/). The documentation also contains a list of all the supported backends. (see [below for nested schema](#nestedatt--spec))

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

- `connector` (Attributes) The 'connector' defines which connector is used. (see [below for nested schema](#nestedatt--spec--connector))

Optional:

- `config_overrides` (Map of String) The 'configOverrides' allow overriding arbitrary Trino settings. For example, for Hive you could add 'hive.metastore.username: trino'.

<a id="nestedatt--spec--connector"></a>
### Nested Schema for `spec.connector`

Optional:

- `black_hole` (Map of String) A [Black Hole](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/black-hole) connector.
- `delta_lake` (Attributes) An [Delta Lake](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/delta-lake) connector. (see [below for nested schema](#nestedatt--spec--connector--delta_lake))
- `generic` (Attributes) A [generic](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/generic) connector. (see [below for nested schema](#nestedatt--spec--connector--generic))
- `google_sheet` (Attributes) A [Google sheets](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/google-sheets) connector. (see [below for nested schema](#nestedatt--spec--connector--google_sheet))
- `hive` (Attributes) An [Apache Hive](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/hive) connector. (see [below for nested schema](#nestedatt--spec--connector--hive))
- `iceberg` (Attributes) An [Apache Iceberg](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/iceberg) connector. (see [below for nested schema](#nestedatt--spec--connector--iceberg))
- `tpcds` (Map of String) A [TPC-DS](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/tpcds) connector.
- `tpch` (Map of String) A [TPC-H](https://docs.stackable.tech/home/nightly/trino/usage-guide/catalogs/tpch) connector.

<a id="nestedatt--spec--connector--delta_lake"></a>
### Nested Schema for `spec.connector.delta_lake`

Required:

- `metastore` (Attributes) Mandatory connection to a Hive Metastore, which will be used as a storage for metadata. (see [below for nested schema](#nestedatt--spec--connector--delta_lake--metastore))

Optional:

- `hdfs` (Attributes) Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS. (see [below for nested schema](#nestedatt--spec--connector--delta_lake--hdfs))
- `s3` (Attributes) Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3). (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3))

<a id="nestedatt--spec--connector--delta_lake--metastore"></a>
### Nested Schema for `spec.connector.delta_lake.metastore`

Required:

- `config_map` (String) Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.


<a id="nestedatt--spec--connector--delta_lake--hdfs"></a>
### Nested Schema for `spec.connector.delta_lake.hdfs`

Required:

- `config_map` (String) Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.


<a id="nestedatt--spec--connector--delta_lake--s3"></a>
### Nested Schema for `spec.connector.delta_lake.s3`

Optional:

- `inline` (Attributes) S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3). (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3--inline))
- `reference` (String)

<a id="nestedatt--spec--connector--delta_lake--s3--inline"></a>
### Nested Schema for `spec.connector.delta_lake.s3.inline`

Required:

- `host` (String) Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.

Optional:

- `access_style` (String) Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).
- `credentials` (Attributes) If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient. (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3--inline--credentials))
- `port` (Number) Port the S3 server listens on. If not specified the product will determine the port to use.
- `tls` (Attributes) Use a TLS connection. If not specified no TLS will be used. (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3--inline--tls))

<a id="nestedatt--spec--connector--delta_lake--s3--inline--credentials"></a>
### Nested Schema for `spec.connector.delta_lake.s3.inline.credentials`

Required:

- `secret_class` (String) [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.

Optional:

- `scope` (Attributes) [Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass). (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3--inline--credentials--scope))

<a id="nestedatt--spec--connector--delta_lake--s3--inline--credentials--scope"></a>
### Nested Schema for `spec.connector.delta_lake.s3.inline.credentials.scope`

Optional:

- `listener_volumes` (List of String) The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.
- `node` (Boolean) The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.
- `pod` (Boolean) The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.
- `services` (List of String) The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.



<a id="nestedatt--spec--connector--delta_lake--s3--inline--tls"></a>
### Nested Schema for `spec.connector.delta_lake.s3.inline.tls`

Required:

- `verification` (Attributes) The verification method used to verify the certificates of the server and/or the client. (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3--inline--tls--verification))

<a id="nestedatt--spec--connector--delta_lake--s3--inline--tls--verification"></a>
### Nested Schema for `spec.connector.delta_lake.s3.inline.tls.verification`

Optional:

- `none` (Map of String) Use TLS but don't verify certificates.
- `server` (Attributes) Use TLS and a CA certificate to verify the server. (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3--inline--tls--verification--server))

<a id="nestedatt--spec--connector--delta_lake--s3--inline--tls--verification--server"></a>
### Nested Schema for `spec.connector.delta_lake.s3.inline.tls.verification.server`

Required:

- `ca_cert` (Attributes) CA cert to verify the server. (see [below for nested schema](#nestedatt--spec--connector--delta_lake--s3--inline--tls--verification--server--ca_cert))

<a id="nestedatt--spec--connector--delta_lake--s3--inline--tls--verification--server--ca_cert"></a>
### Nested Schema for `spec.connector.delta_lake.s3.inline.tls.verification.server.ca_cert`

Optional:

- `secret_class` (String) Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.
- `web_pki` (Map of String) Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.








<a id="nestedatt--spec--connector--generic"></a>
### Nested Schema for `spec.connector.generic`

Required:

- `connector_name` (String) Name of the Trino connector. Will be passed to 'connector.name'.

Optional:

- `properties` (Attributes) A map of properties to put in the connector configuration file. They can be specified either as a raw value or be read from a Secret or ConfigMap. (see [below for nested schema](#nestedatt--spec--connector--generic--properties))

<a id="nestedatt--spec--connector--generic--properties"></a>
### Nested Schema for `spec.connector.generic.properties`

Optional:

- `value` (String)
- `value_from_config_map` (Attributes) Selects a key from a ConfigMap. (see [below for nested schema](#nestedatt--spec--connector--generic--properties--value_from_config_map))
- `value_from_secret` (Attributes) SecretKeySelector selects a key of a Secret. (see [below for nested schema](#nestedatt--spec--connector--generic--properties--value_from_secret))

<a id="nestedatt--spec--connector--generic--properties--value_from_config_map"></a>
### Nested Schema for `spec.connector.generic.properties.value_from_config_map`

Required:

- `key` (String) The key to select.
- `name` (String) Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

Optional:

- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--connector--generic--properties--value_from_secret"></a>
### Nested Schema for `spec.connector.generic.properties.value_from_secret`

Required:

- `key` (String) The key of the secret to select from. Must be a valid secret key.
- `name` (String) Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names

Optional:

- `optional` (Boolean) Specify whether the Secret or its key must be defined




<a id="nestedatt--spec--connector--google_sheet"></a>
### Nested Schema for `spec.connector.google_sheet`

Required:

- `credentials_secret` (String) The Secret containing the Google API JSON key file. The key used from the Secret is 'credentials'.
- `metadata_sheet_id` (String) Sheet ID of the spreadsheet, that contains the table mapping.

Optional:

- `cache` (Attributes) Cache the contents of sheets. This is used to reduce Google Sheets API usage and latency. (see [below for nested schema](#nestedatt--spec--connector--google_sheet--cache))

<a id="nestedatt--spec--connector--google_sheet--cache"></a>
### Nested Schema for `spec.connector.google_sheet.cache`

Optional:

- `sheets_data_expire_after_write` (String) How long to cache spreadsheet data or metadata, defaults to '5m'.
- `sheets_data_max_cache_size` (String) Maximum number of spreadsheets to cache, defaults to 1000.



<a id="nestedatt--spec--connector--hive"></a>
### Nested Schema for `spec.connector.hive`

Required:

- `metastore` (Attributes) Mandatory connection to a Hive Metastore, which will be used as a storage for metadata. (see [below for nested schema](#nestedatt--spec--connector--hive--metastore))

Optional:

- `hdfs` (Attributes) Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS. (see [below for nested schema](#nestedatt--spec--connector--hive--hdfs))
- `s3` (Attributes) Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3). (see [below for nested schema](#nestedatt--spec--connector--hive--s3))

<a id="nestedatt--spec--connector--hive--metastore"></a>
### Nested Schema for `spec.connector.hive.metastore`

Required:

- `config_map` (String) Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.


<a id="nestedatt--spec--connector--hive--hdfs"></a>
### Nested Schema for `spec.connector.hive.hdfs`

Required:

- `config_map` (String) Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.


<a id="nestedatt--spec--connector--hive--s3"></a>
### Nested Schema for `spec.connector.hive.s3`

Optional:

- `inline` (Attributes) S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3). (see [below for nested schema](#nestedatt--spec--connector--hive--s3--inline))
- `reference` (String)

<a id="nestedatt--spec--connector--hive--s3--inline"></a>
### Nested Schema for `spec.connector.hive.s3.inline`

Required:

- `host` (String) Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.

Optional:

- `access_style` (String) Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).
- `credentials` (Attributes) If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient. (see [below for nested schema](#nestedatt--spec--connector--hive--s3--inline--credentials))
- `port` (Number) Port the S3 server listens on. If not specified the product will determine the port to use.
- `tls` (Attributes) Use a TLS connection. If not specified no TLS will be used. (see [below for nested schema](#nestedatt--spec--connector--hive--s3--inline--tls))

<a id="nestedatt--spec--connector--hive--s3--inline--credentials"></a>
### Nested Schema for `spec.connector.hive.s3.inline.credentials`

Required:

- `secret_class` (String) [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.

Optional:

- `scope` (Attributes) [Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass). (see [below for nested schema](#nestedatt--spec--connector--hive--s3--inline--credentials--scope))

<a id="nestedatt--spec--connector--hive--s3--inline--credentials--scope"></a>
### Nested Schema for `spec.connector.hive.s3.inline.credentials.scope`

Optional:

- `listener_volumes` (List of String) The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.
- `node` (Boolean) The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.
- `pod` (Boolean) The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.
- `services` (List of String) The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.



<a id="nestedatt--spec--connector--hive--s3--inline--tls"></a>
### Nested Schema for `spec.connector.hive.s3.inline.tls`

Required:

- `verification` (Attributes) The verification method used to verify the certificates of the server and/or the client. (see [below for nested schema](#nestedatt--spec--connector--hive--s3--inline--tls--verification))

<a id="nestedatt--spec--connector--hive--s3--inline--tls--verification"></a>
### Nested Schema for `spec.connector.hive.s3.inline.tls.verification`

Optional:

- `none` (Map of String) Use TLS but don't verify certificates.
- `server` (Attributes) Use TLS and a CA certificate to verify the server. (see [below for nested schema](#nestedatt--spec--connector--hive--s3--inline--tls--verification--server))

<a id="nestedatt--spec--connector--hive--s3--inline--tls--verification--server"></a>
### Nested Schema for `spec.connector.hive.s3.inline.tls.verification.server`

Required:

- `ca_cert` (Attributes) CA cert to verify the server. (see [below for nested schema](#nestedatt--spec--connector--hive--s3--inline--tls--verification--server--ca_cert))

<a id="nestedatt--spec--connector--hive--s3--inline--tls--verification--server--ca_cert"></a>
### Nested Schema for `spec.connector.hive.s3.inline.tls.verification.server.ca_cert`

Optional:

- `secret_class` (String) Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.
- `web_pki` (Map of String) Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.








<a id="nestedatt--spec--connector--iceberg"></a>
### Nested Schema for `spec.connector.iceberg`

Required:

- `metastore` (Attributes) Mandatory connection to a Hive Metastore, which will be used as a storage for metadata. (see [below for nested schema](#nestedatt--spec--connector--iceberg--metastore))

Optional:

- `hdfs` (Attributes) Connection to an HDFS cluster. Please make sure that the underlying Hive metastore also has access to the HDFS. (see [below for nested schema](#nestedatt--spec--connector--iceberg--hdfs))
- `s3` (Attributes) Connection to an S3 store. Please make sure that the underlying Hive metastore also has access to the S3 store. Learn more about S3 configuration in the [S3 concept docs](https://docs.stackable.tech/home/nightly/concepts/s3). (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3))

<a id="nestedatt--spec--connector--iceberg--metastore"></a>
### Nested Schema for `spec.connector.iceberg.metastore`

Required:

- `config_map` (String) Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the Hive metastore.


<a id="nestedatt--spec--connector--iceberg--hdfs"></a>
### Nested Schema for `spec.connector.iceberg.hdfs`

Required:

- `config_map` (String) Name of the [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) providing information about the HDFS cluster.


<a id="nestedatt--spec--connector--iceberg--s3"></a>
### Nested Schema for `spec.connector.iceberg.s3`

Optional:

- `inline` (Attributes) S3 connection definition as a resource. Learn more on the [S3 concept documentation](https://docs.stackable.tech/home/nightly/concepts/s3). (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3--inline))
- `reference` (String)

<a id="nestedatt--spec--connector--iceberg--s3--inline"></a>
### Nested Schema for `spec.connector.iceberg.s3.inline`

Required:

- `host` (String) Host of the S3 server without any protocol or port. For example: 'west1.my-cloud.com'.

Optional:

- `access_style` (String) Which access style to use. Defaults to virtual hosted-style as most of the data products out there. Have a look at the [AWS documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html).
- `credentials` (Attributes) If the S3 uses authentication you have to specify you S3 credentials. In the most cases a [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) providing 'accessKey' and 'secretKey' is sufficient. (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3--inline--credentials))
- `port` (Number) Port the S3 server listens on. If not specified the product will determine the port to use.
- `tls` (Attributes) Use a TLS connection. If not specified no TLS will be used. (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3--inline--tls))

<a id="nestedatt--spec--connector--iceberg--s3--inline--credentials"></a>
### Nested Schema for `spec.connector.iceberg.s3.inline.credentials`

Required:

- `secret_class` (String) [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) containing the LDAP bind credentials.

Optional:

- `scope` (Attributes) [Scope](https://docs.stackable.tech/home/nightly/secret-operator/scope) of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass). (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3--inline--credentials--scope))

<a id="nestedatt--spec--connector--iceberg--s3--inline--credentials--scope"></a>
### Nested Schema for `spec.connector.iceberg.s3.inline.credentials.scope`

Optional:

- `listener_volumes` (List of String) The listener volume scope allows Node and Service scopes to be inferred from the applicable listeners. This must correspond to Volume names in the Pod that mount Listeners.
- `node` (Boolean) The node scope is resolved to the name of the Kubernetes Node object that the Pod is running on. This will typically be the DNS name of the node.
- `pod` (Boolean) The pod scope is resolved to the name of the Kubernetes Pod. This allows the secret to differentiate between StatefulSet replicas.
- `services` (List of String) The service scope allows Pod objects to specify custom scopes. This should typically correspond to Service objects that the Pod participates in.



<a id="nestedatt--spec--connector--iceberg--s3--inline--tls"></a>
### Nested Schema for `spec.connector.iceberg.s3.inline.tls`

Required:

- `verification` (Attributes) The verification method used to verify the certificates of the server and/or the client. (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3--inline--tls--verification))

<a id="nestedatt--spec--connector--iceberg--s3--inline--tls--verification"></a>
### Nested Schema for `spec.connector.iceberg.s3.inline.tls.verification`

Optional:

- `none` (Map of String) Use TLS but don't verify certificates.
- `server` (Attributes) Use TLS and a CA certificate to verify the server. (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3--inline--tls--verification--server))

<a id="nestedatt--spec--connector--iceberg--s3--inline--tls--verification--server"></a>
### Nested Schema for `spec.connector.iceberg.s3.inline.tls.verification.server`

Required:

- `ca_cert` (Attributes) CA cert to verify the server. (see [below for nested schema](#nestedatt--spec--connector--iceberg--s3--inline--tls--verification--server--ca_cert))

<a id="nestedatt--spec--connector--iceberg--s3--inline--tls--verification--server--ca_cert"></a>
### Nested Schema for `spec.connector.iceberg.s3.inline.tls.verification.server.ca_cert`

Optional:

- `secret_class` (String) Name of the [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) which will provide the CA certificate. Note that a SecretClass does not need to have a key but can also work with just a CA certificate, so if you got provided with a CA cert but don't have access to the key you can still use this method.
- `web_pki` (Map of String) Use TLS and the CA certificates trusted by the common web browsers to verify the server. This can be useful when you e.g. use public AWS S3 or other public available services.
