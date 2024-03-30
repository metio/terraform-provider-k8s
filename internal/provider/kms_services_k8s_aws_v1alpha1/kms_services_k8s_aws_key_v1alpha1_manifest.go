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
						Description:         "A flag to indicate whether to bypass the key policy lockout safety check.Setting this value to true increases the risk that the KMS key becomes unmanageable.Do not set this value to true indiscriminately.For more information, refer to the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)section in the Key Management Service Developer Guide .Use this parameter only when you include a policy in the request and youintend to prevent the principal that is making the request from making asubsequent PutKeyPolicy request on the KMS key.The default value is false.",
						MarkdownDescription: "A flag to indicate whether to bypass the key policy lockout safety check.Setting this value to true increases the risk that the KMS key becomes unmanageable.Do not set this value to true indiscriminately.For more information, refer to the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)section in the Key Management Service Developer Guide .Use this parameter only when you include a policy in the request and youintend to prevent the principal that is making the request from making asubsequent PutKeyPolicy request on the KMS key.The default value is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_key_store_id": schema.StringAttribute{
						Description:         "Creates the KMS key in the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)and the key material in its associated CloudHSM cluster. To create a KMSkey in a custom key store, you must also specify the Origin parameter witha value of AWS_CLOUDHSM. The CloudHSM cluster that is associated with thecustom key store must have at least two active HSMs, each in a differentAvailability Zone in the Region.This parameter is valid only for symmetric encryption KMS keys in a singleRegion. You cannot create any other type of KMS key in a custom key store.To find the ID of a custom key store, use the DescribeCustomKeyStores operation.The response includes the custom key store ID and the ID of the CloudHSMcluster.This operation is part of the custom key store feature (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)feature in KMS, which combines the convenience and extensive integrationof KMS with the isolation and control of a single-tenant key store.",
						MarkdownDescription: "Creates the KMS key in the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)and the key material in its associated CloudHSM cluster. To create a KMSkey in a custom key store, you must also specify the Origin parameter witha value of AWS_CLOUDHSM. The CloudHSM cluster that is associated with thecustom key store must have at least two active HSMs, each in a differentAvailability Zone in the Region.This parameter is valid only for symmetric encryption KMS keys in a singleRegion. You cannot create any other type of KMS key in a custom key store.To find the ID of a custom key store, use the DescribeCustomKeyStores operation.The response includes the custom key store ID and the ID of the CloudHSMcluster.This operation is part of the custom key store feature (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)feature in KMS, which combines the convenience and extensive integrationof KMS with the isolation and control of a single-tenant key store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A description of the KMS key.Use a description that helps you decide whether the KMS key is appropriatefor a task. The default value is an empty string (no description).To set or change the description after the key is created, use UpdateKeyDescription.",
						MarkdownDescription: "A description of the KMS key.Use a description that helps you decide whether the KMS key is appropriatefor a task. The default value is an empty string (no description).To set or change the description after the key is created, use UpdateKeyDescription.",
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
						Description:         "Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT,creates a KMS key with a 256-bit AES-GCM key that is used for encryptionand decryption, except in China Regions, where it creates a 128-bit symmetrickey that uses SM4 encryption. For help choosing a key spec for your KMS key,see Choosing a KMS key type (https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html#symm-asymm-choose)in the Key Management Service Developer Guide .The KeySpec determines whether the KMS key contains a symmetric key or anasymmetric key pair. It also determines the cryptographic algorithms thatthe KMS key supports. You can't change the KeySpec after the KMS key is created.To further restrict the algorithms that can be used with the KMS key, usea condition key in its key policy or IAM policy. For more information, seekms:EncryptionAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm),kms:MacAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-mac-algorithm)or kms:Signing Algorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm)in the Key Management Service Developer Guide .Amazon Web Services services that are integrated with KMS (http://aws.amazon.com/kms/features/#AWS_Service_Integration)use symmetric encryption KMS keys to protect your data. These services donot support asymmetric KMS keys or HMAC KMS keys.KMS supports the following key specs for KMS keys:   * Symmetric encryption key (default) SYMMETRIC_DEFAULT   * HMAC keys (symmetric) HMAC_224 HMAC_256 HMAC_384 HMAC_512   * Asymmetric RSA key pairs RSA_2048 RSA_3072 RSA_4096   * Asymmetric NIST-recommended elliptic curve key pairs ECC_NIST_P256 (secp256r1)   ECC_NIST_P384 (secp384r1) ECC_NIST_P521 (secp521r1)   * Other asymmetric elliptic curve key pairs ECC_SECG_P256K1 (secp256k1),   commonly used for cryptocurrencies.   * SM2 key pairs (China Regions only) SM2",
						MarkdownDescription: "Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT,creates a KMS key with a 256-bit AES-GCM key that is used for encryptionand decryption, except in China Regions, where it creates a 128-bit symmetrickey that uses SM4 encryption. For help choosing a key spec for your KMS key,see Choosing a KMS key type (https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html#symm-asymm-choose)in the Key Management Service Developer Guide .The KeySpec determines whether the KMS key contains a symmetric key or anasymmetric key pair. It also determines the cryptographic algorithms thatthe KMS key supports. You can't change the KeySpec after the KMS key is created.To further restrict the algorithms that can be used with the KMS key, usea condition key in its key policy or IAM policy. For more information, seekms:EncryptionAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm),kms:MacAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-mac-algorithm)or kms:Signing Algorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm)in the Key Management Service Developer Guide .Amazon Web Services services that are integrated with KMS (http://aws.amazon.com/kms/features/#AWS_Service_Integration)use symmetric encryption KMS keys to protect your data. These services donot support asymmetric KMS keys or HMAC KMS keys.KMS supports the following key specs for KMS keys:   * Symmetric encryption key (default) SYMMETRIC_DEFAULT   * HMAC keys (symmetric) HMAC_224 HMAC_256 HMAC_384 HMAC_512   * Asymmetric RSA key pairs RSA_2048 RSA_3072 RSA_4096   * Asymmetric NIST-recommended elliptic curve key pairs ECC_NIST_P256 (secp256r1)   ECC_NIST_P384 (secp384r1) ECC_NIST_P521 (secp521r1)   * Other asymmetric elliptic curve key pairs ECC_SECG_P256K1 (secp256k1),   commonly used for cryptocurrencies.   * SM2 key pairs (China Regions only) SM2",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_usage": schema.StringAttribute{
						Description:         "Determines the cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)for which you can use the KMS key. The default value is ENCRYPT_DECRYPT.This parameter is optional when you are creating a symmetric encryption KMSkey; otherwise, it is required. You can't change the KeyUsage value afterthe KMS key is created.Select only one valid value.   * For symmetric encryption KMS keys, omit the parameter or specify ENCRYPT_DECRYPT.   * For HMAC KMS keys (symmetric), specify GENERATE_VERIFY_MAC.   * For asymmetric KMS keys with RSA key material, specify ENCRYPT_DECRYPT   or SIGN_VERIFY.   * For asymmetric KMS keys with ECC key material, specify SIGN_VERIFY.   * For asymmetric KMS keys with SM2 key material (China Regions only),   specify ENCRYPT_DECRYPT or SIGN_VERIFY.",
						MarkdownDescription: "Determines the cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)for which you can use the KMS key. The default value is ENCRYPT_DECRYPT.This parameter is optional when you are creating a symmetric encryption KMSkey; otherwise, it is required. You can't change the KeyUsage value afterthe KMS key is created.Select only one valid value.   * For symmetric encryption KMS keys, omit the parameter or specify ENCRYPT_DECRYPT.   * For HMAC KMS keys (symmetric), specify GENERATE_VERIFY_MAC.   * For asymmetric KMS keys with RSA key material, specify ENCRYPT_DECRYPT   or SIGN_VERIFY.   * For asymmetric KMS keys with ECC key material, specify SIGN_VERIFY.   * For asymmetric KMS keys with SM2 key material (China Regions only),   specify ENCRYPT_DECRYPT or SIGN_VERIFY.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"multi_region": schema.BoolAttribute{
						Description:         "Creates a multi-Region primary key that you can replicate into other AmazonWeb Services Regions. You cannot change this value after you create the KMSkey.For a multi-Region key, set this parameter to True. For a single-Region KMSkey, omit this parameter or set it to False. The default value is False.This operation supports multi-Region keys, an KMS feature that lets you createmultiple interoperable KMS keys in different Amazon Web Services Regions.Because these KMS keys have the same key ID, key material, and other metadata,you can use them interchangeably to encrypt data in one Amazon Web ServicesRegion and decrypt it in a different Amazon Web Services Region without re-encryptingthe data or making a cross-Region call. For more information about multi-Regionkeys, see Multi-Region keys in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)in the Key Management Service Developer Guide.This value creates a primary key, not a replica. To create a replica key,use the ReplicateKey operation.You can create a multi-Region version of a symmetric encryption KMS key,an HMAC KMS key, an asymmetric KMS key, or a KMS key with imported key material.However, you cannot create a multi-Region key in a custom key store.",
						MarkdownDescription: "Creates a multi-Region primary key that you can replicate into other AmazonWeb Services Regions. You cannot change this value after you create the KMSkey.For a multi-Region key, set this parameter to True. For a single-Region KMSkey, omit this parameter or set it to False. The default value is False.This operation supports multi-Region keys, an KMS feature that lets you createmultiple interoperable KMS keys in different Amazon Web Services Regions.Because these KMS keys have the same key ID, key material, and other metadata,you can use them interchangeably to encrypt data in one Amazon Web ServicesRegion and decrypt it in a different Amazon Web Services Region without re-encryptingthe data or making a cross-Region call. For more information about multi-Regionkeys, see Multi-Region keys in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)in the Key Management Service Developer Guide.This value creates a primary key, not a replica. To create a replica key,use the ReplicateKey operation.You can create a multi-Region version of a symmetric encryption KMS key,an HMAC KMS key, an asymmetric KMS key, or a KMS key with imported key material.However, you cannot create a multi-Region key in a custom key store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"origin": schema.StringAttribute{
						Description:         "The source of the key material for the KMS key. You cannot change the originafter you create the KMS key. The default is AWS_KMS, which means that KMScreates the key material.To create a KMS key with no key material (for imported key material), setthe value to EXTERNAL. For more information about importing key materialinto KMS, see Importing Key Material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html)in the Key Management Service Developer Guide. This value is valid only forsymmetric encryption KMS keys.To create a KMS key in an KMS custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)and create its key material in the associated CloudHSM cluster, set thisvalue to AWS_CLOUDHSM. You must also use the CustomKeyStoreId parameter toidentify the custom key store. This value is valid only for symmetric encryptionKMS keys.",
						MarkdownDescription: "The source of the key material for the KMS key. You cannot change the originafter you create the KMS key. The default is AWS_KMS, which means that KMScreates the key material.To create a KMS key with no key material (for imported key material), setthe value to EXTERNAL. For more information about importing key materialinto KMS, see Importing Key Material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html)in the Key Management Service Developer Guide. This value is valid only forsymmetric encryption KMS keys.To create a KMS key in an KMS custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)and create its key material in the associated CloudHSM cluster, set thisvalue to AWS_CLOUDHSM. You must also use the CustomKeyStoreId parameter toidentify the custom key store. This value is valid only for symmetric encryptionKMS keys.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy": schema.StringAttribute{
						Description:         "The key policy to attach to the KMS key. If you do not specify a key policy,KMS attaches a default key policy to the KMS key. For more information, seeDefault key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default)in the Key Management Service Developer Guide.If you provide a key policy, it must meet the following criteria:   * If you don't set BypassPolicyLockoutSafetyCheck to True, the key policy   must allow the principal that is making the CreateKey request to make   a subsequent PutKeyPolicy request on the KMS key. This reduces the risk   that the KMS key becomes unmanageable. For more information, refer to   the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)   section of the Key Management Service Developer Guide .   * Each statement in the key policy must contain one or more principals.   The principals in the key policy must exist and be visible to KMS. When   you create a new Amazon Web Services principal (for example, an IAM user   or role), you might need to enforce a delay before including the new principal   in a key policy because the new principal might not be immediately visible   to KMS. For more information, see Changes that I make are not always immediately   visible (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)   in the Amazon Web Services Identity and Access Management User Guide.A key policy document can include only the following characters:   * Printable ASCII characters from the space character (u0020) through   the end of the ASCII character range.   * Printable characters in the Basic Latin and Latin-1 Supplement character   set (through u00FF).   * The tab (u0009), line feed (u000A), and carriage return (u000D) special   charactersFor information about key policies, see Key policies in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)in the Key Management Service Developer Guide. For help writing and formattinga JSON policy document, see the IAM JSON Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html)in the Identity and Access Management User Guide .",
						MarkdownDescription: "The key policy to attach to the KMS key. If you do not specify a key policy,KMS attaches a default key policy to the KMS key. For more information, seeDefault key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default)in the Key Management Service Developer Guide.If you provide a key policy, it must meet the following criteria:   * If you don't set BypassPolicyLockoutSafetyCheck to True, the key policy   must allow the principal that is making the CreateKey request to make   a subsequent PutKeyPolicy request on the KMS key. This reduces the risk   that the KMS key becomes unmanageable. For more information, refer to   the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)   section of the Key Management Service Developer Guide .   * Each statement in the key policy must contain one or more principals.   The principals in the key policy must exist and be visible to KMS. When   you create a new Amazon Web Services principal (for example, an IAM user   or role), you might need to enforce a delay before including the new principal   in a key policy because the new principal might not be immediately visible   to KMS. For more information, see Changes that I make are not always immediately   visible (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)   in the Amazon Web Services Identity and Access Management User Guide.A key policy document can include only the following characters:   * Printable ASCII characters from the space character (u0020) through   the end of the ASCII character range.   * Printable characters in the Basic Latin and Latin-1 Supplement character   set (through u00FF).   * The tab (u0009), line feed (u000A), and carriage return (u000D) special   charactersFor information about key policies, see Key policies in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)in the Key Management Service Developer Guide. For help writing and formattinga JSON policy document, see the IAM JSON Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html)in the Identity and Access Management User Guide .",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Assigns one or more tags to the KMS key. Use this parameter to tag the KMSkey when it is created. To tag an existing KMS key, use the TagResource operation.Tagging or untagging a KMS key can allow or deny permission to the KMS key.For details, see ABAC in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html)in the Key Management Service Developer Guide.To use this parameter, you must have kms:TagResource (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)permission in an IAM policy.Each tag consists of a tag key and a tag value. Both the tag key and thetag value are required, but the tag value can be an empty (null) string.You cannot have more than one tag on a KMS key with the same tag key. Ifyou specify an existing tag key with a different tag value, KMS replacesthe current tag value with the specified one.When you add tags to an Amazon Web Services resource, Amazon Web Servicesgenerates a cost allocation report with usage and costs aggregated by tags.Tags can also be used to control access to a KMS key. For details, see TaggingKeys (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).",
						MarkdownDescription: "Assigns one or more tags to the KMS key. Use this parameter to tag the KMSkey when it is created. To tag an existing KMS key, use the TagResource operation.Tagging or untagging a KMS key can allow or deny permission to the KMS key.For details, see ABAC in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html)in the Key Management Service Developer Guide.To use this parameter, you must have kms:TagResource (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)permission in an IAM policy.Each tag consists of a tag key and a tag value. Both the tag key and thetag value are required, but the tag value can be an empty (null) string.You cannot have more than one tag on a KMS key with the same tag key. Ifyou specify an existing tag key with a different tag value, KMS replacesthe current tag value with the specified one.When you add tags to an Amazon Web Services resource, Amazon Web Servicesgenerates a cost allocation report with usage and costs aggregated by tags.Tags can also be used to control access to a KMS key. For details, see TaggingKeys (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).",
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
