---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_storageos_com_storage_os_cluster_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "storageos.com"
description: |-
  StorageOSCluster is the Schema for the storageosclusters API
---

# k8s_storageos_com_storage_os_cluster_v1_manifest (Data Source)

StorageOSCluster is the Schema for the storageosclusters API

## Example Usage

```terraform
data "k8s_storageos_com_storage_os_cluster_v1_manifest" "example" {
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

- `spec` (Attributes) StorageOSClusterSpec defines the desired state of StorageOSCluster (see [below for nested schema](#nestedatt--spec))

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

- `kv_backend` (Attributes) KVBackend defines the key-value store backend used in the cluster. (see [below for nested schema](#nestedatt--spec--kv_backend))
- `secret_ref_name` (String) SecretRefName is the name of the secret object that contains all the sensitive cluster configurations.

Optional:

- `container_resources` (Attributes) ContainerResources is to set the resource requirements of each individual container managed by the operator. (see [below for nested schema](#nestedatt--spec--container_resources))
- `csi` (Attributes) CSI defines the configurations for CSI. (see [below for nested schema](#nestedatt--spec--csi))
- `debug` (Boolean) Debug is to set debug mode of the cluster.
- `disable_cli` (Boolean) Disable StorageOS CLI deployment.
- `disable_fencing` (Boolean) Disable Pod Fencing.  With StatefulSets, Pods are only re-scheduled if the Pod has been marked as killed.  In practice this means that failover of a StatefulSet pod is a manual operation.  By enabling Pod Fencing and setting the 'storageos.com/fenced=true' label on a Pod, StorageOS will enable automated Pod failover (by killing the application Pod on the failed node) if the following conditions exist:  - Pod fencing has not been explicitly disabled. - StorageOS has determined that the node the Pod is running on is offline.  StorageOS uses Gossip and TCP checks and will retry for 30 seconds.  At this point all volumes on the failed node are marked offline (irrespective of whether fencing is enabled) and volume failover starts. - The Pod has the label 'storageos.com/fenced=true' set. - The Pod has at least one StorageOS volume attached. - Each StorageOS volume has at least 1 healthy replica.  When Pod Fencing is disabled, StorageOS will not perform any interaction with Kubernetes when it detects that a node has gone offline. Additionally, the Kubernetes permissions required for Fencing will not be added to the StorageOS role. Deprecated: Not used any more, fencing is enabled/disabled by storageos.com/fenced label on pod.
- `disable_scheduler` (Boolean) Disable StorageOS scheduler extender.
- `disable_tcmu` (Boolean) Disable TCMU can be set to true to disable the TCMU storage driver.  This is required when there are multiple storage systems running on the same node and you wish to avoid conflicts.  Only one TCMU-based storage system can run on a node at a time.  Disabling TCMU will degrade performance. Deprecated: Not used any more.
- `disable_telemetry` (Boolean) Disable Telemetry.
- `enable_portal_manager` (Boolean) EnablePortalManager enables Portal Manager.
- `environment` (Map of String) Environment contains environment variables that are passed to StorageOS.
- `force_tcmu` (Boolean) Force TCMU can be set to true to ensure that TCMU is enabled or cause StorageOS to abort startup.  At startup, StorageOS will automatically fallback to non-TCMU mode if another TCMU-based storage system is running on the node.  Since non-TCMU will degrade performance, this may not always be desired. Deprecated: Not used any more.
- `images` (Attributes) Images defines the various container images used in the cluster. (see [below for nested schema](#nestedatt--spec--images))
- `ingress` (Attributes) Ingress defines the ingress configurations used in the cluster. Deprecated: Not used any more, please create your ingress for dashboard on your own. (see [below for nested schema](#nestedatt--spec--ingress))
- `join` (String) Join is the join token used for service discovery. Deprecated: Not used any more.
- `k8s_distro` (String) K8sDistro is the name of the Kubernetes distribution where the operator is being deployed.  It should be in the format: 'name[-1.0]', where the version is optional and should only be appended if known.  Suitable names include: 'openshift', 'rancher', 'aks', 'gke', 'eks', or the deployment method if using upstream directly, e.g 'minishift' or 'kubeadm'.  Setting k8sDistro is optional, and will be used to simplify cluster configuration by setting appropriate defaults for the distribution.  The distribution information will also be included in the product telemetry (if enabled), to help focus development efforts.
- `metrics` (Attributes) Metrics feature configuration. (see [below for nested schema](#nestedatt--spec--metrics))
- `namespace` (String) Namespace is the kubernetes Namespace where storageos resources are provisioned. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.
- `node_manager_features` (Map of String) Node manager feature list with optional configurations.
- `node_selector_terms` (Attributes List) NodeSelectorTerms is to set the placement of storageos pods using node affinity requiredDuringSchedulingIgnoredDuringExecution. (see [below for nested schema](#nestedatt--spec--node_selector_terms))
- `pause` (Boolean) Pause is to pause the operator for the cluster. Deprecated: Not used any more, operator is always running.
- `resources` (Attributes) Resources is to set the resource requirements of the storageos containers. Deprecated: Set resource requests for individual containers via ContainerResources field in spec. (see [below for nested schema](#nestedatt--spec--resources))
- `secret_ref_namespace` (String) SecretRefNamespace is the namespace of the secret reference. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.
- `service` (Attributes) Service is the Service configuration for the cluster nodes. (see [below for nested schema](#nestedatt--spec--service))
- `shared_dir` (String) SharedDir is the shared directory to be used when the kubelet is running in a container. Typically: '/var/lib/kubelet/plugins/kubernetes.io~storageos'. If not set, defaults will be used.
- `snapshots` (Attributes) Snapshots feature configuration. (see [below for nested schema](#nestedatt--spec--snapshots))
- `storage_class_name` (String) StorageClassName is the name of default StorageClass created for StorageOS volumes.
- `tls_etcd_secret_ref_name` (String) TLSEtcdSecretRefName is the name of the secret object that contains the etcd TLS certs. This secret is shared with etcd, therefore it's not part of the main storageos secret.
- `tls_etcd_secret_ref_namespace` (String) TLSEtcdSecretRefNamespace is the namespace of the etcd TLS secret object. Deprecated: StorageOS uses namespace of storageosclusters.storageos.com resource.
- `tolerations` (Attributes List) Tolerations is to set the placement of storageos pods using pod toleration. (see [below for nested schema](#nestedatt--spec--tolerations))

<a id="nestedatt--spec--kv_backend"></a>
### Nested Schema for `spec.kv_backend`

Required:

- `address` (String)

Optional:

- `backend` (String)


<a id="nestedatt--spec--container_resources"></a>
### Nested Schema for `spec.container_resources`

Optional:

- `api_manager_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--api_manager_container))
- `cli_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--cli_container))
- `csi_external_attacher_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--csi_external_attacher_container))
- `csi_external_provisioner_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--csi_external_provisioner_container))
- `csi_external_resizer_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--csi_external_resizer_container))
- `csi_external_snapshotter_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--csi_external_snapshotter_container))
- `csi_liveness_probe_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--csi_liveness_probe_container))
- `csi_node_driver_registrar_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--csi_node_driver_registrar_container))
- `init_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--init_container))
- `kube_scheduler_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--kube_scheduler_container))
- `metrics_exporter_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--metrics_exporter_container))
- `node_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--node_container))
- `node_manager_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--node_manager_container))
- `portal_manager_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--portal_manager_container))
- `snapshot_controller_container` (Attributes) ResourceRequirements describes the compute resource requirements. (see [below for nested schema](#nestedatt--spec--container_resources--snapshot_controller_container))

<a id="nestedatt--spec--container_resources--api_manager_container"></a>
### Nested Schema for `spec.container_resources.api_manager_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--cli_container"></a>
### Nested Schema for `spec.container_resources.cli_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--csi_external_attacher_container"></a>
### Nested Schema for `spec.container_resources.csi_external_attacher_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--csi_external_provisioner_container"></a>
### Nested Schema for `spec.container_resources.csi_external_provisioner_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--csi_external_resizer_container"></a>
### Nested Schema for `spec.container_resources.csi_external_resizer_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--csi_external_snapshotter_container"></a>
### Nested Schema for `spec.container_resources.csi_external_snapshotter_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--csi_liveness_probe_container"></a>
### Nested Schema for `spec.container_resources.csi_liveness_probe_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--csi_node_driver_registrar_container"></a>
### Nested Schema for `spec.container_resources.csi_node_driver_registrar_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--init_container"></a>
### Nested Schema for `spec.container_resources.init_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--kube_scheduler_container"></a>
### Nested Schema for `spec.container_resources.kube_scheduler_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--metrics_exporter_container"></a>
### Nested Schema for `spec.container_resources.metrics_exporter_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--node_container"></a>
### Nested Schema for `spec.container_resources.node_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--node_manager_container"></a>
### Nested Schema for `spec.container_resources.node_manager_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--portal_manager_container"></a>
### Nested Schema for `spec.container_resources.portal_manager_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--container_resources--snapshot_controller_container"></a>
### Nested Schema for `spec.container_resources.snapshot_controller_container`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/



<a id="nestedatt--spec--csi"></a>
### Nested Schema for `spec.csi`

Optional:

- `attacher_timeout` (String)
- `deployment_strategy` (String)
- `device_dir` (String)
- `driver_registeration_mode` (String)
- `driver_requires_attachment` (String)
- `enable` (Boolean)
- `enable_controller_expand_creds` (Boolean)
- `enable_controller_publish_creds` (Boolean)
- `enable_node_publish_creds` (Boolean)
- `enable_provision_creds` (Boolean)
- `endpoint` (String)
- `kubelet_dir` (String)
- `kubelet_registration_path` (String)
- `plugin_dir` (String)
- `provisioner_timeout` (String)
- `provisioner_worker_count` (Number)
- `registrar_socket_dir` (String)
- `registration_dir` (String)
- `resizer_timeout` (String)
- `snapshotter_timeout` (String)
- `version` (String)


<a id="nestedatt--spec--images"></a>
### Nested Schema for `spec.images`

Optional:

- `api_manager_container` (String)
- `cli_container` (String)
- `csi_cluster_driver_registrar_container` (String)
- `csi_external_attacher_container` (String)
- `csi_external_provisioner_container` (String)
- `csi_external_resizer_container` (String)
- `csi_external_snapshotter_container` (String)
- `csi_liveness_probe_container` (String)
- `csi_node_driver_registrar_container` (String)
- `hyperkube_container` (String)
- `init_container` (String)
- `kube_scheduler_container` (String)
- `metrics_exporter_container` (String)
- `nfs_container` (String)
- `node_container` (String)
- `node_guard_container` (String)
- `node_manager_container` (String)
- `portal_manager_container` (String)
- `snapshot_controller_container` (String)


<a id="nestedatt--spec--ingress"></a>
### Nested Schema for `spec.ingress`

Optional:

- `annotations` (Map of String)
- `enable` (Boolean)
- `hostname` (String)
- `tls` (Boolean)


<a id="nestedatt--spec--metrics"></a>
### Nested Schema for `spec.metrics`

Optional:

- `disabled_collectors` (List of String) DisabledCollectors is a list of collectors that shall be disabled. By default, all are enabled.
- `enabled` (Boolean)
- `log_level` (String) Verbosity of log messages. Accepts go.uber.org/zap log levels.
- `timeout` (Number) Timeout in seconds to serve metrics.


<a id="nestedatt--spec--node_selector_terms"></a>
### Nested Schema for `spec.node_selector_terms`

Optional:

- `match_expressions` (Attributes List) A list of node selector requirements by node's labels. (see [below for nested schema](#nestedatt--spec--node_selector_terms--match_expressions))
- `match_fields` (Attributes List) A list of node selector requirements by node's fields. (see [below for nested schema](#nestedatt--spec--node_selector_terms--match_fields))

<a id="nestedatt--spec--node_selector_terms--match_expressions"></a>
### Nested Schema for `spec.node_selector_terms.match_expressions`

Required:

- `key` (String) The label key that the selector applies to.
- `operator` (String) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.

Optional:

- `values` (List of String) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.


<a id="nestedatt--spec--node_selector_terms--match_fields"></a>
### Nested Schema for `spec.node_selector_terms.match_fields`

Required:

- `key` (String) The label key that the selector applies to.
- `operator` (String) Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.

Optional:

- `values` (List of String) An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.



<a id="nestedatt--spec--resources"></a>
### Nested Schema for `spec.resources`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--service"></a>
### Nested Schema for `spec.service`

Required:

- `name` (String)
- `type` (String)

Optional:

- `annotations` (Map of String)
- `external_port` (Number)
- `internal_port` (Number)


<a id="nestedatt--spec--snapshots"></a>
### Nested Schema for `spec.snapshots`

Optional:

- `volume_snapshot_class_name` (String) VolumeSnapshotClassName is the name of default VolumeSnapshotClass created for StorageOS volumes.


<a id="nestedatt--spec--tolerations"></a>
### Nested Schema for `spec.tolerations`

Optional:

- `effect` (String) Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
- `key` (String) Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
- `operator` (String) Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
- `toleration_seconds` (Number) TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
- `value` (String) Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.