---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2 Resource - terraform-provider-k8s"
subcategory: "flowcontrol.apiserver.k8s.io"
description: |-
  PriorityLevelConfiguration represents the configuration of a priority level.
---

# k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2 (Resource)

PriorityLevelConfiguration represents the configuration of a priority level.

## Example Usage

```terraform
resource "k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    type = "Limited"
    limited = {
      assured_concurrency_shares = 125
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) PriorityLevelConfigurationSpec specifies the configuration of a priority level. (see [below for nested schema](#nestedatt--spec))

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

Required:

- `type` (String) 'type' indicates whether this priority level is subject to limitation on request execution.  A value of ''Exempt'' means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of ''Limited'' means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.

Optional:

- `limited` (Attributes) LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits. It addresses two issues:  - How are requests for this priority level limited?  - What should be done with requests that exceed the limit? (see [below for nested schema](#nestedatt--spec--limited))

<a id="nestedatt--spec--limited"></a>
### Nested Schema for `spec.limited`

Optional:

- `assured_concurrency_shares` (Number) 'assuredConcurrencyShares' (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.  ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:            ACV(l) = ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30.
- `limit_response` (Attributes) LimitResponse defines how to handle requests that can not be executed right now. (see [below for nested schema](#nestedatt--spec--limited--limit_response))

<a id="nestedatt--spec--limited--limit_response"></a>
### Nested Schema for `spec.limited.limit_response`

Required:

- `type` (String) 'type' is 'Queue' or 'Reject'. 'Queue' means that requests that can not be executed upon arrival are held in a queue until they can be executed or a queuing limit is reached. 'Reject' means that requests that can not be executed upon arrival are rejected. Required.

Optional:

- `queuing` (Attributes) QueuingConfiguration holds the configuration parameters for queuing (see [below for nested schema](#nestedatt--spec--limited--limit_response--queuing))

<a id="nestedatt--spec--limited--limit_response--queuing"></a>
### Nested Schema for `spec.limited.limit_response.queuing`

Optional:

- `hand_size` (Number) 'handSize' is a small positive number that configures the shuffle sharding of requests into queues.  When enqueuing a request at this priority level the request's flow identifier (a string pair) is hashed and the hash value is used to shuffle the list of queues and deal a hand of the size specified here.  The request is put into one of the shortest queues in that hand. 'handSize' must be no larger than 'queues', and should be significantly smaller (so that a few heavy flows do not saturate most of the queues).  See the user-facing documentation for more extensive guidance on setting this field.  This field has a default value of 8.
- `queue_length_limit` (Number) 'queueLengthLimit' is the maximum number of requests allowed to be waiting in a given queue of this priority level at a time; excess requests are rejected.  This value must be positive.  If not specified, it will be defaulted to 50.
- `queues` (Number) 'queues' is the number of queues for this priority level. The queues exist independently at each apiserver. The value must be positive.  Setting it to 1 effectively precludes shufflesharding and thus makes the distinguisher method of associated flow schemas irrelevant.  This field has a default value of 64.


