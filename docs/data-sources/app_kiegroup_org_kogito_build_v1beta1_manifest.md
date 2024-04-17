---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_app_kiegroup_org_kogito_build_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "app.kiegroup.org"
description: |-
  KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.
---

# k8s_app_kiegroup_org_kogito_build_v1beta1_manifest (Data Source)

KogitoBuild handles how to build a custom Kogito service in a Kubernetes/OpenShift cluster.

## Example Usage

```terraform
data "k8s_app_kiegroup_org_kogito_build_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    type = "RemoteSource"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) KogitoBuildSpec defines the desired state of KogitoBuild. (see [below for nested schema](#nestedatt--spec))

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

- `type` (String) Sets the type of build that this instance will handle:  Binary - takes an uploaded binary file already compiled and creates a Kogito service image from it.  RemoteSource - pulls the source code from a Git repository, builds the binary and then the final Kogito service image.  LocalSource - takes an uploaded resource file such as DRL (rules), DMN (decision) or BPMN (process), builds the binary and the final Kogito service image.

Optional:

- `artifact` (Attributes) Artifact contains override information for building the Maven artifact (used for Local Source builds).  You might want to override this information when building from decisions, rules or process files. In this scenario the Kogito Images will generate a new Java project for you underneath. This information will be used to generate this project. (see [below for nested schema](#nestedatt--spec--artifact))
- `build_image` (String) Image used to build the Kogito Service from source (Local and Remote).  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.
- `disable_incremental` (Boolean) DisableIncremental indicates that source to image builds should NOT be incremental. Defaults to false.
- `enable_maven_download_output` (Boolean) If set to true will print the logs for downloading/uploading of maven dependencies. Defaults to false.
- `env` (Attributes List) Environment variables used during build time. (see [below for nested schema](#nestedatt--spec--env))
- `git_source` (Attributes) Information about the git repository where the Kogito Service source code resides.  Ignored for binary builds. (see [below for nested schema](#nestedatt--spec--git_source))
- `maven_mirror_url` (String) Maven Mirror URL to be used during source-to-image builds (Local and Remote) to considerably increase build speed.
- `native` (Boolean) Native indicates if the Kogito Service built should be compiled to run on native mode when Runtime is Quarkus (Source to Image build only).  For more information, see https://www.graalvm.org/docs/reference-manual/aot-compilation/.
- `resources` (Attributes) Resources Requirements for builder pods. (see [below for nested schema](#nestedatt--spec--resources))
- `runtime` (String) Which runtime Kogito service base image to use when building the Kogito service. If 'BuildImage' is set, this value is ignored by the operator. Default value: quarkus.
- `runtime_image` (String) Image used as the base image for the final Kogito service. This image only has the required packages to run the application.  For example: quarkus based services will have only JVM installed, native services only the packages required by the OS.  If not defined the operator will use image provided by the Kogito Team based on the 'Runtime' field.  Example: 'quay.io/kiegroup/kogito-jvm-builder:latest'.  On OpenShift an ImageStream will be created in the current namespace pointing to the given image.
- `target_kogito_runtime` (String) Set this field targeting the desired KogitoRuntime when this KogitoBuild instance has a different name than the KogitoRuntime.  By default this KogitoBuild instance will generate a final image named after its own name (.metadata.name).  On OpenShift, an ImageStream will be created causing a redeployment on any KogitoRuntime with the same name. On Kubernetes, the final image will be pushed to the KogitoRuntime deployment.  If you have multiple KogitoBuild instances (let's say BinaryBuildType and Remote Source), you might need that both target the same KogitoRuntime. Both KogitoBuilds will update the same ImageStream or generate a final image to the same KogitoRuntime deployment.
- `web_hooks` (Attributes List) WebHooks secrets for source to image builds based on Git repositories (Remote Sources). (see [below for nested schema](#nestedatt--spec--web_hooks))

<a id="nestedatt--spec--artifact"></a>
### Nested Schema for `spec.artifact`

Optional:

- `artifact_id` (String) Indicates the unique base name of the primary artifact being generated.
- `group_id` (String) Indicates the unique identifier of the organization or group that created the project.
- `version` (String) Indicates the version of the artifact generated by the project.


<a id="nestedatt--spec--env"></a>
### Nested Schema for `spec.env`

Required:

- `name` (String) Name of the environment variable. Must be a C_IDENTIFIER.

Optional:

- `value` (String) Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.
- `value_from` (Attributes) Source for the environment variable's value. Cannot be used if value is not empty. (see [below for nested schema](#nestedatt--spec--env--value_from))

<a id="nestedatt--spec--env--value_from"></a>
### Nested Schema for `spec.env.value_from`

Optional:

- `config_map_key_ref` (Attributes) Selects a key of a ConfigMap. (see [below for nested schema](#nestedatt--spec--env--value_from--config_map_key_ref))
- `field_ref` (Attributes) Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs. (see [below for nested schema](#nestedatt--spec--env--value_from--field_ref))
- `resource_field_ref` (Attributes) Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported. (see [below for nested schema](#nestedatt--spec--env--value_from--resource_field_ref))
- `secret_key_ref` (Attributes) Selects a key of a secret in the pod's namespace (see [below for nested schema](#nestedatt--spec--env--value_from--secret_key_ref))

<a id="nestedatt--spec--env--value_from--config_map_key_ref"></a>
### Nested Schema for `spec.env.value_from.config_map_key_ref`

Required:

- `key` (String) The key to select.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the ConfigMap or its key must be defined


<a id="nestedatt--spec--env--value_from--field_ref"></a>
### Nested Schema for `spec.env.value_from.field_ref`

Required:

- `field_path` (String) Path of the field to select in the specified API version.

Optional:

- `api_version` (String) Version of the schema the FieldPath is written in terms of, defaults to 'v1'.


<a id="nestedatt--spec--env--value_from--resource_field_ref"></a>
### Nested Schema for `spec.env.value_from.resource_field_ref`

Required:

- `resource` (String) Required: resource to select

Optional:

- `container_name` (String) Container name: required for volumes, optional for env vars
- `divisor` (String) Specifies the output format of the exposed resources, defaults to '1'


<a id="nestedatt--spec--env--value_from--secret_key_ref"></a>
### Nested Schema for `spec.env.value_from.secret_key_ref`

Required:

- `key` (String) The key of the secret to select from.  Must be a valid secret key.

Optional:

- `name` (String) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
- `optional` (Boolean) Specify whether the Secret or its key must be defined




<a id="nestedatt--spec--git_source"></a>
### Nested Schema for `spec.git_source`

Required:

- `uri` (String) Git URI for the s2i source.

Optional:

- `context_dir` (String) Context/subdirectory where the code is located, relative to the repo root.
- `reference` (String) Branch to use in the Git repository.


<a id="nestedatt--spec--resources"></a>
### Nested Schema for `spec.resources`

Optional:

- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/


<a id="nestedatt--spec--web_hooks"></a>
### Nested Schema for `spec.web_hooks`

Optional:

- `secret` (String) Secret value for webHook
- `type` (String) WebHook type, either GitHub or Generic.