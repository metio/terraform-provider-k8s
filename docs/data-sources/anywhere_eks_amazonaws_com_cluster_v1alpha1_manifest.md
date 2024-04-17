---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "anywhere.eks.amazonaws.com"
description: |-
  Cluster is the Schema for the clusters API.
---

# k8s_anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest (Data Source)

Cluster is the Schema for the clusters API.

## Example Usage

```terraform
data "k8s_anywhere_eks_amazonaws_com_cluster_v1alpha1_manifest" "example" {
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

- `bundles_ref` (Attributes) BundlesRef contains a reference to the Bundles containing the desired dependencies for the cluster. DEPRECATED: Use EksaVersion instead. (see [below for nested schema](#nestedatt--spec--bundles_ref))
- `cluster_network` (Attributes) (see [below for nested schema](#nestedatt--spec--cluster_network))
- `control_plane_configuration` (Attributes) (see [below for nested schema](#nestedatt--spec--control_plane_configuration))
- `datacenter_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--datacenter_ref))
- `eksa_version` (String) EksaVersion is the semver identifying the release of eks-a used to populate the cluster components.
- `etcd_encryption` (Attributes List) (see [below for nested schema](#nestedatt--spec--etcd_encryption))
- `external_etcd_configuration` (Attributes) ExternalEtcdConfiguration defines the configuration options for using unstacked etcd topology. (see [below for nested schema](#nestedatt--spec--external_etcd_configuration))
- `git_ops_ref` (Attributes) (see [below for nested schema](#nestedatt--spec--git_ops_ref))
- `identity_provider_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--identity_provider_refs))
- `kubernetes_version` (String)
- `machine_health_check` (Attributes) MachineHealthCheck allows to configure timeouts for machine health checks. Machine Health Checks are responsible for remediating unhealthy Machines. Configuring these values will decide how long to wait to remediate unhealthy machine or determine health of nodes' machines. (see [below for nested schema](#nestedatt--spec--machine_health_check))
- `management_cluster` (Attributes) (see [below for nested schema](#nestedatt--spec--management_cluster))
- `packages` (Attributes) PackageConfiguration for installing EKS Anywhere curated packages. (see [below for nested schema](#nestedatt--spec--packages))
- `pod_iam_config` (Attributes) (see [below for nested schema](#nestedatt--spec--pod_iam_config))
- `proxy_configuration` (Attributes) (see [below for nested schema](#nestedatt--spec--proxy_configuration))
- `registry_mirror_configuration` (Attributes) RegistryMirrorConfiguration defines the settings for image registry mirror. (see [below for nested schema](#nestedatt--spec--registry_mirror_configuration))
- `worker_node_group_configurations` (Attributes List) (see [below for nested schema](#nestedatt--spec--worker_node_group_configurations))

<a id="nestedatt--spec--bundles_ref"></a>
### Nested Schema for `spec.bundles_ref`

Required:

- `api_version` (String) APIVersion refers to the Bundles APIVersion
- `name` (String) Name refers to the name of the Bundles object in the cluster
- `namespace` (String) Namespace refers to the Bundles's namespace


<a id="nestedatt--spec--cluster_network"></a>
### Nested Schema for `spec.cluster_network`

Optional:

- `cni` (String) Deprecated. Use CNIConfig
- `cni_config` (Attributes) CNIConfig specifies the CNI plugin to be installed in the cluster (see [below for nested schema](#nestedatt--spec--cluster_network--cni_config))
- `dns` (Attributes) (see [below for nested schema](#nestedatt--spec--cluster_network--dns))
- `nodes` (Attributes) (see [below for nested schema](#nestedatt--spec--cluster_network--nodes))
- `pods` (Attributes) Comma-separated list of CIDR blocks to use for pod and service subnets. Defaults to 192.168.0.0/16 for pod subnet. (see [below for nested schema](#nestedatt--spec--cluster_network--pods))
- `services` (Attributes) (see [below for nested schema](#nestedatt--spec--cluster_network--services))

<a id="nestedatt--spec--cluster_network--cni_config"></a>
### Nested Schema for `spec.cluster_network.cni_config`

Optional:

- `cilium` (Attributes) CiliumConfig contains configuration specific to the Cilium CNI. (see [below for nested schema](#nestedatt--spec--cluster_network--cni_config--cilium))
- `kindnetd` (Map of String) KindnetdConfig contains configuration specific to the Kindnetd CNI.

<a id="nestedatt--spec--cluster_network--cni_config--cilium"></a>
### Nested Schema for `spec.cluster_network.cni_config.cilium`

Optional:

- `egress_masquerade_interfaces` (String) EgressMasquaradeInterfaces determines which network interfaces are used for masquerading. Accepted values are a valid interface name or interface prefix.
- `ipv4_native_routing_cidr` (String) IPv4NativeRoutingCIDR specifies the CIDR to use when RoutingMode is set to direct. When specified, Cilium assumes networking for this CIDR is preconfigured and hands traffic destined for that range to the Linux network stack without applying any SNAT. If this is not set autoDirectNodeRoutes will be set to true
- `ipv6_native_routing_cidr` (String) IPv6NativeRoutingCIDR specifies the IPv6 CIDR to use when RoutingMode is set to direct. When specified, Cilium assumes networking for this CIDR is preconfigured and hands traffic destined for that range to the Linux network stack without applying any SNAT. If this is not set autoDirectNodeRoutes will be set to true
- `policy_enforcement_mode` (String) PolicyEnforcementMode determines communication allowed between pods. Accepted values are default, always, never.
- `routing_mode` (String) RoutingMode indicates the routing tunnel mode to use for Cilium. Accepted values are overlay (geneve tunnel with overlay) or direct (tunneling disabled with direct routing) Defaults to overlay.
- `skip_upgrade` (Boolean) SkipUpgrade indicicates that Cilium maintenance should be skipped during upgrades. This can be used when operators wish to self manage the Cilium installation.



<a id="nestedatt--spec--cluster_network--dns"></a>
### Nested Schema for `spec.cluster_network.dns`

Optional:

- `resolv_conf` (Attributes) ResolvConf refers to the DNS resolver configuration (see [below for nested schema](#nestedatt--spec--cluster_network--dns--resolv_conf))

<a id="nestedatt--spec--cluster_network--dns--resolv_conf"></a>
### Nested Schema for `spec.cluster_network.dns.resolv_conf`

Optional:

- `path` (String) Path defines the path to the file that contains the DNS resolver configuration



<a id="nestedatt--spec--cluster_network--nodes"></a>
### Nested Schema for `spec.cluster_network.nodes`

Optional:

- `cidr_mask_size` (Number) CIDRMaskSize defines the mask size for node cidr in the cluster, default for ipv4 is 24. This is an optional field


<a id="nestedatt--spec--cluster_network--pods"></a>
### Nested Schema for `spec.cluster_network.pods`

Optional:

- `cidr_blocks` (List of String)


<a id="nestedatt--spec--cluster_network--services"></a>
### Nested Schema for `spec.cluster_network.services`

Optional:

- `cidr_blocks` (List of String)



<a id="nestedatt--spec--control_plane_configuration"></a>
### Nested Schema for `spec.control_plane_configuration`

Optional:

- `cert_sans` (List of String) CertSANs is a slice of domain names or IPs to be added as Subject Name Alternatives of the Kube API Servers Certificate.
- `count` (Number) Count defines the number of desired control plane nodes. Defaults to 1.
- `endpoint` (Attributes) Endpoint defines the host ip and port to use for the control plane. (see [below for nested schema](#nestedatt--spec--control_plane_configuration--endpoint))
- `labels` (Map of String) Labels define the labels to assign to the node
- `machine_group_ref` (Attributes) MachineGroupRef defines the machine group configuration for the control plane. (see [below for nested schema](#nestedatt--spec--control_plane_configuration--machine_group_ref))
- `machine_health_check` (Attributes) MachineHealthCheck is a control-plane level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used. (see [below for nested schema](#nestedatt--spec--control_plane_configuration--machine_health_check))
- `skip_load_balancer_deployment` (Boolean) SkipLoadBalancerDeployment skip deploying control plane load balancer. Make sure your infrastructure can handle control plane load balancing when you set this field to true.
- `taints` (Attributes List) Taints define the set of taints to be applied on control plane nodes (see [below for nested schema](#nestedatt--spec--control_plane_configuration--taints))
- `upgrade_rollout_strategy` (Attributes) UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades and related parameters/knobs (see [below for nested schema](#nestedatt--spec--control_plane_configuration--upgrade_rollout_strategy))

<a id="nestedatt--spec--control_plane_configuration--endpoint"></a>
### Nested Schema for `spec.control_plane_configuration.endpoint`

Required:

- `host` (String) Host defines the ip that you want to use to connect to the control plane


<a id="nestedatt--spec--control_plane_configuration--machine_group_ref"></a>
### Nested Schema for `spec.control_plane_configuration.machine_group_ref`

Optional:

- `kind` (String)
- `name` (String)


<a id="nestedatt--spec--control_plane_configuration--machine_health_check"></a>
### Nested Schema for `spec.control_plane_configuration.machine_health_check`

Optional:

- `max_unhealthy` (String) MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.
- `node_startup_timeout` (String) NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.
- `unhealthy_machine_timeout` (String) UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).


<a id="nestedatt--spec--control_plane_configuration--taints"></a>
### Nested Schema for `spec.control_plane_configuration.taints`

Required:

- `effect` (String) Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
- `key` (String) Required. The taint key to be applied to a node.

Optional:

- `time_added` (String) TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
- `value` (String) The taint value corresponding to the taint key.


<a id="nestedatt--spec--control_plane_configuration--upgrade_rollout_strategy"></a>
### Nested Schema for `spec.control_plane_configuration.upgrade_rollout_strategy`

Optional:

- `rolling_update` (Attributes) ControlPlaneRollingUpdateParams is API for rolling update strategy knobs. (see [below for nested schema](#nestedatt--spec--control_plane_configuration--upgrade_rollout_strategy--rolling_update))
- `type` (String) UpgradeRolloutStrategyType defines the types of upgrade rollout strategies.

<a id="nestedatt--spec--control_plane_configuration--upgrade_rollout_strategy--rolling_update"></a>
### Nested Schema for `spec.control_plane_configuration.upgrade_rollout_strategy.rolling_update`

Required:

- `max_surge` (Number)




<a id="nestedatt--spec--datacenter_ref"></a>
### Nested Schema for `spec.datacenter_ref`

Optional:

- `kind` (String)
- `name` (String)


<a id="nestedatt--spec--etcd_encryption"></a>
### Nested Schema for `spec.etcd_encryption`

Required:

- `providers` (Attributes List) (see [below for nested schema](#nestedatt--spec--etcd_encryption--providers))
- `resources` (List of String) Resources defines a list of objects and custom resources definitions that should be encrypted.

<a id="nestedatt--spec--etcd_encryption--providers"></a>
### Nested Schema for `spec.etcd_encryption.providers`

Required:

- `kms` (Attributes) KMS defines the configuration for KMS Encryption provider. (see [below for nested schema](#nestedatt--spec--etcd_encryption--providers--kms))

<a id="nestedatt--spec--etcd_encryption--providers--kms"></a>
### Nested Schema for `spec.etcd_encryption.providers.kms`

Required:

- `name` (String) Name defines the name of KMS plugin to be used.
- `socket_listen_address` (String) SocketListenAddress defines a UNIX socket address that the KMS provider listens on.

Optional:

- `cachesize` (Number) CacheSize defines the maximum number of encrypted objects to be cached in memory. The default value is 1000. You can set this to a negative value to disable caching.
- `timeout` (String) Timeout for kube-apiserver to wait for KMS plugin. Default is 3s.




<a id="nestedatt--spec--external_etcd_configuration"></a>
### Nested Schema for `spec.external_etcd_configuration`

Optional:

- `count` (Number)
- `machine_group_ref` (Attributes) MachineGroupRef defines the machine group configuration for the etcd machines. (see [below for nested schema](#nestedatt--spec--external_etcd_configuration--machine_group_ref))

<a id="nestedatt--spec--external_etcd_configuration--machine_group_ref"></a>
### Nested Schema for `spec.external_etcd_configuration.machine_group_ref`

Optional:

- `kind` (String)
- `name` (String)



<a id="nestedatt--spec--git_ops_ref"></a>
### Nested Schema for `spec.git_ops_ref`

Optional:

- `kind` (String)
- `name` (String)


<a id="nestedatt--spec--identity_provider_refs"></a>
### Nested Schema for `spec.identity_provider_refs`

Optional:

- `kind` (String)
- `name` (String)


<a id="nestedatt--spec--machine_health_check"></a>
### Nested Schema for `spec.machine_health_check`

Optional:

- `max_unhealthy` (String) MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.
- `node_startup_timeout` (String) NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.
- `unhealthy_machine_timeout` (String) UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).


<a id="nestedatt--spec--management_cluster"></a>
### Nested Schema for `spec.management_cluster`

Optional:

- `name` (String)


<a id="nestedatt--spec--packages"></a>
### Nested Schema for `spec.packages`

Optional:

- `controller` (Attributes) Controller package controller configuration (see [below for nested schema](#nestedatt--spec--packages--controller))
- `cronjob` (Attributes) Cronjob for ecr token refresher (see [below for nested schema](#nestedatt--spec--packages--cronjob))
- `disable` (Boolean) Disable package controller on cluster

<a id="nestedatt--spec--packages--controller"></a>
### Nested Schema for `spec.packages.controller`

Optional:

- `digest` (String) Digest package controller digest
- `disable_webhooks` (Boolean) DisableWebhooks on package controller
- `env` (List of String) Env of package controller in the format 'key=value'
- `repository` (String) Repository package controller repository
- `resources` (Attributes) Resources of package controller (see [below for nested schema](#nestedatt--spec--packages--controller--resources))
- `tag` (String) Tag package controller tag

<a id="nestedatt--spec--packages--controller--resources"></a>
### Nested Schema for `spec.packages.controller.resources`

Optional:

- `limits` (Attributes) ImageResource resources for container image. (see [below for nested schema](#nestedatt--spec--packages--controller--tag--limits))
- `requests` (Attributes) Requests for image resources (see [below for nested schema](#nestedatt--spec--packages--controller--tag--requests))

<a id="nestedatt--spec--packages--controller--tag--limits"></a>
### Nested Schema for `spec.packages.controller.tag.limits`

Optional:

- `cpu` (String) CPU image cpu
- `memory` (String) Memory image memory


<a id="nestedatt--spec--packages--controller--tag--requests"></a>
### Nested Schema for `spec.packages.controller.tag.requests`

Optional:

- `cpu` (String) CPU image cpu
- `memory` (String) Memory image memory




<a id="nestedatt--spec--packages--cronjob"></a>
### Nested Schema for `spec.packages.cronjob`

Optional:

- `digest` (String) Digest ecr token refresher digest
- `disable` (Boolean) Disable on cron job
- `repository` (String) Repository ecr token refresher repository
- `tag` (String) Tag ecr token refresher tag



<a id="nestedatt--spec--pod_iam_config"></a>
### Nested Schema for `spec.pod_iam_config`

Required:

- `service_account_issuer` (String)


<a id="nestedatt--spec--proxy_configuration"></a>
### Nested Schema for `spec.proxy_configuration`

Optional:

- `http_proxy` (String)
- `https_proxy` (String)
- `no_proxy` (List of String)


<a id="nestedatt--spec--registry_mirror_configuration"></a>
### Nested Schema for `spec.registry_mirror_configuration`

Optional:

- `authenticate` (Boolean) Authenticate defines if registry requires authentication
- `ca_cert_content` (String) CACertContent defines the contents registry mirror CA certificate
- `endpoint` (String) Endpoint defines the registry mirror endpoint to use for pulling images
- `insecure_skip_verify` (Boolean) InsecureSkipVerify skips the registry certificate verification. Only use this solution for isolated testing or in a tightly controlled, air-gapped environment.
- `oci_namespaces` (Attributes List) OCINamespaces defines the mapping from an upstream registry to a local namespace where upstream artifacts are placed into (see [below for nested schema](#nestedatt--spec--registry_mirror_configuration--oci_namespaces))
- `port` (String) Port defines the port exposed for registry mirror endpoint

<a id="nestedatt--spec--registry_mirror_configuration--oci_namespaces"></a>
### Nested Schema for `spec.registry_mirror_configuration.oci_namespaces`

Required:

- `namespace` (String) Namespace refers to the name of a namespace in the local registry
- `registry` (String) Name refers to the name of the upstream registry



<a id="nestedatt--spec--worker_node_group_configurations"></a>
### Nested Schema for `spec.worker_node_group_configurations`

Optional:

- `autoscaling_configuration` (Attributes) AutoScalingConfiguration defines the auto scaling configuration (see [below for nested schema](#nestedatt--spec--worker_node_group_configurations--autoscaling_configuration))
- `count` (Number) Count defines the number of desired worker nodes. Defaults to 1.
- `kubernetes_version` (String) KuberenetesVersion defines the version for worker nodes. If not set, the top level spec kubernetesVersion will be used.
- `labels` (Map of String) Labels define the labels to assign to the node
- `machine_group_ref` (Attributes) MachineGroupRef defines the machine group configuration for the worker nodes. (see [below for nested schema](#nestedatt--spec--worker_node_group_configurations--machine_group_ref))
- `machine_health_check` (Attributes) MachineHealthCheck is a worker node level override for the timeouts and maxUnhealthy specified in the top-level MHC configuration. If not configured, the defaults in the top-level MHC configuration are used. (see [below for nested schema](#nestedatt--spec--worker_node_group_configurations--machine_health_check))
- `name` (String) Name refers to the name of the worker node group
- `taints` (Attributes List) Taints define the set of taints to be applied on worker nodes (see [below for nested schema](#nestedatt--spec--worker_node_group_configurations--taints))
- `upgrade_rollout_strategy` (Attributes) UpgradeRolloutStrategy determines the rollout strategy to use for rolling upgrades and related parameters/knobs (see [below for nested schema](#nestedatt--spec--worker_node_group_configurations--upgrade_rollout_strategy))

<a id="nestedatt--spec--worker_node_group_configurations--autoscaling_configuration"></a>
### Nested Schema for `spec.worker_node_group_configurations.autoscaling_configuration`

Optional:

- `max_count` (Number) MaxCount defines the maximum number of nodes for the associated resource group.
- `min_count` (Number) MinCount defines the minimum number of nodes for the associated resource group.


<a id="nestedatt--spec--worker_node_group_configurations--machine_group_ref"></a>
### Nested Schema for `spec.worker_node_group_configurations.machine_group_ref`

Optional:

- `kind` (String)
- `name` (String)


<a id="nestedatt--spec--worker_node_group_configurations--machine_health_check"></a>
### Nested Schema for `spec.worker_node_group_configurations.machine_health_check`

Optional:

- `max_unhealthy` (String) MaxUnhealthy is used to configure the maximum number of unhealthy machines in machine health checks. This setting applies to both control plane and worker machines. If the number of unhealthy machines exceeds the limit set by maxUnhealthy, further remediation will not be performed. If not configured, the default value is set to '100%' for controlplane machines and '40%' for worker machines.
- `node_startup_timeout` (String) NodeStartupTimeout is used to configure the node startup timeout in machine health checks. It determines how long a MachineHealthCheck should wait for a Node to join the cluster, before considering a Machine unhealthy. If not configured, the default value is set to '10m0s' (10 minutes) for all providers. For Tinkerbell provider the default is '20m0s'.
- `unhealthy_machine_timeout` (String) UnhealthyMachineTimeout is used to configure the unhealthy machine timeout in machine health checks. If any unhealthy conditions are met for the amount of time specified as the timeout, the machines are considered unhealthy. If not configured, the default value is set to '5m0s' (5 minutes).


<a id="nestedatt--spec--worker_node_group_configurations--taints"></a>
### Nested Schema for `spec.worker_node_group_configurations.taints`

Required:

- `effect` (String) Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
- `key` (String) Required. The taint key to be applied to a node.

Optional:

- `time_added` (String) TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.
- `value` (String) The taint value corresponding to the taint key.


<a id="nestedatt--spec--worker_node_group_configurations--upgrade_rollout_strategy"></a>
### Nested Schema for `spec.worker_node_group_configurations.upgrade_rollout_strategy`

Optional:

- `rolling_update` (Attributes) WorkerNodesRollingUpdateParams is API for rolling update strategy knobs. (see [below for nested schema](#nestedatt--spec--worker_node_group_configurations--upgrade_rollout_strategy--rolling_update))
- `type` (String) UpgradeRolloutStrategyType defines the types of upgrade rollout strategies.

<a id="nestedatt--spec--worker_node_group_configurations--upgrade_rollout_strategy--rolling_update"></a>
### Nested Schema for `spec.worker_node_group_configurations.upgrade_rollout_strategy.rolling_update`

Required:

- `max_surge` (Number)
- `max_unavailable` (Number)