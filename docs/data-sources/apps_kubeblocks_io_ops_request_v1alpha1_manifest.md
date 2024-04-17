---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "apps.kubeblocks.io"
description: |-
  OpsRequest is the Schema for the opsrequests API
---

# k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest (Data Source)

OpsRequest is the Schema for the opsrequests API

## Example Usage

```terraform
data "k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) OpsRequestSpec defines the desired state of OpsRequest (see [below for nested schema](#nestedatt--spec))

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

- `cluster_ref` (String) References the cluster object.
- `type` (String) Defines the operation type.

Optional:

- `backup_spec` (Attributes) Defines how to backup the cluster. (see [below for nested schema](#nestedatt--spec--backup_spec))
- `cancel` (Boolean) Defines the action to cancel the 'Pending/Creating/Running' opsRequest, supported types: 'VerticalScaling/HorizontalScaling'. Once set to true, this opsRequest will be canceled and modifying this property again will not take effect.
- `custom_spec` (Attributes) Specifies a custom operation as defined by OpsDefinition. (see [below for nested schema](#nestedatt--spec--custom_spec))
- `expose` (Attributes List) Defines services the component needs to expose. (see [below for nested schema](#nestedatt--spec--expose))
- `horizontal_scaling` (Attributes List) Defines what component need to horizontal scale the specified replicas. (see [below for nested schema](#nestedatt--spec--horizontal_scaling))
- `reconfigure` (Attributes) Deprecated: replace by reconfigures. Defines the variables that need to input when updating configuration. (see [below for nested schema](#nestedatt--spec--reconfigure))
- `reconfigures` (Attributes List) Defines the variables that need to input when updating configuration. (see [below for nested schema](#nestedatt--spec--reconfigures))
- `restart` (Attributes List) Restarts the specified components. (see [below for nested schema](#nestedatt--spec--restart))
- `restore_from` (Attributes) Cluster RestoreFrom backup or point in time. (see [below for nested schema](#nestedatt--spec--restore_from))
- `restore_spec` (Attributes) Defines how to restore the cluster. Note that this restore operation will roll back cluster services. (see [below for nested schema](#nestedatt--spec--restore_spec))
- `script_spec` (Attributes) Defines the script to be executed. (see [below for nested schema](#nestedatt--spec--script_spec))
- `switchover` (Attributes List) Switches over the specified components. (see [below for nested schema](#nestedatt--spec--switchover))
- `ttl_seconds_after_succeed` (Number) OpsRequest will be deleted after TTLSecondsAfterSucceed second when OpsRequest.status.phase is Succeed.
- `ttl_seconds_before_abort` (Number) OpsRequest will wait at most TTLSecondsBeforeAbort seconds for start-conditions to be met. If not specified, the default value is 0, which means that the start-conditions must be met immediately.
- `upgrade` (Attributes) Specifies the cluster version by specifying clusterVersionRef. (see [below for nested schema](#nestedatt--spec--upgrade))
- `vertical_scaling` (List of Map of String) Note: Quantity struct can not do immutable check by CEL. Defines what component need to vertical scale the specified compute resources.
- `volume_expansion` (Attributes List) Note: Quantity struct can not do immutable check by CEL. Defines what component and volumeClaimTemplate need to expand the specified storage. (see [below for nested schema](#nestedatt--spec--volume_expansion))

<a id="nestedatt--spec--backup_spec"></a>
### Nested Schema for `spec.backup_spec`

Optional:

- `backup_method` (String) Defines the backup method that is defined in backupPolicy.
- `backup_name` (String) Specifies the name of the backup.
- `backup_policy_name` (String) Indicates the backupPolicy applied to perform this backup.
- `deletion_policy` (String) Determines whether the backup contents stored in backup repository should be deleted when the backup custom resource is deleted. Supported values are 'Retain' and 'Delete'. - 'Retain' means that the backup content and its physical snapshot on backup repository are kept. - 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.
- `parent_backup_name` (String) If backupType is incremental, parentBackupName is required.
- `retention_period` (String) Determines a duration up to which the backup should be kept. Controller will remove all backups that are older than the RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m  You can also combine the above durations. For example: 30d12h30m. If not set, the backup will be kept forever.


<a id="nestedatt--spec--custom_spec"></a>
### Nested Schema for `spec.custom_spec`

Required:

- `components` (Attributes List) Defines which components need to perform the actions defined by this OpsDefinition. At least one component is required. The components are identified by their name and can be merged or retained. (see [below for nested schema](#nestedatt--spec--custom_spec--components))
- `ops_definition_ref` (String) Is a reference to an OpsDefinition.

Optional:

- `parallelism` (String) Defines the execution concurrency. By default, all incoming Components will be executed simultaneously. The value can be an absolute number (e.g., 5) or a percentage of desired components (e.g., 10%). The absolute number is calculated from the percentage by rounding up. For instance, if the percentage value is 10% and the components length is 1, the calculated number will be rounded up to 1.
- `service_account_name` (String)

<a id="nestedatt--spec--custom_spec--components"></a>
### Nested Schema for `spec.custom_spec.components`

Required:

- `name` (String) Specifies the unique identifier of the cluster component

Optional:

- `parameters` (Attributes List) Represents the parameters for this operation as declared in the opsDefinition.spec.parametersSchema. (see [below for nested schema](#nestedatt--spec--custom_spec--components--parameters))

<a id="nestedatt--spec--custom_spec--components--parameters"></a>
### Nested Schema for `spec.custom_spec.components.parameters`

Required:

- `name` (String) Specifies the identifier of the parameter as defined in the OpsDefinition.
- `value` (String) Holds the data associated with the parameter. If the parameter type is an array, the format should be 'v1,v2,v3'.




<a id="nestedatt--spec--expose"></a>
### Nested Schema for `spec.expose`

Required:

- `component_name` (String) Specifies the name of the cluster component.
- `services` (Attributes List) A list of services that are to be exposed or removed. If componentNamem is not specified, each 'OpsService' in the list must specify ports and selectors. (see [below for nested schema](#nestedatt--spec--expose--services))
- `switch` (String) Controls the expose operation. If set to Enable, the corresponding service will be exposed. Conversely, if set to Disable, the service will be removed.

<a id="nestedatt--spec--expose--services"></a>
### Nested Schema for `spec.expose.services`

Required:

- `name` (String) Specifies the name of the service. This name is used by others to refer to this service (e.g., connection credential). Note: This field cannot be updated.

Optional:

- `annotations` (Map of String) Contains cloud provider related parameters if ServiceType is LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.
- `ports` (Attributes List) Lists the ports that are exposed by this service. If not provided, the default Services Ports defined in the ClusterDefinition or ComponentDefinition that are neither of NodePort nor LoadBalancer service type will be used. If there is no corresponding Service defined in the ClusterDefinition or ComponentDefinition, the expose operation will fail. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies (see [below for nested schema](#nestedatt--spec--expose--services--ports))
- `role_selector` (String) Allows you to specify a defined role as a selector for the service, extending the ServiceSpec.Selector.
- `selector` (Map of String) Routes service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. This only applies to types ClusterIP, NodePort, and LoadBalancer and is ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/
- `service_type` (String) Determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. - 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.

<a id="nestedatt--spec--expose--services--ports"></a>
### Nested Schema for `spec.expose.services.ports`

Required:

- `port` (Number) The port that will be exposed by this service.

Optional:

- `app_protocol` (String) The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:  * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).  * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540 * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455  * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
- `name` (String) The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.
- `node_port` (Number) The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
- `protocol` (String) The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.
- `target_port` (String) Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service




<a id="nestedatt--spec--horizontal_scaling"></a>
### Nested Schema for `spec.horizontal_scaling`

Required:

- `component_name` (String) Specifies the name of the cluster component.
- `replicas` (Number) Specifies the number of replicas for the workloads.

Optional:

- `instances` (List of String) Defines the names of instances that the rsm should prioritize for scale-down operations. If the RsmTransformPolicy is set to ToPod and the expected number of replicas is less than the current number, the list of Instances will be used.  - 'current replicas - expected replicas > len(Instances)': Scale down from the list of Instances priorly, the others will select from NodeAssignment. - 'current replicas - expected replicas < len(Instances)': Scale down from the list of Instances. - 'current replicas - expected replicas < len(Instances)': Scale down from a part of Instances.
- `nodes` (List of String) Defines the list of nodes where pods can be scheduled during a scale-up operation. If the RsmTransformPolicy is set to ToPod and the expected number of replicas is greater than the current number, the list of Nodes will be used. If the list of Nodes is empty, pods will not be assigned to any specific node. However, if the list of Nodes is populated, pods will be evenly distributed across the nodes in the list during scale-up.


<a id="nestedatt--spec--reconfigure"></a>
### Nested Schema for `spec.reconfigure`

Required:

- `component_name` (String) Specifies the name of the cluster component.
- `configurations` (Attributes List) Specifies the components that will perform the operation. (see [below for nested schema](#nestedatt--spec--reconfigure--configurations))

<a id="nestedatt--spec--reconfigure--configurations"></a>
### Nested Schema for `spec.reconfigure.configurations`

Required:

- `keys` (Attributes List) Sets the parameters to be updated. It should contain at least one item. The keys are merged and retained during patch operations. (see [below for nested schema](#nestedatt--spec--reconfigure--configurations--keys))
- `name` (String) Specifies the name of the configuration template.

Optional:

- `policy` (String) Defines the upgrade policy for the configuration. This field is optional.

<a id="nestedatt--spec--reconfigure--configurations--keys"></a>
### Nested Schema for `spec.reconfigure.configurations.keys`

Required:

- `key` (String) Represents the unique identifier for the ConfigMap.

Optional:

- `file_content` (String) Represents the content of the configuration file. This field is used to update the entire content of the file.
- `parameters` (Attributes List) Defines a list of key-value pairs for a single configuration file. These parameters are used to update the specified configuration settings. (see [below for nested schema](#nestedatt--spec--reconfigure--configurations--policy--parameters))

<a id="nestedatt--spec--reconfigure--configurations--policy--parameters"></a>
### Nested Schema for `spec.reconfigure.configurations.policy.parameters`

Required:

- `key` (String) Represents the name of the parameter that is to be updated.

Optional:

- `value` (String) Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.





<a id="nestedatt--spec--reconfigures"></a>
### Nested Schema for `spec.reconfigures`

Required:

- `component_name` (String) Specifies the name of the cluster component.
- `configurations` (Attributes List) Specifies the components that will perform the operation. (see [below for nested schema](#nestedatt--spec--reconfigures--configurations))

<a id="nestedatt--spec--reconfigures--configurations"></a>
### Nested Schema for `spec.reconfigures.configurations`

Required:

- `keys` (Attributes List) Sets the parameters to be updated. It should contain at least one item. The keys are merged and retained during patch operations. (see [below for nested schema](#nestedatt--spec--reconfigures--configurations--keys))
- `name` (String) Specifies the name of the configuration template.

Optional:

- `policy` (String) Defines the upgrade policy for the configuration. This field is optional.

<a id="nestedatt--spec--reconfigures--configurations--keys"></a>
### Nested Schema for `spec.reconfigures.configurations.keys`

Required:

- `key` (String) Represents the unique identifier for the ConfigMap.

Optional:

- `file_content` (String) Represents the content of the configuration file. This field is used to update the entire content of the file.
- `parameters` (Attributes List) Defines a list of key-value pairs for a single configuration file. These parameters are used to update the specified configuration settings. (see [below for nested schema](#nestedatt--spec--reconfigures--configurations--policy--parameters))

<a id="nestedatt--spec--reconfigures--configurations--policy--parameters"></a>
### Nested Schema for `spec.reconfigures.configurations.policy.parameters`

Required:

- `key` (String) Represents the name of the parameter that is to be updated.

Optional:

- `value` (String) Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.





<a id="nestedatt--spec--restart"></a>
### Nested Schema for `spec.restart`

Required:

- `component_name` (String) Specifies the name of the cluster component.


<a id="nestedatt--spec--restore_from"></a>
### Nested Schema for `spec.restore_from`

Optional:

- `backup` (Attributes List) Refers to the backup name and component name used for restoration. Supports recovery of multiple components. (see [below for nested schema](#nestedatt--spec--restore_from--backup))
- `point_in_time` (Attributes) Refers to the specific point in time for recovery. (see [below for nested schema](#nestedatt--spec--restore_from--point_in_time))

<a id="nestedatt--spec--restore_from--backup"></a>
### Nested Schema for `spec.restore_from.backup`

Optional:

- `ref` (Attributes) Refers to a reference backup that needs to be restored. (see [below for nested schema](#nestedatt--spec--restore_from--backup--ref))

<a id="nestedatt--spec--restore_from--backup--ref"></a>
### Nested Schema for `spec.restore_from.backup.ref`

Optional:

- `name` (String) Refers to the specific name of the resource.
- `namespace` (String) Refers to the specific namespace of the resource.



<a id="nestedatt--spec--restore_from--point_in_time"></a>
### Nested Schema for `spec.restore_from.point_in_time`

Optional:

- `ref` (Attributes) Refers to a reference source cluster that needs to be restored. (see [below for nested schema](#nestedatt--spec--restore_from--point_in_time--ref))
- `time` (String) Refers to the specific time point for restoration, with UTC as the time zone.

<a id="nestedatt--spec--restore_from--point_in_time--ref"></a>
### Nested Schema for `spec.restore_from.point_in_time.ref`

Optional:

- `name` (String) Refers to the specific name of the resource.
- `namespace` (String) Refers to the specific namespace of the resource.




<a id="nestedatt--spec--restore_spec"></a>
### Nested Schema for `spec.restore_spec`

Required:

- `backup_name` (String) Specifies the name of the backup.

Optional:

- `effective_common_component_def` (Boolean) Indicates if this backup will be restored for all components which refer to common ComponentDefinition.
- `restore_time_str` (String) Defines the point in time to restore.
- `volume_restore_policy` (String) Specifies the volume claim restore policy, support values: [Serial, Parallel]


<a id="nestedatt--spec--script_spec"></a>
### Nested Schema for `spec.script_spec`

Required:

- `component_name` (String) Specifies the name of the cluster component.

Optional:

- `image` (String) Specifies the image to be used for the exec command. By default, the image of kubeblocks-datascript is used.
- `script` (List of String) Defines the script to be executed.
- `script_from` (Attributes) Defines the script to be executed from a configMap or secret. (see [below for nested schema](#nestedatt--spec--script_spec--script_from))
- `secret` (Attributes) Defines the secret to be used to execute the script. If not specified, the default cluster root credential secret is used. (see [below for nested schema](#nestedatt--spec--script_spec--secret))
- `selector` (Attributes) By default, KubeBlocks will execute the script on the primary pod with role=leader. Exceptions exist, such as Redis, which does not synchronize account information between primary and secondary. In such cases, the script needs to be executed on all pods matching the selector. Indicates the components on which the script is executed. (see [below for nested schema](#nestedatt--spec--script_spec--selector))

<a id="nestedatt--spec--script_spec--script_from"></a>
### Nested Schema for `spec.script_spec.script_from`

Optional:

- `config_map_ref` (Attributes List) Specifies the configMap that is to be executed. (see [below for nested schema](#nestedatt--spec--script_spec--script_from--config_map_ref))
- `secret_ref` (Attributes List) Specifies the secret that is to be executed. (see [below for nested schema](#nestedatt--spec--script_spec--script_from--secret_ref))

<a id="nestedatt--spec--script_spec--script_from--config_map_ref"></a>
### Nested Schema for `spec.script_spec.script_from.config_map_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--script_spec--script_from--secret_ref"></a>
### Nested Schema for `spec.script_spec.script_from.secret_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined



<a id="nestedatt--spec--script_spec--secret"></a>
### Nested Schema for `spec.script_spec.secret`

Required:

- `name` (String) Specifies the name of the secret.

Optional:

- `password_key` (String) Used to specify the password part of the secret.
- `username_key` (String) Used to specify the username part of the secret.


<a id="nestedatt--spec--script_spec--selector"></a>
### Nested Schema for `spec.script_spec.selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--script_spec--selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--script_spec--selector--match_expressions"></a>
### Nested Schema for `spec.script_spec.selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.




<a id="nestedatt--spec--switchover"></a>
### Nested Schema for `spec.switchover`

Required:

- `component_name` (String) Specifies the name of the cluster component.
- `instance_name` (String) Utilized to designate the candidate primary or leader instance for the switchover process. If assigned '*', it signifies that no specific primary or leader is designated for the switchover, and the switchoverAction defined in 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' will be executed.  It is mandatory that 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' is not left blank.  If assigned a valid instance name other than '*', it signifies that a specific candidate primary or leader is designated for the switchover. The value can be retrieved using 'kbcli cluster list-instances', any other value is considered invalid.  In this scenario, the 'switchoverAction' defined in clusterDefinition.componentDefs[x].switchoverSpec.withCandidate will be executed, and it is mandatory that clusterDefinition.componentDefs[x].switchoverSpec.withCandidate is not left blank.


<a id="nestedatt--spec--upgrade"></a>
### Nested Schema for `spec.upgrade`

Required:

- `cluster_version_ref` (String) A reference to the name of the ClusterVersion.


<a id="nestedatt--spec--volume_expansion"></a>
### Nested Schema for `spec.volume_expansion`

Required:

- `component_name` (String) Specifies the name of the cluster component.
- `volume_claim_templates` (Attributes List) volumeClaimTemplates specifies the storage size and volumeClaimTemplate name. (see [below for nested schema](#nestedatt--spec--volume_expansion--volume_claim_templates))

<a id="nestedatt--spec--volume_expansion--volume_claim_templates"></a>
### Nested Schema for `spec.volume_expansion.volume_claim_templates`

Required:

- `name` (String) A reference to the volumeClaimTemplate name from the cluster components.
- `storage` (String) Specifies the requested storage size for the volume.