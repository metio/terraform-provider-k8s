---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kafka_stackable_tech_kafka_cluster_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "kafka.stackable.tech"
description: |-
  Auto-generated derived type for KafkaClusterSpec via 'CustomResource'
---

# k8s_kafka_stackable_tech_kafka_cluster_v1alpha1_manifest (Data Source)

Auto-generated derived type for KafkaClusterSpec via 'CustomResource'

## Example Usage

```terraform
data "k8s_kafka_stackable_tech_kafka_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    cluster_config = {
      zookeeper_config_map_name = "some-name"
    }
    image = {}
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))
- `spec` (Attributes) A Kafka cluster stacklet. This resource is managed by the Stackable operator for Apache Kafka. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/kafka/). (see [below for nested schema](#nestedatt--spec))

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

- `cluster_config` (Attributes) Kafka settings that affect all roles and role groups. The settings in the 'clusterConfig' are cluster wide settings that do not need to be configurable at role or role group level. (see [below for nested schema](#nestedatt--spec--cluster_config))
- `image` (Attributes) Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details. (see [below for nested schema](#nestedatt--spec--image))

Optional:

- `brokers` (Attributes) This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups). (see [below for nested schema](#nestedatt--spec--brokers))
- `cluster_operation` (Attributes) [Cluster operations](https://docs.stackable.tech/home/nightly/concepts/operations/cluster_operations) properties, allow stopping the product instance as well as pausing reconciliation. (see [below for nested schema](#nestedatt--spec--cluster_operation))

<a id="nestedatt--spec--cluster_config"></a>
### Nested Schema for `spec.cluster_config`

Required:

- `zookeeper_config_map_name` (String) Kafka requires a ZooKeeper cluster connection to run. Provide the name of the ZooKeeper [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) here. When using the [Stackable operator for Apache ZooKeeper](https://docs.stackable.tech/home/nightly/zookeeper/) to deploy a ZooKeeper cluster, this will simply be the name of your ZookeeperCluster resource.

Optional:

- `authentication` (Attributes List) Authentication class settings for Kafka like mTLS authentication. (see [below for nested schema](#nestedatt--spec--cluster_config--authentication))
- `authorization` (Attributes) Authorization settings for Kafka like OPA. (see [below for nested schema](#nestedatt--spec--cluster_config--authorization))
- `tls` (Attributes) TLS encryption settings for Kafka (server, internal). (see [below for nested schema](#nestedatt--spec--cluster_config--tls))
- `vector_aggregator_config_map_name` (String) Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.

<a id="nestedatt--spec--cluster_config--authentication"></a>
### Nested Schema for `spec.cluster_config.authentication`

Required:

- `authentication_class` (String) The AuthenticationClass <https://docs.stackable.tech/home/nightly/concepts/authenticationclass.html> to use. ## TLS provider Only affects client connections. This setting controls: - If clients need to authenticate themselves against the broker via TLS - Which ca.crt to use when validating the provided client certs This will override the server TLS settings (if set) in 'spec.clusterConfig.tls.serverSecretClass'.


<a id="nestedatt--spec--cluster_config--authorization"></a>
### Nested Schema for `spec.cluster_config.authorization`

Optional:

- `opa` (Attributes) Configure the OPA stacklet [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) and the name of the Rego package containing your authorization rules. Consult the [OPA authorization documentation](https://docs.stackable.tech/home/nightly/concepts/opa) to learn how to deploy Rego authorization rules with OPA. (see [below for nested schema](#nestedatt--spec--cluster_config--authorization--opa))

<a id="nestedatt--spec--cluster_config--authorization--opa"></a>
### Nested Schema for `spec.cluster_config.authorization.opa`

Required:

- `config_map_name` (String) The [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery) for the OPA stacklet that should be used for authorization requests.

Optional:

- `package` (String) The name of the Rego package containing the Rego rules for the product.



<a id="nestedatt--spec--cluster_config--tls"></a>
### Nested Schema for `spec.cluster_config.tls`

Optional:

- `internal_secret_class` (String) The [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass.html) to use for internal broker communication. Use mutual verification between brokers (mandatory). This setting controls: - Which cert the brokers should use to authenticate themselves against other brokers - Which ca.crt to use when validating the other brokers Defaults to 'tls'
- `server_secret_class` (String) The [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass.html) to use for client connections. This setting controls: - If TLS encryption is used at all - Which cert the servers should use to authenticate themselves against the client Defaults to 'tls'.



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



<a id="nestedatt--spec--brokers"></a>
### Nested Schema for `spec.brokers`

Required:

- `role_groups` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--role_groups))

Optional:

- `cli_overrides` (Map of String)
- `config` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--config))
- `config_overrides` (Map of Map of String) The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.
- `env_overrides` (Map of String) 'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.
- `pod_overrides` (Map of String) In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.
- `role_config` (Attributes) This is a product-agnostic RoleConfig, which is sufficient for most of the products. (see [below for nested schema](#nestedatt--spec--brokers--role_config))

<a id="nestedatt--spec--brokers--role_groups"></a>
### Nested Schema for `spec.brokers.role_groups`

Optional:

- `cli_overrides` (Map of String)
- `config` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config))
- `config_overrides` (Map of Map of String) The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.
- `env_overrides` (Map of String) 'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.
- `pod_overrides` (Map of String) In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.
- `replicas` (Number)

<a id="nestedatt--spec--brokers--role_groups--config"></a>
### Nested Schema for `spec.brokers.role_groups.config`

Optional:

- `affinity` (Attributes) These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement). (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--affinity))
- `bootstrap_listener_class` (String) The ListenerClass used for bootstrapping new clients. Should use a stable ListenerClass to avoid unnecessary client restarts (such as 'cluster-internal' or 'external-stable').
- `broker_listener_class` (String) The ListenerClass used for connecting to brokers. Should use a direct connection ListenerClass to minimize cost and minimize performance overhead (such as 'cluster-internal' or 'external-unstable').
- `graceful_shutdown_timeout` (String) Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.
- `logging` (Attributes) Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging). (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--logging))
- `resources` (Attributes) Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any. (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--resources))

<a id="nestedatt--spec--brokers--role_groups--config--affinity"></a>
### Nested Schema for `spec.brokers.role_groups.config.affinity`

Optional:

- `node_affinity` (Map of String) Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `node_selector` (Map of String) Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_affinity` (Map of String) Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_anti_affinity` (Map of String) Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)


<a id="nestedatt--spec--brokers--role_groups--config--logging"></a>
### Nested Schema for `spec.brokers.role_groups.config.logging`

Optional:

- `containers` (Attributes) Log configuration per container. (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--logging--containers))
- `enable_vector_agent` (Boolean) Wether or not to deploy a container with the Vector log agent.

<a id="nestedatt--spec--brokers--role_groups--config--logging--containers"></a>
### Nested Schema for `spec.brokers.role_groups.config.logging.containers`

Optional:

- `console` (Attributes) Configuration for the console appender (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--logging--containers--console))
- `custom` (Attributes) Custom log configuration provided in a ConfigMap (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--logging--containers--custom))
- `file` (Attributes) Configuration for the file appender (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--logging--containers--file))
- `loggers` (Attributes) Configuration per logger (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--logging--containers--loggers))

<a id="nestedatt--spec--brokers--role_groups--config--logging--containers--console"></a>
### Nested Schema for `spec.brokers.role_groups.config.logging.containers.console`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--brokers--role_groups--config--logging--containers--custom"></a>
### Nested Schema for `spec.brokers.role_groups.config.logging.containers.custom`

Optional:

- `config_map` (String) ConfigMap containing the log configuration files


<a id="nestedatt--spec--brokers--role_groups--config--logging--containers--file"></a>
### Nested Schema for `spec.brokers.role_groups.config.logging.containers.file`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--brokers--role_groups--config--logging--containers--loggers"></a>
### Nested Schema for `spec.brokers.role_groups.config.logging.containers.loggers`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.




<a id="nestedatt--spec--brokers--role_groups--config--resources"></a>
### Nested Schema for `spec.brokers.role_groups.config.resources`

Optional:

- `cpu` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--resources--cpu))
- `memory` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--resources--memory))
- `storage` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--resources--storage))

<a id="nestedatt--spec--brokers--role_groups--config--resources--cpu"></a>
### Nested Schema for `spec.brokers.role_groups.config.resources.cpu`

Optional:

- `max` (String) The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.
- `min` (String) The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.


<a id="nestedatt--spec--brokers--role_groups--config--resources--memory"></a>
### Nested Schema for `spec.brokers.role_groups.config.resources.memory`

Optional:

- `limit` (String) The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'
- `runtime_limits` (Map of String) Additional options that can be specified.


<a id="nestedatt--spec--brokers--role_groups--config--resources--storage"></a>
### Nested Schema for `spec.brokers.role_groups.config.resources.storage`

Optional:

- `log_dirs` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--resources--storage--log_dirs))

<a id="nestedatt--spec--brokers--role_groups--config--resources--storage--log_dirs"></a>
### Nested Schema for `spec.brokers.role_groups.config.resources.storage.log_dirs`

Optional:

- `capacity` (String) Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ''' <quantity> ::= <signedNumber><suffix> (Note that <suffix> may be empty, from the '' case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= '+' | '-' <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | '' | k | M | G | T | P | E (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> ''' No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi' Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
- `selectors` (Attributes) A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects. (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--resources--storage--log_dirs--selectors))
- `storage_class` (String)

<a id="nestedatt--spec--brokers--role_groups--config--resources--storage--log_dirs--selectors"></a>
### Nested Schema for `spec.brokers.role_groups.config.resources.storage.log_dirs.selectors`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--brokers--role_groups--config--resources--storage--log_dirs--selectors--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--brokers--role_groups--config--resources--storage--log_dirs--selectors--match_expressions"></a>
### Nested Schema for `spec.brokers.role_groups.config.resources.storage.log_dirs.selectors.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.








<a id="nestedatt--spec--brokers--config"></a>
### Nested Schema for `spec.brokers.config`

Optional:

- `affinity` (Attributes) These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement). (see [below for nested schema](#nestedatt--spec--brokers--config--affinity))
- `bootstrap_listener_class` (String) The ListenerClass used for bootstrapping new clients. Should use a stable ListenerClass to avoid unnecessary client restarts (such as 'cluster-internal' or 'external-stable').
- `broker_listener_class` (String) The ListenerClass used for connecting to brokers. Should use a direct connection ListenerClass to minimize cost and minimize performance overhead (such as 'cluster-internal' or 'external-unstable').
- `graceful_shutdown_timeout` (String) Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.
- `logging` (Attributes) Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging). (see [below for nested schema](#nestedatt--spec--brokers--config--logging))
- `resources` (Attributes) Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any. (see [below for nested schema](#nestedatt--spec--brokers--config--resources))

<a id="nestedatt--spec--brokers--config--affinity"></a>
### Nested Schema for `spec.brokers.config.affinity`

Optional:

- `node_affinity` (Map of String) Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `node_selector` (Map of String) Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_affinity` (Map of String) Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)
- `pod_anti_affinity` (Map of String) Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)


<a id="nestedatt--spec--brokers--config--logging"></a>
### Nested Schema for `spec.brokers.config.logging`

Optional:

- `containers` (Attributes) Log configuration per container. (see [below for nested schema](#nestedatt--spec--brokers--config--logging--containers))
- `enable_vector_agent` (Boolean) Wether or not to deploy a container with the Vector log agent.

<a id="nestedatt--spec--brokers--config--logging--containers"></a>
### Nested Schema for `spec.brokers.config.logging.containers`

Optional:

- `console` (Attributes) Configuration for the console appender (see [below for nested schema](#nestedatt--spec--brokers--config--logging--containers--console))
- `custom` (Attributes) Custom log configuration provided in a ConfigMap (see [below for nested schema](#nestedatt--spec--brokers--config--logging--containers--custom))
- `file` (Attributes) Configuration for the file appender (see [below for nested schema](#nestedatt--spec--brokers--config--logging--containers--file))
- `loggers` (Attributes) Configuration per logger (see [below for nested schema](#nestedatt--spec--brokers--config--logging--containers--loggers))

<a id="nestedatt--spec--brokers--config--logging--containers--console"></a>
### Nested Schema for `spec.brokers.config.logging.containers.console`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--brokers--config--logging--containers--custom"></a>
### Nested Schema for `spec.brokers.config.logging.containers.custom`

Optional:

- `config_map` (String) ConfigMap containing the log configuration files


<a id="nestedatt--spec--brokers--config--logging--containers--file"></a>
### Nested Schema for `spec.brokers.config.logging.containers.file`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.


<a id="nestedatt--spec--brokers--config--logging--containers--loggers"></a>
### Nested Schema for `spec.brokers.config.logging.containers.loggers`

Optional:

- `level` (String) The log level threshold. Log events with a lower log level are discarded.




<a id="nestedatt--spec--brokers--config--resources"></a>
### Nested Schema for `spec.brokers.config.resources`

Optional:

- `cpu` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--config--resources--cpu))
- `memory` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--config--resources--memory))
- `storage` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--config--resources--storage))

<a id="nestedatt--spec--brokers--config--resources--cpu"></a>
### Nested Schema for `spec.brokers.config.resources.cpu`

Optional:

- `max` (String) The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.
- `min` (String) The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.


<a id="nestedatt--spec--brokers--config--resources--memory"></a>
### Nested Schema for `spec.brokers.config.resources.memory`

Optional:

- `limit` (String) The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'
- `runtime_limits` (Map of String) Additional options that can be specified.


<a id="nestedatt--spec--brokers--config--resources--storage"></a>
### Nested Schema for `spec.brokers.config.resources.storage`

Optional:

- `log_dirs` (Attributes) (see [below for nested schema](#nestedatt--spec--brokers--config--resources--storage--log_dirs))

<a id="nestedatt--spec--brokers--config--resources--storage--log_dirs"></a>
### Nested Schema for `spec.brokers.config.resources.storage.log_dirs`

Optional:

- `capacity` (String) Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ''' <quantity> ::= <signedNumber><suffix> (Note that <suffix> may be empty, from the '' case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= '+' | '-' <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | '' | k | M | G | T | P | E (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> ''' No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi' Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.
- `selectors` (Attributes) A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects. (see [below for nested schema](#nestedatt--spec--brokers--config--resources--storage--log_dirs--selectors))
- `storage_class` (String)

<a id="nestedatt--spec--brokers--config--resources--storage--log_dirs--selectors"></a>
### Nested Schema for `spec.brokers.config.resources.storage.log_dirs.selectors`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--brokers--config--resources--storage--log_dirs--selectors--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--brokers--config--resources--storage--log_dirs--selectors--match_expressions"></a>
### Nested Schema for `spec.brokers.config.resources.storage.log_dirs.selectors.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.







<a id="nestedatt--spec--brokers--role_config"></a>
### Nested Schema for `spec.brokers.role_config`

Optional:

- `pod_disruption_budget` (Attributes) This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions). (see [below for nested schema](#nestedatt--spec--brokers--role_config--pod_disruption_budget))

<a id="nestedatt--spec--brokers--role_config--pod_disruption_budget"></a>
### Nested Schema for `spec.brokers.role_config.pod_disruption_budget`

Optional:

- `enabled` (Boolean) Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.
- `max_unavailable` (Number) The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.




<a id="nestedatt--spec--cluster_operation"></a>
### Nested Schema for `spec.cluster_operation`

Optional:

- `reconciliation_paused` (Boolean) Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.
- `stopped` (Boolean) Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.
