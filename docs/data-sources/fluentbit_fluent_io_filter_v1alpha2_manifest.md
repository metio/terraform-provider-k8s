---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_fluentbit_fluent_io_filter_v1alpha2_manifest Data Source - terraform-provider-k8s"
subcategory: "fluentbit.fluent.io"
description: |-
  Filter is the Schema for namespace level filter API
---

# k8s_fluentbit_fluent_io_filter_v1alpha2_manifest (Data Source)

Filter is the Schema for namespace level filter API

## Example Usage

```terraform
data "k8s_fluentbit_fluent_io_filter_v1alpha2_manifest" "example" {
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

- `spec` (Attributes) FilterSpec defines the desired state of ClusterFilter (see [below for nested schema](#nestedatt--spec))

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

- `filters` (Attributes List) A set of filter plugins in order. (see [below for nested schema](#nestedatt--spec--filters))
- `log_level` (String)
- `match` (String) A pattern to match against the tags of incoming records. It's case-sensitive and support the star (*) character as a wildcard.
- `match_regex` (String) A regular expression to match against the tags of incoming records. Use this option if you want to use the full regex syntax.

<a id="nestedatt--spec--filters"></a>
### Nested Schema for `spec.filters`

Optional:

- `aws` (Attributes) Aws defines a Aws configuration. (see [below for nested schema](#nestedatt--spec--filters--aws))
- `custom_plugin` (Attributes) CustomPlugin defines a Custom plugin configuration. (see [below for nested schema](#nestedatt--spec--filters--custom_plugin))
- `grep` (Attributes) Grep defines Grep Filter configuration. (see [below for nested schema](#nestedatt--spec--filters--grep))
- `kubernetes` (Attributes) Kubernetes defines Kubernetes Filter configuration. (see [below for nested schema](#nestedatt--spec--filters--kubernetes))
- `lua` (Attributes) Lua defines Lua Filter configuration. (see [below for nested schema](#nestedatt--spec--filters--lua))
- `modify` (Attributes) Modify defines Modify Filter configuration. (see [below for nested schema](#nestedatt--spec--filters--modify))
- `multiline` (Attributes) Multiline defines a Multiline configuration. (see [below for nested schema](#nestedatt--spec--filters--multiline))
- `nest` (Attributes) Nest defines Nest Filter configuration. (see [below for nested schema](#nestedatt--spec--filters--nest))
- `parser` (Attributes) Parser defines Parser Filter configuration. (see [below for nested schema](#nestedatt--spec--filters--parser))
- `record_modifier` (Attributes) RecordModifier defines Record Modifier Filter configuration. (see [below for nested schema](#nestedatt--spec--filters--record_modifier))
- `rewrite_tag` (Attributes) RewriteTag defines a RewriteTag configuration. (see [below for nested schema](#nestedatt--spec--filters--rewrite_tag))
- `throttle` (Attributes) Throttle defines a Throttle configuration. (see [below for nested schema](#nestedatt--spec--filters--throttle))

<a id="nestedatt--spec--filters--aws"></a>
### Nested Schema for `spec.filters.aws`

Optional:

- `account_id` (Boolean) The account ID for current EC2 instance.Default is false.
- `alias` (String) Alias for the plugin
- `ami_id` (Boolean) The EC2 instance image id.Default is false.
- `az` (Boolean) The availability zone; for example, 'us-east-1a'. Default is true.
- `ec2_instance_id` (Boolean) The EC2 instance ID.Default is true.
- `ec2_instance_type` (Boolean) The EC2 instance type.Default is false.
- `host_name` (Boolean) The hostname for current EC2 instance.Default is false.
- `imds_version` (String) Specify which version of the instance metadata service to use. Valid values are 'v1' or 'v2'.
- `private_ip` (Boolean) The EC2 instance private ip.Default is false.
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `vpc_id` (Boolean) The VPC ID for current EC2 instance.Default is false.


<a id="nestedatt--spec--filters--custom_plugin"></a>
### Nested Schema for `spec.filters.custom_plugin`

Optional:

- `config` (String)


<a id="nestedatt--spec--filters--grep"></a>
### Nested Schema for `spec.filters.grep`

Optional:

- `alias` (String) Alias for the plugin
- `exclude` (String) Exclude records which field matches the regular expression. Value Format: FIELD REGEX
- `regex` (String) Keep records which field matches the regular expression. Value Format: FIELD REGEX
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.


<a id="nestedatt--spec--filters--kubernetes"></a>
### Nested Schema for `spec.filters.kubernetes`

Optional:

- `alias` (String) Alias for the plugin
- `annotations` (Boolean) Include Kubernetes resource annotations in the extra metadata.
- `buffer_size` (String) Set the buffer size for HTTP client when reading responses from Kubernetes API server.
- `cache_use_docker_id` (Boolean) When enabled, metadata will be fetched from K8s when docker_id is changed.
- `dns_retries` (Number) DNS lookup retries N times until the network start working
- `dns_wait_time` (Number) DNS lookup interval between network status checks
- `dummy_meta` (Boolean) If set, use dummy-meta data (for test/dev purposes)
- `k8s_logging_exclude` (Boolean) Allow Kubernetes Pods to exclude their logs from the log processor (read more about it in Kubernetes Annotations section).
- `k8s_logging_parser` (Boolean) Allow Kubernetes Pods to suggest a pre-defined Parser (read more about it in Kubernetes Annotations section)
- `keep_log` (Boolean) When Keep_Log is disabled, the log field is removed from the incoming message once it has been successfully merged (Merge_Log must be enabled as well).
- `kube_ca_file` (String) CA certificate file
- `kube_ca_path` (String) Absolute path to scan for certificate files
- `kube_meta_cache_ttl` (String) configurable TTL for K8s cached metadata. By default, it is set to 0 which means TTL for cache entries is disabled and cache entries are evicted at random when capacity is reached. In order to enable this option, you should set the number to a time interval. For example, set this value to 60 or 60s and cache entries which have been created more than 60s will be evicted.
- `kube_meta_preload_cache_dir` (String) If set, Kubernetes meta-data can be cached/pre-loaded from files in JSON format in this directory, named as namespace-pod.meta
- `kube_tag_prefix` (String) When the source records comes from Tail input plugin, this option allows to specify what's the prefix used in Tail configuration.
- `kube_token_file` (String) Token file
- `kube_token_ttl` (String) configurable 'time to live' for the K8s token. By default, it is set to 600 seconds. After this time, the token is reloaded from Kube_Token_File or the Kube_Token_Command.
- `kube_url` (String) API Server end-point
- `kubelet_host` (String) kubelet host using for HTTP request, this only works when Use_Kubelet set to On.
- `kubelet_port` (Number) kubelet port using for HTTP request, this only works when useKubelet is set to On.
- `labels` (Boolean) Include Kubernetes resource labels in the extra metadata.
- `merge_log` (Boolean) When enabled, it checks if the log field content is a JSON string map, if so, it append the map fields as part of the log structure.
- `merge_log_key` (String) When Merge_Log is enabled, the filter tries to assume the log field from the incoming message is a JSON string message and make a structured representation of it at the same level of the log field in the map. Now if Merge_Log_Key is set (a string name), all the new structured fields taken from the original log content are inserted under the new key.
- `merge_log_trim` (Boolean) When Merge_Log is enabled, trim (remove possible n or r) field values.
- `merge_parser` (String) Optional parser name to specify how to parse the data contained in the log key. Recommended use is for developers or testing only.
- `regex_parser` (String) Set an alternative Parser to process record Tag and extract pod_name, namespace_name, container_name and docker_id. The parser must be registered in a parsers file (refer to parser filter-kube-test as an example).
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `tls_debug` (Number) Debug level between 0 (nothing) and 4 (every detail).
- `tls_verify` (Boolean) When enabled, turns on certificate validation when connecting to the Kubernetes API server.
- `use_journal` (Boolean) When enabled, the filter reads logs coming in Journald format.
- `use_kubelet` (Boolean) This is an optional feature flag to get metadata information from kubelet instead of calling Kube Server API to enhance the log. This could mitigate the Kube API heavy traffic issue for large cluster.


<a id="nestedatt--spec--filters--lua"></a>
### Nested Schema for `spec.filters.lua`

Required:

- `call` (String) Lua function name that will be triggered to do filtering. It's assumed that the function is declared inside the Script defined above.
- `script` (Attributes) Path to the Lua script that will be used. (see [below for nested schema](#nestedatt--spec--filters--lua--script))

Optional:

- `alias` (String) Alias for the plugin
- `protected_mode` (Boolean) If enabled, Lua script will be executed in protected mode. It prevents to crash when invalid Lua script is executed. Default is true.
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `time_as_table` (Boolean) By default when the Lua script is invoked, the record timestamp is passed as a Floating number which might lead to loss precision when the data is converted back. If you desire timestamp precision enabling this option will pass the timestamp as a Lua table with keys sec for seconds since epoch and nsec for nanoseconds.
- `type_int_key` (List of String) If these keys are matched, the fields are converted to integer. If more than one key, delimit by space. Note that starting from Fluent Bit v1.6 integer data types are preserved and not converted to double as in previous versions.

<a id="nestedatt--spec--filters--lua--script"></a>
### Nested Schema for `spec.filters.lua.script`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined



<a id="nestedatt--spec--filters--modify"></a>
### Nested Schema for `spec.filters.modify`

Optional:

- `alias` (String) Alias for the plugin
- `conditions` (Attributes List) All conditions have to be true for the rules to be applied. (see [below for nested schema](#nestedatt--spec--filters--modify--conditions))
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `rules` (Attributes List) Rules are applied in the order they appear, with each rule operating on the result of the previous rule. (see [below for nested schema](#nestedatt--spec--filters--modify--rules))

<a id="nestedatt--spec--filters--modify--conditions"></a>
### Nested Schema for `spec.filters.modify.conditions`

Optional:

- `a_key_matches` (String) Is true if a key matches regex KEY
- `key_does_not_exist` (Map of String) Is true if KEY does not exist
- `key_exists` (String) Is true if KEY exists
- `key_value_does_not_equal` (Map of String) Is true if KEY exists and its value is not VALUE
- `key_value_does_not_match` (Map of String) Is true if key KEY exists and its value does not match VALUE
- `key_value_equals` (Map of String) Is true if KEY exists and its value is VALUE
- `key_value_matches` (Map of String) Is true if key KEY exists and its value matches VALUE
- `matching_keys_do_not_have_matching_values` (Map of String) Is true if all keys matching KEY have values that do not match VALUE
- `matching_keys_have_matching_values` (Map of String) Is true if all keys matching KEY have values that match VALUE
- `no_key_matches` (String) Is true if no key matches regex KEY


<a id="nestedatt--spec--filters--modify--rules"></a>
### Nested Schema for `spec.filters.modify.rules`

Optional:

- `add` (Map of String) Add a key/value pair with key KEY and value VALUE if KEY does not exist
- `copy` (Map of String) Copy a key/value pair with key KEY to COPIED_KEY if KEY exists AND COPIED_KEY does not exist
- `hard_copy` (Map of String) Copy a key/value pair with key KEY to COPIED_KEY if KEY exists. If COPIED_KEY already exists, this field is overwritten
- `hard_rename` (Map of String) Rename a key/value pair with key KEY to RENAMED_KEY if KEY exists. If RENAMED_KEY already exists, this field is overwritten
- `remove` (String) Remove a key/value pair with key KEY if it exists
- `remove_regex` (String) Remove all key/value pairs with key matching regexp KEY
- `remove_wildcard` (String) Remove all key/value pairs with key matching wildcard KEY
- `rename` (Map of String) Rename a key/value pair with key KEY to RENAMED_KEY if KEY exists AND RENAMED_KEY does not exist
- `set` (Map of String) Add a key/value pair with key KEY and value VALUE. If KEY already exists, this field is overwritten



<a id="nestedatt--spec--filters--multiline"></a>
### Nested Schema for `spec.filters.multiline`

Required:

- `parser` (String) Specify one or multiple Multiline Parsing definitions to apply to the content. You can specify multiple multiline parsers to detect different formats by separating them with a comma.

Optional:

- `alias` (String) Alias for the plugin
- `key_content` (String) Key name that holds the content to process. Note that a Multiline Parser definition can already specify the key_content to use, but this option allows to overwrite that value for the purpose of the filter.
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.


<a id="nestedatt--spec--filters--nest"></a>
### Nested Schema for `spec.filters.nest`

Optional:

- `add_prefix` (String) Prefix affected keys with this string
- `alias` (String) Alias for the plugin
- `nest_under` (String) Nest records matching the Wildcard under this key
- `nested_under` (String) Lift records nested under the Nested_under key
- `operation` (String) Select the operation nest or lift
- `remove_prefix` (String) Remove prefix from affected keys if it matches this string
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `wildcard` (List of String) Nest records which field matches the wildcard


<a id="nestedatt--spec--filters--parser"></a>
### Nested Schema for `spec.filters.parser`

Optional:

- `alias` (String) Alias for the plugin
- `key_name` (String) Specify field name in record to parse.
- `parser` (String) Specify the parser name to interpret the field. Multiple Parser entries are allowed (split by comma).
- `preserve_key` (Boolean) Keep original Key_Name field in the parsed result. If false, the field will be removed.
- `reserve_data` (Boolean) Keep all other original fields in the parsed result. If false, all other original fields will be removed.
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `unescape_key` (Boolean) If the key is a escaped string (e.g: stringify JSON), unescape the string before to apply the parser.


<a id="nestedatt--spec--filters--record_modifier"></a>
### Nested Schema for `spec.filters.record_modifier`

Optional:

- `alias` (String) Alias for the plugin
- `allowlist_keys` (List of String) If the key is not matched, that field is removed.
- `records` (List of String) Append fields. This parameter needs key and value pair.
- `remove_keys` (List of String) If the key is matched, that field is removed.
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `uuid_keys` (List of String) If set, the plugin appends uuid to each record. The value assigned becomes the key in the map.
- `whitelist_keys` (List of String) An alias of allowlistKeys for backwards compatibility.


<a id="nestedatt--spec--filters--rewrite_tag"></a>
### Nested Schema for `spec.filters.rewrite_tag`

Optional:

- `alias` (String) Alias for the plugin
- `emitter_name` (String) When the filter emits a record under the new Tag, there is an internal emitter plugin that takes care of the job. Since this emitter expose metrics as any other component of the pipeline, you can use this property to configure an optional name for it.
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `rules` (List of String) Defines the matching criteria and the format of the Tag for the matching record. The Rule format have four components: KEY REGEX NEW_TAG KEEP.


<a id="nestedatt--spec--filters--throttle"></a>
### Nested Schema for `spec.filters.throttle`

Optional:

- `alias` (String) Alias for the plugin
- `interval` (String) Interval is the time interval expressed in 'sleep' format. e.g. 3s, 1.5m, 0.5h, etc.
- `print_status` (Boolean) PrintStatus represents whether to print status messages with current rate and the limits to information logs.
- `rate` (Number) Rate is the amount of messages for the time.
- `retry_limit` (String) RetryLimit describes how many times fluent-bit should retry to send data to a specific output. If set to false fluent-bit will try indefinetly. If set to any integer N>0 it will try at most N+1 times. Leading zeros are not allowed (values such as 007, 0150, 01 do not work). If this property is not defined fluent-bit will use the default value: 1.
- `window` (Number) Window is the amount of intervals to calculate average over.