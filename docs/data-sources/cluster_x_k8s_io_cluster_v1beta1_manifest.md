---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cluster_x_k8s_io_cluster_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "cluster.x-k8s.io"
description: |-
  Cluster is the Schema for the clusters API.
---

# k8s_cluster_x_k8s_io_cluster_v1beta1_manifest (Data Source)

Cluster is the Schema for the clusters API.

## Example Usage

```terraform
data "k8s_cluster_x_k8s_io_cluster_v1beta1_manifest" "example" {
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

- `spec` (Attributes) ClusterSpec defines the desired state of Cluster. (see [below for nested schema](#nestedatt--spec))

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

- `availability_gates` (Attributes List) availabilityGates specifies additional conditions to include when evaluating Cluster Available condition. NOTE: this field is considered only for computing v1beta2 conditions. (see [below for nested schema](#nestedatt--spec--availability_gates))
- `cluster_network` (Attributes) Cluster network configuration. (see [below for nested schema](#nestedatt--spec--cluster_network))
- `control_plane_endpoint` (Attributes) ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. (see [below for nested schema](#nestedatt--spec--control_plane_endpoint))
- `control_plane_ref` (Attributes) ControlPlaneRef is an optional reference to a provider-specific resource that holds the details for provisioning the Control Plane for a Cluster. (see [below for nested schema](#nestedatt--spec--control_plane_ref))
- `infrastructure_ref` (Attributes) InfrastructureRef is a reference to a provider-specific resource that holds the details for provisioning infrastructure for a cluster in said provider. (see [below for nested schema](#nestedatt--spec--infrastructure_ref))
- `paused` (Boolean) Paused can be used to prevent controllers from processing the Cluster and all its associated objects.
- `topology` (Attributes) This encapsulates the topology for the cluster. NOTE: It is required to enable the ClusterTopology feature gate flag to activate managed topologies support; this feature is highly experimental, and parts of it might still be not implemented. (see [below for nested schema](#nestedatt--spec--topology))

<a id="nestedatt--spec--availability_gates"></a>
### Nested Schema for `spec.availability_gates`

Required:

- `condition_type` (String) conditionType refers to a positive polarity condition (status true means good) with matching type in the Cluster's condition list. If the conditions doesn't exist, it will be treated as unknown. Note: Both Cluster API conditions or conditions added by 3rd party controllers can be used as availability gates.


<a id="nestedatt--spec--cluster_network"></a>
### Nested Schema for `spec.cluster_network`

Optional:

- `api_server_port` (Number) APIServerPort specifies the port the API Server should bind to. Defaults to 6443.
- `pods` (Attributes) The network ranges from which Pod networks are allocated. (see [below for nested schema](#nestedatt--spec--cluster_network--pods))
- `service_domain` (String) Domain name for services.
- `services` (Attributes) The network ranges from which service VIPs are allocated. (see [below for nested schema](#nestedatt--spec--cluster_network--services))

<a id="nestedatt--spec--cluster_network--pods"></a>
### Nested Schema for `spec.cluster_network.pods`

Required:

- `cidr_blocks` (List of String)


<a id="nestedatt--spec--cluster_network--services"></a>
### Nested Schema for `spec.cluster_network.services`

Required:

- `cidr_blocks` (List of String)



<a id="nestedatt--spec--control_plane_endpoint"></a>
### Nested Schema for `spec.control_plane_endpoint`

Required:

- `host` (String) The hostname on which the API server is serving.
- `port` (Number) The port on which the API server is serving.


<a id="nestedatt--spec--control_plane_ref"></a>
### Nested Schema for `spec.control_plane_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--spec--infrastructure_ref"></a>
### Nested Schema for `spec.infrastructure_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--spec--topology"></a>
### Nested Schema for `spec.topology`

Required:

- `class` (String) The name of the ClusterClass object to create the topology.
- `version` (String) The Kubernetes version of the cluster.

Optional:

- `control_plane` (Attributes) ControlPlane describes the cluster control plane. (see [below for nested schema](#nestedatt--spec--topology--control_plane))
- `rollout_after` (String) RolloutAfter performs a rollout of the entire cluster one component at a time, control plane first and then machine deployments. Deprecated: This field has no function and is going to be removed in the next apiVersion.
- `variables` (Attributes List) Variables can be used to customize the Cluster through patches. They must comply to the corresponding VariableClasses defined in the ClusterClass. (see [below for nested schema](#nestedatt--spec--topology--variables))
- `workers` (Attributes) Workers encapsulates the different constructs that form the worker nodes for the cluster. (see [below for nested schema](#nestedatt--spec--topology--workers))

<a id="nestedatt--spec--topology--control_plane"></a>
### Nested Schema for `spec.topology.control_plane`

Optional:

- `machine_health_check` (Attributes) MachineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this control plane. (see [below for nested schema](#nestedatt--spec--topology--control_plane--machine_health_check))
- `metadata` (Attributes) Metadata is the metadata applied to the ControlPlane and the Machines of the ControlPlane if the ControlPlaneTemplate referenced by the ClusterClass is machine based. If not, it is applied only to the ControlPlane. At runtime this metadata is merged with the corresponding metadata from the ClusterClass. (see [below for nested schema](#nestedatt--spec--topology--control_plane--metadata))
- `node_deletion_timeout` (String) NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.
- `node_drain_timeout` (String) NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'
- `node_volume_detach_timeout` (String) NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.
- `replicas` (Number) Replicas is the number of control plane nodes. If the value is nil, the ControlPlane object is created without the number of Replicas and it's assumed that the control plane controller does not implement support for this field. When specified against a control plane provider that lacks support for this field, this value will be ignored.
- `variables` (Attributes) Variables can be used to customize the ControlPlane through patches. (see [below for nested schema](#nestedatt--spec--topology--control_plane--variables))

<a id="nestedatt--spec--topology--control_plane--machine_health_check"></a>
### Nested Schema for `spec.topology.control_plane.machine_health_check`

Optional:

- `enable` (Boolean) Enable controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.
- `max_unhealthy` (String) Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by 'selector' are not healthy.
- `node_startup_timeout` (String) NodeStartupTimeout allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.
- `remediation_template` (Attributes) RemediationTemplate is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API. (see [below for nested schema](#nestedatt--spec--topology--control_plane--machine_health_check--remediation_template))
- `unhealthy_conditions` (Attributes List) UnhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy. (see [below for nested schema](#nestedatt--spec--topology--control_plane--machine_health_check--unhealthy_conditions))
- `unhealthy_range` (String) Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines

<a id="nestedatt--spec--topology--control_plane--machine_health_check--remediation_template"></a>
### Nested Schema for `spec.topology.control_plane.machine_health_check.remediation_template`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--spec--topology--control_plane--machine_health_check--unhealthy_conditions"></a>
### Nested Schema for `spec.topology.control_plane.machine_health_check.unhealthy_conditions`

Required:

- `status` (String)
- `timeout` (String)
- `type` (String)



<a id="nestedatt--spec--topology--control_plane--metadata"></a>
### Nested Schema for `spec.topology.control_plane.metadata`

Optional:

- `annotations` (Map of String) Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
- `labels` (Map of String) Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels


<a id="nestedatt--spec--topology--control_plane--variables"></a>
### Nested Schema for `spec.topology.control_plane.variables`

Optional:

- `overrides` (Attributes List) Overrides can be used to override Cluster level variables. (see [below for nested schema](#nestedatt--spec--topology--control_plane--variables--overrides))

<a id="nestedatt--spec--topology--control_plane--variables--overrides"></a>
### Nested Schema for `spec.topology.control_plane.variables.overrides`

Required:

- `name` (String) Name of the variable.
- `value` (Map of String) Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111

Optional:

- `definition_from` (String) DefinitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.




<a id="nestedatt--spec--topology--variables"></a>
### Nested Schema for `spec.topology.variables`

Required:

- `name` (String) Name of the variable.
- `value` (Map of String) Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111

Optional:

- `definition_from` (String) DefinitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.


<a id="nestedatt--spec--topology--workers"></a>
### Nested Schema for `spec.topology.workers`

Optional:

- `machine_deployments` (Attributes List) MachineDeployments is a list of machine deployments in the cluster. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments))
- `machine_pools` (Attributes List) MachinePools is a list of machine pools in the cluster. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_pools))

<a id="nestedatt--spec--topology--workers--machine_deployments"></a>
### Nested Schema for `spec.topology.workers.machine_deployments`

Required:

- `class` (String) Class is the name of the MachineDeploymentClass used to create the set of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.
- `name` (String) Name is the unique identifier for this MachineDeploymentTopology. The value is used with other unique identifiers to create a MachineDeployment's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.

Optional:

- `failure_domain` (String) FailureDomain is the failure domain the machines will be created in. Must match a key in the FailureDomains map stored on the cluster object.
- `machine_health_check` (Attributes) MachineHealthCheck allows to enable, disable and override the MachineHealthCheck configuration in the ClusterClass for this MachineDeployment. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--machine_health_check))
- `metadata` (Attributes) Metadata is the metadata applied to the MachineDeployment and the machines of the MachineDeployment. At runtime this metadata is merged with the corresponding metadata from the ClusterClass. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--metadata))
- `min_ready_seconds` (Number) Minimum number of seconds for which a newly created machine should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)
- `node_deletion_timeout` (String) NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machine hosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.
- `node_drain_timeout` (String) NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'
- `node_volume_detach_timeout` (String) NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.
- `replicas` (Number) Replicas is the number of worker nodes belonging to this set. If the value is nil, the MachineDeployment is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.
- `strategy` (Attributes) The deployment strategy to use to replace existing machines with new ones. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--strategy))
- `variables` (Attributes) Variables can be used to customize the MachineDeployment through patches. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--variables))

<a id="nestedatt--spec--topology--workers--machine_deployments--machine_health_check"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.machine_health_check`

Optional:

- `enable` (Boolean) Enable controls if a MachineHealthCheck should be created for the target machines. If false: No MachineHealthCheck will be created. If not set(default): A MachineHealthCheck will be created if it is defined here or in the associated ClusterClass. If no MachineHealthCheck is defined then none will be created. If true: A MachineHealthCheck is guaranteed to be created. Cluster validation will block if 'enable' is true and no MachineHealthCheck definition is available.
- `max_unhealthy` (String) Any further remediation is only allowed if at most 'MaxUnhealthy' machines selected by 'selector' are not healthy.
- `node_startup_timeout` (String) NodeStartupTimeout allows to set the maximum time for MachineHealthCheck to consider a Machine unhealthy if a corresponding Node isn't associated through a 'Spec.ProviderID' field. The duration set in this field is compared to the greatest of: - Cluster's infrastructure ready condition timestamp (if and when available) - Control Plane's initialized condition timestamp (if and when available) - Machine's infrastructure ready condition timestamp (if and when available) - Machine's metadata creation timestamp Defaults to 10 minutes. If you wish to disable this feature, set the value explicitly to 0.
- `remediation_template` (Attributes) RemediationTemplate is a reference to a remediation template provided by an infrastructure provider. This field is completely optional, when filled, the MachineHealthCheck controller creates a new object from the template referenced and hands off remediation of the machine to a controller that lives outside of Cluster API. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--machine_health_check--remediation_template))
- `unhealthy_conditions` (Attributes List) UnhealthyConditions contains a list of the conditions that determine whether a node is considered unhealthy. The conditions are combined in a logical OR, i.e. if any of the conditions is met, the node is unhealthy. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--machine_health_check--unhealthy_conditions))
- `unhealthy_range` (String) Any further remediation is only allowed if the number of machines selected by 'selector' as not healthy is within the range of 'UnhealthyRange'. Takes precedence over MaxUnhealthy. Eg. '[3-5]' - This means that remediation will be allowed only when: (a) there are at least 3 unhealthy machines (and) (b) there are at most 5 unhealthy machines

<a id="nestedatt--spec--topology--workers--machine_deployments--machine_health_check--remediation_template"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.machine_health_check.remediation_template`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--spec--topology--workers--machine_deployments--machine_health_check--unhealthy_conditions"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.machine_health_check.unhealthy_conditions`

Required:

- `status` (String)
- `timeout` (String)
- `type` (String)



<a id="nestedatt--spec--topology--workers--machine_deployments--metadata"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.metadata`

Optional:

- `annotations` (Map of String) Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
- `labels` (Map of String) Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels


<a id="nestedatt--spec--topology--workers--machine_deployments--strategy"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.strategy`

Optional:

- `remediation` (Attributes) Remediation controls the strategy of remediating unhealthy machines and how remediating operations should occur during the lifecycle of the dependant MachineSets. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--strategy--remediation))
- `rolling_update` (Attributes) Rolling update config params. Present only if MachineDeploymentStrategyType = RollingUpdate. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--strategy--rolling_update))
- `type` (String) Type of deployment. Allowed values are RollingUpdate and OnDelete. The default is RollingUpdate.

<a id="nestedatt--spec--topology--workers--machine_deployments--strategy--remediation"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.strategy.remediation`

Optional:

- `max_in_flight` (String) MaxInFlight determines how many in flight remediations should happen at the same time. Remediation only happens on the MachineSet with the most current revision, while older MachineSets (usually present during rollout operations) aren't allowed to remediate. Note: In general (independent of remediations), unhealthy machines are always prioritized during scale down operations over healthy ones. MaxInFlight can be set to a fixed number or a percentage. Example: when this is set to 20%, the MachineSet controller deletes at most 20% of the desired replicas. If not set, remediation is limited to all machines (bounded by replicas) under the active MachineSet's management.


<a id="nestedatt--spec--topology--workers--machine_deployments--strategy--rolling_update"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.strategy.rolling_update`

Optional:

- `delete_policy` (String) DeletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling. Valid values are 'Random, 'Newest', 'Oldest' When no value is supplied, the default DeletePolicy of MachineSet is used
- `max_surge` (String) The maximum number of machines that can be scheduled above the desired number of machines. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 1. Example: when this is set to 30%, the new MachineSet can be scaled up immediately when the rolling update starts, such that the total number of old and new machines do not exceed 130% of desired machines. Once old machines have been killed, new MachineSet can be scaled up further, ensuring that total number of machines running at any time during the update is at most 130% of desired machines.
- `max_unavailable` (String) The maximum number of machines that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired machines (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 0. Example: when this is set to 30%, the old MachineSet can be scaled down to 70% of desired machines immediately when the rolling update starts. Once new machines are ready, old MachineSet can be scaled down further, followed by scaling up the new MachineSet, ensuring that the total number of machines available at all times during the update is at least 70% of desired machines.



<a id="nestedatt--spec--topology--workers--machine_deployments--variables"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.variables`

Optional:

- `overrides` (Attributes List) Overrides can be used to override Cluster level variables. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_deployments--variables--overrides))

<a id="nestedatt--spec--topology--workers--machine_deployments--variables--overrides"></a>
### Nested Schema for `spec.topology.workers.machine_deployments.variables.overrides`

Required:

- `name` (String) Name of the variable.
- `value` (Map of String) Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111

Optional:

- `definition_from` (String) DefinitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.




<a id="nestedatt--spec--topology--workers--machine_pools"></a>
### Nested Schema for `spec.topology.workers.machine_pools`

Required:

- `class` (String) Class is the name of the MachinePoolClass used to create the pool of worker nodes. This should match one of the deployment classes defined in the ClusterClass object mentioned in the 'Cluster.Spec.Class' field.
- `name` (String) Name is the unique identifier for this MachinePoolTopology. The value is used with other unique identifiers to create a MachinePool's Name (e.g. cluster's name, etc). In case the name is greater than the allowed maximum length, the values are hashed together.

Optional:

- `failure_domains` (List of String) FailureDomains is the list of failure domains the machine pool will be created in. Must match a key in the FailureDomains map stored on the cluster object.
- `metadata` (Attributes) Metadata is the metadata applied to the MachinePool. At runtime this metadata is merged with the corresponding metadata from the ClusterClass. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_pools--metadata))
- `min_ready_seconds` (Number) Minimum number of seconds for which a newly created machine pool should be ready. Defaults to 0 (machine will be considered available as soon as it is ready)
- `node_deletion_timeout` (String) NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the MachinePool hosts after the MachinePool is marked for deletion. A duration of 0 will retry deletion indefinitely. Defaults to 10 seconds.
- `node_drain_timeout` (String) NodeDrainTimeout is the total amount of time that the controller will spend on draining a node. The default value is 0, meaning that the node can be drained without any time limitations. NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'
- `node_volume_detach_timeout` (String) NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumes to be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.
- `replicas` (Number) Replicas is the number of nodes belonging to this pool. If the value is nil, the MachinePool is created without the number of Replicas (defaulting to 1) and it's assumed that an external entity (like cluster autoscaler) is responsible for the management of this value.
- `variables` (Attributes) Variables can be used to customize the MachinePool through patches. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_pools--variables))

<a id="nestedatt--spec--topology--workers--machine_pools--metadata"></a>
### Nested Schema for `spec.topology.workers.machine_pools.metadata`

Optional:

- `annotations` (Map of String) Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
- `labels` (Map of String) Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels


<a id="nestedatt--spec--topology--workers--machine_pools--variables"></a>
### Nested Schema for `spec.topology.workers.machine_pools.variables`

Optional:

- `overrides` (Attributes List) Overrides can be used to override Cluster level variables. (see [below for nested schema](#nestedatt--spec--topology--workers--machine_pools--variables--overrides))

<a id="nestedatt--spec--topology--workers--machine_pools--variables--overrides"></a>
### Nested Schema for `spec.topology.workers.machine_pools.variables.overrides`

Required:

- `name` (String) Name of the variable.
- `value` (Map of String) Value of the variable. Note: the value will be validated against the schema of the corresponding ClusterClassVariable from the ClusterClass. Note: We have to use apiextensionsv1.JSON instead of a custom JSON type, because controller-tools has a hard-coded schema for apiextensionsv1.JSON which cannot be produced by another type via controller-tools, i.e. it is not possible to have no type field. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/d0e03a142d0ecdd5491593e941ee1d6b5d91dba6/pkg/crd/known_types.go#L106-L111

Optional:

- `definition_from` (String) DefinitionFrom specifies where the definition of this Variable is from. Deprecated: This field is deprecated, must not be set anymore and is going to be removed in the next apiVersion.
