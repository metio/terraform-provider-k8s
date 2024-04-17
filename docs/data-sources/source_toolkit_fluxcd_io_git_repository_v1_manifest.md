---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_source_toolkit_fluxcd_io_git_repository_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "source.toolkit.fluxcd.io"
description: |-
  GitRepository is the Schema for the gitrepositories API.
---

# k8s_source_toolkit_fluxcd_io_git_repository_v1_manifest (Data Source)

GitRepository is the Schema for the gitrepositories API.

## Example Usage

```terraform
data "k8s_source_toolkit_fluxcd_io_git_repository_v1_manifest" "example" {
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

- `spec` (Attributes) GitRepositorySpec specifies the required configuration to produce anArtifact for a Git repository. (see [below for nested schema](#nestedatt--spec))

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

- `interval` (String) Interval at which the GitRepository URL is checked for updates.This interval is approximate and may be subject to jitter to ensureefficient use of resources.
- `url` (String) URL specifies the Git repository URL, it can be an HTTP/S or SSH address.

Optional:

- `ignore` (String) Ignore overrides the set of excluded patterns in the .sourceignore format(which is the same as .gitignore). If not provided, a default will be used,consult the documentation for your version to find out what those are.
- `include` (Attributes List) Include specifies a list of GitRepository resources which Artifactsshould be included in the Artifact produced for this GitRepository. (see [below for nested schema](#nestedatt--spec--include))
- `proxy_secret_ref` (Attributes) ProxySecretRef specifies the Secret containing the proxy configurationto use while communicating with the Git server. (see [below for nested schema](#nestedatt--spec--proxy_secret_ref))
- `recurse_submodules` (Boolean) RecurseSubmodules enables the initialization of all submodules withinthe GitRepository as cloned from the URL, using their default settings.
- `ref` (Attributes) Reference specifies the Git reference to resolve and monitor forchanges, defaults to the 'master' branch. (see [below for nested schema](#nestedatt--spec--ref))
- `secret_ref` (Attributes) SecretRef specifies the Secret containing authentication credentials forthe GitRepository.For HTTPS repositories the Secret must contain 'username' and 'password'fields for basic auth or 'bearerToken' field for token auth.For SSH repositories the Secret must contain 'identity'and 'known_hosts' fields. (see [below for nested schema](#nestedatt--spec--secret_ref))
- `suspend` (Boolean) Suspend tells the controller to suspend the reconciliation of thisGitRepository.
- `timeout` (String) Timeout for Git operations like cloning, defaults to 60s.
- `verify` (Attributes) Verification specifies the configuration to verify the Git commitsignature(s). (see [below for nested schema](#nestedatt--spec--verify))

<a id="nestedatt--spec--include"></a>
### Nested Schema for `spec.include`

Required:

- `repository` (Attributes) GitRepositoryRef specifies the GitRepository which Artifact contentsmust be included. (see [below for nested schema](#nestedatt--spec--include--repository))

Optional:

- `from_path` (String) FromPath specifies the path to copy contents from, defaults to the rootof the Artifact.
- `to_path` (String) ToPath specifies the path to copy contents to, defaults to the name ofthe GitRepositoryRef.

<a id="nestedatt--spec--include--repository"></a>
### Nested Schema for `spec.include.repository`

Required:

- `name` (String) Name of the referent.



<a id="nestedatt--spec--proxy_secret_ref"></a>
### Nested Schema for `spec.proxy_secret_ref`

Required:

- `name` (String) Name of the referent.


<a id="nestedatt--spec--ref"></a>
### Nested Schema for `spec.ref`

Optional:

- `branch` (String) Branch to check out, defaults to 'master' if no other field is defined.
- `commit` (String) Commit SHA to check out, takes precedence over all reference fields.This can be combined with Branch to shallow clone the branch, in whichthe commit is expected to exist.
- `name` (String) Name of the reference to check out; takes precedence over Branch, Tag and SemVer.It must be a valid Git reference: https://git-scm.com/docs/git-check-ref-format#_descriptionExamples: 'refs/heads/main', 'refs/tags/v0.1.0', 'refs/pull/420/head', 'refs/merge-requests/1/head'
- `semver` (String) SemVer tag expression to check out, takes precedence over Tag.
- `tag` (String) Tag to check out, takes precedence over Branch.


<a id="nestedatt--spec--secret_ref"></a>
### Nested Schema for `spec.secret_ref`

Required:

- `name` (String) Name of the referent.


<a id="nestedatt--spec--verify"></a>
### Nested Schema for `spec.verify`

Required:

- `secret_ref` (Attributes) SecretRef specifies the Secret containing the public keys of trusted Gitauthors. (see [below for nested schema](#nestedatt--spec--verify--secret_ref))

Optional:

- `mode` (String) Mode specifies which Git object(s) should be verified.The variants 'head' and 'HEAD' both imply the same thing, i.e. verifythe commit that the HEAD of the Git repository points to. The variant'head' solely exists to ensure backwards compatibility.

<a id="nestedatt--spec--verify--secret_ref"></a>
### Nested Schema for `spec.verify.secret_ref`

Required:

- `name` (String) Name of the referent.