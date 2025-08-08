/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kms_services_k8s_aws_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KmsServicesK8SAwsKeyV1Alpha1Manifest{}
)

func NewKmsServicesK8SAwsKeyV1Alpha1Manifest() datasource.DataSource {
	return &KmsServicesK8SAwsKeyV1Alpha1Manifest{}
}

type KmsServicesK8SAwsKeyV1Alpha1Manifest struct{}

type KmsServicesK8SAwsKeyV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BypassPolicyLockoutSafetyCheck *bool   `tfsdk:"bypass_policy_lockout_safety_check" json:"bypassPolicyLockoutSafetyCheck,omitempty"`
		CustomKeyStoreID               *string `tfsdk:"custom_key_store_id" json:"customKeyStoreID,omitempty"`
		Description                    *string `tfsdk:"description" json:"description,omitempty"`
		EnableKeyRotation              *bool   `tfsdk:"enable_key_rotation" json:"enableKeyRotation,omitempty"`
		KeySpec                        *string `tfsdk:"key_spec" json:"keySpec,omitempty"`
		KeyUsage                       *string `tfsdk:"key_usage" json:"keyUsage,omitempty"`
		MultiRegion                    *bool   `tfsdk:"multi_region" json:"multiRegion,omitempty"`
		Origin                         *string `tfsdk:"origin" json:"origin,omitempty"`
		Policy                         *string `tfsdk:"policy" json:"policy,omitempty"`
		Tags                           *[]struct {
			TagKey   *string `tfsdk:"tag_key" json:"tagKey,omitempty"`
			TagValue *string `tfsdk:"tag_value" json:"tagValue,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kms_services_k8s_aws_key_v1alpha1_manifest"
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Key is the Schema for the Keys API",
		MarkdownDescription: "Key is the Schema for the Keys API",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "KeySpec defines the desired state of Key.",
				MarkdownDescription: "KeySpec defines the desired state of Key.",
				Attributes: map[string]schema.Attribute{
					"bypass_policy_lockout_safety_check": schema.BoolAttribute{
						Description:         "Skips ('bypasses') the key policy lockout safety check. The default value is false. Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key) in the Key Management Service Developer Guide. Use this parameter only when you intend to prevent the principal that is making the request from making a subsequent PutKeyPolicy (https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key.",
						MarkdownDescription: "Skips ('bypasses') the key policy lockout safety check. The default value is false. Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key) in the Key Management Service Developer Guide. Use this parameter only when you intend to prevent the principal that is making the request from making a subsequent PutKeyPolicy (https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_key_store_id": schema.StringAttribute{
						Description:         "Creates the KMS key in the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html). The ConnectionState of the custom key store must be CONNECTED. To find the CustomKeyStoreID and ConnectionState use the DescribeCustomKeyStores operation. This parameter is valid only for symmetric encryption KMS keys in a single Region. You cannot create any other type of KMS key in a custom key store. When you create a KMS key in an CloudHSM key store, KMS generates a non-exportable 256-bit symmetric key in its associated CloudHSM cluster and associates it with the KMS key. When you create a KMS key in an external key store, you must use the XksKeyId parameter to specify an external key that serves as key material for the KMS key.",
						MarkdownDescription: "Creates the KMS key in the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html). The ConnectionState of the custom key store must be CONNECTED. To find the CustomKeyStoreID and ConnectionState use the DescribeCustomKeyStores operation. This parameter is valid only for symmetric encryption KMS keys in a single Region. You cannot create any other type of KMS key in a custom key store. When you create a KMS key in an CloudHSM key store, KMS generates a non-exportable 256-bit symmetric key in its associated CloudHSM cluster and associates it with the KMS key. When you create a KMS key in an external key store, you must use the XksKeyId parameter to specify an external key that serves as key material for the KMS key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A description of the KMS key. Use a description that helps you decide whether the KMS key is appropriate for a task. The default value is an empty string (no description). Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in CloudTrail logs and other output. To set or change the description after the key is created, use UpdateKeyDescription.",
						MarkdownDescription: "A description of the KMS key. Use a description that helps you decide whether the KMS key is appropriate for a task. The default value is an empty string (no description). Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in CloudTrail logs and other output. To set or change the description after the key is created, use UpdateKeyDescription.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_key_rotation": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_spec": schema.StringAttribute{
						Description:         "Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT, creates a KMS key with a 256-bit AES-GCM key that is used for encryption and decryption, except in China Regions, where it creates a 128-bit symmetric key that uses SM4 encryption. For help choosing a key spec for your KMS key, see Choosing a KMS key type (https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html#symm-asymm-choose) in the Key Management Service Developer Guide . The KeySpec determines whether the KMS key contains a symmetric key or an asymmetric key pair. It also determines the algorithms that the KMS key supports. You can't change the KeySpec after the KMS key is created. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see kms:EncryptionAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm), kms:MacAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-mac-algorithm) or kms:Signing Algorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm) in the Key Management Service Developer Guide . Amazon Web Services services that are integrated with KMS (http://aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support asymmetric KMS keys or HMAC KMS keys. KMS supports the following key specs for KMS keys: * Symmetric encryption key (default) SYMMETRIC_DEFAULT * HMAC keys (symmetric) HMAC_224 HMAC_256 HMAC_384 HMAC_512 * Asymmetric RSA key pairs (encryption and decryption -or- signing and verification) RSA_2048 RSA_3072 RSA_4096 * Asymmetric NIST-recommended elliptic curve key pairs (signing and verification -or- deriving shared secrets) ECC_NIST_P256 (secp256r1) ECC_NIST_P384 (secp384r1) ECC_NIST_P521 (secp521r1) * Other asymmetric elliptic curve key pairs (signing and verification) ECC_SECG_P256K1 (secp256k1), commonly used for cryptocurrencies. * SM2 key pairs (encryption and decryption -or- signing and verification -or- deriving shared secrets) SM2 (China Regions only)",
						MarkdownDescription: "Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT, creates a KMS key with a 256-bit AES-GCM key that is used for encryption and decryption, except in China Regions, where it creates a 128-bit symmetric key that uses SM4 encryption. For help choosing a key spec for your KMS key, see Choosing a KMS key type (https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html#symm-asymm-choose) in the Key Management Service Developer Guide . The KeySpec determines whether the KMS key contains a symmetric key or an asymmetric key pair. It also determines the algorithms that the KMS key supports. You can't change the KeySpec after the KMS key is created. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see kms:EncryptionAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm), kms:MacAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-mac-algorithm) or kms:Signing Algorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm) in the Key Management Service Developer Guide . Amazon Web Services services that are integrated with KMS (http://aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support asymmetric KMS keys or HMAC KMS keys. KMS supports the following key specs for KMS keys: * Symmetric encryption key (default) SYMMETRIC_DEFAULT * HMAC keys (symmetric) HMAC_224 HMAC_256 HMAC_384 HMAC_512 * Asymmetric RSA key pairs (encryption and decryption -or- signing and verification) RSA_2048 RSA_3072 RSA_4096 * Asymmetric NIST-recommended elliptic curve key pairs (signing and verification -or- deriving shared secrets) ECC_NIST_P256 (secp256r1) ECC_NIST_P384 (secp384r1) ECC_NIST_P521 (secp521r1) * Other asymmetric elliptic curve key pairs (signing and verification) ECC_SECG_P256K1 (secp256k1), commonly used for cryptocurrencies. * SM2 key pairs (encryption and decryption -or- signing and verification -or- deriving shared secrets) SM2 (China Regions only)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_usage": schema.StringAttribute{
						Description:         "Determines the cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is ENCRYPT_DECRYPT. This parameter is optional when you are creating a symmetric encryption KMS key; otherwise, it is required. You can't change the KeyUsage value after the KMS key is created. Select only one valid value. * For symmetric encryption KMS keys, omit the parameter or specify ENCRYPT_DECRYPT. * For HMAC KMS keys (symmetric), specify GENERATE_VERIFY_MAC. * For asymmetric KMS keys with RSA key pairs, specify ENCRYPT_DECRYPT or SIGN_VERIFY. * For asymmetric KMS keys with NIST-recommended elliptic curve key pairs, specify SIGN_VERIFY or KEY_AGREEMENT. * For asymmetric KMS keys with ECC_SECG_P256K1 key pairs specify SIGN_VERIFY. * For asymmetric KMS keys with SM2 key pairs (China Regions only), specify ENCRYPT_DECRYPT, SIGN_VERIFY, or KEY_AGREEMENT.",
						MarkdownDescription: "Determines the cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is ENCRYPT_DECRYPT. This parameter is optional when you are creating a symmetric encryption KMS key; otherwise, it is required. You can't change the KeyUsage value after the KMS key is created. Select only one valid value. * For symmetric encryption KMS keys, omit the parameter or specify ENCRYPT_DECRYPT. * For HMAC KMS keys (symmetric), specify GENERATE_VERIFY_MAC. * For asymmetric KMS keys with RSA key pairs, specify ENCRYPT_DECRYPT or SIGN_VERIFY. * For asymmetric KMS keys with NIST-recommended elliptic curve key pairs, specify SIGN_VERIFY or KEY_AGREEMENT. * For asymmetric KMS keys with ECC_SECG_P256K1 key pairs specify SIGN_VERIFY. * For asymmetric KMS keys with SM2 key pairs (China Regions only), specify ENCRYPT_DECRYPT, SIGN_VERIFY, or KEY_AGREEMENT.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"multi_region": schema.BoolAttribute{
						Description:         "Creates a multi-Region primary key that you can replicate into other Amazon Web Services Regions. You cannot change this value after you create the KMS key. For a multi-Region key, set this parameter to True. For a single-Region KMS key, omit this parameter or set it to False. The default value is False. This operation supports multi-Region keys, an KMS feature that lets you create multiple interoperable KMS keys in different Amazon Web Services Regions. Because these KMS keys have the same key ID, key material, and other metadata, you can use them interchangeably to encrypt data in one Amazon Web Services Region and decrypt it in a different Amazon Web Services Region without re-encrypting the data or making a cross-Region call. For more information about multi-Region keys, see Multi-Region keys in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the Key Management Service Developer Guide. This value creates a primary key, not a replica. To create a replica key, use the ReplicateKey operation. You can create a symmetric or asymmetric multi-Region key, and you can create a multi-Region key with imported key material. However, you cannot create a multi-Region key in a custom key store.",
						MarkdownDescription: "Creates a multi-Region primary key that you can replicate into other Amazon Web Services Regions. You cannot change this value after you create the KMS key. For a multi-Region key, set this parameter to True. For a single-Region KMS key, omit this parameter or set it to False. The default value is False. This operation supports multi-Region keys, an KMS feature that lets you create multiple interoperable KMS keys in different Amazon Web Services Regions. Because these KMS keys have the same key ID, key material, and other metadata, you can use them interchangeably to encrypt data in one Amazon Web Services Region and decrypt it in a different Amazon Web Services Region without re-encrypting the data or making a cross-Region call. For more information about multi-Region keys, see Multi-Region keys in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the Key Management Service Developer Guide. This value creates a primary key, not a replica. To create a replica key, use the ReplicateKey operation. You can create a symmetric or asymmetric multi-Region key, and you can create a multi-Region key with imported key material. However, you cannot create a multi-Region key in a custom key store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"origin": schema.StringAttribute{
						Description:         "The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that KMS creates the key material. To create a KMS key with no key material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html) (for imported key material), set this value to EXTERNAL. For more information about importing key material into KMS, see Importing Key Material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in the Key Management Service Developer Guide. The EXTERNAL origin value is valid only for symmetric KMS keys. To create a KMS key in an CloudHSM key store (https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) and create its key material in the associated CloudHSM cluster, set this value to AWS_CLOUDHSM. You must also use the CustomKeyStoreId parameter to identify the CloudHSM key store. The KeySpec value must be SYMMETRIC_DEFAULT. To create a KMS key in an external key store (https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keys.html), set this value to EXTERNAL_KEY_STORE. You must also use the CustomKeyStoreId parameter to identify the external key store and the XksKeyId parameter to identify the associated external key. The KeySpec value must be SYMMETRIC_DEFAULT.",
						MarkdownDescription: "The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that KMS creates the key material. To create a KMS key with no key material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html) (for imported key material), set this value to EXTERNAL. For more information about importing key material into KMS, see Importing Key Material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in the Key Management Service Developer Guide. The EXTERNAL origin value is valid only for symmetric KMS keys. To create a KMS key in an CloudHSM key store (https://docs.aws.amazon.com/kms/latest/developerguide/create-cmk-keystore.html) and create its key material in the associated CloudHSM cluster, set this value to AWS_CLOUDHSM. You must also use the CustomKeyStoreId parameter to identify the CloudHSM key store. The KeySpec value must be SYMMETRIC_DEFAULT. To create a KMS key in an external key store (https://docs.aws.amazon.com/kms/latest/developerguide/create-xks-keys.html), set this value to EXTERNAL_KEY_STORE. You must also use the CustomKeyStoreId parameter to identify the external key store and the XksKeyId parameter to identify the associated external key. The KeySpec value must be SYMMETRIC_DEFAULT.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy": schema.StringAttribute{
						Description:         "The key policy to attach to the KMS key. If you provide a key policy, it must meet the following criteria: * The key policy must allow the calling principal to make a subsequent PutKeyPolicy request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key) in the Key Management Service Developer Guide. (To omit this condition, set BypassPolicyLockoutSafetyCheck to true.) * Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to KMS. When you create a new Amazon Web Services principal, you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to KMS. For more information, see Changes that I make are not always immediately visible (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the Amazon Web Services Identity and Access Management User Guide. If you do not provide a key policy, KMS attaches a default key policy to the KMS key. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the Key Management Service Developer Guide. The key policy size quota is 32 kilobytes (32768 bytes). For help writing and formatting a JSON policy document, see the IAM JSON Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in the Identity and Access Management User Guide . Regex Pattern: '^[u0009u000Au000Du0020-u00FF]+$'",
						MarkdownDescription: "The key policy to attach to the KMS key. If you provide a key policy, it must meet the following criteria: * The key policy must allow the calling principal to make a subsequent PutKeyPolicy request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key) in the Key Management Service Developer Guide. (To omit this condition, set BypassPolicyLockoutSafetyCheck to true.) * Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to KMS. When you create a new Amazon Web Services principal, you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to KMS. For more information, see Changes that I make are not always immediately visible (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the Amazon Web Services Identity and Access Management User Guide. If you do not provide a key policy, KMS attaches a default key policy to the KMS key. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the Key Management Service Developer Guide. The key policy size quota is 32 kilobytes (32768 bytes). For help writing and formatting a JSON policy document, see the IAM JSON Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in the Identity and Access Management User Guide . Regex Pattern: '^[u0009u000Au000Du0020-u00FF]+$'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Assigns one or more tags to the KMS key. Use this parameter to tag the KMS key when it is created. To tag an existing KMS key, use the TagResource operation. Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in CloudTrail logs and other output. Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see ABAC for KMS (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the Key Management Service Developer Guide. To use this parameter, you must have kms:TagResource (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html) permission in an IAM policy. Each tag consists of a tag key and a tag value. Both the tag key and the tag value are required, but the tag value can be an empty (null) string. You cannot have more than one tag on a KMS key with the same tag key. If you specify an existing tag key with a different tag value, KMS replaces the current tag value with the specified one. When you add tags to an Amazon Web Services resource, Amazon Web Services generates a cost allocation report with usage and costs aggregated by tags. Tags can also be used to control access to a KMS key. For details, see Tagging Keys (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).",
						MarkdownDescription: "Assigns one or more tags to the KMS key. Use this parameter to tag the KMS key when it is created. To tag an existing KMS key, use the TagResource operation. Do not include confidential or sensitive information in this field. This field may be displayed in plaintext in CloudTrail logs and other output. Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see ABAC for KMS (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the Key Management Service Developer Guide. To use this parameter, you must have kms:TagResource (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html) permission in an IAM policy. Each tag consists of a tag key and a tag value. Both the tag key and the tag value are required, but the tag value can be an empty (null) string. You cannot have more than one tag on a KMS key with the same tag key. If you specify an existing tag key with a different tag value, KMS replaces the current tag value with the specified one. When you add tags to an Amazon Web Services resource, Amazon Web Services generates a cost allocation report with usage and costs aggregated by tags. Tags can also be used to control access to a KMS key. For details, see Tagging Keys (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"tag_key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tag_value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kms_services_k8s_aws_key_v1alpha1_manifest")

	var model KmsServicesK8SAwsKeyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kms.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Key")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
