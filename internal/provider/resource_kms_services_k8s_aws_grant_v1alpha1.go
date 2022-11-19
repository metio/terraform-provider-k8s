/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type KmsServicesK8SAwsGrantV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*KmsServicesK8SAwsGrantV1Alpha1Resource)(nil)
)

type KmsServicesK8SAwsGrantV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KmsServicesK8SAwsGrantV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Constraints *struct {
			EncryptionContextEquals *map[string]string `tfsdk:"encryption_context_equals" yaml:"encryptionContextEquals,omitempty"`

			EncryptionContextSubset *map[string]string `tfsdk:"encryption_context_subset" yaml:"encryptionContextSubset,omitempty"`
		} `tfsdk:"constraints" yaml:"constraints,omitempty"`

		GrantTokens *[]string `tfsdk:"grant_tokens" yaml:"grantTokens,omitempty"`

		GranteePrincipal *string `tfsdk:"grantee_principal" yaml:"granteePrincipal,omitempty"`

		KeyID *string `tfsdk:"key_id" yaml:"keyID,omitempty"`

		KeyRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"key_ref" yaml:"keyRef,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Operations *[]string `tfsdk:"operations" yaml:"operations,omitempty"`

		RetiringPrincipal *string `tfsdk:"retiring_principal" yaml:"retiringPrincipal,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKmsServicesK8SAwsGrantV1Alpha1Resource() resource.Resource {
	return &KmsServicesK8SAwsGrantV1Alpha1Resource{}
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kms_services_k8s_aws_grant_v1alpha1"
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Grant is the Schema for the Grants API",
		MarkdownDescription: "Grant is the Schema for the Grants API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "GrantSpec defines the desired state of Grant.",
				MarkdownDescription: "GrantSpec defines the desired state of Grant.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"constraints": {
						Description:         "Specifies a grant constraint.  KMS supports the EncryptionContextEquals and EncryptionContextSubset grant constraints. Each constraint value can include up to 8 encryption context pairs. The encryption context value in each constraint cannot exceed 384 characters. For information about grant constraints, see Using grant constraints (https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-constraints) in the Key Management Service Developer Guide. For more information about encryption context, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the Key Management Service Developer Guide .  The encryption context grant constraints allow the permissions in the grant only when the encryption context in the request matches (EncryptionContextEquals) or includes (EncryptionContextSubset) the encryption context specified in this structure.  The encryption context grant constraints are supported only on grant operations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations) that include an EncryptionContext parameter, such as cryptographic operations on symmetric encryption KMS keys. Grants with grant constraints can include the DescribeKey and RetireGrant operations, but the constraint doesn't apply to these operations. If a grant with a grant constraint includes the CreateGrant operation, the constraint requires that any grants created with the CreateGrant permission have an equally strict or stricter encryption context constraint.  You cannot use an encryption context grant constraint for cryptographic operations with asymmetric KMS keys or HMAC KMS keys. These keys don't support an encryption context.",
						MarkdownDescription: "Specifies a grant constraint.  KMS supports the EncryptionContextEquals and EncryptionContextSubset grant constraints. Each constraint value can include up to 8 encryption context pairs. The encryption context value in each constraint cannot exceed 384 characters. For information about grant constraints, see Using grant constraints (https://docs.aws.amazon.com/kms/latest/developerguide/create-grant-overview.html#grant-constraints) in the Key Management Service Developer Guide. For more information about encryption context, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context) in the Key Management Service Developer Guide .  The encryption context grant constraints allow the permissions in the grant only when the encryption context in the request matches (EncryptionContextEquals) or includes (EncryptionContextSubset) the encryption context specified in this structure.  The encryption context grant constraints are supported only on grant operations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations) that include an EncryptionContext parameter, such as cryptographic operations on symmetric encryption KMS keys. Grants with grant constraints can include the DescribeKey and RetireGrant operations, but the constraint doesn't apply to these operations. If a grant with a grant constraint includes the CreateGrant operation, the constraint requires that any grants created with the CreateGrant permission have an equally strict or stricter encryption context constraint.  You cannot use an encryption context grant constraint for cryptographic operations with asymmetric KMS keys or HMAC KMS keys. These keys don't support an encryption context.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"encryption_context_equals": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"encryption_context_subset": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grant_tokens": {
						Description:         "A list of grant tokens.  Use a grant token when your permission to call this operation comes from a new grant that has not yet achieved eventual consistency. For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token) and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token) in the Key Management Service Developer Guide.",
						MarkdownDescription: "A list of grant tokens.  Use a grant token when your permission to call this operation comes from a new grant that has not yet achieved eventual consistency. For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token) and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token) in the Key Management Service Developer Guide.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grantee_principal": {
						Description:         "The identity that gets the permissions specified in the grant.  To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an Amazon Web Services principal. Valid Amazon Web Services principals include Amazon Web Services accounts (root), IAM users, IAM roles, federated users, and assumed role users. For examples of the ARN syntax to use for specifying a principal, see Amazon Web Services Identity and Access Management (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam) in the Example ARNs section of the Amazon Web Services General Reference.",
						MarkdownDescription: "The identity that gets the permissions specified in the grant.  To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an Amazon Web Services principal. Valid Amazon Web Services principals include Amazon Web Services accounts (root), IAM users, IAM roles, federated users, and assumed role users. For examples of the ARN syntax to use for specifying a principal, see Amazon Web Services Identity and Access Management (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam) in the Example ARNs section of the Amazon Web Services General Reference.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"key_id": {
						Description:         "Identifies the KMS key for the grant. The grant gives principals permission to use this KMS key.  Specify the key ID or key ARN of the KMS key. To specify a KMS key in a different Amazon Web Services account, you must use the key ARN.  For example:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.",
						MarkdownDescription: "Identifies the KMS key for the grant. The grant gives principals permission to use this KMS key.  Specify the key ID or key ARN of the KMS key. To specify a KMS key in a different Amazon Web Services account, you must use the key ARN.  For example:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"key_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "A friendly name for the grant. Use this value to prevent the unintended creation of duplicate grants when retrying this request.  When this value is absent, all CreateGrant requests result in a new grant with a unique GrantId even if all the supplied parameters are identical. This can result in unintended duplicates when you retry the CreateGrant request.  When this value is present, you can retry a CreateGrant request with identical parameters; if the grant already exists, the original GrantId is returned without creating a new grant. Note that the returned grant token is unique with every CreateGrant request, even when a duplicate GrantId is returned. All grant tokens for the same grant ID can be used interchangeably.",
						MarkdownDescription: "A friendly name for the grant. Use this value to prevent the unintended creation of duplicate grants when retrying this request.  When this value is absent, all CreateGrant requests result in a new grant with a unique GrantId even if all the supplied parameters are identical. This can result in unintended duplicates when you retry the CreateGrant request.  When this value is present, you can retry a CreateGrant request with identical parameters; if the grant already exists, the original GrantId is returned without creating a new grant. Note that the returned grant token is unique with every CreateGrant request, even when a duplicate GrantId is returned. All grant tokens for the same grant ID can be used interchangeably.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"operations": {
						Description:         "A list of operations that the grant permits.  This list must include only operations that are permitted in a grant. Also, the operation must be supported on the KMS key. For example, you cannot create a grant for a symmetric encryption KMS key that allows the Sign operation, or a grant for an asymmetric KMS key that allows the GenerateDataKey operation. If you try, KMS returns a ValidationError exception. For details, see Grant operations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations) in the Key Management Service Developer Guide.",
						MarkdownDescription: "A list of operations that the grant permits.  This list must include only operations that are permitted in a grant. Also, the operation must be supported on the KMS key. For example, you cannot create a grant for a symmetric encryption KMS key that allows the Sign operation, or a grant for an asymmetric KMS key that allows the GenerateDataKey operation. If you try, KMS returns a ValidationError exception. For details, see Grant operations (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-grant-operations) in the Key Management Service Developer Guide.",

						Type: types.ListType{ElemType: types.StringType},

						Required: true,
						Optional: false,
						Computed: false,
					},

					"retiring_principal": {
						Description:         "The principal that has permission to use the RetireGrant operation to retire the grant.  To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an Amazon Web Services principal. Valid Amazon Web Services principals include Amazon Web Services accounts (root), IAM users, federated users, and assumed role users. For examples of the ARN syntax to use for specifying a principal, see Amazon Web Services Identity and Access Management (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam) in the Example ARNs section of the Amazon Web Services General Reference.  The grant determines the retiring principal. Other principals might have permission to retire the grant or revoke the grant. For details, see RevokeGrant and Retiring and revoking grants (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete) in the Key Management Service Developer Guide.",
						MarkdownDescription: "The principal that has permission to use the RetireGrant operation to retire the grant.  To specify the principal, use the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an Amazon Web Services principal. Valid Amazon Web Services principals include Amazon Web Services accounts (root), IAM users, federated users, and assumed role users. For examples of the ARN syntax to use for specifying a principal, see Amazon Web Services Identity and Access Management (IAM) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam) in the Example ARNs section of the Amazon Web Services General Reference.  The grant determines the retiring principal. Other principals might have permission to retire the grant or revoke the grant. For details, see RevokeGrant and Retiring and revoking grants (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete) in the Key Management Service Developer Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kms_services_k8s_aws_grant_v1alpha1")

	var state KmsServicesK8SAwsGrantV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KmsServicesK8SAwsGrantV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kms.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Grant")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kms_services_k8s_aws_grant_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kms_services_k8s_aws_grant_v1alpha1")

	var state KmsServicesK8SAwsGrantV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KmsServicesK8SAwsGrantV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kms.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Grant")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *KmsServicesK8SAwsGrantV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kms_services_k8s_aws_grant_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
