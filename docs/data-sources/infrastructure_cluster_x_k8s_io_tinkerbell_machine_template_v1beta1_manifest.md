---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "infrastructure.cluster.x-k8s.io"
description: |-
  TinkerbellMachineTemplate is the Schema for the tinkerbellmachinetemplates API.
---

# k8s_infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest (Data Source)

TinkerbellMachineTemplate is the Schema for the tinkerbellmachinetemplates API.

## Example Usage

```terraform
data "k8s_infrastructure_cluster_x_k8s_io_tinkerbell_machine_template_v1beta1_manifest" "example" {
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

- `spec` (Attributes) TinkerbellMachineTemplateSpec defines the desired state of TinkerbellMachineTemplate. (see [below for nested schema](#nestedatt--spec))

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

- `template` (Attributes) TinkerbellMachineTemplateResource describes the data needed to create am TinkerbellMachine from a template. (see [below for nested schema](#nestedatt--spec--template))

<a id="nestedatt--spec--template"></a>
### Nested Schema for `spec.template`

Required:

- `spec` (Attributes) Spec is the specification of the desired behavior of the machine. (see [below for nested schema](#nestedatt--spec--template--spec))

<a id="nestedatt--spec--template--spec"></a>
### Nested Schema for `spec.template.spec`

Optional:

- `hardware_affinity` (Attributes) HardwareAffinity allows filtering for hardware. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity))
- `hardware_name` (String) Those fields are set programmatically, but they cannot be re-constructed from 'state of the world', so we put them in spec instead of status.
- `image_lookup_base_registry` (String) ImageLookupBaseRegistry is the base Registry URL that is used for pulling images, if not set, the default will be to use ghcr.io/tinkerbell/cluster-api-provider-tinkerbell.
- `image_lookup_format` (String) ImageLookupFormat is the URL naming format to use for machine images when a machine does not specify. When set, this will be used for all cluster machines unless a machine specifies a different ImageLookupFormat. Supports substitutions for {{.BaseRegistry}}, {{.OSDistro}}, {{.OSVersion}} and {{.KubernetesVersion}} with the basse URL, OS distribution, OS version, and kubernetes version, respectively. BaseRegistry will be the value in ImageLookupBaseRegistry or ghcr.io/tinkerbell/cluster-api-provider-tinkerbell (the default), OSDistro will be the value in ImageLookupOSDistro or ubuntu (the default), OSVersion will be the value in ImageLookupOSVersion or default based on the OSDistro (if known), and the kubernetes version as defined by the packages produced by kubernetes/release: v1.13.0, v1.12.5-mybuild.1, or v1.17.3. For example, the default image format of {{.BaseRegistry}}/{{.OSDistro}}-{{.OSVersion}}:{{.KubernetesVersion}}.gz will attempt to pull the image from that location. See also: https://golang.org/pkg/text/template/
- `image_lookup_os_distro` (String) ImageLookupOSDistro is the name of the OS distro to use when fetching machine images, if not set it will default to ubuntu.
- `image_lookup_os_version` (String) ImageLookupOSVersion is the version of the OS distribution to use when fetching machine images. If not set it will default based on ImageLookupOSDistro.
- `provider_id` (String)
- `template_override` (String) TemplateOverride overrides the default Tinkerbell template used by CAPT. You can learn more about Tinkerbell templates here: https://tinkerbell.org/docs/concepts/templates/

<a id="nestedatt--spec--template--spec--hardware_affinity"></a>
### Nested Schema for `spec.template.spec.hardware_affinity`

Optional:

- `preferred` (Attributes List) Preferred are the preferred hardware affinity terms. Hardware matching these terms are preferred according to the weights provided, but are not required. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity--preferred))
- `required` (Attributes List) Required are the required hardware affinity terms. The terms are OR'd together, hardware must match one term to be considered. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity--required))

<a id="nestedatt--spec--template--spec--hardware_affinity--preferred"></a>
### Nested Schema for `spec.template.spec.hardware_affinity.preferred`

Required:

- `hardware_affinity_term` (Attributes) HardwareAffinityTerm is the term associated with the corresponding weight. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity--preferred--hardware_affinity_term))
- `weight` (Number) Weight associated with matching the corresponding hardwareAffinityTerm, in the range 1-100.

<a id="nestedatt--spec--template--spec--hardware_affinity--preferred--hardware_affinity_term"></a>
### Nested Schema for `spec.template.spec.hardware_affinity.preferred.hardware_affinity_term`

Required:

- `label_selector` (Attributes) LabelSelector is used to select for particular hardware by label. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity--preferred--hardware_affinity_term--label_selector))

<a id="nestedatt--spec--template--spec--hardware_affinity--preferred--hardware_affinity_term--label_selector"></a>
### Nested Schema for `spec.template.spec.hardware_affinity.preferred.hardware_affinity_term.label_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity--preferred--hardware_affinity_term--label_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--template--spec--hardware_affinity--preferred--hardware_affinity_term--label_selector--match_expressions"></a>
### Nested Schema for `spec.template.spec.hardware_affinity.preferred.hardware_affinity_term.label_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.





<a id="nestedatt--spec--template--spec--hardware_affinity--required"></a>
### Nested Schema for `spec.template.spec.hardware_affinity.required`

Required:

- `label_selector` (Attributes) LabelSelector is used to select for particular hardware by label. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity--required--label_selector))

<a id="nestedatt--spec--template--spec--hardware_affinity--required--label_selector"></a>
### Nested Schema for `spec.template.spec.hardware_affinity.required.label_selector`

Optional:

- `match_expressions` (Attributes List) matchExpressions is a list of label selector requirements. The requirements are ANDed. (see [below for nested schema](#nestedatt--spec--template--spec--hardware_affinity--required--label_selector--match_expressions))
- `match_labels` (Map of String) matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.

<a id="nestedatt--spec--template--spec--hardware_affinity--required--label_selector--match_expressions"></a>
### Nested Schema for `spec.template.spec.hardware_affinity.required.label_selector.match_expressions`

Required:

- `key` (String) key is the label key that the selector applies to.
- `operator` (String) operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.

Optional:

- `values` (List of String) values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
