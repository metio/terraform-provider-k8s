---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_superset_stackable_tech_superset_cluster_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "superset.stackable.tech"
description: |-
  Auto-generated derived type for SupersetClusterSpec via 'CustomResource'
---

# k8s_superset_stackable_tech_superset_cluster_v1alpha1_manifest (Data Source)

Auto-generated derived type for SupersetClusterSpec via 'CustomResource'

## Example Usage

```terraform
data "k8s_superset_stackable_tech_superset_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      credentials_secret = "some-secret"
    }
    image = {}
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) A Superset cluster stacklet. This resource is managed by the Stackable operator for Apache Superset. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/superset/). (see [below for nested schema](#nestedatt--spec))

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

- `cluster_config` (Attributes) Settings that affect all roles and role groups. The settings in the 'clusterConfig' are cluster wide settings that do not need to be configurable at role or role group level. (see [below for nested schema](#nestedatt--spec--cluster_config))
- `image` (Attributes) Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details. (see [below for nested schema](#nestedatt--spec--image))

Optional:

- `nodes` (Attributes) This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups). (see [below for nested schema](#nestedatt--spec--nodes))

<a id="nestedatt--spec--cluster_config"></a>
### Nested Schema for `spec.cluster_config`

Required:

- `credentials_secret` (String) The name of the Secret object containing the admin user credentials and database connection details. Read the [getting started guide first steps](https://docs.stackable.tech/home/nightly/superset/getting_started/first_steps) to find out more.

Optional:

- `authentication` (Attributes List) List of AuthenticationClasses used to authenticate users. (see [below for nested schema](#nestedatt--spec--cluster_config--authentication))
- `cluster_operation` (Attributes) Cluster operations like pause reconciliation or cluster stop. (see [below for nested schema](#nestedatt--spec--cluster_config--cluster_operation))
- `listener_class` (String) This field controls which type of Service the Operator creates for this SupersetCluster: * cluster-internal: Use a ClusterIP service * external-unstable: Use a NodePort service * external-stable: Use a LoadBalancer service This is a temporary solution with the goal to keep yaml manifests forward compatible. In the future, this setting will control which [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass.html) will be used to expose the service, and ListenerClass names will stay the same, allowing for a non-breaking change.
- `mapbox_secret` (String) The name of a Secret object. The Secret should contain a key 'connections.mapboxApiKey'. This is the API key required for map charts to work that use mapbox. The token should be in the JWT format.
- `vector_aggregator_config_map_name` (String) Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.

<a id="nestedatt--spec--cluster_config--authentication"></a>
### Nested Schema for `spec.cluster_config.authentication`

Required:

- `authentication_class` (String) Name of the [AuthenticationClass](https://docs.stackable.tech/home/nightly/concepts/authentication) used to authenticate users.

Optional:

- `oidc` (Attributes) This field contains OIDC-specific configuration. It is only required in case OIDC is used. (see [below for nested schema](#nestedatt--spec--cluster_config--authentication--oidc))
- `sync_roles_at` (String) If we should replace ALL the user's roles each login, or only on registration. Gets mapped to 'AUTH_ROLES_SYNC_AT_LOGIN'
- `user_registration` (Boolean) Allow users who are not already in the FAB DB. Gets mapped to 'AUTH_USER_REGISTRATION'
- `user_registration_role` (String) This role will be given in addition to any AUTH_ROLES_MAPPING. Gets mapped to 'AUTH_USER_REGISTRATION_ROLE'

<a id="nestedatt--spec--cluster_config--authentication--oidc"></a>
### Nested Schema for `spec.cluster_config.authentication.oidc`

Required:

- `client_credentials_secret` (String) A reference to the OIDC client credentials secret. The secret contains the client id and secret.

Optional:

- `extra_scopes` (List of String) An optional list of extra scopes which get merged with the scopes defined in the ['AuthenticationClass'].



<a id="nestedatt--spec--cluster_config--cluster_operation"></a>
### Nested Schema for `spec.cluster_config.cluster_operation`

Optional:

- `reconciliation_paused` (Boolean) Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.
- `stopped` (Boolean) Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.



<a id="nestedatt--spec--image"></a>
### Nested Schema for `spec.image`

Optional:

- `custom` (String) Overwrite the docker image. Specify the full docker image name, e.g. 'docker.stackable.tech/stackable/superset:1.4.1-stackable2.1.0'
- `product_version` (String) Version of the product, e.g. '1.4.1'.
- `pull_policy` (String) [Pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) used when pulling the image.
- `pull_secrets` (Attributes List) [Image pull secrets](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod) to pull images from a private registry. (see [below for nested schema](#nestedatt--spec--image--pull_secrets))
- `repo` (String) Name of the docker repo, e.g. 'docker.stackable.tech/stackable'
- `stackable_version` (String) Stackable version of the product, e.g. '23.4', '23.4.1' or '0.0.0-dev'. If not specified, the operator will use its own version, e.g. '23.4.1'. When using a nightly operator or a pr version, it will use the nightly '0.0.0-dev' image.

<a id="nestedatt--spec--image--pull_secrets"></a>
### Nested Schema for `spec.image.pull_secrets`

Required:

- `name` (String) Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names



<a id="nestedatt--spec--nodes"></a>
### Nested Schema for `spec.nodes`

Required:

- `role_groups` (Attributes) (see [below for nested schema](#nestedatt--spec--nodes--role_groups))

Optional:

- `cli_overrides` (Map of String)
- `config` (Attributes) (see [below for nested schema](#nestedatt--spec--nodes--config))
- `config_overrides` (Map of Map of String) The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.
- `env_overrides` (Map of String) 'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.
- `pod_overrides` (Map of String) In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.
- `role_config` (Attributes) This is a product-agnostic RoleConfig, which is sufficient for most of the products. (see [below for nested schema](#nestedatt--spec--nodes--role_config))

<a id="nestedatt--spec--nodes--role_groups"></a>
### Nested Schema for `spec.nodes.role_groups`

Optional:

- `cli_overrides` (Map of String)
- `config` (Attributes) (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config))
- `config_overrides` (Map of Map of String) The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.
- `env_overrides` (Map of String) 'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.
- `pod_overrides` (Map of String) In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.
- `replicas` (Number)

<a id="nestedatt--spec--nodes--role_groups--config"></a>
### Nested Schema for `spec.nodes.role_groups.config`

Optional:

- `affinity` (Attributes) These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement). (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--affinity))
- `graceful_shutdown_timeout` (String) Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.
- `logging` (Attributes) Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging). (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--logging))
- `resources` (Attributes) CPU and memory limits for Superset pods (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--resources))
- `row_limit` (Number) Row limit when requesting chart data. Corresponds to ROW_LIMIT.
- `webserver_timeout` (Number) Maximum time period a Superset request can take before timing out. This setting affects the maximum duration a query to an underlying datasource can take. If you get timeout errors before your query returns the result you may need to increase this timeout. Corresponds to SUPERSET_WEBSERVER_TIMEOUT.

<a id="nestedatt--spec--nodes--role_groups--config--affinity"></a>
### Nested Schema for `spec.nodes.role_groups.config.affinity`

Optional:

- `node_affinity` (Map of String) Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `node_selector` (Map of String) Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_affinity` (Map of String) Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_anti_affinity` (Map of String) Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)


<a id="nestedatt--spec--nodes--role_groups--config--logging"></a>
### Nested Schema for `spec.nodes.role_groups.config.logging`

Optional:

- `containers` (Attributes) Log configuration per container. (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--logging--containers))
- `enable_vector_agent` (Boolean) Wether or not to deploy a container with the Vector log agent.

<a id="nestedatt--spec--nodes--role_groups--config--logging--containers"></a>
### Nested Schema for `spec.nodes.role_groups.config.logging.containers`

Optional:

- `console` (Attributes) Configuration for the console appender (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--logging--containers--console))
- `custom` (Attributes) Custom log configuration provided in a ConfigMap (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--logging--containers--custom))
- `file` (Attributes) Configuration for the file appender (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--logging--containers--file))
- `loggers` (Attributes) Configuration per logger (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--logging--containers--loggers))

<a id="nestedatt--spec--nodes--role_groups--config--logging--containers--console"></a>
### Nested Schema for `spec.nodes.role_groups.config.logging.containers.console`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--nodes--role_groups--config--logging--containers--custom"></a>
### Nested Schema for `spec.nodes.role_groups.config.logging.containers.custom`

Optional:

- `config_map` (String) ConfigMap containing the log configuration files


<a id="nestedatt--spec--nodes--role_groups--config--logging--containers--file"></a>
### Nested Schema for `spec.nodes.role_groups.config.logging.containers.file`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--nodes--role_groups--config--logging--containers--loggers"></a>
### Nested Schema for `spec.nodes.role_groups.config.logging.containers.loggers`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.




<a id="nestedatt--spec--nodes--role_groups--config--resources"></a>
### Nested Schema for `spec.nodes.role_groups.config.resources`

Optional:

- `cpu` (Attributes) (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--resources--cpu))
- `memory` (Attributes) (see [below for nested schema](#nestedatt--spec--nodes--role_groups--config--resources--memory))
- `storage` (Map of String)

<a id="nestedatt--spec--nodes--role_groups--config--resources--cpu"></a>
### Nested Schema for `spec.nodes.role_groups.config.resources.cpu`

Optional:

- `max` (String) The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.
- `min` (String) The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.


<a id="nestedatt--spec--nodes--role_groups--config--resources--memory"></a>
### Nested Schema for `spec.nodes.role_groups.config.resources.memory`

Optional:

- `limit` (String) The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'
- `runtime_limits` (Map of String) Additional options that can be specified.





<a id="nestedatt--spec--nodes--config"></a>
### Nested Schema for `spec.nodes.config`

Optional:

- `affinity` (Attributes) These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement). (see [below for nested schema](#nestedatt--spec--nodes--config--affinity))
- `graceful_shutdown_timeout` (String) Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.
- `logging` (Attributes) Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging). (see [below for nested schema](#nestedatt--spec--nodes--config--logging))
- `resources` (Attributes) CPU and memory limits for Superset pods (see [below for nested schema](#nestedatt--spec--nodes--config--resources))
- `row_limit` (Number) Row limit when requesting chart data. Corresponds to ROW_LIMIT.
- `webserver_timeout` (Number) Maximum time period a Superset request can take before timing out. This setting affects the maximum duration a query to an underlying datasource can take. If you get timeout errors before your query returns the result you may need to increase this timeout. Corresponds to SUPERSET_WEBSERVER_TIMEOUT.

<a id="nestedatt--spec--nodes--config--affinity"></a>
### Nested Schema for `spec.nodes.config.affinity`

Optional:

- `node_affinity` (Map of String) Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `node_selector` (Map of String) Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_affinity` (Map of String) Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_anti_affinity` (Map of String) Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)


<a id="nestedatt--spec--nodes--config--logging"></a>
### Nested Schema for `spec.nodes.config.logging`

Optional:

- `containers` (Attributes) Log configuration per container. (see [below for nested schema](#nestedatt--spec--nodes--config--logging--containers))
- `enable_vector_agent` (Boolean) Wether or not to deploy a container with the Vector log agent.

<a id="nestedatt--spec--nodes--config--logging--containers"></a>
### Nested Schema for `spec.nodes.config.logging.containers`

Optional:

- `console` (Attributes) Configuration for the console appender (see [below for nested schema](#nestedatt--spec--nodes--config--logging--containers--console))
- `custom` (Attributes) Custom log configuration provided in a ConfigMap (see [below for nested schema](#nestedatt--spec--nodes--config--logging--containers--custom))
- `file` (Attributes) Configuration for the file appender (see [below for nested schema](#nestedatt--spec--nodes--config--logging--containers--file))
- `loggers` (Attributes) Configuration per logger (see [below for nested schema](#nestedatt--spec--nodes--config--logging--containers--loggers))

<a id="nestedatt--spec--nodes--config--logging--containers--console"></a>
### Nested Schema for `spec.nodes.config.logging.containers.console`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--nodes--config--logging--containers--custom"></a>
### Nested Schema for `spec.nodes.config.logging.containers.custom`

Optional:

- `config_map` (String) ConfigMap containing the log configuration files


<a id="nestedatt--spec--nodes--config--logging--containers--file"></a>
### Nested Schema for `spec.nodes.config.logging.containers.file`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--nodes--config--logging--containers--loggers"></a>
### Nested Schema for `spec.nodes.config.logging.containers.loggers`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.




<a id="nestedatt--spec--nodes--config--resources"></a>
### Nested Schema for `spec.nodes.config.resources`

Optional:

- `cpu` (Attributes) (see [below for nested schema](#nestedatt--spec--nodes--config--resources--cpu))
- `memory` (Attributes) (see [below for nested schema](#nestedatt--spec--nodes--config--resources--memory))
- `storage` (Map of String)

<a id="nestedatt--spec--nodes--config--resources--cpu"></a>
### Nested Schema for `spec.nodes.config.resources.cpu`

Optional:

- `max` (String) The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.
- `min` (String) The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.


<a id="nestedatt--spec--nodes--config--resources--memory"></a>
### Nested Schema for `spec.nodes.config.resources.memory`

Optional:

- `limit` (String) The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'
- `runtime_limits` (Map of String) Additional options that can be specified.




<a id="nestedatt--spec--nodes--role_config"></a>
### Nested Schema for `spec.nodes.role_config`

Optional:

- `pod_disruption_budget` (Attributes) This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions). (see [below for nested schema](#nestedatt--spec--nodes--role_config--pod_disruption_budget))

<a id="nestedatt--spec--nodes--role_config--pod_disruption_budget"></a>
### Nested Schema for `spec.nodes.role_config.pod_disruption_budget`

Optional:

- `enabled` (Boolean) Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.
- `max_unavailable` (Number) The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.
