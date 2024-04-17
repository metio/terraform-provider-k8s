---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kustomize_toolkit_fluxcd_io_kustomization_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "kustomize.toolkit.fluxcd.io"
description: |-
  Kustomization is the Schema for the kustomizations API.
---

# k8s_kustomize_toolkit_fluxcd_io_kustomization_v1_manifest (Data Source)

Kustomization is the Schema for the kustomizations API.

## Example Usage

```terraform
data "k8s_kustomize_toolkit_fluxcd_io_kustomization_v1_manifest" "example" {
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

- `spec` (Attributes) KustomizationSpec defines the configuration to calculate the desired statefrom a Source using Kustomize. (see [below for nested schema](#nestedatt--spec))

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

- `interval` (String) The interval at which to reconcile the Kustomization.This interval is approximate and may be subject to jitter to ensureefficient use of resources.
- `prune` (Boolean) Prune enables garbage collection.
- `source_ref` (Attributes) Reference of the source where the kustomization file is. (see [below for nested schema](#nestedatt--spec--source_ref))

Optional:

- `common_metadata` (Attributes) CommonMetadata specifies the common labels and annotations that areapplied to all resources. Any existing label or annotation will beoverridden if its key matches a common one. (see [below for nested schema](#nestedatt--spec--common_metadata))
- `components` (List of String) Components specifies relative paths to specifications of other Components.
- `decryption` (Attributes) Decrypt Kubernetes secrets before applying them on the cluster. (see [below for nested schema](#nestedatt--spec--decryption))
- `depends_on` (Attributes List) DependsOn may contain a meta.NamespacedObjectReference slicewith references to Kustomization resources that must be ready before thisKustomization can be reconciled. (see [below for nested schema](#nestedatt--spec--depends_on))
- `force` (Boolean) Force instructs the controller to recreate resourceswhen patching fails due to an immutable field change.
- `health_checks` (Attributes List) A list of resources to be included in the health assessment. (see [below for nested schema](#nestedatt--spec--health_checks))
- `images` (Attributes List) Images is a list of (image name, new name, new tag or digest)for changing image names, tags or digests. This can also be achieved with apatch, but this operator is simpler to specify. (see [below for nested schema](#nestedatt--spec--images))
- `kube_config` (Attributes) The KubeConfig for reconciling the Kustomization on a remote cluster.When used in combination with KustomizationSpec.ServiceAccountName,forces the controller to act on behalf of that Service Account at thetarget cluster.If the --default-service-account flag is set, its value will be used asa controller level fallback for when KustomizationSpec.ServiceAccountNameis empty. (see [below for nested schema](#nestedatt--spec--kube_config))
- `patches` (Attributes List) Strategic merge and JSON patches, defined as inline YAML objects,capable of targeting objects based on kind, label and annotation selectors. (see [below for nested schema](#nestedatt--spec--patches))
- `path` (String) Path to the directory containing the kustomization.yaml file, or theset of plain YAMLs a kustomization.yaml should be generated for.Defaults to 'None', which translates to the root path of the SourceRef.
- `post_build` (Attributes) PostBuild describes which actions to perform on the YAML manifestgenerated by building the kustomize overlay. (see [below for nested schema](#nestedatt--spec--post_build))
- `retry_interval` (String) The interval at which to retry a previously failed reconciliation.When not specified, the controller uses the KustomizationSpec.Intervalvalue to retry failures.
- `service_account_name` (String) The name of the Kubernetes service account to impersonatewhen reconciling this Kustomization.
- `suspend` (Boolean) This flag tells the controller to suspend subsequent kustomize executions,it does not apply to already started executions. Defaults to false.
- `target_namespace` (String) TargetNamespace sets or overrides the namespace in thekustomization.yaml file.
- `timeout` (String) Timeout for validation, apply and health checking operations.Defaults to 'Interval' duration.
- `wait` (Boolean) Wait instructs the controller to check the health of all the reconciledresources. When enabled, the HealthChecks are ignored. Defaults to false.

<a id="nestedatt--spec--source_ref"></a>
### Nested Schema for `spec.source_ref`

Required:

- `kind` (String) Kind of the referent.
- `name` (String) Name of the referent.

Optional:

- `api_version` (String) API version of the referent.
- `namespace` (String) Namespace of the referent, defaults to the namespace of the Kubernetesresource object that contains the reference.


<a id="nestedatt--spec--common_metadata"></a>
### Nested Schema for `spec.common_metadata`

Optional:

- `annotations` (Map of String) Annotations to be added to the object's metadata.
- `labels` (Map of String) Labels to be added to the object's metadata.


<a id="nestedatt--spec--decryption"></a>
### Nested Schema for `spec.decryption`

Required:

- `provider` (String) Provider is the name of the decryption engine.

Optional:

- `secret_ref` (Attributes) The secret name containing the private OpenPGP keys used for decryption. (see [below for nested schema](#nestedatt--spec--decryption--secret_ref))

<a id="nestedatt--spec--decryption--secret_ref"></a>
### Nested Schema for `spec.decryption.secret_ref`

Required:

- `name` (String) Name of the referent.



<a id="nestedatt--spec--depends_on"></a>
### Nested Schema for `spec.depends_on`

Required:

- `name` (String) Name of the referent.

Optional:

- `namespace` (String) Namespace of the referent, when not specified it acts as LocalObjectReference.


<a id="nestedatt--spec--health_checks"></a>
### Nested Schema for `spec.health_checks`

Required:

- `kind` (String) Kind of the referent.
- `name` (String) Name of the referent.

Optional:

- `api_version` (String) API version of the referent, if not specified the Kubernetes preferred version will be used.
- `namespace` (String) Namespace of the referent, when not specified it acts as LocalObjectReference.


<a id="nestedatt--spec--images"></a>
### Nested Schema for `spec.images`

Required:

- `name` (String) Name is a tag-less image name.

Optional:

- `digest` (String) Digest is the value used to replace the original image tag.If digest is present NewTag value is ignored.
- `new_name` (String) NewName is the value used to replace the original name.
- `new_tag` (String) NewTag is the value used to replace the original tag.


<a id="nestedatt--spec--kube_config"></a>
### Nested Schema for `spec.kube_config`

Required:

- `secret_ref` (Attributes) SecretRef holds the name of a secret that contains a key withthe kubeconfig file as the value. If no key is set, the key will defaultto 'value'.It is recommended that the kubeconfig is self-contained, and the secretis regularly updated if credentials such as a cloud-access-token expire.Cloud specific 'cmd-path' auth helpers will not function without addingbinaries and credentials to the Pod that is responsible for reconcilingKubernetes resources. (see [below for nested schema](#nestedatt--spec--kube_config--secret_ref))

<a id="nestedatt--spec--kube_config--secret_ref"></a>
### Nested Schema for `spec.kube_config.secret_ref`

Required:

- `name` (String) Name of the Secret.

Optional:

- `key` (String) Key in the Secret, when not specified an implementation-specific default key is used.



<a id="nestedatt--spec--patches"></a>
### Nested Schema for `spec.patches`

Required:

- `patch` (String) Patch contains an inline StrategicMerge patch or an inline JSON6902 patch withan array of operation objects.

Optional:

- `target` (Attributes) Target points to the resources that the patch document should be applied to. (see [below for nested schema](#nestedatt--spec--patches--target))

<a id="nestedatt--spec--patches--target"></a>
### Nested Schema for `spec.patches.target`

Optional:

- `annotation_selector` (String) AnnotationSelector is a string that follows the label selection expressionhttps://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#apiIt matches with the resource annotations.
- `group` (String) Group is the API group to select resources from.Together with Version and Kind it is capable of unambiguously identifying and/or selecting resources.https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md
- `kind` (String) Kind of the API Group to select resources from.Together with Group and Version it is capable of unambiguouslyidentifying and/or selecting resources.https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md
- `label_selector` (String) LabelSelector is a string that follows the label selection expressionhttps://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#apiIt matches with the resource labels.
- `name` (String) Name to match resources with.
- `namespace` (String) Namespace to select resources from.
- `version` (String) Version of the API Group to select resources from.Together with Group and Kind it is capable of unambiguously identifying and/or selecting resources.https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md



<a id="nestedatt--spec--post_build"></a>
### Nested Schema for `spec.post_build`

Optional:

- `substitute` (Map of String) Substitute holds a map of key/value pairs.The variables defined in your YAML manifests that match any of the keysdefined in the map will be substituted with the set value.Includes support for bash string replacement functionse.g. ${var:=default}, ${var:position} and ${var/substring/replacement}.
- `substitute_from` (Attributes List) SubstituteFrom holds references to ConfigMaps and Secrets containingthe variables and their values to be substituted in the YAML manifests.The ConfigMap and the Secret data keys represent the var names, and theymust match the vars declared in the manifests for the substitution tohappen. (see [below for nested schema](#nestedatt--spec--post_build--substitute_from))

<a id="nestedatt--spec--post_build--substitute_from"></a>
### Nested Schema for `spec.post_build.substitute_from`

Required:

- `kind` (String) Kind of the values referent, valid values are ('Secret', 'ConfigMap').
- `name` (String) Name of the values referent. Should reside in the same namespace as thereferring resource.

Optional:

- `optional` (Boolean) Optional indicates whether the referenced resource must exist, or whether totolerate its absence. If true and the referenced resource is absent, proceedas if the resource was present but empty, without any variables defined.