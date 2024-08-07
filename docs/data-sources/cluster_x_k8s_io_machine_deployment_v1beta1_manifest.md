---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_cluster_x_k8s_io_machine_deployment_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "cluster.x-k8s.io"
description: |-
  MachineDeployment is the Schema for the machinedeployments API.
---

# k8s_cluster_x_k8s_io_machine_deployment_v1beta1_manifest (Data Source)

MachineDeployment is the Schema for the machinedeployments API.

## Example Usage

```terraform
data "k8s_cluster_x_k8s_io_machine_deployment_v1beta1_manifest" "example" {
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

- `spec` (Attributes) MachineDeploymentSpec defines the desired state of MachineDeployment. (see [below for nested schema](#nestedatt--spec))

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

- `cluster_name` (String) ClusterName is the name of the Cluster this object belongs to.
- `selector` (Attributes) Label selector for machines. Existing MachineSets whose machines areselected by this will be the ones affected by this deployment.It must match the machine template's labels. (see [below for nested schema](#nestedatt--spec--selector))
- `template` (Attributes) Template describes the machines that will be created. (see [below for nested schema](#nestedatt--spec--template))

Optional:

- `min_ready_seconds` (Number) MinReadySeconds is the minimum number of seconds for which a Node for a newly created machine should be ready before considering the replica available.Defaults to 0 (machine will be considered available as soon as the Node is ready)
- `paused` (Boolean) Indicates that the deployment is paused.
- `progress_deadline_seconds` (Number) The maximum time in seconds for a deployment to make progress before itis considered to be failed. The deployment controller will continue toprocess failed deployments and a condition with a ProgressDeadlineExceededreason will be surfaced in the deployment status. Note that progress willnot be estimated during the time a deployment is paused. Defaults to 600s.
- `replicas` (Number) Number of desired machines.This is a pointer to distinguish between explicit zero and not specified.Defaults to:* if the Kubernetes autoscaler min size and max size annotations are set:  - if it's a new MachineDeployment, use min size  - if the replicas field of the old MachineDeployment is < min size, use min size  - if the replicas field of the old MachineDeployment is > max size, use max size  - if the replicas field of the old MachineDeployment is in the (min size, max size) range, keep the value from the oldMD* otherwise use 1Note: Defaulting will be run whenever the replicas field is not set:* A new MachineDeployment is created with replicas not set.* On an existing MachineDeployment the replicas field was first set and is now unset.Those cases are especially relevant for the following Kubernetes autoscaler use cases:* A new MachineDeployment is created and replicas should be managed by the autoscaler* An existing MachineDeployment which initially wasn't controlled by the autoscaler  should be later controlled by the autoscaler
- `revision_history_limit` (Number) The number of old MachineSets to retain to allow rollback.This is a pointer to distinguish between explicit zero and not specified.Defaults to 1.Deprecated: This field is deprecated and is going to be removed in the next apiVersion. Please see https://github.com/kubernetes-sigs/cluster-api/issues/10479 for more details.
- `rollout_after` (String) RolloutAfter is a field to indicate a rollout should be performedafter the specified time even if no changes have been made to theMachineDeployment.Example: In the YAML the time can be specified in the RFC3339 format.To specify the rolloutAfter target as March 9, 2023, at 9 am UTCuse '2023-03-09T09:00:00Z'.
- `strategy` (Attributes) The deployment strategy to use to replace existing machines withnew ones. (see [below for nested schema](#nestedatt--spec--strategy))

<a id="nestedatt--spec--selector"></a>
### Nested Schema for `spec.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--selector--match_expressions"></a>
### Nested Schema for `spec.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.



<a id="nestedatt--spec--template"></a>
### Nested Schema for `spec.template`

Optional:

- `metadata` (Attributes) Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata (see [below for nested schema](#nestedatt--spec--template--metadata))
- `spec` (Attributes) Specification of the desired behavior of the machine.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status (see [below for nested schema](#nestedatt--spec--template--spec))

<a id="nestedatt--spec--template--metadata"></a>
### Nested Schema for `spec.template.metadata`

Optional:

- `annotations` (Map of String) Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations
- `labels` (Map of String) Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels


<a id="nestedatt--spec--template--spec"></a>
### Nested Schema for `spec.template.spec`

Required:

- `bootstrap` (Attributes) Bootstrap is a reference to a local struct which encapsulatesfields to configure the Machine’s bootstrapping mechanism. (see [below for nested schema](#nestedatt--spec--template--spec--bootstrap))
- `cluster_name` (String) ClusterName is the name of the Cluster this object belongs to.
- `infrastructure_ref` (Attributes) InfrastructureRef is a required reference to a custom resourceoffered by an infrastructure provider. (see [below for nested schema](#nestedatt--spec--template--spec--infrastructure_ref))

Optional:

- `failure_domain` (String) FailureDomain is the failure domain the machine will be created in.Must match a key in the FailureDomains map stored on the cluster object.
- `node_deletion_timeout` (String) NodeDeletionTimeout defines how long the controller will attempt to delete the Node that the Machinehosts after the Machine is marked for deletion. A duration of 0 will retry deletion indefinitely.Defaults to 10 seconds.
- `node_drain_timeout` (String) NodeDrainTimeout is the total amount of time that the controller will spend on draining a node.The default value is 0, meaning that the node can be drained without any time limitations.NOTE: NodeDrainTimeout is different from 'kubectl drain --timeout'
- `node_volume_detach_timeout` (String) NodeVolumeDetachTimeout is the total amount of time that the controller will spend on waiting for all volumesto be detached. The default value is 0, meaning that the volumes can be detached without any time limitations.
- `provider_id` (String) ProviderID is the identification ID of the machine provided by the provider.This field must match the provider ID as seen on the node object corresponding to this machine.This field is required by higher level consumers of cluster-api. Example use case is cluster autoscalerwith cluster-api as provider. Clean-up logic in the autoscaler compares machines to nodes to find outmachines at provider which could not get registered as Kubernetes nodes. With cluster-api as ageneric out-of-tree provider for autoscaler, this field is required by autoscaler to beable to have a provider view of the list of machines. Another list of nodes is queried from the k8s apiserverand then a comparison is done to find out unregistered machines and are marked for delete.This field will be set by the actuators and consumed by higher level entities like autoscaler that willbe interfacing with cluster-api as generic provider.
- `version` (String) Version defines the desired Kubernetes version.This field is meant to be optionally used by bootstrap providers.

<a id="nestedatt--spec--template--spec--bootstrap"></a>
### Nested Schema for `spec.template.spec.bootstrap`

Optional:

- `config_ref` (Attributes) ConfigRef is a reference to a bootstrap provider-specific resourcethat holds configuration details. The reference is optional toallow users/operators to specify Bootstrap.DataSecretName withoutthe need of a controller. (see [below for nested schema](#nestedatt--spec--template--spec--bootstrap--config_ref))
- `data_secret_name` (String) DataSecretName is the name of the secret that stores the bootstrap data script.If nil, the Machine should remain in the Pending state.

<a id="nestedatt--spec--template--spec--bootstrap--config_ref"></a>
### Nested Schema for `spec.template.spec.bootstrap.config_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids



<a id="nestedatt--spec--template--spec--infrastructure_ref"></a>
### Nested Schema for `spec.template.spec.infrastructure_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids




<a id="nestedatt--spec--strategy"></a>
### Nested Schema for `spec.strategy`

Optional:

- `remediation` (Attributes) Remediation controls the strategy of remediating unhealthy machinesand how remediating operations should occur during the lifecycle of the dependant MachineSets. (see [below for nested schema](#nestedatt--spec--strategy--remediation))
- `rolling_update` (Attributes) Rolling update config params. Present only ifMachineDeploymentStrategyType = RollingUpdate. (see [below for nested schema](#nestedatt--spec--strategy--rolling_update))
- `type` (String) Type of deployment. Allowed values are RollingUpdate and OnDelete.The default is RollingUpdate.

<a id="nestedatt--spec--strategy--remediation"></a>
### Nested Schema for `spec.strategy.remediation`

Optional:

- `max_in_flight` (String) MaxInFlight determines how many in flight remediations should happen at the same time.Remediation only happens on the MachineSet with the most current revision, whileolder MachineSets (usually present during rollout operations) aren't allowed to remediate.Note: In general (independent of remediations), unhealthy machines are alwaysprioritized during scale down operations over healthy ones.MaxInFlight can be set to a fixed number or a percentage.Example: when this is set to 20%, the MachineSet controller deletes at most 20% ofthe desired replicas.If not set, remediation is limited to all machines (bounded by replicas)under the active MachineSet's management.


<a id="nestedatt--spec--strategy--rolling_update"></a>
### Nested Schema for `spec.strategy.rolling_update`

Optional:

- `delete_policy` (String) DeletePolicy defines the policy used by the MachineDeployment to identify nodes to delete when downscaling.Valid values are 'Random, 'Newest', 'Oldest'When no value is supplied, the default DeletePolicy of MachineSet is used
- `max_surge` (String) The maximum number of machines that can be scheduled above thedesired number of machines.Value can be an absolute number (ex: 5) or a percentage ofdesired machines (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 1.Example: when this is set to 30%, the new MachineSet can be scaledup immediately when the rolling update starts, such that the totalnumber of old and new machines do not exceed 130% of desiredmachines. Once old machines have been killed, new MachineSet canbe scaled up further, ensuring that total number of machines runningat any time during the update is at most 130% of desired machines.
- `max_unavailable` (String) The maximum number of machines that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desiredmachines (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 0.Example: when this is set to 30%, the old MachineSet can be scaleddown to 70% of desired machines immediately when the rolling updatestarts. Once new machines are ready, old MachineSet can be scaleddown further, followed by scaling up the new MachineSet, ensuringthat the total number of machines available at all timesduring the update is at least 70% of desired machines.
