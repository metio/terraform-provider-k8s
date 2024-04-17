---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "infrastructure.cluster.x-k8s.io"
description: |-
  KubevirtClusterTemplate is the Schema for the kubevirtclustertemplates API.
---

# k8s_infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest (Data Source)

KubevirtClusterTemplate is the Schema for the kubevirtclustertemplates API.

## Example Usage

```terraform
data "k8s_infrastructure_cluster_x_k8s_io_kubevirt_cluster_template_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) KubevirtClusterTemplateSpec defines the desired state of KubevirtClusterTemplate. (see [below for nested schema](#nestedatt--spec))

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

- `template` (Attributes) KubevirtClusterTemplateResource describes the data needed to create a KubevirtCluster from a template. (see [below for nested schema](#nestedatt--spec--template))

<a id="nestedatt--spec--template"></a>
### Nested Schema for `spec.template`

Required:

- `spec` (Attributes) KubevirtClusterSpec defines the desired state of KubevirtCluster. (see [below for nested schema](#nestedatt--spec--template--spec))

Optional:

- `metadata` (Attributes) ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create. This is a copy of customizable fields from metav1.ObjectMeta.  ObjectMeta is embedded in 'Machine.Spec', 'MachineDeployment.Template' and 'MachineSet.Template', which are not top-level Kubernetes objects. Given that metav1.ObjectMeta has lots of special cases and read-only fields which end up in the generated CRD validation, having it as a subset simplifies the API and some issues that can impact user experience.  During the [upgrade to controller-tools@v2](https://github.com/kubernetes-sigs/cluster-api/pull/1054) for v1alpha2, we noticed a failure would occur running Cluster API test suite against the new CRDs, specifically 'spec.metadata.creationTimestamp in body must be of type string: 'null''. The investigation showed that 'controller-tools@v2' behaves differently than its previous version when handling types from [metav1](k8s.io/apimachinery/pkg/apis/meta/v1) package.  In more details, we found that embedded (non-top level) types that embedded 'metav1.ObjectMeta' had validation properties, including for 'creationTimestamp' (metav1.Time). The 'metav1.Time' type specifies a custom json marshaller that, when IsZero() is true, returns 'null' which breaks validation because the field isn't marked as nullable.  In future versions, controller-tools@v2 might allow overriding the type and validation for embedded types. When that happens, this hack should be revisited. (see [below for nested schema](#nestedatt--spec--template--metadata))

<a id="nestedatt--spec--template--spec"></a>
### Nested Schema for `spec.template.spec`

Optional:

- `control_plane_endpoint` (Attributes) ControlPlaneEndpoint represents the endpoint used to communicate with the control plane. (see [below for nested schema](#nestedatt--spec--template--spec--control_plane_endpoint))
- `control_plane_service_template` (Attributes) ControlPlaneServiceTemplate can be used to modify service that fronts the control plane nodes to handle the api-server traffic (port 6443). This field is optional, by default control plane nodes will use a service of type ClusterIP, which will make workload cluster only accessible within the same cluster. Note, this does not aim to expose the entire Service spec to users, but only provides capability to modify the service metadata and the service type. (see [below for nested schema](#nestedatt--spec--template--spec--control_plane_service_template))
- `infra_cluster_secret_ref` (Attributes) InfraClusterSecretRef is a reference to a secret with a kubeconfig for external cluster used for infra. (see [below for nested schema](#nestedatt--spec--template--spec--infra_cluster_secret_ref))
- `ssh_keys` (Attributes) SSHKeys is a reference to a local struct for SSH keys persistence. (see [below for nested schema](#nestedatt--spec--template--spec--ssh_keys))

<a id="nestedatt--spec--template--spec--control_plane_endpoint"></a>
### Nested Schema for `spec.template.spec.control_plane_endpoint`

Required:

- `host` (String) Host is the hostname on which the API server is serving.
- `port` (Number) Port is the port on which the API server is serving.


<a id="nestedatt--spec--template--spec--control_plane_service_template"></a>
### Nested Schema for `spec.template.spec.control_plane_service_template`

Optional:

- `metadata` (Map of String) Service metadata allows to set labels, annotations and namespace for the service. When infraClusterSecretRef is used, ControlPlaneService take the kubeconfig namespace by default if metadata.namespace is not specified. This field is optional.
- `spec` (Attributes) Service specification allows to override some fields in the service spec. Note, it does not aim cover all fields of the service spec. (see [below for nested schema](#nestedatt--spec--template--spec--ssh_keys--spec))

<a id="nestedatt--spec--template--spec--ssh_keys--spec"></a>
### Nested Schema for `spec.template.spec.ssh_keys.spec`

Optional:

- `type` (String) Type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types



<a id="nestedatt--spec--template--spec--infra_cluster_secret_ref"></a>
### Nested Schema for `spec.template.spec.infra_cluster_secret_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids


<a id="nestedatt--spec--template--spec--ssh_keys"></a>
### Nested Schema for `spec.template.spec.ssh_keys`

Optional:

- `config_ref` (Attributes) ConfigRef is a reference to a resource containing the keys. The reference is optional to allow users/operators to specify Bootstrap.DataSecretName without the need of a controller. (see [below for nested schema](#nestedatt--spec--template--spec--ssh_keys--config_ref))
- `data_secret_name` (String) DataSecretName is the name of the secret that stores ssh keys.

<a id="nestedatt--spec--template--spec--ssh_keys--config_ref"></a>
### Nested Schema for `spec.template.spec.ssh_keys.config_ref`

Optional:

- `api_version` (String) API version of the referent.
- `field_path` (String) If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.
- `kind` (String) Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
- `namespace` (String) Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
- `resource_version` (String) Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
- `uid` (String) UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids




<a id="nestedatt--spec--template--metadata"></a>
### Nested Schema for `spec.template.metadata`

Optional:

- `annotations` (Map of String) Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
- `labels` (Map of String) Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels