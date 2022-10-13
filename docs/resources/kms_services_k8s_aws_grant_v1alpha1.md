---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_kms_services_k8s_aws_grant_v1alpha1 Resource - terraform-provider-k8s"
subcategory: "kms.services.k8s.aws/v1alpha1"
description: |-
  Grant is the Schema for the Grants API
---

# k8s_kms_services_k8s_aws_grant_v1alpha1 (Resource)

Grant is the Schema for the Grants API

## Example Usage

```terraform
resource "k8s_kms_services_k8s_aws_grant_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) GrantSpec defines the desired state of Grant. (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `api_version` (String) APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
- `id` (Number) The timestamp of the last change to this resource.
- `kind` (String) Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `grantee_principal` (String) The identity that gets the permissions specified in the grant.  To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an Amazon Web Services principal. Valid Amazon Web Services principals include Amazon Web Services accounts (root), IAM users, IAM roles, federated users, and assumed role users. For examples of the ARN syntax to use for specifying a principal, see Amazon Web Services Identity and Access Management (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam) in the Example ARNs section of the Amazon Web Services General Reference.
- `operations` (List of String) A list of operations that the grant permits.  This list must include only operations that are permitted in a grant. Also, the operation must be supported on the KMS key. For example, you cannot create a grant for a symmetric encryption KMS key that allows the Sign operation, or a grant for an asymmetric KMS key that allows the GenerateDataKey operation. If you try, KMS returns a ValidationError exception. For details, see Grant operations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations) in the Key Management Service Developer Guide.

Optional:

- `constraints` (Attributes) Specifies a grant constraint.  KMS supports the EncryptionContextEquals and EncryptionContextSubset grant constraints. Each constraint value can include up to 8 encryption context pairs. The encryption context value in each constraint cannot exceed 384 characters. For information about grant constraints, see Using grant constraints (https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-constraints) in the Key Management Service Developer Guide. For more information about encryption context, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the Key Management Service Developer Guide .  The encryption context grant constraints allow the permissions in the grant only when the encryption context in the request matches (EncryptionContextEquals) or includes (EncryptionContextSubset) the encryption context specified in this structure.  The encryption context grant constraints are supported only on grant operations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations) that include an EncryptionContext parameter, such as cryptographic operations on symmetric encryption KMS keys. Grants with grant constraints can include the DescribeKey and RetireGrant operations, but the constraint doesn't apply to these operations. If a grant with a grant constraint includes the CreateGrant operation, the constraint requires that any grants created with the CreateGrant permission have an equally strict or stricter encryption context constraint.  You cannot use an encryption context grant constraint for cryptographic operations with asymmetric KMS keys or HMAC KMS keys. These keys don't support an encryption context. (see [below for nested schema](#nestedatt--spec--constraints))
- `grant_tokens` (List of String) A list of grant tokens.  Use a grant token when your permission to call this operation comes from a new grant that has not yet achieved eventual consistency. For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token) and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token) in the Key Management Service Developer Guide.
- `key_id` (String) Identifies the KMS key for the grant. The grant gives principals permission to use this KMS key.  Specify the key ID or key ARN of the KMS key. To specify a KMS key in a different Amazon Web Services account, you must use the key ARN.  For example:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
- `key_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api (see [below for nested schema](#nestedatt--spec--key_ref))
- `name` (String) A friendly name for the grant. Use this value to prevent the unintended creation of duplicate grants when retrying this request.  When this value is absent, all CreateGrant requests result in a new grant with a unique GrantId even if all the supplied parameters are identical. This can result in unintended duplicates when you retry the CreateGrant request.  When this value is present, you can retry a CreateGrant request with identical parameters; if the grant already exists, the original GrantId is returned without creating a new grant. Note that the returned grant token is unique with every CreateGrant request, even when a duplicate GrantId is returned. All grant tokens for the same grant ID can be used interchangeably.
- `retiring_principal` (String) The principal that has permission to use the RetireGrant operation to retire the grant.  To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an Amazon Web Services principal. Valid Amazon Web Services principals include Amazon Web Services accounts (root), IAM users, federated users, and assumed role users. For examples of the ARN syntax to use for specifying a principal, see Amazon Web Services Identity and Access Management (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam) in the Example ARNs section of the Amazon Web Services General Reference.  The grant determines the retiring principal. Other principals might have permission to retire the grant or revoke the grant. For details, see RevokeGrant and Retiring and revoking grants (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete) in the Key Management Service Developer Guide.

<a id="nestedatt--spec--constraints"></a>
### Nested Schema for `spec.constraints`

Optional:

- `encryption_context_equals` (Map of String)
- `encryption_context_subset` (Map of String)


<a id="nestedatt--spec--key_ref"></a>
### Nested Schema for `spec.key_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--key_ref--from))

<a id="nestedatt--spec--key_ref--from"></a>
### Nested Schema for `spec.key_ref.from`

Optional:

- `name` (String)

