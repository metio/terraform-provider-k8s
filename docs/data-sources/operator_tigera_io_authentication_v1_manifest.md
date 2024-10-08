---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_operator_tigera_io_authentication_v1_manifest Data Source - terraform-provider-k8s"
subcategory: "operator.tigera.io"
description: |-
  Authentication is the Schema for the authentications API
---

# k8s_operator_tigera_io_authentication_v1_manifest (Data Source)

Authentication is the Schema for the authentications API

## Example Usage

```terraform
data "k8s_operator_tigera_io_authentication_v1_manifest" "example" {
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

- `spec` (Attributes) AuthenticationSpec defines the desired state of Authentication (see [below for nested schema](#nestedatt--spec))

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

Optional:

- `dex_deployment` (Attributes) DexDeployment configures the Dex Deployment. (see [below for nested schema](#nestedatt--spec--dex_deployment))
- `groups_prefix` (String) If specified, GroupsPrefix is prepended to each group obtained from the identity provider. Note that Kibana does not support a groups prefix, so this prefix is removed from Kubernetes Groups when translating log access ClusterRoleBindings into Elastic.
- `ldap` (Attributes) LDAP contains the configuration needed to setup LDAP authentication. (see [below for nested schema](#nestedatt--spec--ldap))
- `manager_domain` (String) ManagerDomain is the domain name of the Manager
- `oidc` (Attributes) OIDC contains the configuration needed to setup OIDC authentication. (see [below for nested schema](#nestedatt--spec--oidc))
- `openshift` (Attributes) Openshift contains the configuration needed to setup Openshift OAuth authentication. (see [below for nested schema](#nestedatt--spec--openshift))
- `username_prefix` (String) If specified, UsernamePrefix is prepended to each user obtained from the identity provider. Note that Kibana does not support a user prefix, so this prefix is removed from Kubernetes User when translating log access ClusterRoleBindings into Elastic.

<a id="nestedatt--spec--dex_deployment"></a>
### Nested Schema for `spec.dex_deployment`

Optional:

- `spec` (Attributes) Spec is the specification of the Dex Deployment. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec))

<a id="nestedatt--spec--dex_deployment--spec"></a>
### Nested Schema for `spec.dex_deployment.spec`

Optional:

- `template` (Attributes) Template describes the Dex Deployment pod that will be created. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template))

<a id="nestedatt--spec--dex_deployment--spec--template"></a>
### Nested Schema for `spec.dex_deployment.spec.template`

Optional:

- `spec` (Attributes) Spec is the Dex Deployment's PodSpec. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template--spec))

<a id="nestedatt--spec--dex_deployment--spec--template--spec"></a>
### Nested Schema for `spec.dex_deployment.spec.template.spec`

Optional:

- `containers` (Attributes List) Containers is a list of Dex containers. If specified, this overrides the specified Dex Deployment containers. If omitted, the Dex Deployment will use its default values for its containers. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template--spec--containers))
- `init_containers` (Attributes List) InitContainers is a list of Dex init containers. If specified, this overrides the specified Dex Deployment init containers. If omitted, the Dex Deployment will use its default values for its init containers. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template--spec--init_containers))

<a id="nestedatt--spec--dex_deployment--spec--template--spec--containers"></a>
### Nested Schema for `spec.dex_deployment.spec.template.spec.containers`

Required:

- `name` (String) Name is an enum which identifies the Dex Deployment container by name. Supported values are: tigera-dex

Optional:

- `resources` (Attributes) Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dex Deployment container's resources. If omitted, the Dex Deployment will use its default value for this container's resources. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template--spec--containers--resources))

<a id="nestedatt--spec--dex_deployment--spec--template--spec--containers--resources"></a>
### Nested Schema for `spec.dex_deployment.spec.template.spec.containers.resources`

Optional:

- `claims` (Attributes List) Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template--spec--containers--resources--claims))
- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

<a id="nestedatt--spec--dex_deployment--spec--template--spec--containers--resources--claims"></a>
### Nested Schema for `spec.dex_deployment.spec.template.spec.containers.resources.claims`

Required:

- `name` (String) Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.




<a id="nestedatt--spec--dex_deployment--spec--template--spec--init_containers"></a>
### Nested Schema for `spec.dex_deployment.spec.template.spec.init_containers`

Required:

- `name` (String) Name is an enum which identifies the Dex Deployment init container by name. Supported values are: tigera-dex-tls-key-cert-provisioner

Optional:

- `resources` (Attributes) Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dex Deployment init container's resources. If omitted, the Dex Deployment will use its default value for this init container's resources. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template--spec--init_containers--resources))

<a id="nestedatt--spec--dex_deployment--spec--template--spec--init_containers--resources"></a>
### Nested Schema for `spec.dex_deployment.spec.template.spec.init_containers.resources`

Optional:

- `claims` (Attributes List) Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers. (see [below for nested schema](#nestedatt--spec--dex_deployment--spec--template--spec--init_containers--resources--claims))
- `limits` (Map of String) Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
- `requests` (Map of String) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

<a id="nestedatt--spec--dex_deployment--spec--template--spec--init_containers--resources--claims"></a>
### Nested Schema for `spec.dex_deployment.spec.template.spec.init_containers.resources.claims`

Required:

- `name` (String) Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.








<a id="nestedatt--spec--ldap"></a>
### Nested Schema for `spec.ldap`

Required:

- `host` (String) The host and port of the LDAP server. Example: ad.example.com:636
- `user_search` (Attributes) User entry search configuration to match the credentials with a user. (see [below for nested schema](#nestedatt--spec--ldap--user_search))

Optional:

- `group_search` (Attributes) Group search configuration to find the groups that a user is in. (see [below for nested schema](#nestedatt--spec--ldap--group_search))
- `start_tls` (Boolean) StartTLS whether to enable the startTLS feature for establishing TLS on an existing LDAP session. If true, the ldap:// protocol is used and then issues a StartTLS command, otherwise, connections will use the ldaps:// protocol.

<a id="nestedatt--spec--ldap--user_search"></a>
### Nested Schema for `spec.ldap.user_search`

Required:

- `base_dn` (String) BaseDN to start the search from. For example 'cn=users,dc=example,dc=com'

Optional:

- `filter` (String) Optional filter to apply when searching the directory. For example '(objectClass=person)'
- `name_attribute` (String) A mapping of the attribute that is used as the username. This attribute can be used to apply RBAC to a user. Default: uid


<a id="nestedatt--spec--ldap--group_search"></a>
### Nested Schema for `spec.ldap.group_search`

Required:

- `base_dn` (String) BaseDN to start the search from. For example 'cn=groups,dc=example,dc=com'
- `name_attribute` (String) The attribute of the group that represents its name. This attribute can be used to apply RBAC to a user group.
- `user_matchers` (Attributes List) Following list contains field pairs that are used to match a user to a group. It adds an additional requirement to the filter that an attribute in the group must match the user's attribute value. (see [below for nested schema](#nestedatt--spec--ldap--group_search--user_matchers))

Optional:

- `filter` (String) Optional filter to apply when searching the directory. For example '(objectClass=posixGroup)'

<a id="nestedatt--spec--ldap--group_search--user_matchers"></a>
### Nested Schema for `spec.ldap.group_search.user_matchers`

Required:

- `group_attribute` (String) The attribute of a group that links it to a user.
- `user_attribute` (String) The attribute of a user that links it to a group.




<a id="nestedatt--spec--oidc"></a>
### Nested Schema for `spec.oidc`

Required:

- `issuer_url` (String) IssuerURL is the URL to the OIDC provider.
- `username_claim` (String) UsernameClaim specifies which claim to use from the OIDC provider as the username.

Optional:

- `email_verification` (String) Some providers do not include the claim 'email_verified' when there is no verification in the user enrollment process or if they are acting as a proxy for another identity provider. By default those tokens are deemed invalid. To skip this check, set the value to 'InsecureSkip'. Default: Verify
- `groups_claim` (String) GroupsClaim specifies which claim to use from the OIDC provider as the group.
- `groups_prefix` (String) Deprecated. Please use Authentication.Spec.GroupsPrefix instead.
- `prompt_types` (List of String) PromptTypes is an optional list of string values that specifies whether the identity provider prompts the end user for re-authentication and consent. See the RFC for more information on prompt types: https://openid.net/specs/openid-connect-core-1_0.html. Default: 'Consent'
- `requested_scopes` (List of String) RequestedScopes is a list of scopes to request from the OIDC provider. If not provided, the following scopes are requested: ['openid', 'email', 'profile', 'groups', 'offline_access'].
- `type` (String) Default: 'Dex'
- `username_prefix` (String) Deprecated. Please use Authentication.Spec.UsernamePrefix instead.


<a id="nestedatt--spec--openshift"></a>
### Nested Schema for `spec.openshift`

Required:

- `issuer_url` (String) IssuerURL is the URL to the Openshift OAuth provider. Ex.: https://api.my-ocp-domain.com:6443
