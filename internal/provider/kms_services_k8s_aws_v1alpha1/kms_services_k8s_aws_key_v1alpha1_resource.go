/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kms_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &KmsServicesK8SAwsKeyV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &KmsServicesK8SAwsKeyV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &KmsServicesK8SAwsKeyV1Alpha1Resource{}
)

func NewKmsServicesK8SAwsKeyV1Alpha1Resource() resource.Resource {
	return &KmsServicesK8SAwsKeyV1Alpha1Resource{}
}

type KmsServicesK8SAwsKeyV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type KmsServicesK8SAwsKeyV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kms_services_k8s_aws_key_v1alpha1"
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Key is the Schema for the Keys API",
		MarkdownDescription: "Key is the Schema for the Keys API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
						Description:         "A flag to indicate whether to bypass the key policy lockout safety check.  Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.  For more information, refer to the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the Key Management Service Developer Guide .  Use this parameter only when you include a policy in the request and you intend to prevent the principal that is making the request from making a subsequent PutKeyPolicy request on the KMS key.  The default value is false.",
						MarkdownDescription: "A flag to indicate whether to bypass the key policy lockout safety check.  Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.  For more information, refer to the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section in the Key Management Service Developer Guide .  Use this parameter only when you include a policy in the request and you intend to prevent the principal that is making the request from making a subsequent PutKeyPolicy request on the KMS key.  The default value is false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_key_store_id": schema.StringAttribute{
						Description:         "Creates the KMS key in the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html) and the key material in its associated CloudHSM cluster. To create a KMS key in a custom key store, you must also specify the Origin parameter with a value of AWS_CLOUDHSM. The CloudHSM cluster that is associated with the custom key store must have at least two active HSMs, each in a different Availability Zone in the Region.  This parameter is valid only for symmetric encryption KMS keys in a single Region. You cannot create any other type of KMS key in a custom key store.  To find the ID of a custom key store, use the DescribeCustomKeyStores operation.  The response includes the custom key store ID and the ID of the CloudHSM cluster.  This operation is part of the custom key store feature (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html) feature in KMS, which combines the convenience and extensive integration of KMS with the isolation and control of a single-tenant key store.",
						MarkdownDescription: "Creates the KMS key in the specified custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html) and the key material in its associated CloudHSM cluster. To create a KMS key in a custom key store, you must also specify the Origin parameter with a value of AWS_CLOUDHSM. The CloudHSM cluster that is associated with the custom key store must have at least two active HSMs, each in a different Availability Zone in the Region.  This parameter is valid only for symmetric encryption KMS keys in a single Region. You cannot create any other type of KMS key in a custom key store.  To find the ID of a custom key store, use the DescribeCustomKeyStores operation.  The response includes the custom key store ID and the ID of the CloudHSM cluster.  This operation is part of the custom key store feature (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html) feature in KMS, which combines the convenience and extensive integration of KMS with the isolation and control of a single-tenant key store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A description of the KMS key.  Use a description that helps you decide whether the KMS key is appropriate for a task. The default value is an empty string (no description).  To set or change the description after the key is created, use UpdateKeyDescription.",
						MarkdownDescription: "A description of the KMS key.  Use a description that helps you decide whether the KMS key is appropriate for a task. The default value is an empty string (no description).  To set or change the description after the key is created, use UpdateKeyDescription.",
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
						Description:         "Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT, creates a KMS key with a 256-bit AES-GCM key that is used for encryption and decryption, except in China Regions, where it creates a 128-bit symmetric key that uses SM4 encryption. For help choosing a key spec for your KMS key, see Choosing a KMS key type (https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html#symm-asymm-choose) in the Key Management Service Developer Guide .  The KeySpec determines whether the KMS key contains a symmetric key or an asymmetric key pair. It also determines the cryptographic algorithms that the KMS key supports. You can't change the KeySpec after the KMS key is created. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see kms:EncryptionAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm), kms:MacAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-mac-algorithm) or kms:Signing Algorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm) in the Key Management Service Developer Guide .  Amazon Web Services services that are integrated with KMS (http://aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support asymmetric KMS keys or HMAC KMS keys.  KMS supports the following key specs for KMS keys:  * Symmetric encryption key (default) SYMMETRIC_DEFAULT  * HMAC keys (symmetric) HMAC_224 HMAC_256 HMAC_384 HMAC_512  * Asymmetric RSA key pairs RSA_2048 RSA_3072 RSA_4096  * Asymmetric NIST-recommended elliptic curve key pairs ECC_NIST_P256 (secp256r1) ECC_NIST_P384 (secp384r1) ECC_NIST_P521 (secp521r1)  * Other asymmetric elliptic curve key pairs ECC_SECG_P256K1 (secp256k1), commonly used for cryptocurrencies.  * SM2 key pairs (China Regions only) SM2",
						MarkdownDescription: "Specifies the type of KMS key to create. The default value, SYMMETRIC_DEFAULT, creates a KMS key with a 256-bit AES-GCM key that is used for encryption and decryption, except in China Regions, where it creates a 128-bit symmetric key that uses SM4 encryption. For help choosing a key spec for your KMS key, see Choosing a KMS key type (https://docs.aws.amazon.com/kms/latest/developerguide/key-types.html#symm-asymm-choose) in the Key Management Service Developer Guide .  The KeySpec determines whether the KMS key contains a symmetric key or an asymmetric key pair. It also determines the cryptographic algorithms that the KMS key supports. You can't change the KeySpec after the KMS key is created. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see kms:EncryptionAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm), kms:MacAlgorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-mac-algorithm) or kms:Signing Algorithm (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm) in the Key Management Service Developer Guide .  Amazon Web Services services that are integrated with KMS (http://aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support asymmetric KMS keys or HMAC KMS keys.  KMS supports the following key specs for KMS keys:  * Symmetric encryption key (default) SYMMETRIC_DEFAULT  * HMAC keys (symmetric) HMAC_224 HMAC_256 HMAC_384 HMAC_512  * Asymmetric RSA key pairs RSA_2048 RSA_3072 RSA_4096  * Asymmetric NIST-recommended elliptic curve key pairs ECC_NIST_P256 (secp256r1) ECC_NIST_P384 (secp384r1) ECC_NIST_P521 (secp521r1)  * Other asymmetric elliptic curve key pairs ECC_SECG_P256K1 (secp256k1), commonly used for cryptocurrencies.  * SM2 key pairs (China Regions only) SM2",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_usage": schema.StringAttribute{
						Description:         "Determines the cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is ENCRYPT_DECRYPT. This parameter is optional when you are creating a symmetric encryption KMS key; otherwise, it is required. You can't change the KeyUsage value after the KMS key is created.  Select only one valid value.  * For symmetric encryption KMS keys, omit the parameter or specify ENCRYPT_DECRYPT.  * For HMAC KMS keys (symmetric), specify GENERATE_VERIFY_MAC.  * For asymmetric KMS keys with RSA key material, specify ENCRYPT_DECRYPT or SIGN_VERIFY.  * For asymmetric KMS keys with ECC key material, specify SIGN_VERIFY.  * For asymmetric KMS keys with SM2 key material (China Regions only), specify ENCRYPT_DECRYPT or SIGN_VERIFY.",
						MarkdownDescription: "Determines the cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is ENCRYPT_DECRYPT. This parameter is optional when you are creating a symmetric encryption KMS key; otherwise, it is required. You can't change the KeyUsage value after the KMS key is created.  Select only one valid value.  * For symmetric encryption KMS keys, omit the parameter or specify ENCRYPT_DECRYPT.  * For HMAC KMS keys (symmetric), specify GENERATE_VERIFY_MAC.  * For asymmetric KMS keys with RSA key material, specify ENCRYPT_DECRYPT or SIGN_VERIFY.  * For asymmetric KMS keys with ECC key material, specify SIGN_VERIFY.  * For asymmetric KMS keys with SM2 key material (China Regions only), specify ENCRYPT_DECRYPT or SIGN_VERIFY.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"multi_region": schema.BoolAttribute{
						Description:         "Creates a multi-Region primary key that you can replicate into other Amazon Web Services Regions. You cannot change this value after you create the KMS key.  For a multi-Region key, set this parameter to True. For a single-Region KMS key, omit this parameter or set it to False. The default value is False.  This operation supports multi-Region keys, an KMS feature that lets you create multiple interoperable KMS keys in different Amazon Web Services Regions. Because these KMS keys have the same key ID, key material, and other metadata, you can use them interchangeably to encrypt data in one Amazon Web Services Region and decrypt it in a different Amazon Web Services Region without re-encrypting the data or making a cross-Region call. For more information about multi-Region keys, see Multi-Region keys in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the Key Management Service Developer Guide.  This value creates a primary key, not a replica. To create a replica key, use the ReplicateKey operation.  You can create a multi-Region version of a symmetric encryption KMS key, an HMAC KMS key, an asymmetric KMS key, or a KMS key with imported key material. However, you cannot create a multi-Region key in a custom key store.",
						MarkdownDescription: "Creates a multi-Region primary key that you can replicate into other Amazon Web Services Regions. You cannot change this value after you create the KMS key.  For a multi-Region key, set this parameter to True. For a single-Region KMS key, omit this parameter or set it to False. The default value is False.  This operation supports multi-Region keys, an KMS feature that lets you create multiple interoperable KMS keys in different Amazon Web Services Regions. Because these KMS keys have the same key ID, key material, and other metadata, you can use them interchangeably to encrypt data in one Amazon Web Services Region and decrypt it in a different Amazon Web Services Region without re-encrypting the data or making a cross-Region call. For more information about multi-Region keys, see Multi-Region keys in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the Key Management Service Developer Guide.  This value creates a primary key, not a replica. To create a replica key, use the ReplicateKey operation.  You can create a multi-Region version of a symmetric encryption KMS key, an HMAC KMS key, an asymmetric KMS key, or a KMS key with imported key material. However, you cannot create a multi-Region key in a custom key store.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"origin": schema.StringAttribute{
						Description:         "The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that KMS creates the key material.  To create a KMS key with no key material (for imported key material), set the value to EXTERNAL. For more information about importing key material into KMS, see Importing Key Material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in the Key Management Service Developer Guide. This value is valid only for symmetric encryption KMS keys.  To create a KMS key in an KMS custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html) and create its key material in the associated CloudHSM cluster, set this value to AWS_CLOUDHSM. You must also use the CustomKeyStoreId parameter to identify the custom key store. This value is valid only for symmetric encryption KMS keys.",
						MarkdownDescription: "The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that KMS creates the key material.  To create a KMS key with no key material (for imported key material), set the value to EXTERNAL. For more information about importing key material into KMS, see Importing Key Material (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in the Key Management Service Developer Guide. This value is valid only for symmetric encryption KMS keys.  To create a KMS key in an KMS custom key store (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html) and create its key material in the associated CloudHSM cluster, set this value to AWS_CLOUDHSM. You must also use the CustomKeyStoreId parameter to identify the custom key store. This value is valid only for symmetric encryption KMS keys.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy": schema.StringAttribute{
						Description:         "The key policy to attach to the KMS key. If you do not specify a key policy, KMS attaches a default key policy to the KMS key. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the Key Management Service Developer Guide.  If you provide a key policy, it must meet the following criteria:  * If you don't set BypassPolicyLockoutSafetyCheck to True, the key policy must allow the principal that is making the CreateKey request to make a subsequent PutKeyPolicy request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, refer to the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section of the Key Management Service Developer Guide .  * Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to KMS. When you create a new Amazon Web Services principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to KMS. For more information, see Changes that I make are not always immediately visible (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the Amazon Web Services Identity and Access Management User Guide.  A key policy document can include only the following characters:  * Printable ASCII characters from the space character (u0020) through the end of the ASCII character range.  * Printable characters in the Basic Latin and Latin-1 Supplement character set (through u00FF).  * The tab (u0009), line feed (u000A), and carriage return (u000D) special characters  For information about key policies, see Key policies in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the Key Management Service Developer Guide. For help writing and formatting a JSON policy document, see the IAM JSON Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in the Identity and Access Management User Guide .",
						MarkdownDescription: "The key policy to attach to the KMS key. If you do not specify a key policy, KMS attaches a default key policy to the KMS key. For more information, see Default key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the Key Management Service Developer Guide.  If you provide a key policy, it must meet the following criteria:  * If you don't set BypassPolicyLockoutSafetyCheck to True, the key policy must allow the principal that is making the CreateKey request to make a subsequent PutKeyPolicy request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, refer to the scenario in the Default Key Policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) section of the Key Management Service Developer Guide .  * Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to KMS. When you create a new Amazon Web Services principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to KMS. For more information, see Changes that I make are not always immediately visible (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the Amazon Web Services Identity and Access Management User Guide.  A key policy document can include only the following characters:  * Printable ASCII characters from the space character (u0020) through the end of the ASCII character range.  * Printable characters in the Basic Latin and Latin-1 Supplement character set (through u00FF).  * The tab (u0009), line feed (u000A), and carriage return (u000D) special characters  For information about key policies, see Key policies in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html) in the Key Management Service Developer Guide. For help writing and formatting a JSON policy document, see the IAM JSON Policy Reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in the Identity and Access Management User Guide .",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Assigns one or more tags to the KMS key. Use this parameter to tag the KMS key when it is created. To tag an existing KMS key, use the TagResource operation.  Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see ABAC in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the Key Management Service Developer Guide.  To use this parameter, you must have kms:TagResource (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html) permission in an IAM policy.  Each tag consists of a tag key and a tag value. Both the tag key and the tag value are required, but the tag value can be an empty (null) string. You cannot have more than one tag on a KMS key with the same tag key. If you specify an existing tag key with a different tag value, KMS replaces the current tag value with the specified one.  When you add tags to an Amazon Web Services resource, Amazon Web Services generates a cost allocation report with usage and costs aggregated by tags. Tags can also be used to control access to a KMS key. For details, see Tagging Keys (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).",
						MarkdownDescription: "Assigns one or more tags to the KMS key. Use this parameter to tag the KMS key when it is created. To tag an existing KMS key, use the TagResource operation.  Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see ABAC in KMS (https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the Key Management Service Developer Guide.  To use this parameter, you must have kms:TagResource (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html) permission in an IAM policy.  Each tag consists of a tag key and a tag value. Both the tag key and the tag value are required, but the tag value can be an empty (null) string. You cannot have more than one tag on a KMS key with the same tag key. If you specify an existing tag key with a different tag value, KMS replaces the current tag value with the specified one.  When you add tags to an Amazon Web Services resource, Amazon Web Services generates a cost allocation report with usage and costs aggregated by tags. Tags can also be used to control access to a KMS key. For details, see Tagging Keys (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html).",
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

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kms_services_k8s_aws_key_v1alpha1")

	var model KmsServicesK8SAwsKeyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("kms.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Key")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kms.services.k8s.aws", Version: "v1alpha1", Resource: "keys"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KmsServicesK8SAwsKeyV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kms_services_k8s_aws_key_v1alpha1")

	var data KmsServicesK8SAwsKeyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kms.services.k8s.aws", Version: "v1alpha1", Resource: "keys"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KmsServicesK8SAwsKeyV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kms_services_k8s_aws_key_v1alpha1")

	var model KmsServicesK8SAwsKeyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kms.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Key")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kms.services.k8s.aws", Version: "v1alpha1", Resource: "keys"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KmsServicesK8SAwsKeyV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kms_services_k8s_aws_key_v1alpha1")

	var data KmsServicesK8SAwsKeyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kms.services.k8s.aws", Version: "v1alpha1", Resource: "keys"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *KmsServicesK8SAwsKeyV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
