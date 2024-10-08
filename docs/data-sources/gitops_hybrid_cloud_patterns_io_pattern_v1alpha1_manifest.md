---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "gitops.hybrid-cloud-patterns.io"
description: |-
  Pattern is the Schema for the patterns API
---

# k8s_gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest (Data Source)

Pattern is the Schema for the patterns API

## Example Usage

```terraform
data "k8s_gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) PatternSpec defines the desired state of Pattern (see [below for nested schema](#nestedatt--spec))

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

- `cluster_group_name` (String)
- `git_spec` (Attributes) (see [below for nested schema](#nestedatt--spec--git_spec))

Optional:

- `analytics_uuid` (String) Analytics UUID. Leave empty to autogenerate a random one. Not PII information
- `experimental_capabilities` (String) Comma separated capabilities to enable certain experimental features
- `extra_parameters` (Attributes List) .Name is dot separated per the helm --set syntax, such as: global.something.field (see [below for nested schema](#nestedatt--spec--extra_parameters))
- `extra_value_files` (List of String) URLs to additional Helm parameter files
- `git_ops_spec` (Attributes) (see [below for nested schema](#nestedatt--spec--git_ops_spec))
- `multi_source_config` (Attributes) (see [below for nested schema](#nestedatt--spec--multi_source_config))

<a id="nestedatt--spec--git_spec"></a>
### Nested Schema for `spec.git_spec`

Optional:

- `hostname` (String) Optional. FQDN of the git server if automatic parsing from TargetRepo is broken
- `in_cluster_git_server` (Boolean) (EXPERIMENTAL) Enable in-cluster git server (avoids the need of forking the upstream repository)
- `origin_repo` (String) Upstream git repo containing the pattern to deploy. Used when in-cluster fork to point to the upstream pattern repository. Takes precedence over TargetRepo
- `origin_revision` (String) (DEPRECATED) Branch, tag or commit in the upstream git repository. Does not support short-sha's. Default to HEAD
- `poll_interval` (Number) Interval in seconds to poll for drifts between origin and target repositories. Default: 180 seconds
- `target_repo` (String) Git repo containing the pattern to deploy. Must use https/http or, for ssh, git@server:foo/bar.git
- `target_revision` (String) Branch, tag, or commit to deploy. Does not support short-sha's. Default: HEAD
- `token_secret` (String) Optional. K8s secret name where the info for connecting to git can be found. The supported secrets are modeled after the private repositories in argo (https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#repositories) currently ssh and username+password are supported
- `token_secret_namespace` (String) Optional. K8s secret namespace where the token for connecting to git can be found


<a id="nestedatt--spec--extra_parameters"></a>
### Nested Schema for `spec.extra_parameters`

Required:

- `name` (String)
- `value` (String)


<a id="nestedatt--spec--git_ops_spec"></a>
### Nested Schema for `spec.git_ops_spec`

Optional:

- `manual_sync` (Boolean) Require manual intervention before Argo will sync new content. Default: False


<a id="nestedatt--spec--multi_source_config"></a>
### Nested Schema for `spec.multi_source_config`

Optional:

- `cluster_group_chart_git_revision` (String) The git reference when deploying the clustergroup helm chart directly from a git repo Defaults to 'main'. (Only used when developing the clustergroup helm chart)
- `cluster_group_chart_version` (String) Which chart version for the clustergroup helm chart. Defaults to '0.8.*'
- `cluster_group_git_repo_url` (String) The url when deploying the clustergroup helm chart directly from a git repo Defaults to '' which means not used (Only used when developing the clustergroup helm chart)
- `enabled` (Boolean) (EXPERIMENTAL) Enable multi-source support when deploying the clustergroup argo application
- `helm_repo_url` (String) The helm chart url to fetch the helm charts from in order to deploy the pattern. Defaults to https://charts.validatedpatterns.io/
