---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_sharedresource_openshift_io_shared_secret_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "sharedresource.openshift.io"
description: |-
  SharedSecret allows a Secret to be shared across namespaces. Pods can mount the shared Secret by adding a CSI volume to the pod specification using the 'csi.sharedresource.openshift.io' CSI driver and a reference to the SharedSecret in the volume attributes:  spec: volumes: - name: shared-secret csi: driver: csi.sharedresource.openshift.io volumeAttributes: sharedSecret: my-share  For the mount to be successful, the pod's service account must be granted permission to 'use' the named SharedSecret object within its namespace with an appropriate Role and RoleBinding. For compactness, here are example 'oc' invocations for creating such Role and RoleBinding objects.  'oc create role shared-resource-my-share --verb=use --resource=sharedsecrets.sharedresource.openshift.io --resource-name=my-share' 'oc create rolebinding shared-resource-my-share --role=shared-resource-my-share --serviceaccount=my-namespace:default'  Shared resource objects, in this case Secrets, have default permissions of list, get, and watch for system authenticated users.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.
---

# k8s_sharedresource_openshift_io_shared_secret_v1alpha1_manifest (Data Source)

SharedSecret allows a Secret to be shared across namespaces. Pods can mount the shared Secret by adding a CSI volume to the pod specification using the 'csi.sharedresource.openshift.io' CSI driver and a reference to the SharedSecret in the volume attributes:  spec: volumes: - name: shared-secret csi: driver: csi.sharedresource.openshift.io volumeAttributes: sharedSecret: my-share  For the mount to be successful, the pod's service account must be granted permission to 'use' the named SharedSecret object within its namespace with an appropriate Role and RoleBinding. For compactness, here are example 'oc' invocations for creating such Role and RoleBinding objects.  'oc create role shared-resource-my-share --verb=use --resource=sharedsecrets.sharedresource.openshift.io --resource-name=my-share' 'oc create rolebinding shared-resource-my-share --role=shared-resource-my-share --serviceaccount=my-namespace:default'  Shared resource objects, in this case Secrets, have default permissions of list, get, and watch for system authenticated users.  Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support. These capabilities should not be used by applications needing long term support.

## Example Usage

```terraform
data "k8s_sharedresource_openshift_io_shared_secret_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) spec is the specification of the desired shared secret (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `secret_ref` (Attributes) secretRef is a reference to the Secret to share (see [below for nested schema](#nestedatt--spec--secret_ref))

Optional:

- `description` (String) description is a user readable explanation of what the backing resource provides.

<a id="nestedatt--spec--secret_ref"></a>
### Nested Schema for `spec.secret_ref`

Required:

- `name` (String) name represents the name of the Secret that is being referenced.
- `namespace` (String) namespace represents the namespace where the referenced Secret is located.