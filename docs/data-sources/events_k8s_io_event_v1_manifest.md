---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_events_k8s_io_event_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "events.k8s.io"
description: |-
  Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
---

# k8s_events_k8s_io_event_v1_manifest (Data Source)

Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.

## Example Usage

```terraform
data "k8s_events_k8s_io_event_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  type   = "Warning"
  reason = "Some not so urgent event has occurred"
  related = {
    kind = "some-kind"
  }
  event_time = "2022-10-16"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `event_time` (String) MicroTime is version of Time with microsecond level precision.
- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `action` (String) action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
- `deprecated_count` (Number) deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
- `deprecated_first_timestamp` (String) Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
- `deprecated_last_timestamp` (String) Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
- `deprecated_source` (Attributes) EventSource contains information for an event. (see [below for nested schema](#nestedatt--deprecated_source))
- `note` (String) note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
- `reason` (String) reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
- `regarding` (Attributes) ObjectReference contains enough information to let you inspect or modify the referred object. (see [below for nested schema](#nestedatt--regarding))
- `related` (Attributes) ObjectReference contains enough information to let you inspect or modify the referred object. (see [below for nested schema](#nestedatt--related))
- `reporting_controller` (String) reportingController is the name of the controller that emitted this Event, e.g. 'kubernetes.io/kubelet'. This field cannot be empty for new Events.
- `reporting_instance` (String) reportingInstance is the ID of the controller instance, e.g. 'kubelet-xyzf'. This field cannot be empty for new Events and it can have at most 128 characters.
- `series` (Attributes) EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time. How often to update the EventSeries is up to the event reporters. The default event reporter in 'k8s.io/client-go/tools/events/event_broadcaster.go' shows how this struct is updated on heartbeats and can guide customized reporter implementations. (see [below for nested schema](#nestedatt--series))
- `type` (String) type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.

### Read-Only

- `id` (String) Contains the value `metadata.namespace/metadata.name`.
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--deprecated_source"></a>
### Nested Schema for `deprecated_source`

Optional:

- `component` (String) Component from which the event is generated.
- `host` (String) Node name on which the event is generated.


<a id="nestedatt--regarding"></a>
### Nested Schema for `regarding`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--related"></a>
### Nested Schema for `related`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--series"></a>
### Nested Schema for `series`

Required:

- `count` (Number) count is the number of occurrences in this series up to the last heartbeat time.
- `last_observed_time` (String) MicroTime is version of Time with microsecond level precision.