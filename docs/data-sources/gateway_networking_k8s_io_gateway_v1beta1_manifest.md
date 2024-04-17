---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_gateway_networking_k8s_io_gateway_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "gateway.networking.k8s.io"
description: |-
  Gateway represents an instance of a service-traffic handling infrastructureby binding Listeners to a set of IP addresses.
---

# k8s_gateway_networking_k8s_io_gateway_v1beta1_manifest (Data Source)

Gateway represents an instance of a service-traffic handling infrastructureby binding Listeners to a set of IP addresses.

## Example Usage

```terraform
data "k8s_gateway_networking_k8s_io_gateway_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    gateway_class_name = "some-class"
    listeners          = []
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) Spec defines the desired state of Gateway. (see [below for nested schema](#nestedatt--spec))

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

- `gateway_class_name` (String) GatewayClassName used for this Gateway. This is the name of aGatewayClass resource.
- `listeners` (Attributes List) Listeners associated with this Gateway. Listeners definelogical endpoints that are bound on this Gateway's addresses.At least one Listener MUST be specified.Each Listener in a set of Listeners (for example, in a single Gateway)MUST be _distinct_, in that a traffic flow MUST be able to be assigned toexactly one listener. (This section uses 'set of Listeners' rather than'Listeners in a single Gateway' because implementations MAY merge configurationfrom multiple Gateways onto a single data plane, and these rules _also_apply in that case).Practically, this means that each listener in a set MUST have a uniquecombination of Port, Protocol, and, if supported by the protocol, Hostname.Some combinations of port, protocol, and TLS settings are consideredCore support and MUST be supported by implementations based on theirtargeted conformance profile:HTTP Profile1. HTTPRoute, Port: 80, Protocol: HTTP2. HTTPRoute, Port: 443, Protocol: HTTPS, TLS Mode: Terminate, TLS keypair providedTLS Profile1. TLSRoute, Port: 443, Protocol: TLS, TLS Mode: Passthrough'Distinct' Listeners have the following property:The implementation can match inbound requests to a single distinctListener. When multiple Listeners share values for fields (forexample, two Listeners with the same Port value), the implementationcan match requests to only one of the Listeners using otherListener fields.For example, the following Listener scenarios are distinct:1. Multiple Listeners with the same Port that all use the 'HTTP'   Protocol that all have unique Hostname values.2. Multiple Listeners with the same Port that use either the 'HTTPS' or   'TLS' Protocol that all have unique Hostname values.3. A mixture of 'TCP' and 'UDP' Protocol Listeners, where no Listener   with the same Protocol has the same Port value.Some fields in the Listener struct have possible values that affectwhether the Listener is distinct. Hostname is particularly relevantfor HTTP or HTTPS protocols.When using the Hostname value to select between same-Port, same-ProtocolListeners, the Hostname value must be different on each Listener for theListener to be distinct.When the Listeners are distinct based on Hostname, inbound requesthostnames MUST match from the most specific to least specific Hostnamevalues to choose the correct Listener and its associated set of Routes.Exact matches must be processed before wildcard matches, and wildcardmatches must be processed before fallback (empty Hostname value)matches. For example, ''foo.example.com'' takes precedence over''*.example.com'', and ''*.example.com'' takes precedence over ''''.Additionally, if there are multiple wildcard entries, more specificwildcard entries must be processed before less specific wildcard entries.For example, ''*.foo.example.com'' takes precedence over ''*.example.com''.The precise definition here is that the higher the number of dots in thehostname to the right of the wildcard character, the higher the precedence.The wildcard character will match any number of characters _and dots_ tothe left, however, so ''*.example.com'' will match both''foo.bar.example.com'' _and_ ''bar.example.com''.If a set of Listeners contains Listeners that are not distinct, then thoseListeners are Conflicted, and the implementation MUST set the 'Conflicted'condition in the Listener Status to 'True'.Implementations MAY choose to accept a Gateway with some ConflictedListeners only if they only accept the partial Listener set that containsno Conflicted Listeners. To put this another way, implementations mayaccept a partial Listener set only if they throw out *all* the conflictingListeners. No picking one of the conflicting listeners as the winner.This also means that the Gateway must have at least one non-conflictingListener in this case, otherwise it violates the requirement that atleast one Listener must be present.The implementation MUST set a 'ListenersNotValid' condition on theGateway Status when the Gateway contains Conflicted Listeners whether ornot they accept the Gateway. That Condition SHOULD clearlyindicate in the Message which Listeners are conflicted, and which areAccepted. Additionally, the Listener status for those listeners SHOULDindicate which Listeners are conflicted and not Accepted.A Gateway's Listeners are considered 'compatible' if:1. They are distinct.2. The implementation can serve them in compliance with the Addresses   requirement that all Listeners are available on all assigned   addresses.Compatible combinations in Extended support are expected to vary acrossimplementations. A combination that is compatible for one implementationmay not be compatible for another.For example, an implementation that cannot serve both TCP and UDP listenerson the same address, or cannot mix HTTPS and generic TLS listens on the same portwould not consider those cases compatible, even though they are distinct.Note that requests SHOULD match at most one Listener. For example, ifListeners are defined for 'foo.example.com' and '*.example.com', arequest to 'foo.example.com' SHOULD only be routed using routes attachedto the 'foo.example.com' Listener (and not the '*.example.com' Listener).This concept is known as 'Listener Isolation'. Implementations that donot support Listener Isolation MUST clearly document this.Implementations MAY merge separate Gateways onto a single set ofAddresses if all Listeners across all Gateways are compatible.Support: Core (see [below for nested schema](#nestedatt--spec--listeners))

Optional:

- `addresses` (Attributes List) Addresses requested for this Gateway. This is optional and behavior candepend on the implementation. If a value is set in the spec and therequested address is invalid or unavailable, the implementation MUSTindicate this in the associated entry in GatewayStatus.Addresses.The Addresses field represents a request for the address(es) on the'outside of the Gateway', that traffic bound for this Gateway will use.This could be the IP address or hostname of an external load balancer orother networking infrastructure, or some other address that traffic willbe sent to.If no Addresses are specified, the implementation MAY schedule theGateway in an implementation-specific manner, assigning an appropriateset of Addresses.The implementation MUST bind all Listeners to every GatewayAddress thatit assigns to the Gateway and add a corresponding entry inGatewayStatus.Addresses.Support: Extended (see [below for nested schema](#nestedatt--spec--addresses))
- `infrastructure` (Attributes) Infrastructure defines infrastructure level attributes about this Gateway instance.Support: Core (see [below for nested schema](#nestedatt--spec--infrastructure))

<a id="nestedatt--spec--listeners"></a>
### Nested Schema for `spec.listeners`

Required:

- `name` (String) Name is the name of the Listener. This name MUST be unique within aGateway.Support: Core
- `port` (Number) Port is the network port. Multiple listeners may use thesame port, subject to the Listener compatibility rules.Support: Core
- `protocol` (String) Protocol specifies the network protocol this listener expects to receive.Support: Core

Optional:

- `allowed_routes` (Attributes) AllowedRoutes defines the types of routes that MAY be attached to aListener and the trusted namespaces where those Route resources MAY bepresent.Although a client request may match multiple route rules, only one rulemay ultimately receive the request. Matching precedence MUST bedetermined in order of the following criteria:* The most specific match as defined by the Route type.* The oldest Route based on creation timestamp. For example, a Route with  a creation timestamp of '2020-09-08 01:02:03' is given precedence over  a Route with a creation timestamp of '2020-09-08 01:02:04'.* If everything else is equivalent, the Route appearing first in  alphabetical order (namespace/name) should be given precedence. For  example, foo/bar is given precedence over foo/baz.All valid rules within a Route attached to this Listener should beimplemented. Invalid Route rules can be ignored (sometimes that will meanthe full Route). If a Route rule transitions from valid to invalid,support for that Route rule should be dropped to ensure consistency. Forexample, even if a filter specified by a Route rule is invalid, the restof the rules within that Route should still be supported.Support: Core (see [below for nested schema](#nestedatt--spec--listeners--allowed_routes))
- `hostname` (String) Hostname specifies the virtual hostname to match for protocol types thatdefine this concept. When unspecified, all hostnames are matched. Thisfield is ignored for protocols that don't require hostname basedmatching.Implementations MUST apply Hostname matching appropriately for each ofthe following protocols:* TLS: The Listener Hostname MUST match the SNI.* HTTP: The Listener Hostname MUST match the Host header of the request.* HTTPS: The Listener Hostname SHOULD match at both the TLS and HTTP  protocol layers as described above. If an implementation does not  ensure that both the SNI and Host header match the Listener hostname,  it MUST clearly document that.For HTTPRoute and TLSRoute resources, there is an interaction with the'spec.hostnames' array. When both listener and route specify hostnames,there MUST be an intersection between the values for a Route to beaccepted. For more information, refer to the Route specific Hostnamesdocumentation.Hostnames that are prefixed with a wildcard label ('*.') are interpretedas a suffix match. That means that a match for '*.example.com' would matchboth 'test.example.com', and 'foo.test.example.com', but not 'example.com'.Support: Core
- `tls` (Attributes) TLS is the TLS configuration for the Listener. This field is required ifthe Protocol field is 'HTTPS' or 'TLS'. It is invalid to set this fieldif the Protocol field is 'HTTP', 'TCP', or 'UDP'.The association of SNIs to Certificate defined in GatewayTLSConfig isdefined based on the Hostname field for this listener.The GatewayClass MUST use the longest matching SNI out of allavailable certificates for any TLS handshake.Support: Core (see [below for nested schema](#nestedatt--spec--listeners--tls))

<a id="nestedatt--spec--listeners--allowed_routes"></a>
### Nested Schema for `spec.listeners.allowed_routes`

Optional:

- `kinds` (Attributes List) Kinds specifies the groups and kinds of Routes that are allowed to bindto this Gateway Listener. When unspecified or empty, the kinds of Routesselected are determined using the Listener protocol.A RouteGroupKind MUST correspond to kinds of Routes that are compatiblewith the application protocol specified in the Listener's Protocol field.If an implementation does not support or recognize this resource type, itMUST set the 'ResolvedRefs' condition to False for this Listener with the'InvalidRouteKinds' reason.Support: Core (see [below for nested schema](#nestedatt--spec--listeners--allowed_routes--kinds))
- `namespaces` (Attributes) Namespaces indicates namespaces from which Routes may be attached to thisListener. This is restricted to the namespace of this Gateway by default.Support: Core (see [below for nested schema](#nestedatt--spec--listeners--allowed_routes--namespaces))

<a id="nestedatt--spec--listeners--allowed_routes--kinds"></a>
### Nested Schema for `spec.listeners.allowed_routes.kinds`

Required:

- `kind` (String) Kind is the kind of the Route.

Optional:

- `group` (String) Group is the group of the Route.


<a id="nestedatt--spec--listeners--allowed_routes--namespaces"></a>
### Nested Schema for `spec.listeners.allowed_routes.namespaces`

Optional:

- `from` (String) From indicates where Routes will be selected for this Gateway. Possiblevalues are:* All: Routes in all namespaces may be used by this Gateway.* Selector: Routes in namespaces selected by the selector may be used by  this Gateway.* Same: Only Routes in the same namespace may be used by this Gateway.Support: Core
- `selector` (Attributes) Selector must be specified when From is set to 'Selector'. In that case,only Routes in Namespaces matching this Selector will be selected by thisGateway. This field is ignored for other values of 'From'.Support: Core (see [below for nested schema](#nestedatt--spec--listeners--allowed_routes--namespaces--selector))

<a id="nestedatt--spec--listeners--allowed_routes--namespaces--selector"></a>
### Nested Schema for `spec.listeners.allowed_routes.namespaces.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--listeners--allowed_routes--namespaces--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--listeners--allowed_routes--namespaces--selector--match_expressions"></a>
### Nested Schema for `spec.listeners.allowed_routes.namespaces.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.





<a id="nestedatt--spec--listeners--tls"></a>
### Nested Schema for `spec.listeners.tls`

Optional:

- `certificate_refs` (Attributes List) CertificateRefs contains a series of references to Kubernetes objects thatcontains TLS certificates and private keys. These certificates are used toestablish a TLS handshake for requests that match the hostname of theassociated listener.A single CertificateRef to a Kubernetes Secret has 'Core' support.Implementations MAY choose to support attaching multiple certificates toa Listener, but this behavior is implementation-specific.References to a resource in different namespace are invalid UNLESS thereis a ReferenceGrant in the target namespace that allows the certificateto be attached. If a ReferenceGrant does not allow this reference, the'ResolvedRefs' condition MUST be set to False for this listener with the'RefNotPermitted' reason.This field is required to have at least one element when the mode is setto 'Terminate' (default) and is optional otherwise.CertificateRefs can reference to standard Kubernetes resources, i.e.Secret, or implementation-specific custom resources.Support: Core - A single reference to a Kubernetes Secret of type kubernetes.io/tlsSupport: Implementation-specific (More than one reference or other resource types) (see [below for nested schema](#nestedatt--spec--listeners--tls--certificate_refs))
- `mode` (String) Mode defines the TLS behavior for the TLS session initiated by the client.There are two possible modes:- Terminate: The TLS session between the downstream client and the  Gateway is terminated at the Gateway. This mode requires certificates  to be specified in some way, such as populating the certificateRefs  field.- Passthrough: The TLS session is NOT terminated by the Gateway. This  implies that the Gateway can't decipher the TLS stream except for  the ClientHello message of the TLS protocol. The certificateRefs field  is ignored in this mode.Support: Core
- `options` (Map of String) Options are a list of key/value pairs to enable extended TLSconfiguration for each implementation. For example, configuring theminimum TLS version or supported cipher suites.A set of common keys MAY be defined by the API in the future. To avoidany ambiguity, implementation-specific definitions MUST usedomain-prefixed names, such as 'example.com/my-custom-option'.Un-prefixed names are reserved for key names defined by Gateway API.Support: Implementation-specific

<a id="nestedatt--spec--listeners--tls--certificate_refs"></a>
### Nested Schema for `spec.listeners.tls.certificate_refs`

Required:

- `name` (String) Name is the name of the referent.

Optional:

- `group` (String) Group is the group of the referent. For example, 'gateway.networking.k8s.io'.When unspecified or empty string, core API group is inferred.
- `kind` (String) Kind is kind of the referent. For example 'Secret'.
- `namespace` (String) Namespace is the namespace of the referenced object. When unspecified, the localnamespace is inferred.Note that when a namespace different than the local namespace is specified,a ReferenceGrant object is required in the referent namespace to allow thatnamespace's owner to accept the reference. See the ReferenceGrantdocumentation for details.Support: Core




<a id="nestedatt--spec--addresses"></a>
### Nested Schema for `spec.addresses`

Required:

- `value` (String) Value of the address. The validity of the values will dependon the type and support by the controller.Examples: '1.2.3.4', '128::1', 'my-ip-address'.

Optional:

- `type` (String) Type of the address.


<a id="nestedatt--spec--infrastructure"></a>
### Nested Schema for `spec.infrastructure`

Optional:

- `annotations` (Map of String) Annotations that SHOULD be applied to any resources created in response to this Gateway.For implementations creating other Kubernetes objects, this should be the 'metadata.annotations' field on resources.For other implementations, this refers to any relevant (implementation specific) 'annotations' concepts.An implementation may chose to add additional implementation-specific annotations as they see fit.Support: Extended
- `labels` (Map of String) Labels that SHOULD be applied to any resources created in response to this Gateway.For implementations creating other Kubernetes objects, this should be the 'metadata.labels' field on resources.For other implementations, this refers to any relevant (implementation specific) 'labels' concepts.An implementation may chose to add additional implementation-specific labels as they see fit.Support: Extended