---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_hyperfoil_io_horreum_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "hyperfoil.io"
description: |-
  Horreum is the object configuring Horreum performance results repository
---

# k8s_hyperfoil_io_horreum_v1alpha1_manifest (Data Source)

Horreum is the object configuring Horreum performance results repository

## Example Usage

```terraform
data "k8s_hyperfoil_io_horreum_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) HorreumSpec defines the desired state of Horreum (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `admin_secret` (String) Name of secret resource with data 'username' and 'password'. This will be the first user that get's created in Horreum with the 'admin' role, therefore it can create other users and teams. Created automatically if it does not exist.
- `database` (Attributes) Database coordinates for Horreum data. Besides 'username' and 'password' the secret must also contain key 'dbsecret' that will be used to sign access to the database. (see [below for nested schema](#nestedatt--spec--database))
- `image` (String) Horreum image. Defaults to quay.io/hyperfoil/horreum:latest
- `keycloak` (Attributes) Keycloak specification (see [below for nested schema](#nestedatt--spec--keycloak))
- `node_host` (String) Host used for NodePort services
- `postgres` (Attributes) PostgreSQL specification (see [below for nested schema](#nestedatt--spec--postgres))
- `route` (Attributes) Route for external access (see [below for nested schema](#nestedatt--spec--route))
- `service_type` (String) Alternative service type when routes are not available (e.g. on vanilla K8s)

<a id="nestedatt--spec--database"></a>
### Nested Schema for `spec.database`

Optional:

- `host` (String) Hostname for the database
- `name` (String) Name of the database
- `port` (Number) Database port; defaults to 5432
- `secret` (String) Name of secret resource with data 'username' and 'password'. Created if does not exist.


<a id="nestedatt--spec--keycloak"></a>
### Nested Schema for `spec.keycloak`

Optional:

- `admin_secret` (String) Secret used for admin access to the deployed Keycloak instance. Created if does not exist. Must contain keys 'username' and 'password'.
- `database` (Attributes) Database coordinates Keycloak should use (see [below for nested schema](#nestedatt--spec--keycloak--database))
- `external` (Attributes) When this is set Keycloak instance will not be deployed and Horreum will use this external instance. (see [below for nested schema](#nestedatt--spec--keycloak--external))
- `image` (String) Image that should be used for Keycloak deployment. Defaults to quay.io/keycloak/keycloak:latest
- `route` (Attributes) Route for external access to the Keycloak instance. (see [below for nested schema](#nestedatt--spec--keycloak--route))
- `service_type` (String) Alternative service type when routes are not available (e.g. on vanilla K8s)

<a id="nestedatt--spec--keycloak--database"></a>
### Nested Schema for `spec.keycloak.database`

Optional:

- `host` (String) Hostname for the database
- `name` (String) Name of the database
- `port` (Number) Database port; defaults to 5432
- `secret` (String) Name of secret resource with data 'username' and 'password'. Created if does not exist.


<a id="nestedatt--spec--keycloak--external"></a>
### Nested Schema for `spec.keycloak.external`

Optional:

- `internal_uri` (String) Internal URI - Horreum will use this for communication but won't disclose that.
- `public_uri` (String) Public facing URI - Horreum will send this URI to the clients.


<a id="nestedatt--spec--keycloak--route"></a>
### Nested Schema for `spec.keycloak.route`

Optional:

- `host` (String) Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com
- `tls` (String) Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'
- `type` (String) Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'



<a id="nestedatt--spec--postgres"></a>
### Nested Schema for `spec.postgres`

Optional:

- `admin_secret` (String) Secret used for unrestricted access to the database. Created if does not exist. Must contain keys 'username' and 'password'.
- `enabled` (Boolean) True (or omitted) to deploy PostgreSQL database
- `image` (String) Image used for PostgreSQL deployment. Defaults to registry.redhat.io/rhel8/postgresql-12:latest
- `persistent_volume_claim` (String) Name of PVC where the database will store the data. If empty, ephemeral storage will be used.
- `user` (Number) Id of the user the container should run as


<a id="nestedatt--spec--route"></a>
### Nested Schema for `spec.route`

Optional:

- `host` (String) Host for the route leading to Controller REST endpoint. Example: horreum.apps.mycloud.example.com
- `tls` (String) Optional for edge and reencrypt routes, required for passthrough; Name of the secret hosting 'tls.crt', 'tls.key' and optionally 'ca.crt'
- `type` (String) Either 'http' (for plain-text routes - not recommended), 'edge', 'reencrypt' or 'passthrough'
