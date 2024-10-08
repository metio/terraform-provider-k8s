---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_networking_istio_io_gateway_v1alpha3_manifest Data Source - terraform-provider-k8s"
subcategory: "networking.istio.io"
description: |-
  
---

# k8s_networking_istio_io_gateway_v1alpha3_manifest (Data Source)



## Example Usage

```terraform
data "k8s_networking_istio_io_gateway_v1alpha3_manifest" "example" {
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

- `spec` (Attributes) Configuration affecting edge load balancer. See more details at: https://istio.io/docs/reference/config/networking/gateway.html (see [below for nested schema](#nestedatt--spec))

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

- `selector` (Map of String) One or more labels that indicate a specific set of pods/VMs on which this gateway configuration should be applied.
- `servers` (Attributes List) A list of server specifications. (see [below for nested schema](#nestedatt--spec--servers))

<a id="nestedatt--spec--servers"></a>
### Nested Schema for `spec.servers`

Required:

- `hosts` (List of String) One or more hosts exposed by this gateway.
- `port` (Attributes) The Port on which the proxy should listen for incoming connections. (see [below for nested schema](#nestedatt--spec--servers--port))

Optional:

- `bind` (String) The ip or the Unix domain socket to which the listener should be bound to.
- `default_endpoint` (String)
- `name` (String) An optional name of the server, when set must be unique across all servers.
- `tls` (Attributes) Set of TLS related options that govern the server's behavior. (see [below for nested schema](#nestedatt--spec--servers--tls))

<a id="nestedatt--spec--servers--port"></a>
### Nested Schema for `spec.servers.port`

Required:

- `name` (String) Label assigned to the port.
- `number` (Number) A valid non-negative integer port number.
- `protocol` (String) The protocol exposed on the port.

Optional:

- `target_port` (Number)


<a id="nestedatt--spec--servers--tls"></a>
### Nested Schema for `spec.servers.tls`

Optional:

- `ca_certificates` (String) REQUIRED if mode is 'MUTUAL' or 'OPTIONAL_MUTUAL'.
- `ca_crl` (String) OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented client side certificate.
- `cipher_suites` (List of String) Optional: If specified, only support the specified cipher list.
- `credential_name` (String) For gateways running on Kubernetes, the name of the secret that holds the TLS certs including the CA certificates.
- `https_redirect` (Boolean) If set to true, the load balancer will send a 301 redirect for all http connections, asking the clients to use HTTPS.
- `max_protocol_version` (String) Optional: Maximum TLS protocol version. Valid Options: TLS_AUTO, TLSV1_0, TLSV1_1, TLSV1_2, TLSV1_3
- `min_protocol_version` (String) Optional: Minimum TLS protocol version. Valid Options: TLS_AUTO, TLSV1_0, TLSV1_1, TLSV1_2, TLSV1_3
- `mode` (String) Optional: Indicates whether connections to this port should be secured using TLS. Valid Options: PASSTHROUGH, SIMPLE, MUTUAL, AUTO_PASSTHROUGH, ISTIO_MUTUAL, OPTIONAL_MUTUAL
- `private_key` (String) REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.
- `server_certificate` (String) REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.
- `subject_alt_names` (List of String) A list of alternate names to verify the subject identity in the certificate presented by the client.
- `verify_certificate_hash` (List of String) An optional list of hex-encoded SHA-256 hashes of the authorized client certificates.
- `verify_certificate_spki` (List of String) An optional list of base64-encoded SHA-256 hashes of the SPKIs of authorized client certificates.
