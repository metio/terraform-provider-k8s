---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1 Resource - terraform-provider-k8s"
subcategory: "graphql.gloo.solo.io"
description: |-
  
---

# k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1 (Resource)



## Example Usage

```terraform
resource "k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Optional:

- `allowed_query_hashes` (List of String)
- `executable_schema` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema))
- `namespaced_statuses` (Attributes) (see [below for nested schema](#nestedatt--spec--namespaced_statuses))
- `options` (Attributes) (see [below for nested schema](#nestedatt--spec--options))
- `persisted_query_cache_config` (Attributes) (see [below for nested schema](#nestedatt--spec--persisted_query_cache_config))
- `stat_prefix` (String)
- `stitched_schema` (Attributes) (see [below for nested schema](#nestedatt--spec--stitched_schema))

<a id="nestedatt--spec--executable_schema"></a>
### Nested Schema for `spec.executable_schema`

Optional:

- `executor` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor))
- `grpc_descriptor_registry` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--grpc_descriptor_registry))
- `schema_definition` (String)

<a id="nestedatt--spec--executable_schema--executor"></a>
### Nested Schema for `spec.executable_schema.executor`

Optional:

- `local` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--local))
- `remote` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote))

<a id="nestedatt--spec--executable_schema--executor--local"></a>
### Nested Schema for `spec.executable_schema.executor.remote`

Optional:

- `enable_introspection` (Boolean)
- `options` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--options))
- `resolutions` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions))

<a id="nestedatt--spec--executable_schema--executor--remote--options"></a>
### Nested Schema for `spec.executable_schema.executor.remote.options`

Optional:

- `max_depth` (Number)


<a id="nestedatt--spec--executable_schema--executor--remote--resolutions"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions`

Optional:

- `grpc_resolver` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--grpc_resolver))
- `mock_resolver` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--mock_resolver))
- `rest_resolver` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--rest_resolver))
- `stat_prefix` (String)

<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--grpc_resolver"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix`

Optional:

- `request_transform` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--request_transform))
- `span_name` (String)
- `upstream_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--upstream_ref))

<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--request_transform"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix.upstream_ref`

Optional:

- `method_name` (String)
- `outgoing_message_json` (Dynamic)
- `request_metadata` (Map of String)
- `service_name` (String)


<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--upstream_ref"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix.upstream_ref`

Optional:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--mock_resolver"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix`

Optional:

- `async_response` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--async_response))
- `error_response` (String)
- `sync_response` (Dynamic)

<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--async_response"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix.sync_response`

Optional:

- `delay` (String)
- `response` (Dynamic)



<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--rest_resolver"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix`

Optional:

- `request` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--request))
- `response` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--response))
- `span_name` (String)
- `upstream_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--upstream_ref))

<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--request"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix.upstream_ref`

Optional:

- `body` (Dynamic)
- `headers` (Map of String)
- `query_params` (Map of String)


<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--response"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix.upstream_ref`

Optional:

- `result_root` (String)
- `setters` (Map of String)


<a id="nestedatt--spec--executable_schema--executor--remote--resolutions--stat_prefix--upstream_ref"></a>
### Nested Schema for `spec.executable_schema.executor.remote.resolutions.stat_prefix.upstream_ref`

Optional:

- `name` (String)
- `namespace` (String)





<a id="nestedatt--spec--executable_schema--executor--remote"></a>
### Nested Schema for `spec.executable_schema.executor.remote`

Optional:

- `headers` (Map of String)
- `query_params` (Map of String)
- `span_name` (String)
- `upstream_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--executor--remote--upstream_ref))

<a id="nestedatt--spec--executable_schema--executor--remote--upstream_ref"></a>
### Nested Schema for `spec.executable_schema.executor.remote.upstream_ref`

Optional:

- `name` (String)
- `namespace` (String)




<a id="nestedatt--spec--executable_schema--grpc_descriptor_registry"></a>
### Nested Schema for `spec.executable_schema.grpc_descriptor_registry`

Optional:

- `proto_descriptor` (String)
- `proto_descriptor_bin` (String)
- `proto_refs_list` (Attributes) (see [below for nested schema](#nestedatt--spec--executable_schema--grpc_descriptor_registry--proto_refs_list))

<a id="nestedatt--spec--executable_schema--grpc_descriptor_registry--proto_refs_list"></a>
### Nested Schema for `spec.executable_schema.grpc_descriptor_registry.proto_refs_list`

Optional:

- `config_map_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--executable_schema--grpc_descriptor_registry--proto_refs_list--config_map_refs))

<a id="nestedatt--spec--executable_schema--grpc_descriptor_registry--proto_refs_list--config_map_refs"></a>
### Nested Schema for `spec.executable_schema.grpc_descriptor_registry.proto_refs_list.config_map_refs`

Optional:

- `name` (String)
- `namespace` (String)





<a id="nestedatt--spec--namespaced_statuses"></a>
### Nested Schema for `spec.namespaced_statuses`

Optional:

- `statuses` (Dynamic)


<a id="nestedatt--spec--options"></a>
### Nested Schema for `spec.options`

Optional:

- `log_sensitive_info` (Boolean)


<a id="nestedatt--spec--persisted_query_cache_config"></a>
### Nested Schema for `spec.persisted_query_cache_config`

Optional:

- `cache_size` (Number)


<a id="nestedatt--spec--stitched_schema"></a>
### Nested Schema for `spec.stitched_schema`

Optional:

- `subschemas` (Attributes List) (see [below for nested schema](#nestedatt--spec--stitched_schema--subschemas))

<a id="nestedatt--spec--stitched_schema--subschemas"></a>
### Nested Schema for `spec.stitched_schema.subschemas`

Optional:

- `name` (String)
- `namespace` (String)
- `type_merge` (Attributes) (see [below for nested schema](#nestedatt--spec--stitched_schema--subschemas--type_merge))

<a id="nestedatt--spec--stitched_schema--subschemas--type_merge"></a>
### Nested Schema for `spec.stitched_schema.subschemas.type_merge`

Optional:

- `args` (Map of String)
- `query_name` (String)
- `selection_set` (String)


