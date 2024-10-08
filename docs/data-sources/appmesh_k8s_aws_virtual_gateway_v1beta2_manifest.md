---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_appmesh_k8s_aws_virtual_gateway_v1beta2_manifest Data Source - terraform-provider-k8s"
subcategory: "appmesh.k8s.aws"
description: |-
  VirtualGateway is the Schema for the virtualgateways API
---

# k8s_appmesh_k8s_aws_virtual_gateway_v1beta2_manifest (Data Source)

VirtualGateway is the Schema for the virtualgateways API

## Example Usage

```terraform
data "k8s_appmesh_k8s_aws_virtual_gateway_v1beta2_manifest" "example" {
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

- `spec` (Attributes) VirtualGatewaySpec defines the desired state of VirtualGateway refers to https://docs.aws.amazon.com/app-mesh/latest/userguide/virtual_gateways.html (see [below for nested schema](#nestedatt--spec))

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

- `aws_name` (String) AWSName is the AppMesh VirtualGateway object's name. If unspecified or empty, it defaults to be '${name}_${namespace}' of k8s VirtualGateway
- `backend_defaults` (Attributes) A reference to an object that represents the defaults for backend GatewayRoutes. (see [below for nested schema](#nestedatt--spec--backend_defaults))
- `gateway_route_selector` (Attributes) GatewayRouteSelector selects GatewayRoutes using labels to designate GatewayRoute membership. If not specified it selects all GatewayRoutes in that namespace. (see [below for nested schema](#nestedatt--spec--gateway_route_selector))
- `listeners` (Attributes List) The listener that the virtual gateway is expected to receive inbound traffic from (see [below for nested schema](#nestedatt--spec--listeners))
- `logging` (Attributes) The inbound and outbound access logging information for the virtual gateway. (see [below for nested schema](#nestedatt--spec--logging))
- `mesh_ref` (Attributes) A reference to k8s Mesh CR that this VirtualGateway belongs to. The admission controller populates it using Meshes's selector, and prevents users from setting this field. Populated by the system. Read-only. (see [below for nested schema](#nestedatt--spec--mesh_ref))
- `namespace_selector` (Attributes) NamespaceSelector selects Namespaces using labels to designate GatewayRoute membership. This field follows standard label selector semantics; if present but empty, it selects all namespaces. (see [below for nested schema](#nestedatt--spec--namespace_selector))
- `pod_selector` (Attributes) PodSelector selects Pods using labels to designate VirtualGateway membership. This field follows standard label selector semantics: if present but empty, it selects all pods within namespace. if absent, it selects no pod. (see [below for nested schema](#nestedatt--spec--pod_selector))

<a id="nestedatt--spec--backend_defaults"></a>
### Nested Schema for `spec.backend_defaults`

Optional:

- `client_policy` (Attributes) A reference to an object that represents a client policy. (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy))

<a id="nestedatt--spec--backend_defaults--client_policy"></a>
### Nested Schema for `spec.backend_defaults.client_policy`

Optional:

- `tls` (Attributes) A reference to an object that represents a Transport Layer Security (TLS) client policy. (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls))

<a id="nestedatt--spec--backend_defaults--client_policy--tls"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls`

Required:

- `validation` (Attributes) A reference to an object that represents a TLS validation context. (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--validation))

Optional:

- `certificate` (Attributes) A reference to an object that represents TLS certificate. (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--certificate))
- `enforce` (Boolean) Whether the policy is enforced. If unspecified, default settings from AWS API will be applied. Refer to AWS Docs for default settings.
- `ports` (List of String) The range of ports that the policy is enforced for.

<a id="nestedatt--spec--backend_defaults--client_policy--tls--validation"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.validation`

Required:

- `trust` (Attributes) A reference to an object that represents a TLS validation context trust (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--validation--trust))

Optional:

- `subject_alternative_names` (Attributes) Possible alternative names to consider (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--validation--subject_alternative_names))

<a id="nestedatt--spec--backend_defaults--client_policy--tls--validation--trust"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.validation.trust`

Optional:

- `acm` (Attributes) A reference to an object that represents a TLS validation context trust for an AWS Certicate Manager (ACM) certificate. (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--validation--trust--acm))
- `file` (Attributes) An object that represents a TLS validation context trust for a local file. (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--validation--trust--file))
- `sds` (Attributes) An object that represents a TLS validation context trust for a SDS certificate (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--validation--trust--sds))

<a id="nestedatt--spec--backend_defaults--client_policy--tls--validation--trust--acm"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.validation.trust.acm`

Required:

- `certificate_authority_ar_ns` (List of String) One or more ACM Amazon Resource Name (ARN)s.


<a id="nestedatt--spec--backend_defaults--client_policy--tls--validation--trust--file"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.validation.trust.file`

Required:

- `certificate_chain` (String) The certificate trust chain for a certificate stored on the file system of the virtual Gateway.


<a id="nestedatt--spec--backend_defaults--client_policy--tls--validation--trust--sds"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.validation.trust.sds`

Required:

- `secret_name` (String) The certificate trust chain for a certificate issued via SDS.



<a id="nestedatt--spec--backend_defaults--client_policy--tls--validation--subject_alternative_names"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.validation.subject_alternative_names`

Required:

- `match` (Attributes) Match is a required field (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--validation--subject_alternative_names--match))

<a id="nestedatt--spec--backend_defaults--client_policy--tls--validation--subject_alternative_names--match"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.validation.subject_alternative_names.match`

Required:

- `exact` (List of String) Exact is a required field




<a id="nestedatt--spec--backend_defaults--client_policy--tls--certificate"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.certificate`

Optional:

- `file` (Attributes) An object that represents a TLS cert via a local file (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--certificate--file))
- `sds` (Attributes) An object that represents a TLS cert via SDS entry (see [below for nested schema](#nestedatt--spec--backend_defaults--client_policy--tls--certificate--sds))

<a id="nestedatt--spec--backend_defaults--client_policy--tls--certificate--file"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.certificate.file`

Required:

- `certificate_chain` (String) The certificate chain for the certificate.
- `private_key` (String) The private key for a certificate stored on the file system of the virtual Gateway.


<a id="nestedatt--spec--backend_defaults--client_policy--tls--certificate--sds"></a>
### Nested Schema for `spec.backend_defaults.client_policy.tls.certificate.sds`

Required:

- `secret_name` (String) The certificate trust chain for a certificate issued via SDS cluster






<a id="nestedatt--spec--gateway_route_selector"></a>
### Nested Schema for `spec.gateway_route_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--gateway_route_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--gateway_route_selector--match_expressions"></a>
### Nested Schema for `spec.gateway_route_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--listeners"></a>
### Nested Schema for `spec.listeners`

Required:

- `port_mapping` (Attributes) The port mapping information for the listener. (see [below for nested schema](#nestedatt--spec--listeners--port_mapping))

Optional:

- `connection_pool` (Attributes) The connection pool settings for the listener (see [below for nested schema](#nestedatt--spec--listeners--connection_pool))
- `health_check` (Attributes) The health check information for the listener. (see [below for nested schema](#nestedatt--spec--listeners--health_check))
- `tls` (Attributes) A reference to an object that represents the Transport Layer Security (TLS) properties for a listener. (see [below for nested schema](#nestedatt--spec--listeners--tls))

<a id="nestedatt--spec--listeners--port_mapping"></a>
### Nested Schema for `spec.listeners.port_mapping`

Required:

- `port` (Number) The port used for the port mapping.
- `protocol` (String) The protocol used for the port mapping.


<a id="nestedatt--spec--listeners--connection_pool"></a>
### Nested Schema for `spec.listeners.connection_pool`

Optional:

- `grpc` (Attributes) Specifies grpc connection pool settings for the virtual gateway listener (see [below for nested schema](#nestedatt--spec--listeners--connection_pool--grpc))
- `http` (Attributes) Specifies http connection pool settings for the virtual gateway listener (see [below for nested schema](#nestedatt--spec--listeners--connection_pool--http))
- `http2` (Attributes) Specifies http2 connection pool settings for the virtual gateway listener (see [below for nested schema](#nestedatt--spec--listeners--connection_pool--http2))

<a id="nestedatt--spec--listeners--connection_pool--grpc"></a>
### Nested Schema for `spec.listeners.connection_pool.grpc`

Required:

- `max_requests` (Number) Represents the maximum number of inflight requests that an envoy can concurrently support across all the hosts in the upstream cluster


<a id="nestedatt--spec--listeners--connection_pool--http"></a>
### Nested Schema for `spec.listeners.connection_pool.http`

Required:

- `max_connections` (Number) Represents the maximum number of outbound TCP connections the envoy can establish concurrently with all the hosts in the upstream cluster.

Optional:

- `max_pending_requests` (Number) Represents the number of overflowing requests after max_connections that an envoy will queue to an upstream cluster.


<a id="nestedatt--spec--listeners--connection_pool--http2"></a>
### Nested Schema for `spec.listeners.connection_pool.http2`

Required:

- `max_requests` (Number) Represents the maximum number of inflight requests that an envoy can concurrently support across all the hosts in the upstream cluster



<a id="nestedatt--spec--listeners--health_check"></a>
### Nested Schema for `spec.listeners.health_check`

Required:

- `interval_millis` (Number) The time period in milliseconds between each health check execution.
- `protocol` (String) The protocol for the health check request
- `timeout_millis` (Number) The amount of time to wait when receiving a response from the health check, in milliseconds.
- `unhealthy_threshold` (Number) The number of consecutive failed health checks that must occur before declaring a virtual Gateway unhealthy.

Optional:

- `healthy_threshold` (Number) The number of consecutive successful health checks that must occur before declaring listener healthy.
- `path` (String) The destination path for the health check request. This value is only used if the specified protocol is http or http2. For any other protocol, this value is ignored.
- `port` (Number) The destination port for the health check request.


<a id="nestedatt--spec--listeners--tls"></a>
### Nested Schema for `spec.listeners.tls`

Required:

- `certificate` (Attributes) A reference to an object that represents a listener's TLS certificate. (see [below for nested schema](#nestedatt--spec--listeners--tls--certificate))
- `mode` (String) ListenerTLS mode

Optional:

- `validation` (Attributes) A reference to an object that represents Validation context (see [below for nested schema](#nestedatt--spec--listeners--tls--validation))

<a id="nestedatt--spec--listeners--tls--certificate"></a>
### Nested Schema for `spec.listeners.tls.certificate`

Optional:

- `acm` (Attributes) A reference to an object that represents an AWS Certificate Manager (ACM) certificate. (see [below for nested schema](#nestedatt--spec--listeners--tls--certificate--acm))
- `file` (Attributes) A reference to an object that represents a local file certificate. (see [below for nested schema](#nestedatt--spec--listeners--tls--certificate--file))
- `sds` (Attributes) A reference to an object that represents an SDS issued certificate (see [below for nested schema](#nestedatt--spec--listeners--tls--certificate--sds))

<a id="nestedatt--spec--listeners--tls--certificate--acm"></a>
### Nested Schema for `spec.listeners.tls.certificate.acm`

Required:

- `certificate_arn` (String) The Amazon Resource Name (ARN) for the certificate.


<a id="nestedatt--spec--listeners--tls--certificate--file"></a>
### Nested Schema for `spec.listeners.tls.certificate.file`

Required:

- `certificate_chain` (String) The certificate chain for the certificate.
- `private_key` (String) The private key for a certificate stored on the file system of the virtual Gateway.


<a id="nestedatt--spec--listeners--tls--certificate--sds"></a>
### Nested Schema for `spec.listeners.tls.certificate.sds`

Required:

- `secret_name` (String) The certificate trust chain for a certificate issued via SDS cluster



<a id="nestedatt--spec--listeners--tls--validation"></a>
### Nested Schema for `spec.listeners.tls.validation`

Required:

- `trust` (Attributes) (see [below for nested schema](#nestedatt--spec--listeners--tls--validation--trust))

Optional:

- `subject_alternative_names` (Attributes) Possible alternate names to consider (see [below for nested schema](#nestedatt--spec--listeners--tls--validation--subject_alternative_names))

<a id="nestedatt--spec--listeners--tls--validation--trust"></a>
### Nested Schema for `spec.listeners.tls.validation.trust`

Optional:

- `acm` (Attributes) A reference to an object that represents a TLS validation context trust for an AWS Certicate Manager (ACM) certificate. (see [below for nested schema](#nestedatt--spec--listeners--tls--validation--trust--acm))
- `file` (Attributes) An object that represents a TLS validation context trust for a local file. (see [below for nested schema](#nestedatt--spec--listeners--tls--validation--trust--file))
- `sds` (Attributes) An object that represents a TLS validation context trust for an SDS system (see [below for nested schema](#nestedatt--spec--listeners--tls--validation--trust--sds))

<a id="nestedatt--spec--listeners--tls--validation--trust--acm"></a>
### Nested Schema for `spec.listeners.tls.validation.trust.acm`

Required:

- `certificate_authority_ar_ns` (List of String) One or more ACM Amazon Resource Name (ARN)s.


<a id="nestedatt--spec--listeners--tls--validation--trust--file"></a>
### Nested Schema for `spec.listeners.tls.validation.trust.file`

Required:

- `certificate_chain` (String) The certificate trust chain for a certificate stored on the file system of the virtual Gateway.


<a id="nestedatt--spec--listeners--tls--validation--trust--sds"></a>
### Nested Schema for `spec.listeners.tls.validation.trust.sds`

Required:

- `secret_name` (String) The certificate trust chain for a certificate issued via SDS.



<a id="nestedatt--spec--listeners--tls--validation--subject_alternative_names"></a>
### Nested Schema for `spec.listeners.tls.validation.subject_alternative_names`

Required:

- `match` (Attributes) Match is a required field (see [below for nested schema](#nestedatt--spec--listeners--tls--validation--subject_alternative_names--match))

<a id="nestedatt--spec--listeners--tls--validation--subject_alternative_names--match"></a>
### Nested Schema for `spec.listeners.tls.validation.subject_alternative_names.match`

Required:

- `exact` (List of String) Exact is a required field






<a id="nestedatt--spec--logging"></a>
### Nested Schema for `spec.logging`

Optional:

- `access_log` (Attributes) The access log configuration for a virtual Gateway. (see [below for nested schema](#nestedatt--spec--logging--access_log))

<a id="nestedatt--spec--logging--access_log"></a>
### Nested Schema for `spec.logging.access_log`

Optional:

- `file` (Attributes) The file object to send virtual gateway access logs to. (see [below for nested schema](#nestedatt--spec--logging--access_log--file))

<a id="nestedatt--spec--logging--access_log--file"></a>
### Nested Schema for `spec.logging.access_log.file`

Required:

- `path` (String) The file path to write access logs to.

Optional:

- `format` (Attributes) Structured access log output format (see [below for nested schema](#nestedatt--spec--logging--access_log--file--format))

<a id="nestedatt--spec--logging--access_log--file--format"></a>
### Nested Schema for `spec.logging.access_log.file.format`

Optional:

- `json` (Attributes List) Output specified fields as a JSON object (see [below for nested schema](#nestedatt--spec--logging--access_log--file--format--json))
- `text` (String) Custom format string

<a id="nestedatt--spec--logging--access_log--file--format--json"></a>
### Nested Schema for `spec.logging.access_log.file.format.json`

Required:

- `key` (String) The name of the field in the JSON object
- `value` (String) The format string






<a id="nestedatt--spec--mesh_ref"></a>
### Nested Schema for `spec.mesh_ref`

Required:

- `name` (String) Name is the name of Mesh CR
- `uid` (String) UID is the UID of Mesh CR


<a id="nestedatt--spec--namespace_selector"></a>
### Nested Schema for `spec.namespace_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--namespace_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--namespace_selector--match_expressions"></a>
### Nested Schema for `spec.namespace_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--pod_selector"></a>
### Nested Schema for `spec.pod_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--pod_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--pod_selector--match_expressions"></a>
### Nested Schema for `spec.pod_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
