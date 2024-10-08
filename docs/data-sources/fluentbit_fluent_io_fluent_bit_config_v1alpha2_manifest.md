---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest Data Source - terraform-provider-k8s"
subcategory: "fluentbit.fluent.io"
description: |-
  FluentBitConfig is the Schema for the API
---

# k8s_fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest (Data Source)

FluentBitConfig is the Schema for the API

## Example Usage

```terraform
data "k8s_fluentbit_fluent_io_fluent_bit_config_v1alpha2_manifest" "example" {
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

- `spec` (Attributes) NamespacedFluentBitCfgSpec defines the desired state of FluentBit (see [below for nested schema](#nestedatt--spec))

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

- `cluster_multiline_parser_selector` (Attributes) Select cluster level multiline parser config (see [below for nested schema](#nestedatt--spec--cluster_multiline_parser_selector))
- `cluster_parser_selector` (Attributes) Select cluster level parser config (see [below for nested schema](#nestedatt--spec--cluster_parser_selector))
- `filter_selector` (Attributes) Select filter plugins (see [below for nested schema](#nestedatt--spec--filter_selector))
- `multiline_parser_selector` (Attributes) Select multiline parser plugins (see [below for nested schema](#nestedatt--spec--multiline_parser_selector))
- `output_selector` (Attributes) Select output plugins (see [below for nested schema](#nestedatt--spec--output_selector))
- `parser_selector` (Attributes) Select parser plugins (see [below for nested schema](#nestedatt--spec--parser_selector))
- `service` (Attributes) Service defines the global behaviour of the Fluent Bit engine. (see [below for nested schema](#nestedatt--spec--service))

<a id="nestedatt--spec--cluster_multiline_parser_selector"></a>
### Nested Schema for `spec.cluster_multiline_parser_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--cluster_multiline_parser_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--cluster_multiline_parser_selector--match_expressions"></a>
### Nested Schema for `spec.cluster_multiline_parser_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--cluster_parser_selector"></a>
### Nested Schema for `spec.cluster_parser_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--cluster_parser_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--cluster_parser_selector--match_expressions"></a>
### Nested Schema for `spec.cluster_parser_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--filter_selector"></a>
### Nested Schema for `spec.filter_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--filter_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--filter_selector--match_expressions"></a>
### Nested Schema for `spec.filter_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--multiline_parser_selector"></a>
### Nested Schema for `spec.multiline_parser_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--multiline_parser_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--multiline_parser_selector--match_expressions"></a>
### Nested Schema for `spec.multiline_parser_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--output_selector"></a>
### Nested Schema for `spec.output_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--output_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--output_selector--match_expressions"></a>
### Nested Schema for `spec.output_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--parser_selector"></a>
### Nested Schema for `spec.parser_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--parser_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--parser_selector--match_expressions"></a>
### Nested Schema for `spec.parser_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--service"></a>
### Nested Schema for `spec.service`

Optional:

- `daemon` (Boolean) If true go to background on start
- `emitter_mem_buf_limit` (String)
- `emitter_name` (String) Per-namespace re-emitter configuration
- `emitter_storage_type` (String)
- `flush_seconds` (Number) Interval to flush output
- `grace_seconds` (Number) Wait time on exit
- `hc_errors_count` (Number) the error count to meet the unhealthy requirement, this is a sum for all output plugins in a defined HC_Period, example for output error: [2022/02/16 10:44:10] [ warn] [engine] failed to flush chunk '1-1645008245.491540684.flb', retry in 7 seconds: task_id=0, input=forward.1 > output=cloudwatch_logs.3 (out_id=3)
- `hc_period` (Number) The time period by second to count the error and retry failure data point
- `hc_retry_failure_count` (Number) the retry failure count to meet the unhealthy requirement, this is a sum for all output plugins in a defined HC_Period, example for retry failure: [2022/02/16 20:11:36] [ warn] [engine] chunk '1-1645042288.260516436.flb' cannot be retried: task_id=0, input=tcp.3 > output=cloudwatch_logs.1
- `health_check` (Boolean) enable Health check feature at http://127.0.0.1:2020/api/v1/health Note: Enabling this will not automatically configure kubernetes to use fluentbit's healthcheck endpoint
- `hot_reload` (Boolean) If true enable reloading via HTTP
- `http_listen` (String) Address to listen
- `http_port` (Number) Port to listen
- `http_server` (Boolean) If true enable statistics HTTP server
- `log_file` (String) File to log diagnostic output
- `log_level` (String) Diagnostic level (error/warning/info/debug/trace)
- `parsers_file` (String) Optional 'parsers' config file (can be multiple)
- `parsers_files` (List of String) backward compatible
- `storage` (Attributes) Configure a global environment for the storage layer in Service. It is recommended to configure the volume and volumeMount separately for this storage. The hostPath type should be used for that Volume in Fluentbit daemon set. (see [below for nested schema](#nestedatt--spec--service--storage))

<a id="nestedatt--spec--service--storage"></a>
### Nested Schema for `spec.service.storage`

Optional:

- `backlog_mem_limit` (String) This option configure a hint of maximum value of memory to use when processing these records
- `checksum` (String) Enable the data integrity check when writing and reading data from the filesystem
- `delete_irrecoverable_chunks` (String) When enabled, irrecoverable chunks will be deleted during runtime, and any other irrecoverable chunk located in the configured storage path directory will be deleted when Fluent-Bit starts.
- `max_chunks_up` (Number) If the input plugin has enabled filesystem storage type, this property sets the maximum number of Chunks that can be up in memory
- `metrics` (String) If http_server option has been enabled in the Service section, this option registers a new endpoint where internal metrics of the storage layer can be consumed
- `path` (String) Select an optional location in the file system to store streams and chunks of data/
- `sync` (String) Configure the synchronization mode used to store the data into the file system
