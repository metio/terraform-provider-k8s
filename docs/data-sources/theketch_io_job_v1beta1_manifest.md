---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_theketch_io_job_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "theketch.io"
description: |-
  Job is the Schema for the jobs API
---

# k8s_theketch_io_job_v1beta1_manifest (Data Source)

Job is the Schema for the jobs API

## Example Usage

```terraform
data "k8s_theketch_io_job_v1beta1_manifest" "example" {
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

- `spec` (Attributes) JobSpec defines the desired state of Job (see [below for nested schema](#nestedatt--spec))

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

- `name` (String)
- `namespace` (String)
- `type` (String)

Optional:

- `backoff_limit` (Number)
- `completions` (Number)
- `containers` (Attributes List) (see [below for nested schema](#nestedatt--spec--containers))
- `description` (String)
- `failed_jobs_history_limit` (Number)
- `parallelism` (Number)
- `policy` (Attributes) Policy represents the policy types a job can have (see [below for nested schema](#nestedatt--spec--policy))
- `schedule` (String) CronJob-specific
- `starting_deadline_seconds` (Number)
- `successful_jobs_history_limit` (Number)
- `suspend` (Boolean)
- `version` (String)

<a id="nestedatt--spec--containers"></a>
### Nested Schema for `spec.containers`

Required:

- `command` (List of String)
- `image` (String)
- `name` (String)


<a id="nestedatt--spec--policy"></a>
### Nested Schema for `spec.policy`

Optional:

- `concurrency_policy` (String) CronJob-specific
- `restart_policy` (String) RestartPolicy describes how the container should be restarted. Only one of the following restart policies may be specified. If none of the following policies is specified, the default one is RestartPolicyAlways.