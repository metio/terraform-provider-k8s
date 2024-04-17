---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_resources_teleport_dev_teleport_saml_connector_v2_manifest Data Source - terraform-provider-k8s"
subcategory: "resources.teleport.dev"
description: |-
  SAMLConnector is the Schema for the samlconnectors API
---

# k8s_resources_teleport_dev_teleport_saml_connector_v2_manifest (Data Source)

SAMLConnector is the Schema for the samlconnectors API

## Example Usage

```terraform
data "k8s_resources_teleport_dev_teleport_saml_connector_v2_manifest" "example" {
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

- `spec` (Attributes) SAMLConnector resource definition v2 from Teleport (see [below for nested schema](#nestedatt--spec))

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

- `acs` (String) AssertionConsumerService is a URL for assertion consumer service on the service provider (Teleport's side).
- `allow_idp_initiated` (Boolean) AllowIDPInitiated is a flag that indicates if the connector can be used for IdP-initiated logins.
- `assertion_key_pair` (Attributes) EncryptionKeyPair is a key pair used for decrypting SAML assertions. (see [below for nested schema](#nestedatt--spec--assertion_key_pair))
- `attributes_to_roles` (Attributes List) AttributesToRoles is a list of mappings of attribute statements to roles. (see [below for nested schema](#nestedatt--spec--attributes_to_roles))
- `audience` (String) Audience uniquely identifies our service provider.
- `cert` (String) Cert is the identity provider certificate PEM. IDP signs <Response> responses using this certificate.
- `display` (String) Display controls how this connector is displayed.
- `entity_descriptor` (String) EntityDescriptor is XML with descriptor. It can be used to supply configuration parameters in one XML file rather than supplying them in the individual elements.
- `entity_descriptor_url` (String) EntityDescriptorURL is a URL that supplies a configuration XML.
- `issuer` (String) Issuer is the identity provider issuer.
- `provider` (String) Provider is the external identity provider.
- `service_provider_issuer` (String) ServiceProviderIssuer is the issuer of the service provider (Teleport).
- `signing_key_pair` (Attributes) SigningKeyPair is an x509 key pair used to sign AuthnRequest. (see [below for nested schema](#nestedatt--spec--signing_key_pair))
- `sso` (String) SSO is the URL of the identity provider's SSO service.

<a id="nestedatt--spec--assertion_key_pair"></a>
### Nested Schema for `spec.assertion_key_pair`

Optional:

- `cert` (String) Cert is a PEM-encoded x509 certificate.
- `private_key` (String) PrivateKey is a PEM encoded x509 private key.


<a id="nestedatt--spec--attributes_to_roles"></a>
### Nested Schema for `spec.attributes_to_roles`

Optional:

- `name` (String) Name is an attribute statement name.
- `roles` (List of String) Roles is a list of static teleport roles to map to.
- `value` (String) Value is an attribute statement value to match.


<a id="nestedatt--spec--signing_key_pair"></a>
### Nested Schema for `spec.signing_key_pair`

Optional:

- `cert` (String) Cert is a PEM-encoded x509 certificate.
- `private_key` (String) PrivateKey is a PEM encoded x509 private key.